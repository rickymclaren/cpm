package z80

import (
	"fmt"
	"math/bits"
	"os"
)

const NOOP = 0x00
const LD_BC_XXXX = 0x01
const LD_ADDR_BC_A = 0x02
const INC_BC = 0x03
const INC_B = 0x04
const DEC_B = 0x05
const LD_B_XX = 0x06
const RLCA = 0x07
const EX_AF_ALT_AF = 0x08
const ADD_HL_BC = 0x09
const LD_A_ADDR_BC = 0x0a
const DEC_BC = 0x0b
const INC_C = 0x0c
const DEC_C = 0x0d
const LD_C_XX = 0x0e
const RRCA = 0x0f
const DJNZ_XX = 0x10
const LD_DE_XXXX = 0x11
const LD_ADDR_DE_A = 0x12
const INC_DE = 0x13
const INC_D = 0x14
const DEC_D = 0x15
const LD_D_XX = 0x16
const RLA = 0x17
const JR = 0x18
const ADD_HL_DE = 0x19
const LD_A_ADDR_DE = 0x1a
const DEC_DE = 0x1b
const INC_E = 0x1c
const DEC_E = 0x1d
const LD_E_XX = 0x1e
const RRA = 0x1f
const JRNZ = 0x20
const LD_HL_XXXX = 0x21
const LD_ADDR_XXXX_HL = 0x22
const INC_HL = 0x23
const INC_H = 0x24
const DEC_H = 0x25
const LD_H_XX = 0x26
const DAA = 0x27
const JRZ = 0x28
const ADD_HL_HL = 0x29
const LD_HL_ADDR = 0x2a
const DEC_HL = 0x2b
const INC_L = 0x2c
const DEC_L = 0x2d
const LD_L_XX = 0x2e
const CPL = 0x2f
const JRNC = 0x30
const LD_SP_XXXX = 0x31
const LD_ADDR_XXXX_A = 0x32
const INC_SP = 0x33
const INC_ADDR_HL = 0x34
const DEC_ADDR_HL = 0x35
const LD_ADDR_HL_XX = 0x36
const SCF = 0x37
const JRC = 0x38
const ADD_HL_SP = 0x39
const LD_A_ADDR = 0x3a
const DEC_SP = 0x3b
const INC_A = 0x3c
const DEC_A = 0x3d
const LD_A_XX = 0x3e
const CCF = 0x3f
const LD_B_B = 0x40
const LD_B_C = 0x41
const LD_B_D = 0x42
const LD_B_E = 0x43
const LD_B_H = 0x44
const LD_B_L = 0x45
const LD_B_ADDR_HL = 0x46
const LD_B_A = 0x47
const LD_C_B = 0x48
const LD_C_C = 0x49
const LD_C_D = 0x4a
const LD_C_E = 0x4b
const LD_C_H = 0x4c
const LD_C_L = 0x4d
const LD_C_ADDR_HL = 0x4e
const LD_C_A = 0x4f
const LD_D_B = 0x50
const LD_D_C = 0x51
const LD_D_D = 0x52
const LD_D_E = 0x53
const LD_D_H = 0x54
const LD_D_L = 0x55
const LD_D_ADDR_HL = 0x56
const LD_D_A = 0x57
const LD_E_B = 0x58
const LD_E_C = 0x59
const LD_E_D = 0x5a
const LD_E_E = 0x5b
const LD_E_H = 0x5c
const LD_E_L = 0x5d
const LD_E_ADDR_HL = 0x5e
const LD_E_A = 0x5f
const LD_H_B = 0x60
const LD_H_C = 0x61
const LD_H_D = 0x62
const LD_H_E = 0x63
const LD_H_H = 0x64
const LD_H_L = 0x65
const LD_H_ADDR_HL = 0x66
const LD_H_A = 0x67
const LD_L_B = 0x68
const LD_L_C = 0x69
const LD_L_D = 0x6a
const LD_L_E = 0x6b
const LD_L_H = 0x6c
const LD_L_L = 0x6d
const LD_L_ADDR_HL = 0x6e
const LD_L_A = 0x6f
const LD_ADDR_HL_B = 0x70
const LD_ADDR_HL_C = 0x71
const LD_ADDR_HL_D = 0x72
const LD_ADDR_HL_E = 0x73
const LD_ADDR_HL_H = 0x74
const LD_ADDR_HL_L = 0x75
const HALT = 0x76
const LD_ADDR_HL_A = 0x77
const LD_A_B = 0x78
const LD_A_C = 0x79
const LD_A_D = 0x7a
const LD_A_E = 0x7b
const LD_A_H = 0x7c
const LD_A_L = 0x7d
const LD_A_ADDR_HL = 0x7e
const LD_A_A = 0x7f
const ADD_B = 0x80
const ADD_C = 0x81
const ADD_D = 0x82
const ADD_E = 0x83
const ADD_H = 0x84
const ADD_L = 0x85
const ADD_ADDR_HL = 0x86
const ADD_A = 0x87
const ADC_B = 0x88
const ADC_C = 0x89
const ADC_D = 0x8a
const ADC_E = 0x8b
const ADC_H = 0x8c
const ADC_L = 0x8d
const ADC_ADDR_HL = 0x8e
const ADC_A = 0x8f
const SUB_B = 0x90
const SUB_C = 0x91
const SUB_D = 0x92
const SUB_E = 0x93
const SUB_H = 0x94
const SUB_L = 0x95
const SUB_ADDR_HL = 0x96
const SUB_A = 0x97
const SBC_B = 0x98
const SBC_C = 0x99
const SBC_D = 0x9a
const SBC_E = 0x9b
const SBC_H = 0x9c
const SBC_L = 0x9d
const SBC_ADDR_HL = 0x9e
const SBC_A = 0x9f
const AND_B = 0xa0
const AND_C = 0xa1
const AND_D = 0xa2
const AND_E = 0xa3
const AND_H = 0xa4
const AND_L = 0xa5
const AND_ADDR_HL = 0xa6
const AND_A = 0xa7
const XOR_B = 0xa8
const XOR_C = 0xa9
const XOR_D = 0xaa
const XOR_E = 0xab
const XOR_H = 0xac
const XOR_L = 0xad
const XOR_ADDR_HL = 0xae
const XOR_A = 0xaf
const OR_B = 0xb0
const OR_C = 0xb1
const OR_D = 0xb2
const OR_E = 0xb3
const OR_H = 0xb4
const OR_L = 0xb5
const OR_ADDR_HL = 0xb6
const OR_A = 0xb7
const CP_B = 0xb8
const CP_C = 0xb9
const CP_D = 0xba
const CP_E = 0xbb
const CP_H = 0xbc
const CP_L = 0xbd
const CP_ADDR_HL = 0xbe
const CP_A = 0xbf
const RETNZ = 0xc0
const POP_BC = 0xc1
const JPNZ_XXXX = 0xc2
const JP_XXXX = 0xc3
const CALLNZ_XXXX = 0xc4
const PUSH_BC = 0xc5
const ADD_XX = 0xc6
const RST_00H = 0xc7
const RETZ = 0xc8
const RET = 0xc9
const JPZ_XXXX = 0xca
const BITS = 0xcb
const CALLZ_XXXX = 0xcc
const CALL_XXXX = 0xcd
const ADC_XX = 0xce
const RST_08H = 0xcf
const RETNC = 0xd0
const POP_DE = 0xd1
const JPNC_XXXX = 0xd2
const OUT_A_XX = 0xd3
const CALLNC_XXXX = 0xd4
const PUSH_DE = 0xd5
const SUB_XX = 0xd6
const RST_10H = 0xd7
const RETC = 0xd8
const EXX = 0xd9
const JPC_XXXX = 0xda
const IN_A_XX = 0xdb
const CALLC_XXXX = 0xdc
const IX = 0xdd
const SBC_A_XX = 0xde
const RST_18H = 0xdf
const RETPO = 0xe0
const POP_HL = 0xe1
const JPPO_XXXX = 0xe2
const EX_ADDR_SP_HL = 0xe3
const CALLPO_XXXX = 0xe4
const PUSH_HL = 0xe5
const AND_XX = 0xe6
const RST_20H = 0xe7
const RETPE = 0xe8
const JP_ADDR_HL = 0xe9
const JPPE_XXXX = 0xea
const EX_DE_HL = 0xeb
const CALLPE_XXXX = 0xec
const EXTD = 0xed
const XOR_XX = 0xee
const RST_28H = 0xef
const RETP = 0xf0
const POP_AF = 0xf1
const JPP_XXXX = 0xf2
const DI = 0xf3
const CALLP_XXXX = 0xf4
const PUSH_AF = 0xf5
const OR_XX = 0xf6
const RST_30H = 0xf7
const RETM = 0xf8
const LD_SP_HL = 0xf9
const JPM_XXXX = 0xfa
const EI = 0xfb
const CALLM_XXXX = 0xfc
const IY = 0xfd
const CP_XX = 0xfe
const RST_38H = 0xff

