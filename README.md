# cpm

This is really an excuse to practice some Go.
I always wanted to build a Z80 computer that runs CP/M. The next best thing is to create a Z80 emulator that can run CP/M.

A top-level Makefile is provided to run the lower level make files in the correct order to build the disk images.
After that it is just a case of running "go run cpm".

This uses the boot and bios from z80pack so the disk layout is as per z80pack:
i.e. Four floppies a:, b:, c:, and d: plus two hard drives i:, and j:

I have cribbed many ideas from :

* https://github.com/koron-go/z80
* https://github.com/udo-munk/z80pack
* https://github.com/johnwinans
* https://www.youtube.com/c/JohnsBasement

Note: Although I describe this as a Z80 emulator it is only required to run CP/M so ignores some
extended instructions to do with IO and interrupts since I only need the 8080 subset of instructions.

It passes the tests for documented opcodes from the FUSE testpack.
