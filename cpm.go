package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/koron-go/z80"
)

// Start is an address where a program starts.
const Start = 0x0100

// Memory provides 64K bytes array memory.
type Memory struct {
	buf [65536]uint8
}

// NewMemory creates a new memory which includes minimal CP/M.
func NewMemory() *Memory {
	m := new(Memory)
	return m
}

func (m *Memory) Set(addr uint16, data uint8) {
	m.buf[addr] = data
}

func (m *Memory) Get(addr uint16) uint8 {
	return m.buf[addr]
}

// put puts "data" block from addr.
func (m *Memory) put(addr uint16, data ...uint8) {
	copy(m.buf[int(addr):int(addr)+len(data)], data)
}

// LoadFile loads a file from "Start" (0x0100) as program.
func (m *Memory) LoadFile(name string) error {
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

type IO struct {
	stdout io.Writer
	warnl  *log.Logger
}

func NewIO() *IO {
	return &IO{
		stdout: os.Stdout,
		warnl:  log.New(os.Stderr, "[WARN][IO]", 0),
	}
}

// In inputs a value from "addr" port.
// However this shows warning message always.
func (io *IO) In(addr uint8) uint8 {
	io.warnl.Printf("not impl. I/O In addr=0x%02x", addr)
	return 0
}

// Out outputs "value" to "addr" port.
// Only supports addr=0x0000, otherwise show warning message.
func (io *IO) Out(addr uint8, value uint8) {
	if addr != 0 {
		io.warnl.Printf("not impl. I/O Out addr=0x%02x value=0x%02x", addr, value)
		return
	}
	b := []byte{value}
	io.stdout.Write(b)
}

// SetStdout overrides stdout for this I/O.
func (io *IO) SetStdout(w io.Writer) {
	io.stdout = w
}

// SetWarnLogger overrides a warning logger.
func (io *IO) SetWarnLogger(l *log.Logger) {
	io.warnl = l
}

func main() {
	print("Hello\n")

	io := NewIO()
	mem := NewMemory()
	mem.LoadFile("a.bin")
	states := z80.States{SPR: z80.SPR{PC: 0x000}}
	cpu := z80.CPU{
		States: states,
		Memory: mem,
		IO:     io,
	}

	fmt.Printf("CPU=%v\n", cpu)

	err := cpu.Run(context.Background())

	fmt.Printf("Error %v\n", err)

}
