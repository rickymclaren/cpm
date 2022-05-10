# cpm

This is really an excuse to practice some Go.
I always wanted to build a Z80 computer that runs CPM. The next best thing is to create a Z80 emulator that can run CPM.

I have chosen to implement my own Z80 emulator but given how hard it is to figure out the undocumented workings of the Z80
I have cribbed many ideas from and CPM disks from :

* https://github.com/koron-go/z80
* https://github.com/udo-munk/z80pack

John Winans also has an excellent YouTube channel with videos on building a Z80 card to run CPM
* https://github.com/johnwinans
* https://www.youtube.com/c/JohnsBasement

Note: Although I describe this as a Z80 emulator it is only required to run CPM so ignores some
extended instructions to do with IO and interrupts since I only need the 8080 subset of instructions.

It passes ZEXDOC.
It also passes the tests for documented opcodes from the FUSE testpack.
The FUSE testpack was the best resource for testing individual opcodes.
