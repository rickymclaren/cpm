package main

import (
	"cpm/internal/z80"
	"testing"

	"github.com/stretchr/testify/assert"
)

const FLG_UNDOC = z80.FLG_3 | z80.FLG_5

func TestADD(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0x01, z80.FLG_NONE},
		{0x7f, z80.FLG_NONE, 0x80, z80.FLG_S + z80.FLG_H + z80.FLG_P_V},
		{0x80, z80.FLG_NONE, 0x81, z80.FLG_S},
		{0xff, z80.FLG_NONE, 0x00, z80.FLG_C + z80.FLG_H + z80.FLG_Z},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.ADD_XX
	cpu.Memory[1] = 0x01

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "ADD %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "ADD %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestADC(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0x01, z80.FLG_NONE},
		{0x00, z80.FLG_C, 0x02, z80.FLG_NONE},
		{0x7f, z80.FLG_NONE, 0x80, z80.FLG_S + z80.FLG_P_V + z80.FLG_H},
		{0x7f, z80.FLG_C, 0x81, z80.FLG_S + z80.FLG_P_V + z80.FLG_H},
		{0x80, z80.FLG_NONE, 0x81, z80.FLG_S},
		{0x80, z80.FLG_C, 0x82, z80.FLG_S},
		{0xff, z80.FLG_NONE, 0x00, z80.FLG_C + z80.FLG_H + z80.FLG_Z},
		{0xff, z80.FLG_C, 0x01, z80.FLG_C + z80.FLG_H},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.ADC_XX
	cpu.Memory[1] = 0x01

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "ADC %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "ADC %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestADD_HL_BC(t *testing.T) {
	var tests = []struct {
		old_h, old_l, old_b, old_c, old_f, new_h, new_l, new_f uint8
	}{
		{0x00, 0x00, 0x00, 0x00, z80.FLG_NONE, 0x00, 0x00, z80.FLG_NONE},
		{0x00, 0x00, 0x00, 0x08, z80.FLG_NONE, 0x00, 0x08, z80.FLG_NONE},
		{0x00, 0x00, 0x00, 0x08, z80.FLG_C, 0x00, 0x08, z80.FLG_NONE},
		{0x7f, 0xff, 0x00, 0x01, z80.FLG_NONE, 0x80, 0x00, z80.FLG_NONE | z80.FLG_H},
		{0xff, 0xff, 0x00, 0x01, z80.FLG_NONE, 0x00, 0x00, z80.FLG_C | z80.FLG_H},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.ADD_HL_BC

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.HL.Hi = tt.old_h
		cpu.HL.Lo = tt.old_l
		cpu.BC.Hi = tt.old_b
		cpu.BC.Lo = tt.old_c
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_h), cpu.HL.Hi, "ADD_HL_BC %02x %02x %02x", tt.old_h, tt.old_l, tt.old_f)
		assert.Equal(t, uint8(tt.new_l), cpu.HL.Lo, "ADD_HL_BC %02x %02x %02x", tt.old_h, tt.old_l, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "ADD_HL_BC %02x %02x %02x", tt.old_h, tt.old_l, tt.old_f)
	}

}

