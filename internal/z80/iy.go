package z80

import (
	"fmt"
)

// IY Instructions - FD
const ADD_IY_BC = 0x09
const ADD_IY_DE = 0x19
const LD_IY_XXXX = 0x21
const LD_ADDR_XXXX_IY = 0x22
const INC_IY = 0x23
const INC_IYH = 0x24
const DEC_IYH = 0x25
const LD_IYH_XX = 0x26
const ADD_IY_IY = 0x29
const LD_IY_ADDR_XXXX = 0x2a
const DEC_IY = 0x2b
const INC_IYL = 0x2c
const DEC_IYL = 0x2d
const LD_IYL_XX = 0x2e
const INC_ADDR_IY_PLUS_XX = 0x34
const DEC_ADDR_IY_PLUS_XX = 0x35
const LD_ADDR_IY_PLUS_XX_XX = 0x36
const ADD_IY_SP = 0x39
const IY_40H = 0x40
const IY_41H = 0x41
const IY_42H = 0x42
const IY_43H = 0x43
const LD_B_IYH = 0x44
const LD_B_IYL = 0x45
const LD_B_ADDR_IY_PLUS_XX = 0x46
const IY_47H = 0x47
const IY_48H = 0x48
const IY_49H = 0x49
const IY_4AH = 0x4a
const IY_4BH = 0x4b
const LD_C_IYH = 0x4c
const LD_C_IYL = 0x4d
const LD_C_ADDR_IY_PLUS_XX = 0x4e
const IY_4FH = 0x4f
const IY_50H = 0x50
const IY_51H = 0x51
const IY_52H = 0x52
const IY_53H = 0x53
const LD_D_IYH = 0x54
const LD_D_IYL = 0x55
const LD_D_ADDR_IY_PLUS_XX = 0x56
const IY_57H = 0x57
const IY_58H = 0x58
const IY_59H = 0x59
const IY_5AH = 0x5a
const IY_5BH = 0x5b
const LD_E_IYH = 0x5c
const LD_E_IYL = 0x5d
const LD_E_ADDR_IY_PLUS_XX = 0x5e
const IY_5FH = 0x5f
const LD_IYH_B = 0x60
const LD_IYH_C = 0x61
const LD_IYH_D = 0x62
const LD_IYH_E = 0x63
const LD_IYH_IYH = 0x64
const LD_IYH_IYL = 0x65
const LD_H_ADDR_IY_PLUS_XX = 0x66
const LD_IYH_A = 0x67
const LD_IYL_B = 0x68
const LD_IYL_C = 0x69
const LD_IYL_D = 0x6a
const LD_IYL_E = 0x6b
const LD_IYL_IYH = 0x6c
const LD_IYL_IYL = 0x6d
const LD_L_ADDR_IY_PLUS_XX = 0x6e
const LD_IYL_A = 0x6f
const LD_ADDR_IY_PLUS_XX_B = 0x70
const LD_ADDR_IY_PLUS_XX_C = 0x71
const LD_ADDR_IY_PLUS_XX_D = 0x72
const LD_ADDR_IY_PLUS_XX_E = 0x73
const LD_ADDR_IY_PLUS_XX_H = 0x74
const LD_ADDR_IY_PLUS_XX_L = 0x75
const IY_76H = 0x76
const LD_ADDR_IY_PLUS_XX_A = 0x77
const IY_78H = 0x78
const IY_79H = 0x79
const IY_7AH = 0x7a
const IY_7BH = 0x7b
const LD_A_IYH = 0x7c
const LD_A_IYL = 0x7d
const LD_A_ADDR_IY_PLUS_XX = 0x7e
const IY_7FH = 0x7f
const ADD_A_IYH = 0x84
const ADD_A_IYL = 0x85
const ADD_A_ADDR_IY_PLUS_XX = 0x86
const ADC_A_IYH = 0x8c
const ADC_A_IYL = 0x8d
const ADC_A_ADDR_IY_PLUS_XX = 0x8e
const SUB_A_IYH = 0x94
const SUB_A_IYL = 0x95
const SUB_A_ADDR_IY_PLUS_XX = 0x96
const SBC_A_IYH = 0x9c
const SBC_A_IYL = 0x9d
const SBC_A_ADDR_IY_PLUS_XX = 0x9e
const AND_A_IYH = 0xa4
const AND_A_IYL = 0xa5
const AND_A_ADDR_IY_PLUS_XX = 0xa6
const XOR_A_IYH = 0xac
const XOR_A_IYL = 0xad
const XOR_A_ADDR_IY_PLUS_XX = 0xae
const OR_A_IYH = 0xb4
const OR_A_IYL = 0xb5
const OR_A_ADDR_IY_PLUS_XX = 0xb6
const CP_A_IYH = 0xbc
const CP_A_IYL = 0xbd
const CP_A_ADDR_IY_PLUS_XX = 0xbe
const BITS_IY = 0xcb
const POP_IY = 0xe1
const EX_ADDR_SP_IY = 0xe3
const PUSH_IY = 0xe5
const JP_IY = 0xe9
const LD_SP_IY = 0xf9

