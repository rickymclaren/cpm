# Build a filesystem from the cpm22 distro files, retro.asm BIOS, and some
# sample downloads like adventure.

all: disk_i_img

disk_i_img: ../cpm22/filesystem/* adventure/adv-B03
	rm -f DISK.IMG
	mkfs.cpm -f z80pack-hd DISK.IMG
	cpmcp -f z80pack-hd -T raw DISK.IMG ../cpm22/filesystem/* 0:
	cpmcp -f z80pack-hd DISK.IMG adventure/adv-A02/* 0:
	cp DISK.IMG i/DISK.IMG

adventure/adv-B03: adventure/cpm-advent.zip
	unzip -d adventure adventure/cpm-advent.zip

adventure/cpm-advent.zip:
	wget -P adventure https://ifarchive.org/if-archive/games/cpm/cpm-advent.zip

clean:
	rm -rf `find adventure/* -name \*.zip -o -print`
	rm -f DISK.IMG

lsimg:
	cpmls -f z80pack-hd -T raw i/DISK.IMG
