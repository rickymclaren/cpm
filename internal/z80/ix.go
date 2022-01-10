package z80

import (
	"fmt"
)

// IX Instructions - DD
const ADD_IX_BC = 0x09
const ADD_IX_DE = 0x19
const LD_IX_XXXX = 0x21
const LD_ADDR_XXXX_IX = 0x22
const INC_IX = 0x23
const INC_IXH = 0x24
const DEC_IXH = 0x25
const LD_IXH_XX = 0x26
const ADD_IX_IX = 0x29
const LD_IX_ADDR_XXXX = 0x2a
const DEC_IX = 0x2b
const INC_IXL = 0x2c
const DEC_IXL = 0x2d
const LD_IXL_XX = 0x2e
const INC_ADDR_IX_PLUS_XX = 0x34
const DEC_ADDR_IX_PLUS_XX = 0x35
const LD_ADDR_IX_PLUS_XX_XX = 0x36
const ADD_IX_SP = 0x39
const IX_40H = 0x40
const IX_41H = 0x41
const IX_42H = 0x42
const IX_43H = 0x43
const LD_B_IXH = 0x44
const LD_B_IXL = 0x45
const LD_B_ADDR_IX_PLUS_XX = 0x46
const IX_47H = 0x47
const IX_48H = 0x48
const IX_49H = 0x49
const IX_4AH = 0x4a
const IX_4BH = 0x4b
const LD_C_IXH = 0x4c
const LD_C_IXL = 0x4d
const LD_C_ADDR_IX_PLUS_XX = 0x4e
const IX_4FH = 0x4f
const IX_50H = 0x50
const IX_51H = 0x51
const IX_52H = 0x52
const IX_53H = 0x53
const LD_D_IXH = 0x54
const LD_D_IXL = 0x55
const LD_D_ADDR_IX_PLUS_XX = 0x56
const IX_57H = 0x57
const IX_58H = 0x58
const IX_59H = 0x59
const IX_5AH = 0x5a
const IX_5BH = 0x5b
const LD_E_IXH = 0x5c
const LD_E_IXL = 0x5d
const LD_E_ADDR_IX_PLUS_XX = 0x5e
const IX_5FH = 0x5f
const LD_IXH_B = 0x60
const LD_IXH_C = 0x61
const LD_IXH_D = 0x62
const LD_IXH_E = 0x63
const LD_IXH_IXH = 0x64
const LD_IXH_IXL = 0x65
const LD_H_ADDR_IX_PLUS_XX = 0x66
const LD_IXH_A = 0x67
const LD_IXL_B = 0x68
const LD_IXL_C = 0x69
const LD_IXL_D = 0x6a
const LD_IXL_E = 0x6b
const LD_IXL_IXH = 0x6c
const LD_IXL_IXL = 0x6d
const LD_L_ADDR_IX_PLUS_XX = 0x6e
const LD_IXL_A = 0x6f
const LD_ADDR_IX_PLUS_XX_B = 0x70
const LD_ADDR_IX_PLUS_XX_C = 0x71
const LD_ADDR_IX_PLUS_XX_D = 0x72
const LD_ADDR_IX_PLUS_XX_E = 0x73
const LD_ADDR_IX_PLUS_XX_H = 0x74
const LD_ADDR_IX_PLUS_XX_L = 0x75
const IX_76H = 0x76
const LD_ADDR_IX_PLUS_XX_A = 0x77
const IX_78H = 0x78
const IX_79H = 0x79
const IX_7AH = 0x7a
const IX_7BH = 0x7b
const LD_A_IXH = 0x7c
const LD_A_IXL = 0x7d
const LD_A_ADDR_IX_PLUS_XX = 0x7e
const IX_7FH = 0x7f
const ADD_A_IXH = 0x84
const ADD_A_IXL = 0x85
const ADD_A_ADDR_IX_PLUS_XX = 0x86
const ADC_A_IXH = 0x8c
const ADC_A_IXL = 0x8d
const ADC_A_ADDR_IX_PLUS_XX = 0x8e
const SUB_A_IXH = 0x94
const SUB_A_IXL = 0x95
const SUB_A_ADDR_IX_PLUS_XX = 0x96
const SBC_A_IXH = 0x9c
const SBC_A_IXL = 0x9d
const SBC_A_ADDR_IX_PLUS_XX = 0x9e
const AND_A_IXH = 0xa4
const AND_A_IXL = 0xa5
const AND_A_ADDR_IX_PLUS_XX = 0xa6
const XOR_A_IXH = 0xac
const XOR_A_IXL = 0xad
const XOR_A_ADDR_IX_PLUS_XX = 0xae
const OR_A_IXH = 0xb4
const OR_A_IXL = 0xb5
const OR_A_ADDR_IX_PLUS_XX = 0xb6
const CP_A_IXH = 0xbc
const CP_A_IXL = 0xbd
const CP_A_ADDR_IX_PLUS_XX = 0xbe
const BITS_IX = 0xcb
const POP_IX = 0xe1
const EX_ADDR_SP_IX = 0xe3
const PUSH_IX = 0xe5
const JP_IX = 0xe9
const LD_SP_IX = 0xf9