const FLG_S = uint8(0x80)    // Sign
const FLG_Z = uint8(0x40)    // Zero
const FLG_5 = uint8(0x20)    // Unused
const FLG_H = uint8(0x10)    // Half Carry for BCD
const FLG_3 = uint8(0x08)    // Unused
const FLG_P_V = uint8(0x04)  // Parity or Overflow
const FLG_N = uint8(0x02)    // Add/Subtract for BCD
const FLG_C = uint8(0x01)    // Carry
const FLG_NONE = uint8(0x00) // No flags set
const FLG_STD = FLG_S + FLG_3 + FLG_5

func (cpu *CPU) opcode(code uint8) {
	switch code {
	case HALT:
		cpu.Halt = true

	case NOOP:
		// do nothing

	// Load instructions
	case LD_A_A:
		// do nothing

	case LD_A_B:
		cpu.AF.Hi = cpu.BC.Hi

	case LD_A_C:
		cpu.AF.Hi = cpu.BC.Lo

	case LD_A_D:
		cpu.AF.Hi = cpu.DE.Hi

	case LD_A_E:
		cpu.AF.Hi = cpu.DE.Lo

	case LD_A_H:
		cpu.AF.Hi = cpu.HL.Hi

	case LD_A_L:
		cpu.AF.Hi = cpu.HL.Lo

	case LD_A_ADDR_BC:
		cpu.AF.Hi = cpu.Memory[cpu.BC.getValue()]

	case LD_A_ADDR_DE:
		cpu.AF.Hi = cpu.Memory[cpu.DE.getValue()]

	case LD_A_ADDR:
		value := cpu.fetch16()
		cpu.AF.Hi = cpu.Memory[value]

	case LD_B_A:
		cpu.BC.Hi = cpu.AF.Hi

	case LD_B_B:
		// do nothing

	case LD_B_C:
		cpu.BC.Hi = cpu.BC.Lo

	case LD_B_D:
		cpu.BC.Hi = cpu.DE.Hi

	case LD_B_E:
		cpu.BC.Hi = cpu.DE.Lo

	case LD_B_H:
		cpu.BC.Hi = cpu.HL.Hi

	case LD_B_L:
		cpu.BC.Hi = cpu.HL.Lo

	case LD_C_A:
		cpu.BC.Lo = cpu.AF.Hi

	case LD_C_B:
		cpu.BC.Lo = cpu.BC.Hi

	case LD_C_C:
		// do nothing

	case LD_C_D:
		cpu.BC.Lo = cpu.DE.Hi

	case LD_C_E:
		cpu.BC.Lo = cpu.DE.Lo

	case LD_C_H:
		cpu.BC.Lo = cpu.HL.Hi

	case LD_C_L:
		cpu.BC.Lo = cpu.HL.Lo

	case LD_D_A:
		cpu.DE.Hi = cpu.AF.Hi

	case LD_D_B:
		cpu.DE.Hi = cpu.BC.Hi

	case LD_D_C:
		cpu.DE.Hi = cpu.BC.Lo

	case LD_D_D:
		// do nothing

	case LD_D_E:
		cpu.DE.Hi = cpu.DE.Lo

	case LD_D_H:
		cpu.DE.Hi = cpu.HL.Hi

	case LD_D_L:
		cpu.DE.Hi = cpu.HL.Lo

	case LD_E_A:
		cpu.DE.Lo = cpu.AF.Hi

	case LD_E_B:
		cpu.DE.Lo = cpu.BC.Hi

	case LD_E_C:
		cpu.DE.Lo = cpu.BC.Lo

	case LD_E_D:
		cpu.DE.Lo = cpu.DE.Hi

	case LD_E_E:
		// do nothing

	case LD_E_H:
		cpu.DE.Lo = cpu.HL.Hi

	case LD_E_L:
		cpu.DE.Lo = cpu.HL.Lo

	case LD_H_A:
		cpu.HL.Hi = cpu.AF.Hi

	case LD_H_B:
		cpu.HL.Hi = cpu.BC.Hi

	case LD_H_C:
		cpu.HL.Hi = cpu.BC.Lo

	case LD_H_D:
		cpu.HL.Hi = cpu.DE.Hi

	case LD_H_E:
		cpu.HL.Hi = cpu.DE.Lo

	case LD_H_H:
		// do nothing

	case LD_H_L:
		cpu.HL.Hi = cpu.HL.Lo

	case LD_L_A:
		cpu.HL.Lo = cpu.AF.Hi

	case LD_L_B:
		cpu.HL.Lo = cpu.BC.Hi

	case LD_L_C:
		cpu.HL.Lo = cpu.BC.Lo

	case LD_L_D:
		cpu.HL.Lo = cpu.DE.Hi

	case LD_L_E:
		cpu.HL.Lo = cpu.DE.Lo

	case LD_L_H:
		cpu.HL.Lo = cpu.HL.Hi

	case LD_L_L:
		// do nothing

	case LD_A_XX:
		cpu.AF.Hi = cpu.fetch()

	case LD_B_XX:
		cpu.BC.Hi = cpu.fetch()

	case LD_C_XX:
		cpu.BC.Lo = cpu.fetch()

	case LD_D_XX:
		cpu.DE.Hi = cpu.fetch()

	case LD_E_XX:
		cpu.DE.Lo = cpu.fetch()

	case LD_H_XX:
		cpu.HL.Hi = cpu.fetch()

	case LD_L_XX:
		cpu.HL.Lo = cpu.fetch()

	case LD_A_ADDR_HL:
		cpu.AF.Hi = cpu.Memory[cpu.HL.getValue()]

	case LD_B_ADDR_HL:
		cpu.BC.Hi = cpu.Memory[cpu.HL.getValue()]

	case LD_C_ADDR_HL:
		cpu.BC.Lo = cpu.Memory[cpu.HL.getValue()]

	case LD_D_ADDR_HL:
		cpu.DE.Hi = cpu.Memory[cpu.HL.getValue()]

	case LD_E_ADDR_HL:
		cpu.DE.Lo = cpu.Memory[cpu.HL.getValue()]

	case LD_H_ADDR_HL:
		cpu.HL.Hi = cpu.Memory[cpu.HL.getValue()]

	case LD_L_ADDR_HL:
		cpu.HL.Lo = cpu.Memory[cpu.HL.getValue()]

	case LD_BC_XXXX:
		cpu.BC = cpu.fetchRegister()

	case LD_DE_XXXX:
		cpu.DE = cpu.fetchRegister()

	case LD_HL_XXXX:
		cpu.HL = cpu.fetchRegister()

	case LD_HL_ADDR:
		address := cpu.fetch16()
		cpu.HL.Lo = cpu.Memory[address]
		cpu.HL.Hi = cpu.Memory[address+1]

	case LD_SP_HL:
		cpu.SP = cpu.HL.getValue()

	case LD_ADDR_XXXX_HL:
		address := cpu.fetch16()
		cpu.Memory[address] = cpu.HL.Lo
		cpu.Memory[address+1] = cpu.HL.Hi

	case LD_ADDR_DE_A:
		address := cpu.DE.getValue()
		cpu.Memory[address] = cpu.AF.Hi

	case LD_ADDR_XXXX_A:
		address := cpu.fetch16()
		cpu.Memory[address] = cpu.AF.Hi

	case LD_SP_XXXX:
		cpu.SP = cpu.fetch16()

	case LD_ADDR_BC_A:
		cpu.Memory[cpu.BC.getValue()] = cpu.AF.Hi

	case LD_ADDR_HL_A:
		cpu.Memory[cpu.HL.getValue()] = cpu.AF.Hi

	case LD_ADDR_HL_B:
		cpu.Memory[cpu.HL.getValue()] = cpu.BC.Hi

	case LD_ADDR_HL_C:
		cpu.Memory[cpu.HL.getValue()] = cpu.BC.Lo

	case LD_ADDR_HL_D:
		cpu.Memory[cpu.HL.getValue()] = cpu.DE.Hi

	case LD_ADDR_HL_E:
		cpu.Memory[cpu.HL.getValue()] = cpu.DE.Lo

	case LD_ADDR_HL_H:
		cpu.Memory[cpu.HL.getValue()] = cpu.HL.Hi

	case LD_ADDR_HL_L:
		cpu.Memory[cpu.HL.getValue()] = cpu.HL.Lo

	case LD_ADDR_HL_XX:
		cpu.Memory[cpu.HL.getValue()] = cpu.fetch()

	// Exchange instructions
	case EX_DE_HL:
		tmp := cpu.DE
		cpu.DE = cpu.HL
		cpu.HL = tmp

	case EX_ADDR_SP_HL:
		lo := cpu.Memory[cpu.SP]
		hi := cpu.Memory[cpu.SP+1]
		cpu.Memory[cpu.SP] = cpu.HL.Lo
		cpu.Memory[cpu.SP+1] = cpu.HL.Hi
		cpu.HL.Lo = lo
		cpu.HL.Hi = hi

	// INC/DEC instructions
	case INC_A:
		cpu.AF.Hi = cpu.inc(cpu.AF.Hi)

	case DEC_A:
		cpu.AF.Hi = cpu.dec(cpu.AF.Hi)

	case INC_B:
		cpu.BC.Hi = cpu.inc(cpu.BC.Hi)

	case DEC_B:
		cpu.BC.Hi = cpu.dec(cpu.BC.Hi)

	case INC_C:
		cpu.BC.Lo = cpu.inc(cpu.BC.Lo)

	case DEC_C:
		cpu.BC.Lo = cpu.dec(cpu.BC.Lo)

	case INC_D:
		cpu.DE.Hi = cpu.inc(cpu.DE.Hi)

	case DEC_D:
		cpu.DE.Hi = cpu.dec(cpu.DE.Hi)

	case INC_E:
		cpu.DE.Lo = cpu.inc(cpu.DE.Lo)

	case DEC_E:
		cpu.DE.Lo = cpu.dec(cpu.DE.Lo)

	case INC_H:
		cpu.HL.Hi = cpu.inc(cpu.HL.Hi)

	case DEC_H:
		cpu.HL.Hi = cpu.dec(cpu.HL.Hi)

	case INC_L:
		cpu.HL.Lo = cpu.inc(cpu.HL.Lo)

	case DEC_L:
		cpu.HL.Lo = cpu.dec(cpu.HL.Lo)

	case INC_ADDR_HL:
		address := cpu.HL.getValue()
		cpu.Memory[address] = cpu.inc(cpu.Memory[address])

	case DEC_ADDR_HL:
		address := cpu.HL.getValue()
		cpu.Memory[address] = cpu.dec(cpu.Memory[address])

	case INC_BC:
		v := cpu.BC.getValue()
		v++
		cpu.BC.setValue(v)

	case DEC_BC:
		v := cpu.BC.getValue()
		v--
		cpu.BC.setValue(v)

	case INC_DE:
		v := cpu.DE.getValue()
		v++
		cpu.DE.setValue(v)

	case DEC_DE:
		v := cpu.DE.getValue()
		v--
		cpu.DE.setValue(v)

	case INC_HL:
		v := cpu.HL.getValue()
		v++
		cpu.HL.setValue(v)

	case DEC_HL:
		v := cpu.HL.getValue()
		v--
		cpu.HL.setValue(v)

	case INC_SP:
		v := cpu.SP
		v++
		cpu.SP = v

	case DEC_SP:
		v := cpu.SP
		v--
		cpu.SP = v

	case DJNZ_XX:
		rel := cpu.fetch()
		cpu.BC.Hi = cpu.BC.Hi - 1
		if cpu.BC.Hi != uint8(0x00) {
			cpu.PC = addrOff(cpu.PC, rel)
		}

	// ADD instructions
	case ADD_XX:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.fetch(), false)

	case ADD_A:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.AF.Hi, false)

	case ADD_B:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.BC.Hi, false)

	case ADD_C:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.BC.Lo, false)

	case ADD_D:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.DE.Hi, false)

	case ADD_E:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.DE.Lo, false)

	case ADD_H:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.HL.Hi, false)

	case ADD_L:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.HL.Lo, false)

	case ADD_HL_BC:
		cpu.HL.setValue(cpu.add16(cpu.HL.getValue(), cpu.BC.getValue()))

	case ADD_HL_DE:
		cpu.HL.setValue(cpu.add16(cpu.HL.getValue(), cpu.DE.getValue()))

	case ADD_HL_HL:
		cpu.HL.setValue(cpu.add16(cpu.HL.getValue(), cpu.HL.getValue()))

	case ADD_HL_SP:
		cpu.HL.setValue(cpu.add16(cpu.HL.getValue(), cpu.SP))

	case ADD_ADDR_HL:
		value := cpu.Memory[cpu.HL.getValue()]
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, value, false)

	case ADC_XX:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.fetch(), cpu.flagC())

	case ADC_A:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.AF.Hi, cpu.flagC())

	case ADC_B:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.BC.Hi, cpu.flagC())

	case ADC_C:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.BC.Lo, cpu.flagC())

	case ADC_D:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.DE.Hi, cpu.flagC())

	case ADC_E:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.DE.Lo, cpu.flagC())

	case ADC_H:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.HL.Hi, cpu.flagC())

	case ADC_L:
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, cpu.HL.Lo, cpu.flagC())

	case ADC_ADDR_HL:
		value := cpu.Memory[cpu.HL.getValue()]
		cpu.AF.Hi = cpu.add(cpu.AF.Hi, value, cpu.flagC())

	case SUB_XX:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.fetch(), false)

	case SUB_A:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.AF.Hi, false)

	case SUB_B:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.BC.Hi, false)

	case SUB_C:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.BC.Lo, false)

	case SUB_D:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.DE.Hi, false)

	case SUB_E:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.DE.Lo, false)

	case SUB_H:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.HL.Hi, false)

	case SUB_L:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.HL.Lo, false)

	case SUB_ADDR_HL:
		value := cpu.Memory[cpu.HL.getValue()]
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, value, false)

	case SBC_A_XX:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.fetch(), cpu.flagC())

	case SBC_A:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.AF.Hi, cpu.flagC())

	case SBC_B:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.BC.Hi, cpu.flagC())

	case SBC_C:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.BC.Lo, cpu.flagC())

	case SBC_D:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.DE.Hi, cpu.flagC())

	case SBC_E:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.DE.Lo, cpu.flagC())

	case SBC_H:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.HL.Hi, cpu.flagC())

	case SBC_L:
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, cpu.HL.Lo, cpu.flagC())

	case SBC_ADDR_HL:
		value := cpu.Memory[cpu.HL.getValue()]
		cpu.AF.Hi = cpu.sub(cpu.AF.Hi, value, cpu.flagC())

	// Comparison instructions
	case CP_XX:
		cpu.sub(cpu.AF.Hi, cpu.fetch(), false)

	case CP_A:
		cpu.sub(cpu.AF.Hi, cpu.AF.Hi, false)

	case CP_B:
		cpu.sub(cpu.AF.Hi, cpu.BC.Hi, false)

	case CP_C:
		cpu.sub(cpu.AF.Hi, cpu.BC.Lo, false)

	case CP_D:
		cpu.sub(cpu.AF.Hi, cpu.DE.Hi, false)

	case CP_E:
		cpu.sub(cpu.AF.Hi, cpu.DE.Lo, false)

	case CP_H:
		cpu.sub(cpu.AF.Hi, cpu.HL.Hi, false)

	case CP_L:
		cpu.sub(cpu.AF.Hi, cpu.HL.Lo, false)

	case CP_ADDR_HL:
		cpu.sub(cpu.AF.Hi, cpu.Memory[cpu.HL.getValue()], false)

	// Bitwise instructions
	case AND_XX:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.fetch())

	case AND_A:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.AF.Hi)

	case AND_B:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.BC.Hi)

	case AND_C:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.BC.Lo)

	case AND_D:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.DE.Hi)

	case AND_E:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.DE.Lo)

	case AND_H:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.HL.Hi)

	case AND_L:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.HL.Lo)

	case AND_ADDR_HL:
		cpu.AF.Hi = cpu.and(cpu.AF.Hi, cpu.Memory[cpu.HL.getValue()])

	case OR_XX:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.fetch())

	case OR_A:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.AF.Hi)

	case OR_B:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.BC.Hi)

	case OR_C:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.BC.Lo)

	case OR_D:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.DE.Hi)

	case OR_E:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.DE.Lo)

	case OR_H:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.HL.Hi)

	case OR_L:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.HL.Lo)

	case OR_ADDR_HL:
		cpu.AF.Hi = cpu.or(cpu.AF.Hi, cpu.Memory[cpu.HL.getValue()])

	case XOR_XX:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.fetch())

	case XOR_A:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.AF.Hi)

	case XOR_B:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.BC.Hi)

	case XOR_C:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.BC.Lo)

	case XOR_D:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.DE.Hi)

	case XOR_E:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.DE.Lo)

	case XOR_H:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.HL.Hi)

	case XOR_L:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.HL.Lo)

	case XOR_ADDR_HL:
		cpu.AF.Hi = cpu.xor(cpu.AF.Hi, cpu.Memory[cpu.HL.getValue()])

	case CPL:
		cpu.AF.Hi = ^cpu.AF.Hi
		cpu.AF.Lo |= FLG_H + FLG_N

	// Flag instructions
	case CCF:
		var nand uint8 = FLG_H | FLG_N | FLG_C
		var or uint8
		if cpu.flagC() {
			or |= FLG_H
		} else {
			or |= FLG_C
		}
		cpu.AF.Lo = cpu.AF.Lo&^nand | or

	case SCF:
		var nand uint8 = FLG_H | FLG_N | FLG_C
		var or uint8 = FLG_C
		cpu.AF.Lo = cpu.AF.Lo&^nand | or

	// BCD instructions
	case DAA:
		r := cpu.AF.Hi
		c := cpu.flagC()
		if cpu.flagN() {
			if cpu.flagH() || (cpu.AF.Hi&0x0f) > 9 {
				r -= 0x06
			}
			if c || (cpu.AF.Hi > 0x99) {
				r -= 0x60
			}
		} else {
			if cpu.flagH() || (cpu.AF.Hi&0x0f) > 9 {
				r += 0x06
			}
			if c || (cpu.AF.Hi > 0x99) {
				r += 0x60
			}
		}

		var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V
		var or uint8
		or |= uint8(r) & FLG_STD
		if uint8(r) == 0 {
			or |= FLG_Z
		}
		or |= (cpu.AF.Hi ^ r) & FLG_H
		or |= (uint8(bits.OnesCount8(r)%2) - 1) & FLG_P_V
		if cpu.AF.Hi > 0x99 {
			or |= FLG_C
		}
		cpu.AF.Lo = cpu.AF.Lo&^nand | or

		cpu.AF.Hi = r

	// JUMP Instructions
	case JP_XXXX:
		cpu.PC = cpu.fetch16()

	case JPZ_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_Z != 0 {
			cpu.PC = target
		}

	case JPNZ_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_Z == 0 {
			cpu.PC = target
		}

	case JPC_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_C != 0 {
			cpu.PC = target
		}

	case JPNC_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_C == 0 {
			cpu.PC = target
		}

	case JPM_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_S != 0 {
			cpu.PC = target
		}

	case JPP_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_S == 0 {
			cpu.PC = target
		}

	case JPPE_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_P_V != 0 {
			cpu.PC = target
		}

	case JPPO_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_P_V == 0 {
			cpu.PC = target
		}

	case JP_ADDR_HL:
		cpu.PC = cpu.HL.getValue()

	case JR:
		r := cpu.fetch()
		cpu.PC = addrOff(cpu.PC, r)

	case JRZ:
		r := cpu.fetch()
		if cpu.AF.Lo&FLG_Z != 0 {
			cpu.PC = addrOff(cpu.PC, r)
		}

	case JRNZ:
		r := cpu.fetch()
		if cpu.AF.Lo&FLG_Z == 0 {
			cpu.PC = addrOff(cpu.PC, r)
		}

	case JRC:
		r := cpu.fetch()
		if cpu.AF.Lo&FLG_C != 0 {
			cpu.PC = addrOff(cpu.PC, r)
		}

	case JRNC:
		r := cpu.fetch()
		if cpu.AF.Lo&FLG_C == 0 {
			cpu.PC = addrOff(cpu.PC, r)
		}

	// Call instructions
	case CALL_XXXX:
		target := cpu.fetch16()
		if target == 0x0005 && cpu.Bdos != nil {
			cpu.Bdos.TraceBDOS()
		}
		cpu.push(cpu.PC)
		cpu.PC = target

	case CALLC_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_C != 0 {
			cpu.push(cpu.PC)
			cpu.PC = target
		}

	case CALLNC_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_C == 0 {
			cpu.push(cpu.PC)
			cpu.PC = target
		}

	case CALLZ_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_Z != 0 {
			cpu.push(cpu.PC)
			cpu.PC = target
		}

	case CALLNZ_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_Z == 0 {
			cpu.push(cpu.PC)
			cpu.PC = target
		}

	case CALLPE_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_P_V != 0 {
			cpu.push(cpu.PC)
			cpu.PC = target
		}

	case CALLPO_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_P_V == 0 {
			cpu.push(cpu.PC)
			cpu.PC = target
		}

	case CALLM_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_S != 0 {
			cpu.push(cpu.PC)
			cpu.PC = target
		}

	case CALLP_XXXX:
		target := cpu.fetch16()
		if cpu.AF.Lo&FLG_S == 0 {
			cpu.push(cpu.PC)
			cpu.PC = target
		}

	// Return instructions
	case RET:
		cpu.PC = cpu.pop()

	case RETZ:
		if cpu.AF.Lo&FLG_Z != 0 {
			cpu.PC = cpu.pop()
		}

	case RETNZ:
		if cpu.AF.Lo&FLG_Z == 0 {
			cpu.PC = cpu.pop()
		}

	case RETC:
		if cpu.AF.Lo&FLG_C != 0 {
			cpu.PC = cpu.pop()
		}

	case RETNC:
		if cpu.AF.Lo&FLG_C == 0 {
			cpu.PC = cpu.pop()
		}

	case RETM:
		if cpu.AF.Lo&FLG_S != 0 {
			cpu.PC = cpu.pop()
		}

	case RETP:
		if cpu.AF.Lo&FLG_S == 0 {
			cpu.PC = cpu.pop()
		}

	case RETPE:
		if cpu.AF.Lo&FLG_P_V != 0 {
			cpu.PC = cpu.pop()
		}

	case RETPO:
		if cpu.AF.Lo&FLG_P_V == 0 {
			cpu.PC = cpu.pop()
		}

	// Push/Pop
	case PUSH_AF:
		cpu.push(cpu.AF.getValue())

	case PUSH_BC:
		cpu.push(cpu.BC.getValue())

	case PUSH_DE:
		cpu.push(cpu.DE.getValue())

	case PUSH_HL:
		cpu.push(cpu.HL.getValue())

	case POP_AF:
		cpu.AF.setValue(cpu.pop())

	case POP_BC:
		cpu.BC.setValue(cpu.pop())

	case POP_DE:
		cpu.DE.setValue(cpu.pop())

	case POP_HL:
		cpu.HL.setValue(cpu.pop())

	// Shift and Rotate
	case RLA: // Bit 7 -> C, C -> Bit 0
		a := cpu.AF.Hi
		a2 := a<<1 | cpu.AF.Lo&FLG_C
		cpu.AF.Hi = a2
		cpu.updateFlagRL(a)

	case RLCA: // Bit 7 -> C and Bit 0
		a := cpu.AF.Hi
		a2 := a<<1 | a>>7
		cpu.AF.Hi = a2
		cpu.updateFlagRL(a)

	case RRA: // Bit 0 -> C, C -> Bit 7
		a := cpu.AF.Hi
		a2 := a >> 1
		if cpu.flagC() {
			a2 |= 0x80
		}
		cpu.AF.Hi = a2
		cpu.updateFlagRR(a)

	case RRCA: // Bit 0 -> C and Bit 7
		a := cpu.AF.Hi
		a2 := a>>1 | a<<7
		cpu.AF.Hi = a2
		cpu.updateFlagRR(a)

	// IO Instructions
	case IN_A_XX:
		port := cpu.fetch()
		cpu.AF.Hi = cpu.Io.In(port)

	case OUT_A_XX:
		port := cpu.fetch()
		cpu.Io.Out(port, cpu.AF.Hi)

	// Interrupts
	case DI:
		// do nothing

	case EI:
		// do nothing

	// Reset instructions
	case RST_00H:
		cpu.push(cpu.PC)
		cpu.PC = 0x0000

	case RST_08H:
		cpu.push(cpu.PC)
		cpu.PC = 0x0008

	case RST_10H:
		cpu.push(cpu.PC)
		cpu.PC = 0x0010

	case RST_18H:
		cpu.push(cpu.PC)
		cpu.PC = 0x0018

	case RST_20H:
		cpu.push(cpu.PC)
		cpu.PC = 0x0020

	case RST_28H:
		cpu.push(cpu.PC)
		cpu.PC = 0x0028

	case RST_30H:
		cpu.push(cpu.PC)
		cpu.PC = 0x0030

	case RST_38H:
		cpu.push(cpu.PC)
		cpu.PC = 0x0038

	case BITS:
		cpu.bits(cpu.fetch())

	case EXTD:
		cpu.extended(cpu.fetch())

	case IX:
		cpu.ix(cpu.fetch())

	case IY:
		cpu.iy(cpu.fetch())

	// exchange instructions
	case EX_AF_ALT_AF:
		tmp := cpu.AF.getValue()
		cpu.AF = cpu.AF_
		cpu.AF_.setValue(tmp)

	case EXX:
		tmp := cpu.BC.getValue()
		cpu.BC = cpu.BC_
		cpu.BC_.setValue(tmp)
		tmp = cpu.DE.getValue()
		cpu.DE = cpu.DE_
		cpu.DE_.setValue(tmp)
		tmp = cpu.HL.getValue()
		cpu.HL = cpu.HL_
		cpu.HL_.setValue(tmp)

	default:
		fmt.Printf("Unknown opcode 0x%02x at %04x\n", code, cpu.PC)
		os.Exit(-1)

	}
}

