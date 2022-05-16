package z80

type operation func(uint8) uint8

func (cpu *CPU) bits(code uint8) {
	bitop := code >> 6
	bit := code >> 3 & 0x7
	register := code & 0x7
	switch bitop {
	case 0:
		switch code >> 3 & 0x07 {
		case 0: // rlc
			cpu.rotate_shift(code, func(a uint8) uint8 {
				return a<<1 | a>>7
			}, 1)
		case 1: // rrc
			cpu.rotate_shift(code, func(a uint8) uint8 {
				return a>>1 | a<<7
			}, 0)
		case 2: // rl
			cpu.rotate_shift(code, func(a uint8) uint8 {
				return a<<1 | cpu.AF.Lo&FLG_C
			}, 1)
		case 3: // rr
			cpu.rotate_shift(code, func(a uint8) uint8 {
				return a>>1 | (cpu.AF.Lo&FLG_C)<<7
			}, 0)
		case 4: //sla
			cpu.rotate_shift(code, func(a uint8) uint8 {
				return a << 1
			}, 1)
		case 5: // sra
			cpu.rotate_shift(code, func(a uint8) uint8 {
				return a&0x80 | a>>1
			}, 0)
		case 6: // sll
			cpu.rotate_shift(code, func(a uint8) uint8 {
				return a<<1 + 1
			}, 1)
		case 7: // srl
			cpu.rotate_shift(code, func(a uint8) uint8 {
				return a >> 1
			}, 0)
		}
	case 1:
		cpu.testBit(bit, register)

	case 2:
		cpu.resetBit(bit, register)

	case 3:
		cpu.setBit(bit, register)
	}

}

func (cpu *CPU) rotate_shift(code uint8, op operation, flags int) {
	var a, r uint8
	switch code & 0x07 {
	case 0:
		a = cpu.BC.Hi
		r = op(a)
		cpu.BC.Hi = r

	case 1:
		a = cpu.BC.Lo
		r = op(a)
		cpu.BC.Lo = r

	case 2:
		a = cpu.DE.Hi
		r = op(a)
		cpu.DE.Hi = r

	case 3:
		a = cpu.DE.Lo
		r = op(a)
		cpu.DE.Lo = r

	case 4:
		a = cpu.HL.Hi
		r = op(a)
		cpu.HL.Hi = r

	case 5:
		a = cpu.HL.Lo
		r = op(a)
		cpu.HL.Lo = r

	case 6:
		a = cpu.Memory[cpu.HL.getValue()]
		r = op(a)
		cpu.Memory[cpu.HL.getValue()] = r

	case 7:
		a = cpu.AF.Hi
		r = op(a)
		cpu.AF.Hi = r

	}
	if flags == 0 {
		cpu.updateFlagBitop(r, a)
	} else {
		cpu.updateFlagBitop(r, a>>7)
	}

}

func (cpu *CPU) rotate_shift_address(address uint16, op operation, flags int) {
	a := cpu.Memory[address]
	r := op(a)
	cpu.Memory[address] = r
	if flags == 0 {
		cpu.updateFlagBitop(r, a)
	} else {
		cpu.updateFlagBitop(r, a>>7)
	}
}

