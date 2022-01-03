  org 0x0000

main: ld a, 0
loop: out(0x10), a
      inc a
      jp loop
