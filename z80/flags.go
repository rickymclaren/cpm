package z80

import (
	"math/bits"
)

func (cpu *CPU) updateFlagArith8(r, a, b uint16, subtract bool) {
	c := r ^ a ^ b
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N | FLG_C
	var or uint8
	or |= uint8(r) & FLG_STD
	if uint8(r) == 0 {
		or |= FLG_Z
	}
	or |= uint8(c) & FLG_H
	or |= uint8((c>>6)^(c>>5)) & FLG_P_V
	if subtract {
		or |= FLG_N
	}
	or |= uint8(r>>8) & FLG_C
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func (cpu *CPU) updateFlagLogic8(r uint8, and bool) {
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N | FLG_C
	var or uint8
	or |= r & FLG_STD
	if r == 0 {
		or |= FLG_Z
	}
	if and {
		or |= FLG_H
	}
	or |= (uint8(bits.OnesCount8(r)%2) - 1) & FLG_P_V
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func (cpu *CPU) updateFlagCPx(r, a, b uint8) {
	c := r ^ a ^ b
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N
	var or uint8
	or |= r & FLG_S
	if r == 0 {
		or |= FLG_Z
	}
	or |= c & FLG_H
	if cpu.BC.Lo != 0 || cpu.BC.Hi != 0 {
		or |= FLG_P_V
	}
	or |= FLG_N

	r2 := r - (c & 0x10 >> 4)
	or |= r2 & 0x02 << 4 // mask5
	or |= r2 & FLG_3
	//if r&0x0f == 8 && c&maskH != 0 {
	//	or &= ^uint8(mask3)
	//}

	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func (cpu *CPU) updateFlagLDID(a uint8) {
	var nand uint8 = FLG_H | FLG_P_V | FLG_N | FLG_5 | FLG_3
	var or uint8
	if cpu.BC.Lo != 0 || cpu.BC.Hi != 0 {
		or |= FLG_P_V
	}
	a += cpu.AF.Hi
	or |= a&FLG_3 | (a<<4)&FLG_5
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func (cpu *CPU) updateFlagBitop(r uint8, carry uint8) {
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N | FLG_C
	var or uint8
	or |= r & FLG_STD
	if r == 0 {
		or |= FLG_Z
	}
	or |= (uint8(bits.OnesCount8(r)%2) - 1) & FLG_P_V
	or |= carry & FLG_C
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func (cpu *CPU) updateFlagRxD(r uint8) {
	var nand uint8 = FLG_STD | FLG_Z | FLG_H | FLG_P_V | FLG_N
	var or uint8
	or |= uint8(r) & FLG_STD
	if uint8(r) == 0 {
		or |= FLG_Z
	}
	or |= (uint8(bits.OnesCount8(r)%2) - 1) & FLG_P_V
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func (cpu *CPU) updateFlagRL(r uint8) {
	var nand uint8 = FLG_H | FLG_N | FLG_C | FLG_5 | FLG_3
	var or = (r >> 7) & FLG_C
	or |= r << 1 & (FLG_5 | FLG_3)
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func (cpu *CPU) updateFlagRR(r uint8) {
	var nand uint8 = FLG_H | FLG_N | FLG_C | FLG_5 | FLG_3
	var or = r & FLG_C
	or |= r >> 1 & (FLG_5 | FLG_3)
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}