func TestSUB(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0xff, z80.FLG_C | z80.FLG_S | z80.FLG_H | z80.FLG_N},
		{0xff, z80.FLG_NONE, 0xfe, z80.FLG_S | z80.FLG_N},
		{0x80, z80.FLG_NONE, 0x7f, z80.FLG_P_V | z80.FLG_H | z80.FLG_N},
		{0x7f, z80.FLG_NONE, 0x7e, z80.FLG_NONE | z80.FLG_N},
		{0x01, z80.FLG_NONE, 0x00, z80.FLG_Z | z80.FLG_N},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.SUB_XX
	cpu.Memory[1] = 0x01

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "SUB %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "SUB %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestSBC(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0xff, z80.FLG_C | z80.FLG_S | z80.FLG_H | z80.FLG_N},
		{0x00, z80.FLG_C, 0xfe, z80.FLG_C | z80.FLG_S | z80.FLG_H | z80.FLG_N},
		{0xff, z80.FLG_NONE, 0xfe, z80.FLG_S | z80.FLG_N},
		{0xff, z80.FLG_C, 0xfd, z80.FLG_S | z80.FLG_N},
		{0x80, z80.FLG_NONE, 0x7f, z80.FLG_P_V | z80.FLG_H | z80.FLG_N},
		{0x80, z80.FLG_C, 0x7e, z80.FLG_P_V | z80.FLG_H | z80.FLG_N},
		{0x7f, z80.FLG_NONE, 0x7e, z80.FLG_NONE | z80.FLG_N},
		{0x7f, z80.FLG_C, 0x7d, z80.FLG_NONE | z80.FLG_N},
		{0x01, z80.FLG_NONE, 0x00, z80.FLG_Z | z80.FLG_N},
		{0x01, z80.FLG_C, 0xff, z80.FLG_S | z80.FLG_C | z80.FLG_H | z80.FLG_N},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.SBC_A_XX
	cpu.Memory[1] = 0x01

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "SBC %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "SBC %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestCP(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0x00, z80.FLG_C | z80.FLG_S | z80.FLG_H | z80.FLG_N},
		{0xff, z80.FLG_NONE, 0xff, z80.FLG_S | z80.FLG_N},
		{0x80, z80.FLG_NONE, 0x80, z80.FLG_P_V | z80.FLG_H | z80.FLG_N},
		{0x7f, z80.FLG_NONE, 0x7f, z80.FLG_NONE | z80.FLG_N},
		{0x01, z80.FLG_NONE, 0x01, z80.FLG_Z | z80.FLG_N},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.CP_XX
	cpu.Memory[1] = 0x01

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "CP %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "CP %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestINC(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0x01, z80.FLG_NONE},
		{0x7f, z80.FLG_NONE, 0x80, z80.FLG_S + z80.FLG_P_V + z80.FLG_H},
		{0x80, z80.FLG_NONE, 0x81, z80.FLG_S},
		{0xff, z80.FLG_NONE, 0x00, z80.FLG_Z + z80.FLG_H},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.INC_A

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "INC %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "INC %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestDEC(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0xff, z80.FLG_S + z80.FLG_H + z80.FLG_N},
		{0xff, z80.FLG_NONE, 0xfe, z80.FLG_S + z80.FLG_N},
		{0x80, z80.FLG_NONE, 0x7f, z80.FLG_P_V + z80.FLG_H + z80.FLG_N},
		{0x7f, z80.FLG_NONE, 0x7e, z80.FLG_NONE + z80.FLG_N},
		{0x01, z80.FLG_NONE, 0x00, z80.FLG_Z + z80.FLG_N},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.DEC_A

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "DEC %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "DEC %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestAND(t *testing.T) {
	var tests = []struct {
		old_a, old_f, value, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0xff, 0x00, z80.FLG_Z | z80.FLG_H | z80.FLG_P_V},
		{0xff, z80.FLG_NONE, 0x55, 0x55, z80.FLG_H | z80.FLG_P_V},
		{0xff, z80.FLG_NONE, 0xaa, 0xaa, z80.FLG_S | z80.FLG_H | z80.FLG_P_V},
		{0xff, z80.FLG_NONE, 0x00, 0x00, z80.FLG_Z | z80.FLG_H | z80.FLG_P_V},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.AND_XX

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.Memory[1] = tt.value
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "AND %02x %02x %02x", tt.old_a, tt.old_f, tt.value)
		assert.Equal(t, uint8(tt.new_f), flags, "AND %02x %02x %02x", tt.old_a, tt.old_f, tt.value)
	}

}

func TestOR(t *testing.T) {
	var tests = []struct {
		old_a, old_f, value, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0xff, 0xff, z80.FLG_S | z80.FLG_P_V},
		{0xff, z80.FLG_NONE, 0x55, 0xff, z80.FLG_S | z80.FLG_P_V},
		{0x00, z80.FLG_NONE, 0xaa, 0xaa, z80.FLG_S | z80.FLG_P_V},
		{0x00, z80.FLG_NONE, 0x55, 0x55, z80.FLG_P_V},
		{0x00, z80.FLG_NONE, 0x00, 0x00, z80.FLG_Z | z80.FLG_P_V},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.OR_XX

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.Memory[1] = tt.value
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "OR %02x %02x %02x", tt.old_a, tt.old_f, tt.value)
		assert.Equal(t, uint8(tt.new_f), flags, "OR %02x %02x %02x", tt.old_a, tt.old_f, tt.value)
	}

}

func TestXOR(t *testing.T) {
	var tests = []struct {
		old_a, old_f, value, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0xff, 0xff, z80.FLG_S | z80.FLG_P_V},
		{0xff, z80.FLG_NONE, 0x55, 0xaa, z80.FLG_S | z80.FLG_P_V},
		{0x00, z80.FLG_NONE, 0xaa, 0xaa, z80.FLG_S | z80.FLG_P_V},
		{0x00, z80.FLG_NONE, 0x55, 0x55, z80.FLG_P_V},
		{0x00, z80.FLG_NONE, 0x00, 0x00, z80.FLG_Z | z80.FLG_P_V},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.XOR_XX

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.Memory[1] = tt.value
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "AND %02x %02x %02x", tt.old_a, tt.old_f, tt.value)
		assert.Equal(t, uint8(tt.new_f), flags, "AND %02x %02x %02x", tt.old_a, tt.old_f, tt.value)
	}

}

