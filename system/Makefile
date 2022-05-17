all: system.bin boot.bin

system.bin: boot.asm cpm22.asm bios.asm
	z80asm -i system.asm -o $@ --list=system.lst --label=system.sym

cpm22.asm: ../cpm22/CPM22.Z80
	expand $^ |\
	sed -e 's/\(SECTRN\)/SECTRAN/' \
	    -e 's/\(^MEM \)/;\1/' -e 's/,M/,(HL)/g' \
		-e 's/^\([[:alpha:]]*:[[:space:]]*JP[[:space:]]*0\)/;\1/' \
		-e 's/^\([[:alpha:]][^: ]*\)\([[:space:]]\)/\1:\2/' \
	> $@

bios.asm: ../z80pack/bios.asm
	expand $^ |\
	sed -e 's/\(^IOBYTE \)/;\1/' \
	    -e 's/,M/,(HL)/g' \
		-e 's/^\([[:alpha:]]*:[[:space:]]*JP[[:space:]]*0\)/;\1/' \
		-e 's/^\([[:alpha:]][^: ]*\)\([[:space:]]\)/\1:\2/' \
	> $@

boot.bin: boot.asm
	z80asm -i boot.asm -o $@ --list=boot.lst --label=boot.sym $(ASM_FLAGS)

boot.asm: ../z80pack/boot.asm
	expand $^ | \
	sed -e 's/,M/,(HL)/g' \
	    -e 's/^\([[:alpha:]]*:[[:space:]]*JP[[:space:]]*0\)/;\1/' \
		-e 's/^\([[:alpha:]][^: ]*\)\([[:space:]]\)/\1:\2/' \
	> $@

clean:
	rm -f *.lst *.bin *.hex *.sym cpm22.asm bios.asm boot.asm



