package main

import (
	"bufio"
	"cpm/z80"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const HALT = 0x76
const FLG_UNDOC = z80.FLG_3 | z80.FLG_5

func TestFuse(t *testing.T) {

	inFile, err := os.Open("fuse/tests/tests.in")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	expFile, err := os.Open("fuse/tests/tests.expected")
	if err != nil {
		panic(err)
	}
	defer expFile.Close()

	scannerIn := bufio.NewScanner(inFile)
	scannerExp := bufio.NewScanner(expFile)
	for scannerIn.Scan() {
		scannerExp.Scan()
		for scannerIn.Text() == "" {
			scannerIn.Scan()
		}
		for scannerExp.Text() == "" {
			scannerExp.Scan()
		}
		testName := scannerIn.Text()
		expectedName := scannerExp.Text()
		if testName != expectedName {
			os.Exit(-1)
		}

		fmt.Printf("Test: %s\n", testName)

		scannerIn.Scan()
		mainRegs := strings.Split(scannerIn.Text(), " ")
		af, _ := strconv.ParseUint(mainRegs[0], 16, 0)
		bc, _ := strconv.ParseUint(mainRegs[1], 16, 0)
		de, _ := strconv.ParseUint(mainRegs[2], 16, 0)
		hl, _ := strconv.ParseUint(mainRegs[3], 16, 0)
		af_, _ := strconv.ParseUint(mainRegs[4], 16, 0)
		bc_, _ := strconv.ParseUint(mainRegs[5], 16, 0)
		de_, _ := strconv.ParseUint(mainRegs[6], 16, 0)
		hl_, _ := strconv.ParseUint(mainRegs[7], 16, 0)
		ix, _ := strconv.ParseUint(mainRegs[8], 16, 0)
		iy, _ := strconv.ParseUint(mainRegs[9], 16, 0)
		sp, _ := strconv.ParseUint(mainRegs[10], 16, 0)
		pc, _ := strconv.ParseUint(mainRegs[11], 16, 0)

		cpu := newCPU()

		cpu.AF.Hi = uint8(af >> 8)
		cpu.AF.Lo = uint8(af & 0xff)
		cpu.BC.Hi = uint8(bc >> 8)
		cpu.BC.Lo = uint8(bc & 0xff)
		cpu.DE.Hi = uint8(de >> 8)
		cpu.DE.Lo = uint8(de & 0xff)
		cpu.HL.Hi = uint8(hl >> 8)
		cpu.HL.Lo = uint8(hl & 0xff)
		cpu.AF_.Hi = uint8(af_ >> 8)
		cpu.AF_.Lo = uint8(af_ & 0xff)
		cpu.BC_.Hi = uint8(bc_ >> 8)
		cpu.BC_.Lo = uint8(bc_ & 0xff)
		cpu.DE_.Hi = uint8(de_ >> 8)
		cpu.DE_.Lo = uint8(de_ & 0xff)
		cpu.HL_.Hi = uint8(hl_ >> 8)
		cpu.HL_.Lo = uint8(hl_ & 0xff)
		cpu.IX.Hi = uint8(ix >> 8)
		cpu.IX.Lo = uint8(ix & 0xff)
		cpu.IY.Hi = uint8(iy >> 8)
		cpu.IY.Lo = uint8(iy & 0xff)
		cpu.SP = uint16(sp)
		cpu.PC = uint16(pc)

		scannerIn.Scan() // other regs
		scannerIn.Scan() // memory set

		for scannerIn.Text() != "-1" {
			memWrites := strings.Split(scannerIn.Text(), " ")
			addr, _ := strconv.ParseUint(memWrites[0], 16, 0)
			for i := 1; i < (len(memWrites)); i++ {
				byte := memWrites[i]
				if byte != "-1" {
					value, _ := strconv.ParseUint(byte, 16, 0)
					cpu.Memory[uint16(addr)] = uint8(value)
					addr++
				}
			}
			scannerIn.Scan()
		}

		steps := stepsFor(testName)
		for i := 0; i < steps; i++ {
			cpu.Step()
		}

		scannerExp.Scan()
		for len(scannerExp.Text()) > 0 && strings.HasPrefix(scannerExp.Text(), " ") {
			scannerExp.Scan() // skip events
		}
		expectedRegs := strings.Split(scannerExp.Text(), " ")
		af, _ = strconv.ParseUint(expectedRegs[0], 16, 0)
		bc, _ = strconv.ParseUint(expectedRegs[1], 16, 0)
		de, _ = strconv.ParseUint(expectedRegs[2], 16, 0)
		hl, _ = strconv.ParseUint(expectedRegs[3], 16, 0)
		af_, _ = strconv.ParseUint(expectedRegs[4], 16, 0)
		bc_, _ = strconv.ParseUint(expectedRegs[5], 16, 0)
		de_, _ = strconv.ParseUint(expectedRegs[6], 16, 0)
		hl_, _ = strconv.ParseUint(expectedRegs[7], 16, 0)
		ix, _ = strconv.ParseUint(expectedRegs[8], 16, 0)
		iy, _ = strconv.ParseUint(expectedRegs[9], 16, 0)
		sp, _ = strconv.ParseUint(expectedRegs[10], 16, 0)
		pc, _ = strconv.ParseUint(expectedRegs[11], 16, 0)

		if steps > 0 {
			assert.Equal(t, uint8(af>>8), cpu.AF.Hi, "%s %s", testName, "A")
			assert.Equal(t, uint8(af&0xff)&^FLG_UNDOC, cpu.AF.Lo&^FLG_UNDOC, "%s %s", testName, "F")
			assert.Equal(t, uint8(bc>>8), cpu.BC.Hi, "%s %s", testName, "B")
			assert.Equal(t, uint8(bc&0xff), cpu.BC.Lo, "%s %s", testName, "C")
			assert.Equal(t, uint8(de>>8), cpu.DE.Hi, "%s %s", testName, "D")
			assert.Equal(t, uint8(de&0xff), cpu.DE.Lo, "%s %s", testName, "E")
			assert.Equal(t, uint8(hl>>8), cpu.HL.Hi, "%s %s", testName, "H")
			assert.Equal(t, uint8(hl&0xff), cpu.HL.Lo, "%s %s", testName, "L")
			assert.Equal(t, uint8(af_>>8), cpu.AF_.Hi, "%s %s", testName, "A'")
			assert.Equal(t, uint8(af_&0xff), cpu.AF_.Lo, "%s %s", testName, "F'")
			assert.Equal(t, uint8(bc_>>8), cpu.BC_.Hi, "%s %s", testName, "B'")
			assert.Equal(t, uint8(bc_&0xff), cpu.BC_.Lo, "%s %s", testName, "C'")
			assert.Equal(t, uint8(de_>>8), cpu.DE_.Hi, "%s %s", testName, "D'")
			assert.Equal(t, uint8(de_&0xff), cpu.DE_.Lo, "%s %s", testName, "E'")
			assert.Equal(t, uint8(hl_>>8), cpu.HL_.Hi, "%s %s", testName, "H'")
			assert.Equal(t, uint8(hl_&0xff), cpu.HL_.Lo, "%s %s", testName, "L'")
			assert.Equal(t, uint8(ix>>8), cpu.IX.Hi, "%s %s", testName, "IXH")
			assert.Equal(t, uint8(ix&0xff), cpu.IX.Lo, "%s %s", testName, "IXL")
			assert.Equal(t, uint8(iy>>8), cpu.IY.Hi, "%s %s", testName, "IYH")
			assert.Equal(t, uint8(iy&0xff), cpu.IY.Lo, "%s %s", testName, "IHL")
			assert.Equal(t, uint16(sp), cpu.SP, "%s %s", testName, "SP")
			assert.Equal(t, uint16(pc), cpu.PC, "%s %s", testName, "PC")
		}

		scannerExp.Scan() // extra regs
		scannerExp.Scan() // memory checks
		for scannerExp.Text() != "" {
			memWrites := strings.Split(scannerExp.Text(), " ")
			addr, _ := strconv.ParseUint(memWrites[0], 16, 0)
			for i := 1; i < (len(memWrites)); i++ {
				byte := memWrites[i]
				if byte != "-1" {
					value, _ := strconv.ParseUint(byte, 16, 0)
					if steps > 0 {
						assert.Equal(t, uint8(value), cpu.Memory[uint16(addr)], testName)
					}
					addr++
				}
			}
			scannerExp.Scan()
		}

	}
}

// Zero steps = skip test, mostly for undocumented instructions
func stepsFor(name string) int {
	steps := 1
	switch name {
	case "10":
		steps = 17
	case "76":
		steps = 0
	case "d3", "d3_1", "d3_2", "d3_3", "d3_4":
		steps = 0
	case "db", "db_1", "db_2", "db_3":
		steps = 0
	case "dd00":
		steps = 0
	case "ddcb00", "ddcb01", "ddcb02", "ddcb03", "ddcb04", "ddcb05", "ddcb07":
		steps = 0
	case "ddcb08", "ddcb09", "ddcb0a", "ddcb0b", "ddcb0c", "ddcb0d", "ddcb0f":
		steps = 0
	case "ddcb10", "ddcb11", "ddcb12", "ddcb13", "ddcb14", "ddcb15", "ddcb17":
		steps = 0
	case "ddcb18", "ddcb19", "ddcb1a", "ddcb1b", "ddcb1c", "ddcb1d", "ddcb1f":
		steps = 0
	case "ddcb20", "ddcb21", "ddcb22", "ddcb23", "ddcb24", "ddcb25", "ddcb27":
		steps = 0
	case "ddcb28", "ddcb29", "ddcb2a", "ddcb2b", "ddcb2c", "ddcb2d", "ddcb2f":
		steps = 0
	case "ddcb30", "ddcb31", "ddcb32", "ddcb33", "ddcb34", "ddcb35", "ddcb37":
		steps = 0
	case "ddcb38", "ddcb39", "ddcb3a", "ddcb3b", "ddcb3c", "ddcb3d", "ddcb3f":
		steps = 0
	case "ddcb40", "ddcb41", "ddcb42", "ddcb43", "ddcb44", "ddcb45", "ddcb47":
		steps = 0
	case "ddcb48", "ddcb49", "ddcb4a", "ddcb4b", "ddcb4c", "ddcb4d", "ddcb4f":
		steps = 0
	case "ddcb50", "ddcb51", "ddcb52", "ddcb53", "ddcb54", "ddcb55", "ddcb57":
		steps = 0
	case "ddcb58", "ddcb59", "ddcb5a", "ddcb5b", "ddcb5c", "ddcb5d", "ddcb5f":
		steps = 0
	case "ddcb60", "ddcb61", "ddcb62", "ddcb63", "ddcb64", "ddcb65", "ddcb67":
		steps = 0
	case "ddcb68", "ddcb69", "ddcb6a", "ddcb6b", "ddcb6c", "ddcb6d", "ddcb6f":
		steps = 0
	case "ddcb70", "ddcb71", "ddcb72", "ddcb73", "ddcb74", "ddcb75", "ddcb77":
		steps = 0
	case "ddcb78", "ddcb79", "ddcb7a", "ddcb7b", "ddcb7c", "ddcb7d", "ddcb7f":
		steps = 0
	case "ddcb80", "ddcb81", "ddcb82", "ddcb83", "ddcb84", "ddcb85", "ddcb87":
		steps = 0
	case "ddcb88", "ddcb89", "ddcb8a", "ddcb8b", "ddcb8c", "ddcb8d", "ddcb8f":
		steps = 0
	case "ddcb90", "ddcb91", "ddcb92", "ddcb93", "ddcb94", "ddcb95", "ddcb97":
		steps = 0
	case "ddcb98", "ddcb99", "ddcb9a", "ddcb9b", "ddcb9c", "ddcb9d", "ddcb9f":
		steps = 0
	case "ddcba0", "ddcba1", "ddcba2", "ddcba3", "ddcba4", "ddcba5", "ddcba7":
		steps = 0
	case "ddcba8", "ddcba9", "ddcbaa", "ddcbab", "ddcbac", "ddcbad", "ddcbaf":
		steps = 0
	case "ddcbb0", "ddcbb1", "ddcbb2", "ddcbb3", "ddcbb4", "ddcbb5", "ddcbb7":
		steps = 0
	case "ddcbb8", "ddcbb9", "ddcbba", "ddcbbb", "ddcbbc", "ddcbbd", "ddcbbf":
		steps = 0
	case "ddcbc0", "ddcbc1", "ddcbc2", "ddcbc3", "ddcbc4", "ddcbc5", "ddcbc7":
		steps = 0
	case "ddcbc8", "ddcbc9", "ddcbca", "ddcbcb", "ddcbcc", "ddcbcd", "ddcbcf":
		steps = 0
	case "ddcbd0", "ddcbd1", "ddcbd2", "ddcbd3", "ddcbd4", "ddcbd5", "ddcbd7":
		steps = 0
	case "ddcbd8", "ddcbd9", "ddcbda", "ddcbdb", "ddcbdc", "ddcbdd", "ddcbdf":
		steps = 0
	case "ddcbe0", "ddcbe1", "ddcbe2", "ddcbe3", "ddcbe4", "ddcbe5", "ddcbe7":
		steps = 0
	case "ddcbe8", "ddcbe9", "ddcbea", "ddcbeb", "ddcbec", "ddcbed", "ddcbef":
		steps = 0
	case "ddcbf0", "ddcbf1", "ddcbf2", "ddcbf3", "ddcbf4", "ddcbf5", "ddcbf7":
		steps = 0
	case "ddcbf8", "ddcbf9", "ddcbfa", "ddcbfb", "ddcbfc", "ddcbfd", "ddcbff":
		steps = 0
	case "ddfd00":
		steps = 0
	case "ed40", "ed45", "ed48", "ed4c", "ed4d":
		steps = 0
	case "ed50", "ed54", "ed55", "ed57", "ed58", "ed5c", "ed5d", "ed5f":
		steps = 0
	case "ed60", "ed63", "ed64", "ed65", "ed68", "ed6c", "ed6d":
		steps = 0
	case "ed74", "ed75", "ed78", "ed7c", "ed7d":
		steps = 0
	case "eda2", "eda2_01", "eda2_02", "eda2_03":
		steps = 0
	case "eda3", "eda3_01", "eda3_02", "eda3_03", "eda3_04", "eda3_05", "eda3_06", "eda3_07", "eda3_08", "eda3_09", "eda3_10", "eda3_11":
		steps = 0
	case "edaa", "edaa_01", "edaa_02", "edaa_03", "edab", "edab_01", "edab_02":
		steps = 0
	case "edb0":
		steps = 16
	case "edb0_1":
		steps = 2
	case "edb1", "edb1_1":
		steps = 4
	case "edb2", "edb2_1", "edb3", "edb3_1":
		steps = 0
	case "edb8":
		steps = 8
	case "edb8_1":
		steps = 2
	case "edb9", "edb9_1":
		steps = 8
	case "edba", "edba_1", "edbb", "edbb_1":
		steps = 0
	case "fd25", "fd2c", "fd2d":
		steps = 0
	case "fdcb00", "fdcb01", "fdcb02", "fdcb03", "fdcb04", "fdcb05", "fdcb07":
		steps = 0
	case "fdcb08", "fdcb09", "fdcb0a", "fdcb0b", "fdcb0c", "fdcb0d", "fdcb0f":
		steps = 0
	case "fdcb10", "fdcb11", "fdcb12", "fdcb13", "fdcb14", "fdcb15", "fdcb17":
		steps = 0
	case "fdcb18", "fdcb19", "fdcb1a", "fdcb1b", "fdcb1c", "fdcb1d", "fdcb1f":
		steps = 0
	case "fdcb20", "fdcb21", "fdcb22", "fdcb23", "fdcb24", "fdcb25", "fdcb27":
		steps = 0
	case "fdcb28", "fdcb29", "fdcb2a", "fdcb2b", "fdcb2c", "fdcb2d", "fdcb2f":
		steps = 0
	case "fdcb30", "fdcb31", "fdcb32", "fdcb33", "fdcb34", "fdcb35", "fdcb37":
		steps = 0
	case "fdcb38", "fdcb39", "fdcb3a", "fdcb3b", "fdcb3c", "fdcb3d", "fdcb3f":
		steps = 0
	case "fdcb40", "fdcb41", "fdcb42", "fdcb43", "fdcb44", "fdcb45", "fdcb47":
		steps = 0
	case "fdcb48", "fdcb49", "fdcb4a", "fdcb4b", "fdcb4c", "fdcb4d", "fdcb4f":
		steps = 0
	case "fdcb50", "fdcb51", "fdcb52", "fdcb53", "fdcb54", "fdcb55", "fdcb57":
		steps = 0
	case "fdcb58", "fdcb59", "fdcb5a", "fdcb5b", "fdcb5c", "fdcb5d", "fdcb5f":
		steps = 0
	case "fdcb60", "fdcb61", "fdcb62", "fdcb63", "fdcb64", "fdcb65", "fdcb67":
		steps = 0
	case "fdcb68", "fdcb69", "fdcb6a", "fdcb6b", "fdcb6c", "fdcb6d", "fdcb6f":
		steps = 0
	case "fdcb70", "fdcb71", "fdcb72", "fdcb73", "fdcb74", "fdcb75", "fdcb77":
		steps = 0
	case "fdcb78", "fdcb79", "fdcb7a", "fdcb7b", "fdcb7c", "fdcb7d", "fdcb7f":
		steps = 0
	case "fdcb80", "fdcb81", "fdcb82", "fdcb83", "fdcb84", "fdcb85", "fdcb87":
		steps = 0
	case "fdcb88", "fdcb89", "fdcb8a", "fdcb8b", "fdcb8c", "fdcb8d", "fdcb8f":
		steps = 0
	case "fdcb90", "fdcb91", "fdcb92", "fdcb93", "fdcb94", "fdcb95", "fdcb97":
		steps = 0
	case "fdcb98", "fdcb99", "fdcb9a", "fdcb9b", "fdcb9c", "fdcb9d", "fdcb9f":
		steps = 0
	case "fdcba0", "fdcba1", "fdcba2", "fdcba3", "fdcba4", "fdcba5", "fdcba7":
		steps = 0
	case "fdcba8", "fdcba9", "fdcbaa", "fdcbab", "fdcbac", "fdcbad", "fdcbaf":
		steps = 0
	case "fdcbb0", "fdcbb1", "fdcbb2", "fdcbb3", "fdcbb4", "fdcbb5", "fdcbb7":
		steps = 0
	case "fdcbb8", "fdcbb9", "fdcbba", "fdcbbb", "fdcbbc", "fdcbbd", "fdcbbf":
		steps = 0
	case "fdcbc0", "fdcbc1", "fdcbc2", "fdcbc3", "fdcbc4", "fdcbc5", "fdcbc7":
		steps = 0
	case "fdcbc8", "fdcbc9", "fdcbca", "fdcbcb", "fdcbcc", "fdcbcd", "fdcbcf":
		steps = 0
	case "fdcbd0", "fdcbd1", "fdcbd2", "fdcbd3", "fdcbd4", "fdcbd5", "fdcbd7":
		steps = 0
	case "fdcbd8", "fdcbd9", "fdcbda", "fdcbdb", "fdcbdc", "fdcbdd", "fdcbdf":
		steps = 0
	case "fdcbe0", "fdcbe1", "fdcbe2", "fdcbe3", "fdcbe4", "fdcbe5", "fdcbe7":
		steps = 0
	case "fdcbe8", "fdcbe9", "fdcbea", "fdcbeb", "fdcbec", "fdcbed", "fdcbef":
		steps = 0
	case "fdcbf0", "fdcbf1", "fdcbf2", "fdcbf3", "fdcbf4", "fdcbf5", "fdcbf7":
		steps = 0
	case "fdcbf8", "fdcbf9", "fdcbfa", "fdcbfb", "fdcbfc", "fdcbfd", "fdcbff":
		steps = 0
	}
	return steps
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