func TestRLA(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		// Rotate A left. Bit 7 -> Carry, Carry -> Bit 0
		{0x00, z80.FLG_NONE, 0x00, z80.FLG_NONE},
		{0x00, z80.FLG_C, 0x01, z80.FLG_NONE},
		{0x01, z80.FLG_NONE, 0x02, z80.FLG_NONE},
		{0x01, z80.FLG_C, 0x03, z80.FLG_NONE},
		{0x80, z80.FLG_NONE, 0x00, z80.FLG_C},
		{0x80, z80.FLG_C, 0x01, z80.FLG_C},
		{0xff, z80.FLG_NONE, 0xfe, z80.FLG_C},
		{0xff, z80.FLG_C, 0xff, z80.FLG_C},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.RLA

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "RLA %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "RLA %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestRLCA(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		// Rotate A left. Bit 7 -> Bit 0 and Carry
		{0x00, z80.FLG_NONE, 0x00, z80.FLG_NONE},
		{0x00, z80.FLG_C, 0x00, z80.FLG_NONE},
		{0x01, z80.FLG_NONE, 0x02, z80.FLG_NONE},
		{0x01, z80.FLG_C, 0x02, z80.FLG_NONE},
		{0x80, z80.FLG_NONE, 0x01, z80.FLG_C},
		{0x80, z80.FLG_C, 0x01, z80.FLG_C},
		{0xff, z80.FLG_NONE, 0xff, z80.FLG_C},
		{0xff, z80.FLG_C, 0xff, z80.FLG_C},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.RLCA

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "RLCA %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "RLCA %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestRRA(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		// Rotate A right. Bit 0 -> Carry, Carry -> Bit 7
		{0x00, z80.FLG_NONE, 0x00, z80.FLG_NONE},
		{0x00, z80.FLG_C, 0x80, z80.FLG_NONE},
		{0x01, z80.FLG_NONE, 0x00, z80.FLG_C},
		{0x01, z80.FLG_C, 0x80, z80.FLG_C},
		{0x80, z80.FLG_NONE, 0x40, z80.FLG_NONE},
		{0x80, z80.FLG_C, 0xc0, z80.FLG_NONE},
		{0xff, z80.FLG_NONE, 0x7f, z80.FLG_C},
		{0xff, z80.FLG_C, 0xff, z80.FLG_C},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.RRA

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "RRA %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "RRA %02x %02x", tt.old_a, tt.old_f)
	}

}

func TestRRCA(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		// Rotate A right. Bit 0 -> Bit 7 and Carry
		{0x00, z80.FLG_NONE, 0x00, z80.FLG_NONE},
		{0x00, z80.FLG_C, 0x00, z80.FLG_NONE},
		{0x01, z80.FLG_NONE, 0x80, z80.FLG_C},
		{0x01, z80.FLG_C, 0x80, z80.FLG_C},
		{0x80, z80.FLG_NONE, 0x40, z80.FLG_NONE},
		{0x80, z80.FLG_C, 0x40, z80.FLG_NONE},
		{0xff, z80.FLG_NONE, 0xff, z80.FLG_C},
		{0xff, z80.FLG_C, 0xff, z80.FLG_C},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.RRCA

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "RRCA %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "RRCA %02x %02x", tt.old_a, tt.old_f)
	}
}

func TestSCF(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0x00, z80.FLG_C},
		{0x00, z80.FLG_C, 0x00, z80.FLG_C},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.SCF

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "SCF %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "SCF %02x %02x", tt.old_a, tt.old_f)
	}
}

func TestCCF(t *testing.T) {
	var tests = []struct {
		old_a, old_f, new_a, new_f uint8
	}{
		{0x00, z80.FLG_NONE, 0x00, z80.FLG_C},
		{0x00, z80.FLG_C, 0x00, z80.FLG_H},
		{0x00, z80.FLG_C | z80.FLG_N | z80.FLG_H, 0x00, z80.FLG_H},
	}

	cpu := newCPU()
	cpu.Memory[0] = z80.CCF

	for _, tt := range tests {
		cpu.PC = 0x0000
		cpu.AF.Hi = tt.old_a
		cpu.AF.Lo = tt.old_f
		cpu.Step()
		flags := cpu.AF.Lo &^ FLG_UNDOC
		assert.Equal(t, uint8(tt.new_a), cpu.AF.Hi, "CCF %02x %02x", tt.old_a, tt.old_f)
		assert.Equal(t, uint8(tt.new_f), flags, "CCF %02x %02x", tt.old_a, tt.old_f)
	}
}

func TestPUSH_POP(t *testing.T) {
	cpu := newCPU()

	cpu.Memory[0] = z80.PUSH_BC
	cpu.Memory[1] = z80.POP_DE
	cpu.SP = 0x100
	cpu.BC.Hi = 0x12
	cpu.BC.Lo = 0x34

	cpu.Step()
	assert.Equal(t, uint16(0x00fe), cpu.SP)
	assert.Equal(t, uint8(0x12), cpu.Memory[0x00ff])
	assert.Equal(t, uint8(0x34), cpu.Memory[0x00fe])

	cpu.Step()
	assert.Equal(t, uint16(0x0100), cpu.SP)
	assert.Equal(t, uint8(0x12), cpu.DE.Hi)
	assert.Equal(t, uint8(0x34), cpu.DE.Lo)

}

func newCPU() z80.CPU {
	cpu := z80.CPU{
		PC: uint16(0),
		AF: z80.Register{Hi: 0x00, Lo: 0x00},
		BC: z80.Register{Hi: 0x00, Lo: 0x00},
		DE: z80.Register{Hi: 0x00, Lo: 0x00},
		HL: z80.Register{Hi: 0x00, Lo: 0x00},
	}
	return cpu
}
