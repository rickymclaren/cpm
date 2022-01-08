package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/koron-go/z80"
)

// Start is an address where a program starts.
const Start = 0x0100

type Machine struct {
	Memory      [65536]uint8
	Cpu         z80.CPU
	fdc_drive   uint8
	fdc_track   uint8
	fdc_sector  uint8
	fdc_command uint8
	fdc_status  uint8
	fdc_dma_low uint8
	fdc_dma_hi  uint8
}

func NewMachine(base int) *Machine {
	m := new(Machine)
	m.Cpu = z80.CPU{
		States: z80.States{
			SPR: z80.SPR{PC: uint16(base)},
			GPR: z80.GPR{
				AF: z80.Register{Hi: 0x00, Lo: 0x00},
				BC: z80.Register{Hi: 0x00, Lo: 0x00},
				DE: z80.Register{Hi: 0x00, Lo: 0x00},
				HL: z80.Register{Hi: 0x00, Lo: 0x00},
			},
		},
		Memory: m,
		IO:     m,
	}
	return m
}

func (m *Machine) In(addr uint8) uint8 {
	value := uint8(0x00)
	switch addr {
	case 0x00:
		// Console Input Status - is there a character available
		// 0xff yes, 0x00 no
		value = 0x00
	case 0x01:
		// Console Data
		var buffer = make([]byte, 1)
		_, err := os.Stdin.Read(buffer)
		if err != nil {
			panic(err)
		}
		value = buffer[0]
	case 0x02:
		// Printer Status
		value = 0x00
	case 0x03:
		// Printer Data
		fmt.Printf("not impl. I/O In from Printer Status\n", addr)
	case 0x05:
		// Aux Data
		fmt.Printf("not impl. I/O In from Aux data\n", addr)
	case 0x0a:
		value = m.fdc_drive
	case 0x0b:
		value = m.fdc_track
	case 0x0c:
		value = m.fdc_sector
	case 0x0d:
		value = m.fdc_command
	case 0x0e:
		value = m.fdc_status
	case 0x0f:
		value = m.fdc_dma_low
	case 0x10:
		value = m.fdc_dma_hi
	default:
		fmt.Printf("not impl. I/O In addr=0x%02x\n", addr)
	}

	return uint8(value)
}

func (m *Machine) Out(addr uint8, value uint8) {
	switch addr {
	case 0x00:
		// Console Input Status - not implemented
	case 0x01:
		// Console Data
		fmt.Printf("%c", value)
	case 0x02:
		// Printer Status - not implemented
	case 0x03:
		// Printer Data - not implemented
	case 0x05:
		// Aux Data - not implemeneted
	case 0x0a:
		m.fdc_drive = value
	case 0x0b:
		m.fdc_track = value
	case 0x0c:
		m.fdc_sector = value
	case 0x0d:
		m.fdc_command = value
		switch value {
		case 0: // Disk Read
			image := diskImage(m.fdc_drive)
			data, err := ioutil.ReadFile(image)
			if err != nil {
				panic(err)
			}
			dma := (int(m.fdc_dma_hi) << 8) | int(m.fdc_dma_low)
			sector := m.fdc_track*26 + m.fdc_sector - 1
			offset := int(sector) * 128
			sector_data := data[offset : offset+128]
			m.put(dma, sector_data...)
		default:
			fmt.Printf("FDC Drive:%d Track:%d Sector:%02d Command:%02x \n",
				m.fdc_drive,
				m.fdc_track,
				m.fdc_sector,
				m.fdc_command)
		}
	case 0x0e:
		m.fdc_status = value
	case 0x0f:
		m.fdc_dma_low = value
	case 0x10:
		m.fdc_dma_hi = value
	default:
		fmt.Printf("not impl. I/O Out addr=0x%02x\n", addr)
	}

}

func (m *Machine) Set(addr uint16, data uint8) {
	m.Memory[addr] = data
}

func (m *Machine) Get(addr uint16) uint8 {
	return m.Memory[addr]
}

// put puts "data" block from addr.
func (m *Machine) put(addr int, data ...uint8) {
	copy(m.Memory[addr:addr+len(data)], data)
}

// LoadFile loads a section of a file into address.
func (m *Machine) LoadFile(name string, address int, offset int, length int) error {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	m.put(address, data[offset:(offset+length)]...)
	return nil
}

func diskImage(disk uint8) string {
	letters := "abcdefghijklmnop"
	return fmt.Sprintf("disks/%c/DISK.IMG", letters[disk])
}

func main() {
	b := (64 - 20) * 1024
	base := 0x3400 + b
	bdos := 0x3c00 + b
	bios := 0x4a00 + b
	warmboot := uint16(bios + 3)

	m := NewMachine(base)
	m.LoadFile("disks/a/DISK.IMG", base, 0x0080, 0x1980)

	m.Cpu.Memory.Set(0, 0xc3) // JP to WARMBOOT
	m.Cpu.Memory.Set(1, uint8(warmboot&0xff))
	m.Cpu.Memory.Set(2, uint8(warmboot>>8&0xff))
	m.Cpu.Memory.Set(3, 0x00) // IOBYTE
	m.Cpu.Memory.Set(4, 0x00) // Current Drive
	m.Cpu.Memory.Set(5, 0xc3) // JP to BDOS
	m.Cpu.Memory.Set(6, uint8(bdos&0xff))
	m.Cpu.Memory.Set(7, uint8(bdos>>8&0xff))

	err := m.Cpu.Run(context.Background())

	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

}
