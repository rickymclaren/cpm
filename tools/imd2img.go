package main

import (
	"fmt"
	"os"
)

func process(data []byte) {
	i := 0
	c := data[i]
	for c != 0x1A {
		fmt.Printf("%c", data[i])
		i++
		c = data[i]
	}
	i++

	for i < len(data) {
		mode := data[i]
		i++
		cylinder := data[i]
		i++
		head := data[i]
		i++
		sectors := data[i]
		i++
		sectorSize := data[i]
		i++

		sectorSizeReal := 1 << (7 + sectorSize)

		sectorNumberingMap := data[i : i+int(sectors)]
		fmt.Printf("Mode %d, Cylinder %d, Head %d, Sectors %d, Size %d:%d, Map %v\n",
			mode, cylinder, head, sectors, sectorSize, sectorSizeReal, sectorNumberingMap)
		i += int(sectors)

		if head&0x80 != 0 {
			sectorCylinderMap := data[i : i+int(sectors)]
			fmt.Printf("  Sector Cylinder Map %v\n", sectorCylinderMap)
			i += int(sectors)

		}

		if head&0x40 != 0 {
			sectorHeadMap := data[i : i+int(sectors)]
			fmt.Printf("  Sector Head Map %v\n", sectorHeadMap)
			i += int(sectors)

		}

		for s := 0; s < int(sectors); s++ {
			sectorDataType := data[i]
			i++
			switch sectorDataType {
			case 0:
				fmt.Printf("X")
			case 1:
				fmt.Printf("N")
				i += int(sectorSizeReal)
			case 2:
				fmt.Printf("C")
				i++

			default:
				fmt.Printf("?%d", sectorDataType)
				os.Exit(-1)
			}
		}
		fmt.Println()

	}

}

func main() {
	if len(os.Args) != 2 {
		println("Usage 'imd2img <imd file>'")
		os.Exit(-1)

	}
	args := os.Args[1:]
	data, err := os.ReadFile(args[0])
	if err != nil {
		println("Cannot find '" + args[0] + "'")
		os.Exit(-1)
	}

	if data[0] != 'I' || data[1] != 'M' || data[2] != 'D' {
		println("Not an IMD file")
		os.Exit(-1)

	}
	process(data)

}
