package main

import (
	"cpm/z80"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const CR = 0x0d
const LF = 0x0a
const LOGGING = false

type Machine struct {
	Cpu         z80.CPU
	Console     chan byte
	fdc_drive   int
	fdc_track   int
	fdc_sector  int
	fdc_command uint8
	fdc_status  uint8
	fdc_dma_low uint8
	fdc_dma_hi  uint8
}

func NewMachine(base int) *Machine {
	m := new(Machine)
	m.Cpu = z80.CPU{
		PC: uint16(base),
		AF: z80.Register{Hi: 0x00, Lo: 0x00},
		BC: z80.Register{Hi: 0x00, Lo: 0x00},
		DE: z80.Register{Hi: 0x00, Lo: 0x00},
		HL: z80.Register{Hi: 0x00, Lo: 0x00},
	}
	m.Cpu.Io = m
	m.Console = make(chan byte, 256)
	return m
}

func (m *Machine) In(addr uint8) uint8 {
	value := uint8(0x00)
	switch addr {
	case 0x00:
		// Console Input Status - is there a character available
		// 0xff yes, 0x00 no
		if len(m.Console) > 0 {
			// fmt.Println("Console status 0xff")
			value = 0xff
		} else {
			value = 0x00
		}

	case 0x01:
		// Console Data
		value = <-m.Console

	case 0x02:
		// Printer Status
		value = 0x00
	case 0x03:
		// Printer Data
		fmt.Printf("not impl. I/O In from Printer Status\n")
	case 0x05:
		// Aux Data
		fmt.Printf("not impl. I/O In from Aux data\n")
	case 0x0a:
		value = uint8(m.fdc_drive)
	case 0x0b:
		value = uint8(m.fdc_track)
	case 0x0c:
		value = uint8(m.fdc_sector)
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
		m.fdc_drive = int(value)
	case 0x0b:
		m.fdc_track = int(value)
	case 0x0c:
		m.fdc_sector = int(value)
	case 0x0d:
		m.fdc_command = value
		sectors_per_track := 26
		if m.fdc_drive > 3 {
			sectors_per_track = 128
		}
		switch value {

		case 0: // Disk Read
			image := diskImage(m.fdc_drive)
			file, err := os.OpenFile(image, os.O_RDONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			dma := (int(m.fdc_dma_hi) << 8) | int(m.fdc_dma_low)
			sector := m.fdc_track*sectors_per_track + m.fdc_sector - 1
			offset := int(sector) * 128
			sector_data := m.Cpu.Memory[dma : dma+128]

			file.Seek(int64(offset), io.SeekStart)
			num_bytes, err := file.Read(sector_data)
			if err != nil {
				panic(err)
			}
			if num_bytes != 128 {
				panic(fmt.Sprintf("Bytes read is %d instead of 128\n", num_bytes))
			}

		case 1: // Disk Write
			image := diskImage(m.fdc_drive)
			file, err := os.OpenFile(image, os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			dma := (int(m.fdc_dma_hi) << 8) | int(m.fdc_dma_low)
			sector := m.fdc_track*sectors_per_track + m.fdc_sector - 1
			offset := int(sector) * 128
			sector_data := m.Cpu.Memory[dma : dma+128]

			file.Seek(int64(offset), io.SeekStart)
			file.Write(sector_data)

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

// LoadFile loads a section of a file into address.
func (m *Machine) LoadFile(name string, address int, offset int, length int) error {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	num_bytes := copy(m.Cpu.Memory[address:address+length], data[offset:(offset+length)])
	if num_bytes != length {
		panic(fmt.Sprintf("Loadfile copied %d bytes instead of %d\n", num_bytes, length))
	}
	return nil
}

func diskImage(disk int) string {
	letters := "abcdefghijklmnop"
	return fmt.Sprintf("disks/%c/DISK.IMG", letters[disk])
}

func keyboard(m *Machine) {
	var buffer = make([]byte, 1)
	for !m.Cpu.Halt {
		_, err := os.Stdin.Read(buffer)
		if err != nil {
			panic(err)
		}
		c := buffer[0]
		if c == LF {
			c = CR
		}
		m.Console <- c
	}
}

func main() {
	if LOGGING {
		file, err := os.OpenFile("./debug/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	m := NewMachine(0)
	go keyboard(m)

	err := m.LoadFile("disks/a/DISK.IMG", 0, 0x0000, 0x0080)
	if err != nil {
		log.Fatal(err)
	} else {
		m.Cpu.Run()
	}

}