func (cpu *CPU) ix(code uint8) {
	switch code {

	case LD_A_IXH:
		cpu.AF.Hi = cpu.IX.Hi

	case LD_A_IXL:
		cpu.AF.Hi = cpu.IX.Lo

	case LD_IX_XXXX:
		cpu.IX.setValue(cpu.fetch16())

	case LD_IXH_XX:
		cpu.IX.Hi = cpu.fetch()

	case LD_IXL_XX:
		cpu.IX.Lo = cpu.fetch()

	case LD_ADDR_XXXX_IX:
		address := cpu.fetch16()
		cpu.Memory[address] = cpu.IX.Lo
		cpu.Memory[address+1] = cpu.IX.Hi

	case LD_IX_ADDR_XXXX:
		address := cpu.fetch16()
		cpu.IX.Lo = cpu.Memory[address]
		cpu.IX.Hi = cpu.Memory[address+1]

	case LD_ADDR_IX_PLUS_XX_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		value := cpu.fetch()
		cpu.Memory[address] = value

	case LD_ADDR_IX_PLUS_XX_A:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.AF.Hi

	case LD_ADDR_IX_PLUS_XX_B:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.BC.Hi

	case LD_ADDR_IX_PLUS_XX_C:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.BC.Lo

	case LD_ADDR_IX_PLUS_XX_D:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.DE.Hi

	case LD_ADDR_IX_PLUS_XX_E:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.DE.Lo

	case LD_ADDR_IX_PLUS_XX_H:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.HL.Hi

	case LD_ADDR_IX_PLUS_XX_L:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.HL.Lo

	case LD_A_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.Memory[address]

	case LD_B_IXH:
		cpu.BC.Hi = cpu.IX.Hi

	case LD_B_IXL:
		cpu.BC.Hi = cpu.IX.Lo

	case LD_B_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.BC.Hi = cpu.Memory[address]

	case LD_C_IXH:
		cpu.BC.Lo = cpu.IX.Hi

	case LD_C_IXL:
		cpu.BC.Lo = cpu.IX.Lo

	case LD_C_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.BC.Lo = cpu.Memory[address]

	case LD_D_IXH:
		cpu.DE.Hi = cpu.IX.Hi

	case LD_D_IXL:
		cpu.DE.Hi = cpu.IX.Lo

	case LD_D_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.DE.Hi = cpu.Memory[address]

	case LD_E_IXH:
		cpu.DE.Lo = cpu.IX.Hi

	case LD_E_IXL:
		cpu.DE.Lo = cpu.IX.Lo

	case LD_E_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.DE.Lo = cpu.Memory[address]

	case LD_IXH_A:
		cpu.IX.Hi = cpu.AF.Hi

	case LD_IXH_B:
		cpu.IX.Hi = cpu.BC.Hi

	case LD_IXH_C:
		cpu.IX.Hi = cpu.BC.Lo

	case LD_IXH_D:
		cpu.IX.Hi = cpu.DE.Hi

	case LD_IXH_E:
		cpu.IX.Hi = cpu.DE.Lo

	case LD_IXH_IXH:
		// noop

	case LD_IXH_IXL:
		cpu.IX.Hi = cpu.IX.Lo

	case LD_IXL_A:
		cpu.IX.Lo = cpu.AF.Hi

	case LD_IXL_B:
		cpu.IX.Lo = cpu.BC.Hi

	case LD_IXL_C:
		cpu.IX.Lo = cpu.BC.Lo

	case LD_IXL_D:
		cpu.IX.Lo = cpu.DE.Hi

	case LD_IXL_E:
		cpu.IX.Lo = cpu.DE.Lo

	case LD_IXL_IXH:
		cpu.IX.Lo = cpu.IX.Hi

	case LD_IXL_IXL:
		// noop

	case LD_H_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.HL.Hi = cpu.Memory[address]

	case LD_L_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.HL.Lo = cpu.Memory[address]

	case INC_IX:
		cpu.IX.setValue(cpu.IX.getValue() + 1)

	case INC_IXH:
		cpu.IX.Hi = cpu.inc(cpu.IX.Hi)

	case INC_IXL:
		cpu.IX.Lo = cpu.inc(cpu.IX.Lo)

	case DEC_IX:
		cpu.IX.setValue(cpu.IX.getValue() - 1)

	case DEC_IXH:
		cpu.IX.Hi = cpu.dec(cpu.IX.Hi)

	case DEC_IXL:
		cpu.IX.Lo = cpu.dec(cpu.IX.Lo)

	case INC_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.inc(cpu.Memory[address])

	case DEC_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.Memory[address] = cpu.dec(cpu.Memory[address])

	case ADD_IX_BC:
		cpu.IX.setValue(cpu.add16(cpu.IX.getValue(), cpu.BC.getValue()))

	case ADD_IX_DE:
		cpu.IX.setValue(cpu.add16(cpu.IX.getValue(), cpu.DE.getValue()))

	case ADD_IX_IX:
		cpu.IX.setValue(cpu.add16(cpu.IX.getValue(), cpu.IX.getValue()))

	case ADD_IX_SP:
		cpu.IX.setValue(cpu.add16(cpu.IX.getValue(), cpu.SP))

	case ADD_A_IXH:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.IX.Hi, false)

	case ADD_A_IXL:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.IX.Lo, false)

	case ADD_A_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.Memory[address], false)

	case ADC_A_IXH:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.IX.Hi, cpu.flagC())

	case ADC_A_IXL:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.IX.Lo, cpu.flagC())

	case ADC_A_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.Memory[address], cpu.flagC())

	case SUB_A_IXH:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.IX.Hi, false)

	case SUB_A_IXL:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.IX.Lo, false)

	case SUB_A_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.Memory[address], false)

	case SBC_A_IXH:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.IX.Hi, cpu.flagC())

	case SBC_A_IXL:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.IX.Lo, cpu.flagC())

	case SBC_A_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.Memory[address], cpu.flagC())

	case CP_A_IXH:
		cpu.sub(cpu.AF.Hi, cpu.IX.Hi, false)

	case CP_A_IXL:
		cpu.sub(cpu.AF.Hi, cpu.IX.Lo, false)

	case CP_A_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.sub(cpu.AF.Hi, cpu.Memory[address], false)

	case AND_A_IXH:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.IX.Hi)

	case AND_A_IXL:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.IX.Lo)

	case AND_A_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.Memory[address])

	case OR_A_IXH:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.IX.Hi)

	case OR_A_IXL:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.IX.Lo)

	case OR_A_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.Memory[address])

	case XOR_A_IXH:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.IX.Hi)

	case XOR_A_IXL:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.IX.Lo)

	case XOR_A_ADDR_IX_PLUS_XX:
		address := addrOff(cpu.IX.getValue(), cpu.fetch())
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.Memory[address])

	case PUSH_IX:
		cpu.push(cpu.IX.getValue())

	case POP_IX:
		cpu.IX.setValue(cpu.pop())

	case BITS_IX:
		cpu.ixyBits(cpu.IX.getValue())

	case IX_40H:
		//cpu.BC.Hi = cpu.BC.Hi

	case IX_41H:
		cpu.BC.Hi = cpu.BC.Lo

	case IX_42H:
		cpu.BC.Hi = cpu.DE.Hi

	case IX_43H:
		cpu.BC.Hi = cpu.DE.Lo

	case IX_47H:
		cpu.BC.Hi = cpu.AF.Hi

	case IX_48H:
		cpu.BC.Lo = cpu.BC.Hi

	case IX_49H:
		//cpu.BC.Lo = cpu.BC.Lo

	case IX_4AH:
		cpu.BC.Lo = cpu.DE.Hi

	case IX_4BH:
		cpu.BC.Lo = cpu.DE.Lo

	case IX_4FH:
		cpu.BC.Lo = cpu.AF.Hi

	case IX_50H:
		cpu.DE.Hi = cpu.BC.Hi

	case IX_51H:
		cpu.DE.Hi = cpu.BC.Lo

	case IX_52H:
		//cpu.DE.Hi = cpu.DE.Hi

	case IX_53H:
		cpu.DE.Hi = cpu.DE.Lo

	case IX_57H:
		cpu.DE.Hi = cpu.AF.Hi

	case IX_58H:
		cpu.DE.Lo = cpu.BC.Hi

	case IX_59H:
		cpu.DE.Lo = cpu.BC.Lo

	case IX_5AH:
		cpu.DE.Lo = cpu.DE.Hi

	case IX_5BH:
		//cpu.DE.Lo = cpu.DE.Lo

	case IX_5FH:
		cpu.DE.Lo = cpu.AF.Hi

	case IX_78H:
		cpu.AF.Hi = cpu.BC.Hi

	case IX_79H:
		cpu.AF.Hi = cpu.BC.Lo

	case IX_7AH:
		cpu.AF.Hi = cpu.DE.Hi

	case IX_7BH:
		cpu.AF.Hi = cpu.DE.Lo

	case IX_7FH:
		//cpu.AF.Hi = cpu.AF.Hi

	case EX_ADDR_SP_IX:
		addr := cpu.SP
		tmp := cpu.IX.Lo
		cpu.IX.Lo = cpu.Memory[addr]
		cpu.Memory[addr] = tmp
		addr++
		tmp = cpu.IX.Hi
		cpu.IX.Hi = cpu.Memory[addr]
		cpu.Memory[addr] = tmp

	case JP_IX:
		cpu.PC = cpu.IX.getValue()

	case LD_SP_IX:
		cpu.SP = cpu.IX.getValue()

	default:
		fmt.Printf("Unknown IX opcode 0x%02x at %04x\n", code, cpu.PC)
	}
}