func (cpu *CPU) iy(code uint8) {
	switch code {

	case LD_A_IYH:
		cpu.AF.Hi = cpu.IY.Hi

	case LD_A_IYL:
		cpu.AF.Hi = cpu.IY.Lo

	case LD_IY_XXXX:
		cpu.IY.setValue(cpu.fetch16())

	case LD_IYH_XX:
		cpu.IY.Hi = cpu.fetch()

	case LD_IYL_XX:
		cpu.IY.Lo = cpu.fetch()

	case LD_IYH_A:
		cpu.IY.Hi = cpu.AF.Hi

	case LD_IYH_B:
		cpu.IY.Hi = cpu.BC.Hi

	case LD_IYH_C:
		cpu.IY.Hi = cpu.BC.Lo

	case LD_IYH_D:
		cpu.IY.Hi = cpu.DE.Hi

	case LD_IYH_E:
		cpu.IY.Hi = cpu.DE.Lo

	case LD_IYH_IYH:
		// noop

	case LD_IYH_IYL:
		cpu.IY.Hi = cpu.IY.Lo

	case LD_IYL_A:
		cpu.IY.Lo = cpu.AF.Hi

	case LD_IYL_B:
		cpu.IY.Lo = cpu.BC.Hi

	case LD_IYL_C:
		cpu.IY.Lo = cpu.BC.Lo

	case LD_IYL_D:
		cpu.IY.Lo = cpu.DE.Hi

	case LD_IYL_E:
		cpu.IY.Lo = cpu.DE.Lo

	case LD_IYL_IYH:
		cpu.IY.Lo = cpu.IY.Hi

	case LD_IYL_IYL:
		// noop

	case LD_ADDR_IY_PLUS_XX_A:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.AF.Hi

	case LD_ADDR_IY_PLUS_XX_B:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.BC.Hi

	case LD_ADDR_IY_PLUS_XX_C:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.BC.Lo

	case LD_ADDR_IY_PLUS_XX_D:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.DE.Hi

	case LD_ADDR_IY_PLUS_XX_E:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.DE.Lo

	case LD_ADDR_IY_PLUS_XX_H:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.HL.Hi

	case LD_ADDR_IY_PLUS_XX_L:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.HL.Lo

	case LD_ADDR_XXXX_IY:
		address := cpu.fetch16()
		cpu.Memory[address] = cpu.IY.Lo
		cpu.Memory[address+1] = cpu.IY.Hi

	case LD_IY_ADDR_XXXX:
		address := cpu.fetch16()
		cpu.IY.Lo = cpu.Memory[address]
		cpu.IY.Hi = cpu.Memory[address+1]

	case LD_ADDR_IY_PLUS_XX_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		value := cpu.fetch()
		cpu.Memory[address] = value

	case LD_A_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.Memory[address]

	case LD_B_IYH:
		cpu.BC.Hi = cpu.IY.Hi

	case LD_B_IYL:
		cpu.BC.Hi = cpu.IY.Lo

	case LD_B_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.BC.Hi = cpu.Memory[address]

	case LD_C_IYH:
		cpu.BC.Lo = cpu.IY.Hi

	case LD_C_IYL:
		cpu.BC.Lo = cpu.IY.Lo

	case LD_C_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.BC.Lo = cpu.Memory[address]

	case LD_D_IYH:
		cpu.DE.Hi = cpu.IY.Hi

	case LD_D_IYL:
		cpu.DE.Hi = cpu.IY.Lo

	case LD_D_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.DE.Hi = cpu.Memory[address]

	case LD_E_IYH:
		cpu.DE.Lo = cpu.IY.Hi

	case LD_E_IYL:
		cpu.DE.Lo = cpu.IY.Lo

	case LD_E_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.DE.Lo = cpu.Memory[address]

	case LD_H_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.HL.Hi = cpu.Memory[address]

	case LD_L_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.HL.Lo = cpu.Memory[address]

	case INC_IY:
		cpu.IY.setValue(cpu.IY.getValue() + 1)

	case INC_IYH:
		cpu.IY.Hi = cpu.add(cpu.IY.Hi, 1, false)

	case INC_IYL:
		cpu.IY.Lo = cpu.add(cpu.IY.Lo, 1, false)

	case DEC_IY:
		cpu.IY.setValue(cpu.IY.getValue() - 1)

	case DEC_IYH:
		cpu.IY.Hi = cpu.sub(cpu.IY.Hi, 1, false)

	case DEC_IYL:
		cpu.IY.Lo = cpu.sub(cpu.IY.Lo, 1, false)

	case INC_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.inc(cpu.Memory[address])

	case DEC_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.dec(cpu.Memory[address])

	case ADD_IY_BC:
		cpu.IY.setValue(cpu.add16(cpu.IY.getValue(), cpu.BC.getValue()))

	case ADD_IY_DE:
		cpu.IY.setValue(cpu.add16(cpu.IY.getValue(), cpu.DE.getValue()))

	case ADD_IY_IY:
		cpu.IY.setValue(cpu.add16(cpu.IY.getValue(), cpu.IY.getValue()))

	case ADD_IY_SP:
		cpu.IY.setValue(cpu.add16(cpu.IY.getValue(), cpu.SP))

	case ADD_A_IYH:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.IY.Hi, false)

	case ADD_A_IYL:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.IY.Lo, false)

	case ADD_A_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.Memory[address], false)

	case ADC_A_IYH:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.IY.Hi, cpu.flagC())

	case ADC_A_IYL:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.IY.Lo, cpu.flagC())

	case ADC_A_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.Memory[address], cpu.flagC())

	case SUB_A_IYH:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.IY.Hi, false)

	case SUB_A_IYL:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.IY.Lo, false)

	case SUB_A_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.Memory[address], false)

	case SBC_A_IYH:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.IY.Hi, cpu.flagC())

	case SBC_A_IYL:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.IY.Lo, cpu.flagC())

	case SBC_A_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.Memory[address], cpu.flagC())

	case CP_A_IYH:
		cpu.sub(cpu.AF.Hi, cpu.IY.Hi, false)

	case CP_A_IYL:
		cpu.sub(cpu.AF.Hi, cpu.IY.Lo, false)

	case CP_A_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.sub(cpu.AF.Hi, cpu.Memory[address], false)

	case AND_A_IYH:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.IY.Hi)

	case AND_A_IYL:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.IY.Lo)

	case AND_A_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.Memory[address])

	case OR_A_IYH:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.IY.Hi)

	case OR_A_IYL:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.IY.Lo)

	case OR_A_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.Memory[address])

	case XOR_A_IYH:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.IY.Hi)

	case XOR_A_IYL:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.IY.Lo)

	case XOR_A_ADDR_IY_PLUS_XX:
		address := addrOff(cpu.IY.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.Memory[address])

	case PUSH_IY:
		cpu.push(cpu.IY.getValue())

	case POP_IY:
		cpu.IY.setValue(cpu.pop())

	case BITS_IY:
		cpu.ixyBits(cpu.IY.getValue())

	case IY_40H:
		//cpu.BC.Hi = cpu.BC.Hi

	case IY_41H:
		cpu.BC.Hi = cpu.BC.Lo

	case IY_42H:
		cpu.BC.Hi = cpu.DE.Hi

	case IY_43H:
		cpu.BC.Hi = cpu.DE.Lo

	case IY_47H:
		cpu.BC.Hi = cpu.AF.Hi

	case IY_48H:
		cpu.BC.Lo = cpu.BC.Hi

	case IY_49H:
		//cpu.BC.Lo = cpu.BC.Lo

	case IY_4AH:
		cpu.BC.Lo = cpu.DE.Hi

	case IY_4BH:
		cpu.BC.Lo = cpu.DE.Lo

	case IY_4FH:
		cpu.BC.Lo = cpu.AF.Hi

	case IY_50H:
		cpu.DE.Hi = cpu.BC.Hi

	case IY_51H:
		cpu.DE.Hi = cpu.BC.Lo

	case IY_52H:
		//cpu.DE.Hi = cpu.DE.Hi

	case IY_53H:
		cpu.DE.Hi = cpu.DE.Lo

	case IY_57H:
		cpu.DE.Hi = cpu.AF.Hi

	case IY_58H:
		cpu.DE.Lo = cpu.BC.Hi

	case IY_59H:
		cpu.DE.Lo = cpu.BC.Lo

	case IY_5AH:
		cpu.DE.Lo = cpu.DE.Hi

	case IY_5BH:
		//cpu.DE.Lo = cpu.DE.Lo

	case IY_5FH:
		cpu.DE.Lo = cpu.AF.Hi

	case IY_78H:
		cpu.AF.Hi = cpu.BC.Hi

	case IY_79H:
		cpu.AF.Hi = cpu.BC.Lo

	case IY_7AH:
		cpu.AF.Hi = cpu.DE.Hi

	case IY_7BH:
		cpu.AF.Hi = cpu.DE.Lo

	case IY_7FH:
		//cpu.AF.Hi = cpu.AF.Hi

	case EX_ADDR_SP_IY:
		addr := cpu.SP
		tmp := cpu.IY.Lo
		cpu.IY.Lo = cpu.Memory[addr]
		cpu.Memory[addr] = tmp
		addr++
		tmp = cpu.IY.Hi
		cpu.IY.Hi = cpu.Memory[addr]
		cpu.Memory[addr] = tmp

	case JP_IY:
		cpu.PC = cpu.IY.getValue()

	case LD_SP_IY:
		cpu.SP = cpu.IY.getValue()

	default:
		fmt.Printf("Unknown IY opcode 0x%02x at %04x\n", code, cpu.PC)
	}
}
