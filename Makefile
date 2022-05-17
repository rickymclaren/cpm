all: make_all clean

make_all:
	make -C cpm22
	make -C system
	make -C disks

clean:
	make -C cpm22 clean
	make -C system clean
	make -C disks clean
