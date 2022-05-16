package z80

import (
	"fmt"
)

// Extended opcodes - ED
const SBC_HL_BC = 0x42
const LD_ADDR_XXXX_BC = 0x43
const NEG = 0x44
const ADC_HL_BC = 0x4a
const LD_BC_ADDR_XXXX = 0x4b
const SBC_HL_DE = 0x52
const LD_ADDR_XXXX_DE = 0x53
const ADC_HL_DE = 0x5a
const LD_DE_ADDR_XXXX = 0x5b
const SBC_HL_HL = 0x62
const RRD = 0x67
const ADC_HL_HL = 0x6a
const LD_HL_ADDR_XXXX = 0x6b
const RLD = 0x6f
const SBC_HL_SP = 0x72
const ADC_HL_SP = 0x7a
const LD_ADDR_XXXX_SP = 0x73
const LD_SP_ADDR_XXXX = 0x7b
const LDI = 0xa0
const CPI = 0xa1
const LDD = 0xa8
const CPD = 0xa9
const LDIR = 0xb0
const CPIR = 0xb1
const LDDR = 0xb8
const CPDR = 0xb9

func (cpu *CPU) extended(code uint8) {
	switch code {

	case NEG:
		cpu.AF.Hi = cpu.sub(0, cpu.AF.Hi, false)

	case RRD:
		address := cpu.HL.getValue()
		a := cpu.AF.Hi & 0x0f
		h := (cpu.Memory[address] >> 4) & 0x0f
		l := cpu.Memory[address] & 0x0f
		cpu.Memory[address] = (a << 4) + h
		cpu.AF.Hi = (cpu.AF.Hi & 0xf0) + l
		cpu.updateFlagRxD(cpu.AF.Hi)

	case RLD:
		address := cpu.HL.getValue()
		a := cpu.AF.Hi & 0x0f
		h := (cpu.Memory[address] >> 4) & 0x0f
		l := cpu.Memory[address] & 0x0f
		cpu.Memory[address] = (l << 4) + a
		cpu.AF.Hi = (cpu.AF.Hi & 0xf0) + h
		cpu.updateFlagRxD(cpu.AF.Hi)

	case LD_ADDR_XXXX_BC:
		address := cpu.fetch16()
		cpu.Memory[address] = cpu.BC.Lo
		cpu.Memory[address+1] = cpu.BC.Hi

	case LD_ADDR_XXXX_DE:
		address := cpu.fetch16()
		cpu.Memory[address] = cpu.DE.Lo
		cpu.Memory[address+1] = cpu.DE.Hi

	case ADC_HL_BC:
		cpu.HL.setValue(uint16(cpu.adc16(cpu.HL.getValue(), cpu.BC.getValue())))

	case ADC_HL_DE:
		cpu.HL.setValue(uint16(cpu.adc16(cpu.HL.getValue(), cpu.DE.getValue())))

	case ADC_HL_HL:
		cpu.HL.setValue(uint16(cpu.adc16(cpu.HL.getValue(), cpu.HL.getValue())))

	case ADC_HL_SP:
		cpu.HL.setValue(uint16(cpu.adc16(cpu.HL.getValue(), cpu.SP)))

	case SBC_HL_BC:
		cpu.HL.setValue(uint16(cpu.sub16(cpu.HL.getValue(), cpu.BC.getValue())))

	case SBC_HL_DE:
		cpu.HL.setValue(uint16(cpu.sub16(cpu.HL.getValue(), cpu.DE.getValue())))

	case SBC_HL_HL:
		cpu.HL.setValue(uint16(cpu.sub16(cpu.HL.getValue(), cpu.HL.getValue())))

	case SBC_HL_SP:
		cpu.HL.setValue(uint16(cpu.sub16(cpu.HL.getValue(), cpu.SP)))

	case LDI, LDIR:
		de := cpu.DE.getValue()
		hl := cpu.HL.getValue()
		bc := cpu.BC.getValue()
		a := cpu.Memory[hl]
		cpu.Memory[de] = a
		cpu.DE.setValue(de + 1)
		cpu.HL.setValue(hl + 1)
		cpu.BC.setValue(bc - 1)
		cpu.updateFlagLDID(a)
		if code == LDIR && cpu.AF.Lo&FLG_P_V != 0 { // cpu.BC != 0
			cpu.PC -= 2
		}

	case LDD, LDDR:
		de := cpu.DE.getValue()
		hl := cpu.HL.getValue()
		bc := cpu.BC.getValue()
		a := cpu.Memory[hl]
		cpu.Memory[de] = a
		cpu.DE.setValue(de - 1)
		cpu.HL.setValue(hl - 1)
		cpu.BC.setValue(bc - 1)
		cpu.updateFlagLDID(a)
		if code == LDDR && cpu.AF.Lo&FLG_P_V != 0 { // cpu.BC != 0
			cpu.PC -= 2
		}

	case LD_ADDR_XXXX_SP:
		addr := cpu.fetch16()
		cpu.Memory[addr] = uint8(cpu.SP)
		cpu.Memory[addr+1] = uint8(cpu.SP >> 8)

	case LD_BC_ADDR_XXXX:
		addr := cpu.fetch16()
		v := uint16(cpu.Memory[addr])
		v += uint16(cpu.Memory[addr+1]) << 8
		cpu.BC.setValue(v)

	case LD_DE_ADDR_XXXX:
		addr := cpu.fetch16()
		v := uint16(cpu.Memory[addr])
		v += uint16(cpu.Memory[addr+1]) << 8
		cpu.DE.setValue(v)

	case LD_HL_ADDR_XXXX:
		addr := cpu.fetch16()
		v := uint16(cpu.Memory[addr])
		v += uint16(cpu.Memory[addr+1]) << 8
		cpu.HL.setValue(v)

	case LD_SP_ADDR_XXXX:
		addr := cpu.fetch16()
		cpu.SP = uint16(cpu.Memory[addr])
		cpu.SP += uint16(cpu.Memory[addr+1]) << 8

	case CPI:
		a := cpu.AF.Hi
		b := cpu.Memory[cpu.HL.getValue()]
		r := a - b
		cpu.HL.setValue(cpu.HL.getValue() + 1)
		cpu.BC.setValue(cpu.BC.getValue() - 1)
		cpu.updateFlagCPx(r, a, b)

	case CPIR:
		a := cpu.AF.Hi
		b := cpu.Memory[cpu.HL.getValue()]
		r := a - b
		cpu.HL.setValue(cpu.HL.getValue() + 1)
		cpu.BC.setValue(cpu.BC.getValue() - 1)
		cpu.updateFlagCPx(r, a, b)
		if cpu.AF.Lo&FLG_P_V != 0 && cpu.AF.Lo&FLG_Z == 0 {
			cpu.PC -= 2
		}

	case CPD:
		a := cpu.AF.Hi
		b := cpu.Memory[cpu.HL.getValue()]
		r := a - b
		cpu.HL.setValue(cpu.HL.getValue() - 1)
		cpu.BC.setValue(cpu.BC.getValue() - 1)
		cpu.updateFlagCPx(r, a, b)

	case CPDR:
		a := cpu.AF.Hi
		b := cpu.Memory[cpu.HL.getValue()]
		r := a - b
		cpu.HL.setValue(cpu.HL.getValue() - 1)
		cpu.BC.setValue(cpu.BC.getValue() - 1)
		cpu.updateFlagCPx(r, a, b)
		if cpu.AF.Lo&FLG_P_V != 0 && cpu.AF.Lo&FLG_Z == 0 {
			cpu.PC -= 2
		}

	default:
		fmt.Printf("Unknown extended opcode 0x%02x at %04x\n", code, cpu.PC)
	}
}
