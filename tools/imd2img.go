package main

import (
	"fmt"
	"os"
)

func process(data []byte) []byte {

	if data[0] != 'I' || data[1] != 'M' || data[2] != 'D' {
		println("Not an IMD file")
		os.Exit(-1)

	}

	imgData := make([]byte, 0, 80*26*128)
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

		sectorSizeReal := int(1 << (7 + sectorSize))

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
				for x := 0; x < sectorSizeReal; x++ {
					imgData = append(imgData, 0)
				}
			case 1:
				fmt.Printf("N")
				imgData = append(imgData, data[i:i+sectorSizeReal]...)
				i += int(sectorSizeReal)
			case 2:
				fmt.Printf("C")
				for x := 0; x < sectorSizeReal; x++ {
					imgData = append(imgData, data[i])
				}
				i++

			default:
				fmt.Printf("?%d", sectorDataType)
				os.Exit(-1)
			}
		}
		fmt.Println()

	}

	return imgData

}

func main() {
	if len(os.Args) != 3 {
		println("Usage 'imd2img <imd file> <img file>'")
		os.Exit(-1)

	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Cannot find '%v'\n", os.Args[1])
		os.Exit(-1)
	}

	imgData := process(data)

	fmt.Printf("Image Data = %v bytes\n", len(imgData))

	err = os.WriteFile(os.Args[2], imgData, 0644)
	if err != nil {
		fmt.Printf("Cannot write '%v'\n", os.Args[2])
		os.Exit(-1)

	}

}
