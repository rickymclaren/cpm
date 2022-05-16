ASM_FLAGS=-I../lib
DATE := $(date -R)
GIT_VERSION := $(git describe --long --dirty; git show -s --format='%ci')

all: system.bin

system.bin: cpm22.asm bios.asm
	z80asm -i system.asm -o $@ --list=system.lst --label=system.sym $(ASM_FLAGS)

cpm22.asm: ../cpm22/CPM22.Z80
	expand $^ | sed -e 's/\(SECTRN\)/SECTRAN/' -e 's/\(^MEM \)/;\1/' -e 's/,M/,(HL)/g' -e 's/^\([[:alpha:]]*:[[:space:]]*JP[[:space:]]*0\)/;\1/' -e 's/^\([[:alpha:]][^: ]*\)\([[:space:]]\)/\1:\2/' | sed -e "s/@@DATE@@/$(DATE)/g" -e "s/@@GIT_VERSION@@/$(GIT_VERSION)/g"  > $@

bios.asm: ../z80pack/bios.asm
	expand $^ | sed -e 's/\(^IOBYTE \)/;\1/' -e 's/,M/,(HL)/g' -e 's/^\([[:alpha:]]*:[[:space:]]*JP[[:space:]]*0\)/;\1/' -e 's/^\([[:alpha:]][^: ]*\)\([[:space:]]\)/\1:\2/' | sed -e "s/@@DATE@@/$(DATE)/g" -e "s/@@GIT_VERSION@@/$(GIT_VERSION)/g"  > $@

clean:
	rm -f *.lst *.bin *.hex *.sym cpm22.asm bios.asm



