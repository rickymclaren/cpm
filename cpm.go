package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/koron-go/z80"
)

// Start is an address where a program starts.
const Start = 0x0100

type Machine struct {
	Memory [65536]uint8
	Cpu    z80.CPU
}

func NewMachine() *Machine {
	m := new(Machine)
	m.Cpu = z80.CPU{
		States: z80.States{SPR: z80.SPR{PC: 0x000}},
		Memory: m,
		IO:     m,
	}
	return m
}

func (m *Machine) In(addr uint8) uint8 {
	fmt.Printf("not impl. I/O In addr=0x%02x", addr)
	return 0
}

func (m *Machine) Out(addr uint8, value uint8) {
	fmt.Printf("%c", value)
}

func (m *Machine) Set(addr uint16, data uint8) {
	m.Memory[addr] = data
}

func (m *Machine) Get(addr uint16) uint8 {
	return m.Memory[addr]
}

// put puts "data" block from addr.
func (m *Machine) put(addr uint16, data ...uint8) {
	copy(m.Memory[int(addr):int(addr)+len(data)], data)
}

// LoadFile loads a file from "Start" (0x0100) as program.
func (m *Machine) LoadFile(name string) error {
	prog, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	if len(prog) > 65535 {
		m.put(0, prog[0:65535]...)
	} else {
		m.put(0, prog...)

	}
	return nil
}

func main() {
	print("Hello\n")

	m := NewMachine()
	m.LoadFile("a.bin")

	fmt.Printf("CPU=%v\n", m.Cpu)

	err := m.Cpu.Run(context.Background())

	fmt.Printf("Error %v\n", err)

}
