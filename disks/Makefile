# Build a filesystem from the cpm22 distro files, retro.asm BIOS, and some
# sample downloads like adventure.

all: disk_a disks_bcd disk_i disk_j

disk_a: ../cpm22/filesystem/*
	rm -f DISK.IMG
	mkfs.cpm -b ../system/boot.bin -b ../system/system.bin DISK.IMG
	cpmcp -T raw DISK.IMG ../cpm22/filesystem/* 0:
	cp DISK.IMG a/DISK.IMG

disks_bcd: 
	rm -f DISK.IMG
	mkfs.cpm DISK.IMG
	cp DISK.IMG b/DISK.IMG
	cp DISK.IMG c/DISK.IMG
	cp DISK.IMG d/DISK.IMG

disk_i: ../cpm22/filesystem/* adventure mbasic sargon zork startrek
	rm -f DISK.IMG
	mkfs.cpm -f z80pack-hd DISK.IMG
	cpmcp -f z80pack-hd -T raw DISK.IMG ../cpm22/filesystem/* 0:
	cpmcp -f z80pack-hd -T raw DISK.IMG adventure/adv-A02/* 0:
	cpmcp -f z80pack-hd -T raw DISK.IMG mbasic/mbasic.com 0:
	cpmcp -f z80pack-hd -T raw DISK.IMG sargon/sargon.com 0:
	cpmcp -f z80pack-hd -T raw DISK.IMG zork/ZORK*.COM 0:
	cpmcp -f z80pack-hd -T raw DISK.IMG zork/ZORK*.DAT 0:
	cpmcp -f z80pack-hd -T raw DISK.IMG startrek/startrek.bas 0:
	cp DISK.IMG i/DISK.IMG

disk_j: 
	rm -f DISK.IMG
	mkfs.cpm -f z80pack-hd DISK.IMG
	cpmcp -f z80pack-hd -T raw DISK.IMG ../cpm22/filesystem/* 0:
	cp DISK.IMG j/DISK.IMG

adventure:
	wget -P adventure https://ifarchive.org/if-archive/games/cpm/cpm-advent.zip
	unzip -d adventure adventure/cpm-advent.zip

mbasic:
	wget -P mbasic http://www.retroarchive.org/cpm/lang/mbasic.zip
	unzip -d mbasic mbasic/mbasic.zip

sargon:
	wget -P sargon http://www.retroarchive.org/cpm/games/sargon.zip
	unzip -d sargon sargon/sargon.zip

zork:
	wget -P zork http://www.retroarchive.org/cpm/games/zork123_80.zip
	unzip -d zork zork/zork123_80.zip

startrek:
	wget -P startrek "https://raw.githubusercontent.com/RC2014Z80/RC2014/master/BASIC-Programs/Super Startrek/startrek.bas"

clean:
	rm -rf adventure
	rm -rf mbasic
	rm -rf sargon
	rm -rf zork
	rm -rf startrek
	rm -f DISK.IMG

lsimg:
	cpmls -f z80pack-hd -T raw i/DISK.IMG
