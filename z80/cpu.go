package z80

type CPU struct {
	Memory [65536]uint8
	AF     Register
	BC     Register
	DE     Register
	HL     Register
	IX     Register
	IY     Register
	PC     uint16
	SP     uint16
	AF_    Register
	BC_    Register
	DE_    Register
	HL_    Register
	Io     IO
	Halt   bool
}

type Register struct {
	Lo uint8
	Hi uint8
}

type IO interface {
	In(port uint8) uint8
	Out(port uint8, value uint8)
}

func (cpu *CPU) Run() {

	for !cpu.Halt {
		// time.Sleep(1 * time.Microsecond)
		cpu.Step()
	}
}

func (cpu *CPU) Step() {
	opcode := cpu.fetch()
	cpu.opcode(opcode)
}

func (cpu *CPU) fetch() uint8 {
	data := cpu.Memory[cpu.PC]
	cpu.PC++
	return data
}

func (cpu *CPU) fetch16() uint16 {
	lo := cpu.Memory[cpu.PC]
	cpu.PC++
	hi := cpu.Memory[cpu.PC]
	cpu.PC++
	return uint16(hi)<<8 + uint16(lo)
}

func (cpu *CPU) fetchRegister() Register {
	register := Register{Lo: cpu.Memory[cpu.PC], Hi: cpu.Memory[cpu.PC+1]}
	cpu.PC += 2
	return register
}

func (cpu *CPU) flagC() bool {
	return cpu.AF.Lo&FLG_C != 0
}

func (cpu *CPU) flagN() bool {
	return cpu.AF.Lo&FLG_N != 0
}

func (cpu *CPU) flagH() bool {
	return cpu.AF.Lo&FLG_H != 0
}

func (cpu *CPU) flagZ() bool {
	return cpu.AF.Lo&FLG_Z != 0
}

func (cpu *CPU) add(o uint8, v uint8, c bool) uint8 {
	n := uint16(o) + uint16(v)
	if c {
		n++
	}
	r := uint8(n)
	cpu.updateFlagArith8(n, uint16(o), uint16(v), false)
	return r
}

func (cpu *CPU) add16(o uint16, v uint16) uint16 {
	n := uint32(o) + uint32(v)
	c := n ^ uint32(o) ^ uint32(v)
	var nand uint8 = FLG_H | FLG_N | FLG_C | FLG_5 | FLG_3
	var or uint8
	or |= uint8(n>>8) & (FLG_5 | FLG_3)
	or |= uint8(c>>8) & FLG_H
	or |= uint8(n>>16) & FLG_C
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return uint16(n)
}

func (cpu *CPU) adc16(o uint16, v uint16) uint16 {
	n := uint32(o) + uint32(v)
	if cpu.flagC() {
		n++
	}
	c := n ^ uint32(o) ^ uint32(v)
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N | FLG_C
	var or uint8
	or |= uint8(n>>8) & FLG_STD
	if uint16(n) == 0 {
		or |= FLG_Z
	}
	or |= uint8(c>>8) & FLG_H
	or |= uint8((c>>14)^(c>>13)) & FLG_P_V
	or |= uint8(n>>16) & FLG_C
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return uint16(n)
}

func (cpu *CPU) sub(o uint8, v uint8, c bool) uint8 {
	n := uint16(o) - uint16(v)
	if c {
		n--
	}
	r := uint8(n)
	cpu.updateFlagArith8(n, uint16(o), uint16(v), true)
	return r
}

func (cpu *CPU) sub16(a uint16, b uint16) uint16 {
	a32, b32 := uint32(a), uint32(b)
	r := a32 - b32 - uint32(cpu.AF.Lo&FLG_C)
	c := r ^ a32 ^ b32
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N | FLG_C
	var or uint8
	or |= uint8(r>>8) & FLG_STD
	if uint16(r) == 0 {
		or |= FLG_Z
	}
	or |= uint8(c>>8) & FLG_H
	or |= uint8((c>>14)^(c>>13)) & FLG_P_V
	or |= FLG_N
	or |= uint8(r>>16) & FLG_C
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return uint16(r)
}

func (cpu *CPU) inc(o uint8) uint8 {
	r := o + 1
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N
	var or uint8
	or |= r & FLG_STD
	if r == 0 {
		or |= FLG_Z
	}
	if r&0x0f == 0 {
		or |= FLG_H
	}
	if o == 0x7f {
		or |= FLG_P_V
	}
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return r
}

func (cpu *CPU) dec(o uint8) uint8 {
	r := o - 1
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N
	var or uint8
	or |= r & FLG_STD
	if r == 0 {
		or |= FLG_Z
	}
	if r&0x0f == 0x0f {
		or |= FLG_H
	}
	if o == 0x80 {
		or |= FLG_P_V
	}
	or |= FLG_N
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return r
}

func (cpu *CPU) and(o uint8, v uint8) uint8 {
	n := o & v
	cpu.updateFlagLogic8(n, true)
	return n
}

func (cpu *CPU) or(o uint8, v uint8) uint8 {
	n := o | v
	cpu.updateFlagLogic8(n, false)
	return n
}

func (cpu *CPU) xor(o uint8, v uint8) uint8 {
	n := o ^ v
	cpu.updateFlagLogic8(n, false)
	return n
}

func (register *Register) getValue() uint16 {
	return uint16(register.Hi)<<8 + uint16(register.Lo)
}

func (register *Register) setValue(value uint16) {
	register.Hi = uint8(value >> 8)
	register.Lo = uint8(value & 0xff)
}