func (cpu *CPU) flags() string {
	s := ""
	if cpu.AF.Lo&FLG_S != 0 {
		s += "S"
	} else {
		s += "_"
	}
	if cpu.AF.Lo&FLG_Z != 0 {
		s += "Z"
	} else {
		s += "_"
	}
	if cpu.AF.Lo&FLG_H != 0 {
		s += "H"
	} else {
		s += "_"
	}
	if cpu.AF.Lo&FLG_P_V != 0 {
		s += "V"
	} else {
		s += "_"
	}
	if cpu.AF.Lo&FLG_N != 0 {
		s += "N"
	} else {
		s += "_"
	}
	if cpu.AF.Lo&FLG_C != 0 {
		s += "C"
	} else {
		s += "_"
	}
	return s

}

func (cpu *CPU) push(value uint16) {
	cpu.SP--
	cpu.Memory[cpu.SP] = uint8(value >> 8)
	cpu.SP--
	cpu.Memory[cpu.SP] = uint8(value & 0xff)
}

func (cpu *CPU) pop() uint16 {
	value := uint16(cpu.Memory[cpu.SP])
	cpu.SP++
	value = value + (uint16(cpu.Memory[cpu.SP]) << 8)
	cpu.SP++
	return value
}

func addrOff(addr uint16, off uint8) uint16 {
	return addr + uint16(int16(int8(off)))
}
