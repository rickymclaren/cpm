package z80

import (
	"fmt"
)

var Opcodes = [256]string{
	"noop", "ld bc", "ld (bc), a", "inc bc", "inc b", "dec b", "ld b", "rlca", "ex af,af'", "add hl,bc", "ld a,(bc)", "dec bc", "inc c", "dec c", "ld c", "rrca",
	"djnz", "ld de", "ld (de),a", "inc de", "inc d", "dec d", "ld d", "rla", "jr", "add hl,de", "ld a,(de)", "dec de", "inc e", "dec e", "ld e", "rra",
	"jr nz", "ld hl", "ld (**) hl", "inc hl", "inc h", "dec h", "ld h", "daa", "jr z", "add hl,hl", "ld hl", "dec hl", "inc l", "dec l", "ld l", "cpl",
	"jr nc", "ld sp", "ld (**) a", "inc sp", "inc (hl)", "dec (hl)", "ld (hl)", "scf", "jr c", "add hl,sp", "ld a ()", "dec sp", "inc a", "dec a", "ld a", "ccf",
	"ld b,b", "ld b,c", "ld b,d", "ld b,e", "ld b,h", "ld b,l", "ld b,(hl)", "ld b,a", "ld c,b", "ld c,c", "ld c,d", "ld c,e", "ld c,h", "ld c,l", "ld c,(hl)", "ld c,a",
	"ld d,b", "ld d,c", "ld d,d", "ld d,e", "ld d,h", "ld d,l", "ld d,(hl)", "ld d,a", "ld e,b", "ld e,c", "ld e,d", "ld e,e", "ld e,h", "ld e,l", "ld e,(hl)", "ld e,a",
	"ld h,b", "ld h,c", "ld h,d", "ld h,e", "ld h,h", "ld h,l", "ld h,(hl)", "ld h,a", "ld l,b", "ld l,c", "ld l,d", "ld l,e", "ld l,h", "ld l,l", "ld l,(hl)", "ld l,a",
	"ld (hl),b", "ld (hl),c", "ld (hl),d", "ld (hl),e", "ld (hl),h", "ld (hl),l", "halt", "ld (hl),a", "ld a,b", "ld a,c", "ld a,d", "ld a,e", "ld a,h", "ld a,l", "ld a,(hl)", "ld a,a",
	"add a,b", "add a,c", "add a,d", "add a,e", "add a,h", "add a,l", "add a,(hl)", "add a,a", "adc a,b", "adc a,c", "adc a,d", "adc a,e", "adc a,h", "adc a,l", "adc a,(hl)", "adc a,a",
	"sub b", "sub c", "sub d", "sub e", "sub h", "sub l", "sub (hl)", "sub a", "sbc a,b", "sbc a,c", "sbc a,d", "sbc a,e", "sbc a,h", "sbc a,l", "sbc a,(hl)", "sbc a,a",
	"and b", "and c", "and d", "and e", "and h", "and l", "and (hl)", "and a", "xor b", "xor c", "xor d", "xor e", "xor h", "xor l", "xor (hl)", "xor a",
	"or b", "or c", "or d", "or e", "or h", "or l", "or (hl)", "or a", "cp b", "cp c", "cp d", "cp e", "cp h", "cp l", "cp (hl)", "cp a",
	"ret nz", "pop bc", "jp nz", "jp", "call nz", "push bc", "add a", "rst 00h", "ret z", "ret", "jp z", "BITS", "call z", "call", "adc a", "rst 08h",
	"ret nc", "pop de", "jp nc", "out (*),a", "call nc", "push de", "sub", "rst 10h", "ret c", "exx", "jp c", "in a", "call c", "IX", "sbc a", "rst 18h",
	"ret po", "pop hl", "jp po", "ex (sp),hl", "call po", "push hl", "and", "rst 20h", "ret pe", "jp (hl)", "jp pe", "ex de,hl", "call pe", "EXTD", "xor", "rst 28h",
	"ret p", "pop af", "jp p", "di", "call p", "push af", "or", "rst 30", "ret m", "ld sp,hl", "jp m", "ei", "call m", "IY", "cp", "rst 38",
}

var OpcodeSize = [256]int{
	1, 3, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	2, 3, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1,
	2, 3, 3, 1, 1, 1, 2, 1, 2, 1, 3, 1, 1, 1, 2, 1,
	2, 3, 3, 1, 1, 1, 2, 1, 2, 1, 3, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 1, 2, 1, 1, 1, 3, 1, 3, 3, 2, 1,
	1, 1, 3, 2, 3, 1, 2, 1, 1, 1, 3, 2, 3, 1, 2, 1,
	1, 1, 3, 1, 3, 1, 2, 1, 1, 1, 3, 1, 3, 1, 2, 1,
	1, 1, 3, 1, 3, 1, 2, 1, 1, 1, 3, 1, 3, 1, 2, 1,
}

type CPU struct {
	Memory    [65536]uint8
	AF        Register
	BC        Register
	DE        Register
	HL        Register
	IX        Register
	IY        Register
	PC        uint16
	SP        uint16
	AF_       Register
	BC_       Register
	DE_       Register
	HL_       Register
	Io        IO
	Bdos      BDOS
	Halt      bool
	Debug     bool
	DebugFrom uint16
}

type Register struct {
	Lo uint8
	Hi uint8
}

type IO interface {
	In(port uint8) uint8
	Out(port uint8, value uint8)
}

type BDOS interface {
	TraceBDOS()
}

func (cpu *CPU) Run() {

	const BOOT = 0x0000
	const BDOS = 0xec00
	const BIOS = 0xfa00
	const TPA = 0x0100
	const NO_DEBUG = 0xffff

	cpu.Debug = false
	cpu.DebugFrom = NO_DEBUG

	for !cpu.Halt {
		// time.Sleep(1 * time.Microsecond)

		v := cpu.Memory[cpu.PC]
		opcode := Opcodes[v]
		if opcode == "?" {
			opcode = fmt.Sprintf("0x%02x", v)
		}
		data := "??    "
		switch OpcodeSize[v] {
		case 1:
			data = "      "
		case 2:
			data = fmt.Sprintf("0x%02x  ", cpu.Memory[cpu.PC+1])
		case 3:
			data = fmt.Sprintf("0x%04x", int(cpu.Memory[cpu.PC+1])+int(cpu.Memory[cpu.PC+2])<<8)
		}
		af := cpu.AF.getValue()
		bc := cpu.BC.getValue()
		de := cpu.DE.getValue()
		hl := cpu.HL.getValue()
		sp := cpu.SP
		return_address := uint16(cpu.Memory[cpu.SP]) + uint16(cpu.Memory[cpu.SP+1])<<8
		if cpu.PC == cpu.DebugFrom {
			cpu.Debug = true
		}
		if cpu.Debug && cpu.PC < BDOS {
			fmt.Printf("%04x: %-10s %s - AF:%04x BC:%04x DE:%04x HL:%04x SP:%04x (%04x) %s\n",
				cpu.PC, opcode, data, af, bc, de, hl,
				sp, return_address,
				cpu.flags())
		}
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