func (cpu *CPU) testBit(bit uint8, register uint8) {
	var value uint8

	switch register {
	case 0:
		value = cpu.BC.Hi

	case 1:
		value = cpu.BC.Lo

	case 2:
		value = cpu.DE.Hi

	case 3:
		value = cpu.DE.Lo

	case 4:
		value = cpu.HL.Hi

	case 5:
		value = cpu.HL.Lo

	case 6:
		value = cpu.Memory[cpu.HL.getValue()]

	case 7:
		value = cpu.AF.Hi

	}

	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N
	var or uint8
	if value&(0x01<<bit) == 0 {
		or |= FLG_Z | FLG_P_V
	} else if bit == 7 {
		or |= FLG_S
	}
	or |= FLG_H
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func (cpu *CPU) resetBit(bit uint8, register uint8) {
	mask := uint8(0x01) << bit

	switch register {
	case 0:
		cpu.BC.Hi = cpu.BC.Hi &^ mask

	case 1:
		cpu.BC.Lo = cpu.BC.Lo &^ mask

	case 2:
		cpu.DE.Hi = cpu.DE.Hi &^ mask

	case 3:
		cpu.DE.Lo = cpu.DE.Lo &^ mask

	case 4:
		cpu.HL.Hi = cpu.HL.Hi &^ mask

	case 5:
		cpu.HL.Lo = cpu.HL.Lo &^ mask

	case 6:
		cpu.Memory[cpu.HL.getValue()] = cpu.Memory[cpu.HL.getValue()] &^ mask

	case 7:
		cpu.AF.Hi = cpu.AF.Hi &^ mask

	}

}

func (cpu *CPU) setBit(bit uint8, register uint8) {
	mask := uint8(0x01) << bit

	switch register {
	case 0:
		cpu.BC.Hi = cpu.BC.Hi | mask

	case 1:
		cpu.BC.Lo = cpu.BC.Lo | mask

	case 2:
		cpu.DE.Hi = cpu.DE.Hi | mask

	case 3:
		cpu.DE.Lo = cpu.DE.Lo | mask

	case 4:
		cpu.HL.Hi = cpu.HL.Hi | mask

	case 5:
		cpu.HL.Lo = cpu.HL.Lo | mask

	case 6:
		cpu.Memory[cpu.HL.getValue()] = cpu.Memory[cpu.HL.getValue()] | mask

	case 7:
		cpu.AF.Hi = cpu.AF.Hi | mask

	}

}

func (cpu *CPU) ixyBits(ixy uint16) {
	address := addrOff(ixy, cpu.fetch())
	value := cpu.Memory[address]
	opcode := cpu.fetch()
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N
	var or uint8 = FLG_H
	switch opcode {
	case 0x06, 0x0e, 0x16, 0x1e, 0x26, 0x2e, 0x36, 0x3e:
		code := (opcode - 0x06) >> 3 & 0x07
		switch code {
		case 0: // rlc
			cpu.rotate_shift_address(address, func(a uint8) uint8 {
				return a<<1 | a>>7
			}, 1)
		case 1: // rrc
			cpu.rotate_shift_address(address, func(a uint8) uint8 {
				return a>>1 | a<<7
			}, 0)
		case 2: // rl
			cpu.rotate_shift_address(address, func(a uint8) uint8 {
				return a<<1 | cpu.AF.Lo&FLG_C
			}, 1)
		case 3: // rr
			cpu.rotate_shift_address(address, func(a uint8) uint8 {
				return a>>1 | (cpu.AF.Lo&FLG_C)<<7
			}, 0)
		case 4: //sla
			cpu.rotate_shift_address(address, func(a uint8) uint8 {
				return a << 1
			}, 1)
		case 5: // sra
			cpu.rotate_shift_address(address, func(a uint8) uint8 {
				return a&0x80 | a>>1
			}, 0)
		case 6: // sll
			cpu.rotate_shift_address(address, func(a uint8) uint8 {
				return a<<1 + 1
			}, 1)
		case 7: // srl
			cpu.rotate_shift_address(address, func(a uint8) uint8 {
				return a >> 1
			}, 0)
		}

	case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x76, 0x7e:
		// test bit (ix+xx)
		bit := ((opcode - 0x46) >> 3) & 0x07
		if value&(0x01<<bit) == 0 {
			or |= FLG_Z | FLG_P_V
		} else if bit == 7 {
			or |= FLG_S
		}
		cpu.AF.Lo = cpu.AF.Lo&^nand | or

	case 0x86, 0x8e, 0x96, 0x9e, 0xa6, 0xae, 0xb6, 0xbe:
		// reset bit (ix+xx)
		bit := ((opcode - 0x86) >> 3) & 0x07
		mask := uint8(1) << bit
		cpu.Memory[address] = value &^ mask

	case 0xc6, 0xce, 0xd6, 0xde, 0xe6, 0xee, 0xf6, 0xfe:
		// set bit (ix+xx)
		bit := ((opcode - 0xc6) >> 3) & 0x07
		mask := uint8(1) << bit
		cpu.Memory[address] = value | mask

	}
}
