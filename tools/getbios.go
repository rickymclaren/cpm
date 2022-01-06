package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		println("Usage 'getbios <img file> <bios file>'")
		os.Exit(-1)

	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Cannot find '%v'\n", os.Args[1])
		os.Exit(-1)
	}

	// BIOS is in track 1 sectors 20 - 26
	sectorSize := 128
	start := (26 + 19) * sectorSize
	end := (26+26)*sectorSize - 1

	err = os.WriteFile(os.Args[2], data[start:end], 0644)
	if err != nil {
		fmt.Printf("Cannot write '%v'\n", os.Args[2])
		os.Exit(-1)

	}

}
