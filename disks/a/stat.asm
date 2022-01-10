; z80dasm 1.1.6
; command line: z80dasm -al -b stat.blk -o stat.asm stat.com

	org	00100h


; BLOCK 'start' (start 0x0100 end 0x0104)
start_start:
	jp data1_end		;0100
	defb 020h		;0103
start_end:

; BLOCK 'data1' (start 0x0104 end 0x0433)
data1_start:
	defw 02020h		;0104
	defw 06f43h		;0106
	defw 07970h		;0108
	defw 06972h		;010a
	defw 06867h		;010c
	defw 02074h		;010e
	defw 06328h		;0110
	defw 02029h		;0112
	defw 03931h		;0114
	defw 03937h		;0116
	defw 0202ch		;0118
	defw 06944h		;011a
	defw 06967h		;011c
	defw 06174h		;011e
	defw 0206ch		;0120
	defw 06552h		;0122
	defw 06573h		;0124
	defw 07261h		;0126
	defw 06863h		;0128
l012ah:
	defw 03f3fh		;012a
	defw 03f3fh		;012c
	defw 03f3fh		;012e
	defw 03f3fh		;0130
	defw 03f3fh		;0132
	defw 03f3fh		;0134
	defw 00000h		;0136
	defw 04300h		;0138
	defw 04e4fh		;013a
	defw 0523ah		;013c
	defw 05244h		;013e
	defw 0503ah		;0140
	defw 04e55h		;0142
	defw 04c3ah		;0144
	defw 05453h		;0146
	defw 0443ah		;0148
	defw 05645h		;014a
	defw 0563ah		;014c
	defw 04c41h		;014e
	defw 0553ah		;0150
	defw 05253h		;0152
	defw 0443ah		;0154
	defw 04b53h		;0156
	defw 0543ah		;0158
	defw 05954h		;015a
	defw 0433ah		;015c
	defw 05452h		;015e
	defw 0423ah		;0160
	defw 05441h		;0162
	defw 0553ah		;0164
	defw 03143h		;0166
	defw 0543ah		;0168
	defw 05954h		;016a
	defw 0503ah		;016c
	defw 05254h		;016e
	defw 0553ah		;0170
	defw 03152h		;0172
	defw 0553ah		;0174
	defw 03252h		;0176
	defw 0543ah		;0178
	defw 05954h		;017a
	defw 0503ah		;017c
	defw 05054h		;017e
	defw 0553ah		;0180
	defw 03150h		;0182
	defw 0553ah		;0184
	defw 03250h		;0186
	defw 0543ah		;0188
	defw 05954h		;018a
	defw 0433ah		;018c
	defw 05452h		;018e
	defw 04c3ah		;0190
	defw 05450h		;0192
	defw 0553ah		;0194
	defw 0314ch		;0196
	defw 0523ah		;0198
	defw 04f2fh		;019a
	defw 05200h		;019c
	defw 0572fh		;019e
	defw 05300h		;01a0
	defw 05359h		;01a2
	defw 04400h		;01a4
	defw 05249h		;01a6
	defw 05200h		;01a8
	defw 04f2fh		;01aa
	defw 05220h		;01ac
	defw 0572fh		;01ae
	defw 05320h		;01b0
	defw 05359h		;01b2
	defw 04420h		;01b4
	defw 05249h		;01b6
	defw 02a20h		;01b8
	defw 0202ah		;01ba
	defw 06241h		;01bc
	defw 0726fh		;01be
	defw 06574h		;01c0
	defw 02064h		;01c2
	defw 02a2ah		;01c4
	defw 04100h		;01c6
	defw 07463h		;01c8
	defw 07669h		;01ca
	defw 02065h		;01cc
	defw 07355h		;01ce
	defw 07265h		;01d0
	defw 03a20h		;01d2
	defw 04100h		;01d4
	defw 07463h		;01d6
	defw 07669h		;01d8
	defw 02065h		;01da
	defw 06946h		;01dc
	defw 0656ch		;01de
	defw 03a73h		;01e0
	defw 02000h		;01e2
	defw 02020h		;01e4
	defw 00020h		;01e6
l01e8h:
	defw 04420h		;01e8
	defw 06972h		;01ea
	defw 06576h		;01ec
	defw 04320h		;01ee
	defw 06168h		;01f0
	defw 06172h		;01f2
	defw 07463h		;01f4
	defw 07265h		;01f6
	defw 07369h		;01f8
	defw 06974h		;01fa
	defw 07363h		;01fc
	defw 03600h		;01fe
l0200h:
	defw 03535h		;0200
	defw 03633h		;0202
	defw 0203ah		;0204
	defw 03100h		;0206
	defw 03832h		;0208
	defw 04220h		;020a
	defw 07479h		;020c
	defw 02065h		;020e
	defw 06552h		;0210
	defw 06f63h		;0212
	defw 06472h		;0214
	defw 04320h		;0216
	defw 07061h		;0218
	defw 06361h		;021a
	defw 07469h		;021c
	defw 00079h		;021e
l0220h:
	defw 0694bh		;0220
	defw 06f6ch		;0222
	defw 07962h		;0224
	defw 06574h		;0226
	defw 04420h		;0228
	defw 06972h		;022a
	defw 06576h		;022c
	defw 02020h		;022e
	defw 06143h		;0230
	defw 06170h		;0232
	defw 06963h		;0234
	defw 07974h		;0236
	defw 03300h		;0238
	defw 02032h		;023a
	defw 04220h		;023c
	defw 07479h		;023e
	defw 02065h		;0240
	defw 06944h		;0242
	defw 06572h		;0244
	defw 07463h		;0246
	defw 0726fh		;0248
	defw 02079h		;024a
	defw 06e45h		;024c
	defw 07274h		;024e
	defw 06569h		;0250
	defw 00073h		;0252
l0254h:
	defw 06843h		;0254
	defw 06365h		;0256
	defw 0656bh		;0258
	defw 02064h		;025a
	defw 04420h		;025c
	defw 07269h		;025e
	defw 06365h		;0260
	defw 06f74h		;0262
	defw 07972h		;0264
	defw 04520h		;0266
	defw 0746eh		;0268
	defw 06972h		;026a
	defw 07365h		;026c
	defw 05200h		;026e
	defw 06365h		;0270
	defw 0726fh		;0272
	defw 07364h		;0274
	defw 0202fh		;0276
	defw 07845h		;0278
	defw 06574h		;027a
	defw 0746eh		;027c
	defw 05200h		;027e
	defw 06365h		;0280
	defw 0726fh		;0282
	defw 07364h		;0284
	defw 0202fh		;0286
	defw 06c42h		;0288
	defw 0636fh		;028a
	defw 0006bh		;028c
l028eh:
	defw 06553h		;028e
	defw 07463h		;0290
	defw 0726fh		;0292
	defw 02f73h		;0294
	defw 05420h		;0296
	defw 06172h		;0298
	defw 06b63h		;029a
	defw 05200h		;029c
	defw 07365h		;029e
	defw 07265h		;02a0
	defw 06576h		;02a2
	defw 02064h		;02a4
	defw 07254h		;02a6
	defw 06361h		;02a8
	defw 0736bh		;02aa
	defw 02000h		;02ac
	defw 07369h		;02ae
	defw 00020h		;02b0
l02b2h:
	defw 06554h		;02b2
	defw 0706dh		;02b4
	defw 05220h		;02b6
	defw 04f2fh		;02b8
	defw 04420h		;02ba
	defw 07369h		;02bc
	defw 03a6bh		;02be
	defw 06420h		;02c0
	defw 03d3ah		;02c2
	defw 02f52h		;02c4
	defw 0004fh		;02c6
l02c8h:
	defw 06553h		;02c8
	defw 02074h		;02ca
	defw 06e49h		;02cc
	defw 06964h		;02ce
	defw 06163h		;02d0
	defw 06f74h		;02d2
	defw 03a72h		;02d4
	defw 06420h		;02d6
	defw 0663ah		;02d8
	defw 06c69h		;02da
	defw 06e65h		;02dc
	defw 06d61h		;02de
	defw 02e65h		;02e0
	defw 07974h		;02e2
	defw 02070h		;02e4
	defw 05224h		;02e6
	defw 04f2fh		;02e8
	defw 02420h		;02ea
	defw 02f52h		;02ec
	defw 02057h		;02ee
	defw 05324h		;02f0
	defw 05359h		;02f2
	defw 02420h		;02f4
	defw 04944h		;02f6
	defw 00052h		;02f8
l02fah:
	defw 06944h		;02fa
	defw 06b73h		;02fc
	defw 05320h		;02fe
	defw 06174h		;0300
	defw 07574h		;0302
	defw 02073h		;0304
	defw 03a20h		;0306
	defw 04420h		;0308
	defw 04b53h		;030a
	defw 0203ah		;030c
	defw 03a64h		;030e
	defw 05344h		;0310
	defw 03a4bh		;0312
	defw 05500h		;0314
	defw 06573h		;0316
	defw 02072h		;0318
	defw 07453h		;031a
	defw 07461h		;031c
	defw 07375h		;031e
	defw 02020h		;0320
	defw 0203ah		;0322
	defw 05355h		;0324
	defw 03a52h		;0326
	defw 04900h		;0328
	defw 0626fh		;032a
	defw 07479h		;032c
	defw 02065h		;032e
	defw 07341h		;0330
	defw 06973h		;0332
	defw 06e67h		;0334
	defw 0003ah		;0336
l0338h:
	defw 03d20h		;0338
	defw 04200h		;033a
	defw 06461h		;033c
	defw 04420h		;033e
	defw 06c65h		;0340
	defw 06d69h		;0342
	defw 07469h		;0344
	defw 07265h		;0346
	defw 04900h		;0348
	defw 0766eh		;034a
	defw 06c61h		;034c
	defw 06469h		;034e
	defw 04120h		;0350
	defw 07373h		;0352
	defw 06769h		;0354
	defw 06d6eh		;0356
	defw 06e65h		;0358
	defw 00074h		;035a
l035ch:
	defw 06142h		;035c
	defw 02064h		;035e
	defw 06544h		;0360
	defw 0696ch		;0362
	defw 0696dh		;0364
	defw 06574h		;0366
	defw 00072h		;0368
l036ah:
	defw 0203ah		;036a
	defw 04200h		;036c
	defw 07479h		;036e
	defw 07365h		;0370
	defw 05220h		;0372
	defw 06d65h		;0374
	defw 06961h		;0376
	defw 0696eh		;0378
	defw 0676eh		;037a
	defw 04f20h		;037c
	defw 0206eh		;037e
	defw 05200h		;0380
	defw 0002fh		;0382
l0384h:
	defw 0202ch		;0384
	defw 07053h		;0386
	defw 06361h		;0388
	defw 03a65h		;038a
	defw 00020h		;038c
l038eh:
	defw 06e49h		;038e
	defw 06176h		;0390
	defw 0696ch		;0392
	defw 02064h		;0394
	defw 06946h		;0396
	defw 0656ch		;0398
	defw 04920h		;039a
	defw 0646eh		;039c
	defw 06369h		;039e
	defw 07461h		;03a0
	defw 0726fh		;03a2
	defw 02a00h		;03a4
	defw 0202ah		;03a6
	defw 06f54h		;03a8
	defw 0206fh		;03aa
	defw 0614dh		;03ac
	defw 0796eh		;03ae
	defw 04620h		;03b0
	defw 06c69h		;03b2
	defw 07365h		;03b4
	defw 02a20h		;03b6
	defw 0002ah		;03b8
l03bah:
	defw 06946h		;03ba
	defw 0656ch		;03bc
	defw 04e20h		;03be
	defw 0746fh		;03c0
	defw 04620h		;03c2
	defw 0756fh		;03c4
	defw 0646eh		;03c6
	defw 02000h		;03c8
	defw 06953h		;03ca
	defw 0657ah		;03cc
	defw 00020h		;03ce
l03d0h:
	defw 05220h		;03d0
	defw 06365h		;03d2
	defw 02073h		;03d4
	defw 04220h		;03d6
	defw 07479h		;03d8
	defw 07365h		;03da
	defw 02020h		;03dc
	defw 07845h		;03de
	defw 02074h		;03e0
	defw 06341h		;03e2
	defw 00063h		;03e4
l03e6h:
	defw 03536h		;03e6
l03e8h:
	defw 03335h		;03e8
	defw 00036h		;03ea
l03ech:
	defw 07320h		;03ec
	defw 07465h		;03ee
	defw 07420h		;03f0
	defw 0206fh		;03f2
	defw 05200h		;03f4
	defw 04f2fh		;03f6
	defw 04920h		;03f8
	defw 0766eh		;03fa
	defw 06c61h		;03fc
	defw 06469h		;03fe
l0400h:
	defw 04420h		;0400
	defw 07369h		;0402
	defw 0206bh		;0404
	defw 07341h		;0406
	defw 06973h		;0408
	defw 06e67h		;040a
	defw 0656dh		;040c
	defw 0746eh		;040e
	defw 05700h		;0410
	defw 06f72h		;0412
	defw 0676eh		;0414
	defw 04320h		;0416
	defw 02f50h		;0418
	defw 0204dh		;041a
	defw 06556h		;041c
	defw 07372h		;041e
	defw 06f69h		;0420
	defw 0206eh		;0422
	defw 05228h		;0424
	defw 07165h		;0426
	defw 06975h		;0428
	defw 06572h		;042a
	defw 02073h		;042c
	defw 02e32h		;042e
	defw 02930h		;0430
	defb 000h		;0432
data1_end:

; BLOCK 'code1' (start 0x0433 end 0x1300)
code1_start:
	ld hl,00000h		;0433
	add hl,sp			;0436
	ld (0156bh),hl		;0437
	ld hl,0158dh		;043a
	ld sp,hl			;043d
	call sub_04edh		;043e
	cp 020h		;0441
	jp nc,l044fh		;0443
	ld bc,00411h		;0446
	call sub_04d2h		;0449
	jp l048bh		;044c
l044fh:
	ld hl,01596h		;044f
	ld (hl),001h		;0452
	ld a,(0005ch)		;0454
	sub 000h		;0457
	sub 001h		;0459
	sbc a,a			;045b
	push af			;045c
	ld a,(0005dh)		;045d
	sub 020h		;0460
	sub 001h		;0462
	sbc a,a			;0464
	pop bc			;0465
	ld c,b			;0466
	and c			;0467
	rra			;0468
	jp nc,l0472h		;0469
	call sub_0d33h		;046c
	jp l048bh		;046f
l0472h:
	ld a,(0005ch)		;0472
	cp 000h		;0475
	jp z,l0480h		;0477
	call sub_1402h		;047a
	jp l048bh		;047d
l0480h:
	call sub_0a87h		;0480
	cpl			;0483
	rra			;0484
	jp nc,l048bh		;0485
	call sub_0dbah		;0488
l048bh:
	ld hl,(0156bh)		;048b
	ld sp,hl			;048e
	ret			;048f
sub_0490h:
	ld hl,0155bh		;0490
	ld (hl),c			;0493
	ld hl,(0155bh)		;0494
	ld h,000h		;0497
	ex de,hl			;0499
	ld c,002h		;049a
	call 00005h		;049c
	ret			;049f
sub_04a0h:
	ld c,00dh		;04a0
	call sub_0490h		;04a2
	ld c,00ah		;04a5
	call sub_0490h		;04a7
	ret			;04aa
sub_04abh:
	ld c,020h		;04ab
	call sub_0490h		;04ad
	ret			;04b0
sub_04b1h:
	ld hl,0155dh		;04b1
	ld (hl),b			;04b4
	dec hl			;04b5
	ld (hl),c			;04b6
l04b7h:
	ld hl,(0155ch)		;04b7
	ld a,(hl)			;04ba
	cp 000h		;04bb
	jp z,l04d1h		;04bd
	ld hl,(0155ch)		;04c0
	ld c,(hl)			;04c3
	call sub_0490h		;04c4
	ld hl,(0155ch)		;04c7
	inc hl			;04ca
	ld (0155ch),hl		;04cb
	jp l04b7h		;04ce
l04d1h:
	ret			;04d1
sub_04d2h:
	ld hl,0155fh		;04d2
	ld (hl),b			;04d5
	dec hl			;04d6
	ld (hl),c			;04d7
	call sub_04a0h		;04d8
	ld hl,(0155eh)		;04db
	ld b,h			;04de
	ld c,l			;04df
	call sub_04b1h		;04e0
	ret			;04e3
sub_04e4h:
	ld de,00000h		;04e4
	ld c,00bh		;04e7
	call 00005h		;04e9
	ret			;04ec
sub_04edh:
	ld de,00000h		;04ed
	ld c,00ch		;04f0
	call 00005h		;04f2
	ret			;04f5
sub_04f6h:
	ld hl,01561h		;04f6
	ld (hl),c			;04f9
	ld hl,(01561h)		;04fa
	ld h,000h		;04fd
	ex de,hl			;04ff
	ld c,00eh		;0500
	call 00005h		;0502
	ret			;0505
	ld hl,01563h		;0506
	ld (hl),b			;0509
	dec hl			;050a
	ld (hl),c			;050b
	ld hl,(01562h)		;050c
	ex de,hl			;050f
	ld c,00fh		;0510
	call 00005h		;0512
	ld (01560h),a		;0515
	ret			;0518
sub_0519h:
	ld hl,01565h		;0519
	ld (hl),b			;051c
	dec hl			;051d
	ld (hl),c			;051e
	ld hl,(01564h)		;051f
	ex de,hl			;0522
	ld c,011h		;0523
	call 00005h		;0525
	ld (01560h),a		;0528
	ret			;052b
sub_052ch:
	ld de,00000h		;052c
	ld c,012h		;052f
	call 00005h		;0531
	ld (01560h),a		;0534
	ret			;0537
sub_0538h:
	ld de,00000h		;0538
	ld c,019h		;053b
	call 00005h		;053d
	ret			;0540
sub_0541h:
	ld hl,01567h		;0541
	ld (hl),b			;0544
	dec hl			;0545
	ld (hl),c			;0546
	ld hl,(01566h)		;0547
	ex de,hl			;054a
	ld c,01ah		;054b
	call 00005h		;054d
	ret			;0550
sub_0551h:
	ld de,00000h		;0551
	ld c,01bh		;0554
	call 00005h		;0556
	ret			;0559
sub_055ah:
	ld de,00000h		;055a
	ld c,018h		;055d
	call 00005h		;055f
	ret			;0562
sub_0563h:
	ld de,00000h		;0563
	ld c,01ch		;0566
	call 00005h		;0568
	ret			;056b
sub_056ch:
	ld de,00000h		;056c
	ld c,01dh		;056f
	call 00005h		;0571
	ret			;0574
sub_0575h:
	ld de,0005ch		;0575
	ld c,01eh		;0578
	call 00005h		;057a
	ret			;057d
sub_057eh:
	ld de,00000h		;057e
	ld c,01fh		;0581
	call 00005h		;0583
	ld (01557h),hl		;0586
	ret			;0589
sub_058ah:
	ld de,000ffh		;058a
	ld c,020h		;058d
	call 00005h		;058f
	ret			;0592
	ld hl,01568h		;0593
	ld (hl),c			;0596
	ld hl,(01568h)		;0597
	ld h,000h		;059a
	ex de,hl			;059c
	ld c,020h		;059d
	call 00005h		;059f
	ret			;05a2
sub_05a3h:
	ld hl,0156ah		;05a3
	ld (hl),b			;05a6
	dec hl			;05a7
	ld (hl),c			;05a8
	ld hl,(01569h)		;05a9
	ex de,hl			;05ac
	ld c,023h		;05ad
	call 00005h		;05af
	ret			;05b2
sub_05b3h:
	call sub_057eh		;05b3
	ld hl,(01557h)		;05b6
	inc hl			;05b9
	inc hl			;05ba
	ld c,(hl)			;05bb
	ld hl,00001h		;05bc
	call sub_14aeh		;05bf
	ld de,00080h		;05c2
	call sub_148fh		;05c5
	ld (0158dh),hl		;05c8
	ret			;05cb
sub_05cch:
	ld hl,0158fh		;05cc
	ld (hl),c			;05cf
	ld hl,(0158fh)		;05d0
	ld c,l			;05d3
	call sub_04f6h		;05d4
	call sub_05b3h		;05d7
	ret			;05da
sub_05dbh:
	ld hl,01591h		;05db
	ld (hl),b			;05de
	dec hl			;05df
	ld (hl),c			;05e0
	ld c,003h		;05e1
	ld hl,01590h		;05e3
	call sub_14b4h		;05e6
	ex de,hl			;05e9
	ld hl,(01559h)		;05ea
	add hl,de			;05ed
	ld a,007h		;05ee
	ld de,01590h		;05f0
	push hl			;05f3
	call sub_1463h		;05f4
	inc hl			;05f7
	ld c,l			;05f8
	pop hl			;05f9
	call sub_14a3h		;05fa
	ret			;05fd
sub_05feh:
	ld hl,01598h		;05fe
	ld (hl),b			;0601
	dec hl			;0602
	ld (hl),c			;0603
	ld hl,01599h		;0604
	ld (hl),000h		;0607
l0609h:
	ld a,003h		;0609
	ld hl,01599h		;060b
	cp (hl)			;060e
	jp c,l0636h		;060f
	ld hl,(01599h)		;0612
	ld h,000h		;0615
	ex de,hl			;0617
	ld hl,(01597h)		;0618
	add hl,de			;061b
	push hl			;061c
	ld hl,(01599h)		;061d
	ld h,000h		;0620
	ld bc,01592h		;0622
	add hl,bc			;0625
	pop de			;0626
	ld a,(de)			;0627
	cp (hl)			;0628
	jp z,l062fh		;0629
	ld a,000h		;062c
	ret			;062e
l062fh:
	ld hl,01599h		;062f
	inc (hl)			;0632
	jp nz,l0609h		;0633
l0636h:
	ld a,001h		;0636
	ret			;0638
l0639h:
	ld hl,(01596h)		;0639
	ld h,000h		;063c
	ld bc,00080h		;063e
	add hl,bc			;0641
	ld a,(hl)			;0642
	cp 020h		;0643
	jp nz,l064fh		;0645
	ld hl,01596h		;0648
	inc (hl)			;064b
	jp l0639h		;064c
l064fh:
	ld hl,0159ah		;064f
	ld (hl),000h		;0652
l0654h:
	ld a,(0159ah)		;0654
	cp 004h		;0657
	jp nc,l06e6h		;0659
	ld hl,(01596h)		;065c
	ld h,000h		;065f
	ld bc,00080h		;0661
	add hl,bc			;0664
	ld a,(hl)			;0665
	ld (0159bh),a		;0666
	ld c,a			;0669
	ld a,001h		;066a
	cp c			;066c
	jp nc,l067ah		;066d
	ld hl,(0159bh)		;0670
	ld c,l			;0673
	call sub_06ebh		;0674
	jp l067fh		;0677
l067ah:
	ld c,020h		;067a
	call sub_06ebh		;067c
l067fh:
	ld a,001h		;067f
	ld hl,0159bh		;0681
	sub (hl)			;0684
	sbc a,a			;0685
	cpl			;0686
	push af			;0687
	ld a,(hl)			;0688
	sub 02ch		;0689
	sub 001h		;068b
	sbc a,a			;068d
	pop bc			;068e
	ld c,b			;068f
	or c			;0690
	push af			;0691
	ld a,(hl)			;0692
	sub 03ah		;0693
	sub 001h		;0695
	sbc a,a			;0697
	pop bc			;0698
	ld c,b			;0699
	or c			;069a
	push af			;069b
	ld a,(hl)			;069c
	sub 02ah		;069d
	sub 001h		;069f
	sbc a,a			;06a1
	pop bc			;06a2
	ld c,b			;06a3
	or c			;06a4
	push af			;06a5
	ld a,(hl)			;06a6
	sub 02eh		;06a7
	sub 001h		;06a9
	sbc a,a			;06ab
	pop bc			;06ac
	ld c,b			;06ad
	or c			;06ae
	push af			;06af
	ld a,(hl)			;06b0
	sub 03eh		;06b1
	sub 001h		;06b3
	sbc a,a			;06b5
	pop bc			;06b6
	ld c,b			;06b7
	or c			;06b8
	push af			;06b9
	ld a,(hl)			;06ba
	sub 03ch		;06bb
	sub 001h		;06bd
	sbc a,a			;06bf
	pop bc			;06c0
	ld c,b			;06c1
	or c			;06c2
	push af			;06c3
	ld a,(hl)			;06c4
	sub 03dh		;06c5
	sub 001h		;06c7
	sbc a,a			;06c9
	pop bc			;06ca
	ld c,b			;06cb
	or c			;06cc
	rra			;06cd
	jp nc,l06dfh		;06ce
	ld hl,(01596h)		;06d1
	ld h,000h		;06d4
	ld bc,00080h		;06d6
	add hl,bc			;06d9
	ld (hl),001h		;06da
	jp l06e3h		;06dc
l06dfh:
	ld hl,01596h		;06df
	inc (hl)			;06e2
l06e3h:
	jp l0654h		;06e3
l06e6h:
	ld hl,01596h		;06e6
	inc (hl)			;06e9
	ret			;06ea
sub_06ebh:
	ld hl,0159ch		;06eb
	ld (hl),c			;06ee
	ld hl,(0159ah)		;06ef
	ld h,000h		;06f2
	ld bc,01592h		;06f4
	add hl,bc			;06f7
	ld a,(0159ch)		;06f8
	ld (hl),a			;06fb
	ld hl,0159ah		;06fc
	inc (hl)			;06ff
	ret			;0700
sub_0701h:
	ld hl,015a0h		;0701
	ld (hl),d			;0704
	dec hl			;0705
	ld (hl),e			;0706
	dec hl			;0707
	ld (hl),b			;0708
	dec hl			;0709
	ld (hl),c			;070a
	ld hl,015a1h		;070b
	ld (hl),001h		;070e
l0710h:
	ld a,000h		;0710
	ld de,0159fh		;0712
	call sub_14dbh		;0715
	or l			;0718
	jp z,l0772h		;0719
	ld hl,(0159dh)		;071c
	ex de,hl			;071f
	ld hl,(0159fh)		;0720
	call sub_1470h		;0723
	ld hl,015a2h		;0726
	ld (hl),e			;0729
	ld hl,(0159dh)		;072a
	ex de,hl			;072d
	call sub_1472h		;072e
	ld (0159dh),hl		;0731
	ld d,b			;0734
	ld e,c			;0735
	ld hl,0000ah		;0736
	call sub_1470h		;0739
	ex de,hl			;073c
	ld (0159fh),hl		;073d
	ld a,000h		;0740
	call sub_14c4h		;0742
	or l			;0745
	add a,0ffh		;0746
	sbc a,a			;0748
	ld hl,015a1h		;0749
	and (hl)			;074c
	inc hl			;074d
	push af			;074e
	ld a,(hl)			;074f
	sub e			;0750
	sub 001h		;0751
	sbc a,a			;0753
	pop bc			;0754
	ld c,b			;0755
	and c			;0756
	rra			;0757
	jp nc,l0761h		;0758
	call sub_04abh		;075b
	jp l076fh		;075e
l0761h:
	ld hl,015a1h		;0761
	ld (hl),000h		;0764
	ld a,(015a2h)		;0766
	add a,030h		;0769
	ld c,a			;076b
	call sub_0490h		;076c
l076fh:
	jp l0710h		;076f
l0772h:
	ret			;0772
sub_0773h:
	ld hl,015a6h		;0773
	ld (hl),d			;0776
	dec hl			;0777
	ld (hl),e			;0778
	dec hl			;0779
	ld (hl),b			;077a
	dec hl			;077b
	ld (hl),c			;077c
	ld hl,(015a5h)		;077d
	ld de,0158dh		;0780
	push hl			;0783
	call sub_144bh		;0784
	ex de,hl			;0787
	pop hl			;0788
	ld (hl),e			;0789
	inc hl			;078a
	ld (hl),d			;078b
l078ch:
	ld hl,(015a5h)		;078c
	ex de,hl			;078f
	ld bc,l0400h		;0790
	call sub_14d3h		;0793
	jp c,l07b6h		;0796
	ld hl,(015a5h)		;0799
	ex de,hl			;079c
	ld bc,l0400h		;079d
	call sub_14d3h		;07a0
	ex de,hl			;07a3
	dec hl			;07a4
	ld (hl),e			;07a5
	inc hl			;07a6
	ld (hl),d			;07a7
	ld hl,(015a3h)		;07a8
	ld c,(hl)			;07ab
	inc hl			;07ac
	ld b,(hl)			;07ad
	inc bc			;07ae
	dec hl			;07af
	ld (hl),c			;07b0
	inc hl			;07b1
	ld (hl),b			;07b2
	jp l078ch		;07b3
l07b6h:
	ret			;07b6
sub_07b7h:
	ld hl,015a7h		;07b7
	ld (hl),c			;07ba
	ld hl,00000h		;07bb
	ld (015a8h),hl		;07be
	ld (015aah),hl		;07c1
	ld a,l			;07c4
	ld (015aeh),a		;07c5
	ld l,a			;07c8
	ld h,000h		;07c9
	ld (015ach),hl		;07cb
l07ceh:
	ld bc,00005h		;07ce
	ld hl,(01557h)		;07d1
	add hl,bc			;07d4
	ex de,hl			;07d5
	ld bc,015ach		;07d6
	call sub_14ceh		;07d9
	jp c,l080eh		;07dc
	ld a,(015a7h)		;07df
	rra			;07e2
	jp nc,l07f1h		;07e3
	ld hl,(015ach)		;07e6
	ld b,h			;07e9
	ld c,l			;07ea
	call sub_05dbh		;07eb
	ld (015aeh),a		;07ee
l07f1h:
	ld a,(015aeh)		;07f1
	rra			;07f4
	jp c,l0801h		;07f5
	ld de,015aah		;07f8
	ld bc,015a8h		;07fb
	call sub_0773h		;07fe
l0801h:
	ld de,00001h		;0801
	ld hl,(015ach)		;0804
	add hl,de			;0807
	ld (015ach),hl		;0808
	jp nc,l07ceh		;080b
l080eh:
	ld hl,(015a8h)		;080e
	ret			;0811
sub_0812h:
	ld bc,001b9h		;0812
	call sub_04d2h		;0815
	ret			;0818
sub_0819h:
	ld bc,001c7h		;0819
	call sub_04d2h		;081c
	call sub_058ah		;081f
	ld c,a			;0822
	ld b,000h		;0823
	ld de,0000ah		;0825
	call sub_0701h		;0828
	ld bc,001d5h		;082b
	call sub_04d2h		;082e
	ld hl,015afh		;0831
	ld (hl),000h		;0834
l0836h:
	ld a,01fh		;0836
	ld hl,015afh		;0838
	cp (hl)			;083b
	jp c,l0851h		;083c
	ld hl,(015afh)		;083f
	ld h,000h		;0842
	ld bc,015b0h		;0844
	add hl,bc			;0847
	ld (hl),000h		;0848
	ld h,b			;084a
	ld l,c			;084b
	dec hl			;084c
	inc (hl)			;084d
	jp nz,l0836h		;084e
l0851h:
	ld bc,02a08h		;0851
	call sub_0541h		;0854
	ld bc,l012ah		;0857
	call sub_0519h		;085a
l085dh:
	ld a,(01560h)		;085d
	cp 0ffh		;0860
	jp z,l0893h		;0862
	ld a,(01560h)		;0865
	and 003h		;0868
	add a,a			;086a
	add a,a			;086b
	add a,a			;086c
	add a,a			;086d
	add a,a			;086e
	ld c,a			;086f
	ld b,000h		;0870
	ld hl,02a08h		;0872
	add hl,bc			;0875
	ld a,(hl)			;0876
	ld (015afh),a		;0877
	cp 0e5h		;087a
	jp z,l088dh		;087c
	ld a,(015afh)		;087f
	and 01fh		;0882
	ld c,a			;0884
	ld b,000h		;0885
	ld hl,015b0h		;0887
	add hl,bc			;088a
	ld (hl),001h		;088b
l088dh:
	call sub_052ch		;088d
	jp l085dh		;0890
l0893h:
	ld hl,015afh		;0893
	ld (hl),000h		;0896
l0898h:
	ld a,01fh		;0898
	ld hl,015afh		;089a
	cp (hl)			;089d
	jp c,l08c2h		;089e
	ld hl,(015afh)		;08a1
	ld h,000h		;08a4
	ld bc,015b0h		;08a6
	add hl,bc			;08a9
	ld a,(hl)			;08aa
	rra			;08ab
	jp nc,l08bbh		;08ac
	ld hl,(015afh)		;08af
	ld c,l			;08b2
	ld b,000h		;08b3
	ld de,0000ah		;08b5
	call sub_0701h		;08b8
l08bbh:
	ld hl,015afh		;08bb
	inc (hl)			;08be
	jp nz,l0898h		;08bf
l08c2h:
	ret			;08c2
sub_08c3h:
	ld bc,001e3h		;08c3
	call sub_04d2h		;08c6
	call sub_0538h		;08c9
	add a,041h		;08cc
	ld c,a			;08ce
	call sub_0490h		;08cf
	ld c,03ah		;08d2
	call sub_0490h		;08d4
	ld bc,l01e8h		;08d7
	call sub_04b1h		;08da
	ld hl,(01557h)		;08dd
	inc hl			;08e0
	inc hl			;08e1
	ld c,(hl)			;08e2
	ld hl,00001h		;08e3
	call sub_14aeh		;08e6
	ld (015d0h),hl		;08e9
	ld bc,00005h		;08ec
	ld hl,(01557h)		;08ef
	add hl,bc			;08f2
	ld c,(hl)			;08f3
	inc hl			;08f4
	ld b,(hl)			;08f5
	inc bc			;08f6
	ld hl,(015d0h)		;08f7
	ex de,hl			;08fa
	call sub_1491h		;08fb
	ld (015d2h),hl		;08fe
	ld a,000h		;0901
	call sub_14c4h		;0903
	or l			;0906
	sub 001h		;0907
	sbc a,a			;0909
	ld hl,015d0h		;090a
	push af			;090d
	call sub_14e9h		;090e
	or l			;0911
	add a,0ffh		;0912
	sbc a,a			;0914
	pop bc			;0915
	ld c,b			;0916
	and c			;0917
	rra			;0918
	jp nc,l0925h		;0919
	ld bc,001ffh		;091c
	call sub_04d2h		;091f
	jp l092dh		;0922
l0925h:
	ld hl,(015d2h)		;0925
	ld b,h			;0928
	ld c,l			;0929
	call sub_09c0h		;092a
l092dh:
	ld bc,00207h		;092d
	call sub_04b1h		;0930
	ld c,000h		;0933
	call sub_07b7h		;0935
	ld b,h			;0938
	ld c,l			;0939
	call sub_09c0h		;093a
	ld bc,l0220h		;093d
	call sub_04b1h		;0940
	ld bc,00007h		;0943
	ld hl,(01557h)		;0946
	add hl,bc			;0949
	ld c,(hl)			;094a
	inc hl			;094b
	ld b,(hl)			;094c
	inc bc			;094d
	call sub_09c0h		;094e
	ld bc,00239h		;0951
	call sub_04b1h		;0954
	ld bc,0000bh		;0957
	ld hl,(01557h)		;095a
	add hl,bc			;095d
	ld e,(hl)			;095e
	inc hl			;095f
	ld d,(hl)			;0960
	ex de,hl			;0961
	add hl,hl			;0962
	add hl,hl			;0963
	ld b,h			;0964
	ld c,l			;0965
	call sub_09c0h		;0966
	ld bc,l0254h		;0969
	call sub_04b1h		;096c
	ld bc,00004h		;096f
	ld hl,(01557h)		;0972
	add hl,bc			;0975
	ld a,(hl)			;0976
	inc a			;0977
	ld l,a			;0978
	ld h,000h		;0979
	ld de,00080h		;097b
	call sub_148fh		;097e
	ld b,h			;0981
	ld c,l			;0982
	call sub_09c0h		;0983
	ld bc,0026fh		;0986
	call sub_04b1h		;0989
	ld hl,(015d0h)		;098c
	ld b,h			;098f
	ld c,l			;0990
	call sub_09c0h		;0991
	ld bc,0027fh		;0994
	call sub_04b1h		;0997
	ld hl,(01557h)		;099a
	ld c,(hl)			;099d
	inc hl			;099e
	ld b,(hl)			;099f
	call sub_09c0h		;09a0
	ld bc,l028eh		;09a3
	call sub_04b1h		;09a6
	ld bc,0000dh		;09a9
	ld hl,(01557h)		;09ac
	add hl,bc			;09af
	ld c,(hl)			;09b0
	inc hl			;09b1
	ld b,(hl)			;09b2
	call sub_09c0h		;09b3
	ld bc,0029dh		;09b6
	call sub_04b1h		;09b9
	call sub_04a0h		;09bc
	ret			;09bf
sub_09c0h:
	ld hl,015d5h		;09c0
	ld (hl),b			;09c3
	dec hl			;09c4
	ld (hl),c			;09c5
	call sub_04a0h		;09c6
	ld hl,(015d4h)		;09c9
	ld b,h			;09cc
	ld c,l			;09cd
	ld de,02710h		;09ce
	call sub_0701h		;09d1
	ld c,03ah		;09d4
	call sub_0490h		;09d6
	call sub_04abh		;09d9
	ret			;09dc
sub_09ddh:
	call sub_055ah		;09dd
	ld (015d6h),hl		;09e0
	ld hl,015d8h		;09e3
	ld (hl),000h		;09e6
l09e8h:
	ld a,000h		;09e8
	ld de,015d6h		;09ea
	call sub_14dbh		;09ed
	or l			;09f0
	jp z,l0a18h		;09f1
	ld hl,(015d6h)		;09f4
	ld a,l			;09f7
	rra			;09f8
	jp nc,l0a06h		;09f9
	ld hl,(015d8h)		;09fc
	ld c,l			;09ff
	call sub_05cch		;0a00
	call sub_08c3h		;0a03
l0a06h:
	ld c,001h		;0a06
	ld hl,015d6h		;0a08
	call sub_14b4h		;0a0b
	ex de,hl			;0a0e
	dec hl			;0a0f
	ld (hl),e			;0a10
	inc hl			;0a11
	ld (hl),d			;0a12
	inc hl			;0a13
	inc (hl)			;0a14
	jp l09e8h		;0a15
l0a18h:
	ret			;0a18
sub_0a19h:
	ld hl,015dbh		;0a19
	ld (hl),e			;0a1c
	dec hl			;0a1d
	ld (hl),b			;0a1e
	dec hl			;0a1f
	ld (hl),c			;0a20
	ld hl,015ddh		;0a21
	ld (hl),000h		;0a24
	ld hl,015dfh		;0a26
	ld (hl),000h		;0a29
	ld (hl),001h		;0a2b
l0a2dh:
	ld a,(015dbh)		;0a2d
	ld hl,015dfh		;0a30
	cp (hl)			;0a33
	jp c,l0a84h		;0a34
	ld hl,015deh		;0a37
	ld (hl),001h		;0a3a
	ld hl,015dch		;0a3c
	ld (hl),000h		;0a3f
l0a41h:
	ld a,003h		;0a41
	ld hl,015dch		;0a43
	cp (hl)			;0a46
	jp c,l0a72h		;0a47
	ld hl,(015ddh)		;0a4a
	ld h,000h		;0a4d
	ex de,hl			;0a4f
	ld hl,(015d9h)		;0a50
	add hl,de			;0a53
	push hl			;0a54
	ld hl,(015dch)		;0a55
	ld h,000h		;0a58
	ld bc,01592h		;0a5a
	add hl,bc			;0a5d
	pop de			;0a5e
	ld a,(de)			;0a5f
	cp (hl)			;0a60
	jp z,l0a69h		;0a61
	ld hl,015deh		;0a64
	ld (hl),000h		;0a67
l0a69h:
	ld hl,015ddh		;0a69
	inc (hl)			;0a6c
	dec hl			;0a6d
	inc (hl)			;0a6e
	jp nz,l0a41h		;0a6f
l0a72h:
	ld a,(015deh)		;0a72
	rra			;0a75
	jp nc,l0a7dh		;0a76
	ld a,(015dfh)		;0a79
	ret			;0a7c
l0a7dh:
	ld hl,015dfh		;0a7d
	inc (hl)			;0a80
	jp nz,l0a2dh		;0a81
l0a84h:
	ld a,000h		;0a84
	ret			;0a86
sub_0a87h:
	ld hl,015e3h		;0a87
	ld (hl),000h		;0a8a
l0a8ch:
	call l0639h		;0a8c
	ld e,008h		;0a8f
	ld bc,00139h		;0a91
	call sub_0a19h		;0a94
	ld (015e0h),a		;0a97
	cp 000h		;0a9a
	jp nz,l0aa8h		;0a9c
	ld a,(015e3h)		;0a9f
	sub 000h		;0aa2
	add a,0ffh		;0aa4
	sbc a,a			;0aa6
	ret			;0aa7
l0aa8h:
	ld hl,015e3h		;0aa8
	inc (hl)			;0aab
	ld a,(015e0h)		;0aac
	cp 005h		;0aaf
	jp nz,l0b18h		;0ab1
	ld a,(00003h)		;0ab4
	ld (015e2h),a		;0ab7
	ld hl,015e1h		;0aba
	ld (hl),000h		;0abd
	dec hl			;0abf
	ld (hl),000h		;0ac0
l0ac2h:
	ld a,003h		;0ac2
	ld hl,015e0h		;0ac4
	cp (hl)			;0ac7
	jp c,l0b15h		;0ac8
	ld a,(015e0h)		;0acb
	add a,a			;0ace
	add a,a			;0acf
	ld c,a			;0ad0
	ld b,000h		;0ad1
	ld hl,00139h		;0ad3
	add hl,bc			;0ad6
	ld b,h			;0ad7
	ld c,l			;0ad8
	call sub_0c69h		;0ad9
	ld bc,002adh		;0adc
	call sub_04b1h		;0adf
	ld a,(015e2h)		;0ae2
	and 003h		;0ae5
	add a,a			;0ae7
	add a,a			;0ae8
	ld hl,015e1h		;0ae9
	add a,(hl)			;0aec
	ld c,a			;0aed
	ld b,000h		;0aee
	ld hl,00159h		;0af0
	add hl,bc			;0af3
	ld b,h			;0af4
	ld c,l			;0af5
	call sub_0c69h		;0af6
	ld a,(015e1h)		;0af9
	add a,010h		;0afc
	ld (015e1h),a		;0afe
	ld a,(015e2h)		;0b01
	and 0feh		;0b04
	rra			;0b06
	rra			;0b07
	ld (015e2h),a		;0b08
	call sub_04a0h		;0b0b
	ld hl,015e0h		;0b0e
	inc (hl)			;0b11
	jp nz,l0ac2h		;0b12
l0b15h:
	jp l0c46h		;0b15
l0b18h:
	ld a,(015e0h)		;0b18
	cp 006h		;0b1b
	jp nz,l0baeh		;0b1d
	ld bc,l02b2h		;0b20
	call sub_04d2h		;0b23
	ld bc,l02c8h		;0b26
	call sub_04d2h		;0b29
	ld bc,l02fah		;0b2c
	call sub_04d2h		;0b2f
	ld bc,00315h		;0b32
	call sub_04d2h		;0b35
	ld bc,00329h		;0b38
	call sub_04d2h		;0b3b
	ld hl,015e0h		;0b3e
	ld (hl),000h		;0b41
l0b43h:
	ld a,003h		;0b43
	ld hl,015e0h		;0b45
	cp (hl)			;0b48
	jp c,l0babh		;0b49
	call sub_04a0h		;0b4c
	ld a,(015e0h)		;0b4f
	add a,a			;0b52
	add a,a			;0b53
	ld c,a			;0b54
	ld b,000h		;0b55
	ld hl,00139h		;0b57
	add hl,bc			;0b5a
	ld b,h			;0b5b
	ld c,l			;0b5c
	call sub_0c69h		;0b5d
	ld bc,l0338h		;0b60
	call sub_04b1h		;0b63
	ld hl,015e1h		;0b66
	ld (hl),000h		;0b69
l0b6bh:
	ld a,00ch		;0b6b
	ld hl,015e1h		;0b6d
	cp (hl)			;0b70
	jp c,l0ba4h		;0b71
	jp l0b85h		;0b74
l0b77h:
	ld a,(015e1h)		;0b77
	add a,004h		;0b7a
	ld (015e1h),a		;0b7c
	jp nc,l0b6bh		;0b7f
	jp l0ba4h		;0b82
l0b85h:
	ld c,020h		;0b85
	call sub_0490h		;0b87
	ld a,(015e0h)		;0b8a
	add a,a			;0b8d
	add a,a			;0b8e
	add a,a			;0b8f
	add a,a			;0b90
	ld hl,015e1h		;0b91
	add a,(hl)			;0b94
	ld c,a			;0b95
	ld b,000h		;0b96
	ld hl,00159h		;0b98
	add hl,bc			;0b9b
	ld b,h			;0b9c
	ld c,l			;0b9d
	call sub_0c69h		;0b9e
	jp l0b77h		;0ba1
l0ba4h:
	ld hl,015e0h		;0ba4
	inc (hl)			;0ba7
	jp nz,l0b43h		;0ba8
l0babh:
	jp l0c46h		;0bab
l0baeh:
	ld a,(015e0h)		;0bae
	cp 007h		;0bb1
	jp nz,l0bbfh		;0bb3
	call sub_0819h		;0bb6
	ld a,001h		;0bb9
	ret			;0bbb
	jp l0c46h		;0bbc
l0bbfh:
	ld a,(015e0h)		;0bbf
	cp 008h		;0bc2
	jp nz,l0bcdh		;0bc4
	call sub_09ddh		;0bc7
	jp l0c46h		;0bca
l0bcdh:
	ld a,(015e0h)		;0bcd
	dec a			;0bd0
	ld (015e0h),a		;0bd1
	add a,a			;0bd4
	add a,a			;0bd5
	add a,a			;0bd6
	add a,a			;0bd7
	ld (015e1h),a		;0bd8
	call l0639h		;0bdb
	ld a,(01592h)		;0bde
	cp 03dh		;0be1
	jp z,l0befh		;0be3
	ld bc,0033bh		;0be6
	call sub_04d2h		;0be9
	ld a,001h		;0bec
	ret			;0bee
l0befh:
	call l0639h		;0bef
	ld hl,(015e1h)		;0bf2
	ld h,000h		;0bf5
	ld bc,00159h		;0bf7
	add hl,bc			;0bfa
	ld b,h			;0bfb
	ld c,l			;0bfc
	ld e,004h		;0bfd
	call sub_0a19h		;0bff
	dec a			;0c02
	ld (015e1h),a		;0c03
	cp 0ffh		;0c06
	jp nz,l0c14h		;0c08
	ld bc,00349h		;0c0b
	call sub_04d2h		;0c0e
	ld a,001h		;0c11
	ret			;0c13
l0c14h:
	ld hl,015e2h		;0c14
	ld (hl),0fch		;0c17
l0c19h:
	ld a,(015e0h)		;0c19
	dec a			;0c1c
	ld (015e0h),a		;0c1d
	cp 0ffh		;0c20
	jp z,l0c38h		;0c22
	ld a,(015e2h)		;0c25
	rlca			;0c28
	rlca			;0c29
	ld (015e2h),a		;0c2a
	ld a,(015e1h)		;0c2d
	add a,a			;0c30
	add a,a			;0c31
	ld (015e1h),a		;0c32
	jp l0c19h		;0c35
l0c38h:
	ld a,(015e2h)		;0c38
	ld hl,00003h		;0c3b
	and (hl)			;0c3e
	ld hl,015e1h		;0c3f
	or (hl)			;0c42
	ld (00003h),a		;0c43
l0c46h:
	call l0639h		;0c46
	ld a,(01592h)		;0c49
	cp 020h		;0c4c
	jp nz,l0c54h		;0c4e
	ld a,001h		;0c51
	ret			;0c53
l0c54h:
	ld a,(01592h)		;0c54
	cp 02ch		;0c57
	jp z,l0c65h		;0c59
	ld bc,l035ch		;0c5c
	call sub_04d2h		;0c5f
	ld a,001h		;0c62
	ret			;0c64
l0c65h:
	jp l0a8ch		;0c65
	ret			;0c68
sub_0c69h:
	ld hl,015e5h		;0c69
	ld (hl),b			;0c6c
	dec hl			;0c6d
	ld (hl),c			;0c6e
l0c6fh:
	ld hl,(015e4h)		;0c6f
	ld a,(hl)			;0c72
	cp 03ah		;0c73
	jp z,l0c89h		;0c75
	ld hl,(015e4h)		;0c78
	ld c,(hl)			;0c7b
	call sub_0490h		;0c7c
	ld hl,(015e4h)		;0c7f
	inc hl			;0c82
	ld (015e4h),hl		;0c83
	jp l0c6fh		;0c86
l0c89h:
	ld c,03ah		;0c89
	call sub_0490h		;0c8b
	ret			;0c8e
sub_0c8fh:
	ld hl,015e7h		;0c8f
	ld (hl),b			;0c92
	dec hl			;0c93
	ld (hl),c			;0c94
	ld hl,02710h		;0c95
	ld (015eah),hl		;0c98
	ld hl,015e9h		;0c9b
	ld (hl),000h		;0c9e
l0ca0h:
	ld a,000h		;0ca0
	ld de,015eah		;0ca2
	call sub_14dbh		;0ca5
	or l			;0ca8
	jp z,l0cfch		;0ca9
	ld hl,(015e6h)		;0cac
	ex de,hl			;0caf
	ld hl,(015eah)		;0cb0
	call sub_1470h		;0cb3
	ld a,e			;0cb6
	ld (015e8h),a		;0cb7
	ld hl,(015e6h)		;0cba
	ex de,hl			;0cbd
	call sub_1472h		;0cbe
	ld (015e6h),hl		;0cc1
	ld d,b			;0cc4
	ld e,c			;0cc5
	ld hl,0000ah		;0cc6
	call sub_1470h		;0cc9
	ex de,hl			;0ccc
	ld (015eah),hl		;0ccd
	ld a,000h		;0cd0
	call sub_14c4h		;0cd2
	or l			;0cd5
	sub 001h		;0cd6
	sbc a,a			;0cd8
	ld hl,015e9h		;0cd9
	or (hl)			;0cdc
	dec hl			;0cdd
	push af			;0cde
	ld a,(hl)			;0cdf
	sub e			;0ce0
	add a,0ffh		;0ce1
	sbc a,a			;0ce3
	pop bc			;0ce4
	ld c,b			;0ce5
	or c			;0ce6
	rra			;0ce7
	jp nc,l0cf9h		;0ce8
	ld hl,015e9h		;0ceb
	ld (hl),001h		;0cee
	ld a,(015e8h)		;0cf0
	add a,030h		;0cf3
	ld c,a			;0cf5
	call sub_0490h		;0cf6
l0cf9h:
	jp l0ca0h		;0cf9
l0cfch:
	ld c,06bh		;0cfc
	call sub_0490h		;0cfe
	call sub_04a0h		;0d01
	ret			;0d04
sub_0d05h:
	call sub_0551h		;0d05
	ld (01559h),hl		;0d08
	call sub_0538h		;0d0b
	add a,041h		;0d0e
	ld c,a			;0d10
	call sub_0490h		;0d11
	ld bc,l036ah		;0d14
	call sub_04b1h		;0d17
	ret			;0d1a
sub_0d1bh:
	ld c,001h		;0d1b
	call sub_07b7h		;0d1d
	ld b,h			;0d20
	ld c,l			;0d21
	call sub_0c8fh		;0d22
	ret			;0d25
sub_0d26h:
	ld bc,0036dh		;0d26
	call sub_04d2h		;0d29
	call sub_0d05h		;0d2c
	call sub_0d1bh		;0d2f
	ret			;0d32
sub_0d33h:
	call sub_055ah		;0d33
	ld (015ech),hl		;0d36
	call sub_056ch		;0d39
	ld (015eeh),hl		;0d3c
	ld hl,015f0h		;0d3f
	ld (hl),000h		;0d42
l0d44h:
	ld a,000h		;0d44
	ld de,015ech		;0d46
	call sub_14dbh		;0d49
	or l			;0d4c
	jp z,l0da5h		;0d4d
	ld hl,(015ech)		;0d50
	ld a,l			;0d53
	rra			;0d54
	jp nc,l0d86h		;0d55
	ld hl,(015f0h)		;0d58
	ld c,l			;0d5b
	call sub_05cch		;0d5c
	call sub_0d05h		;0d5f
	ld bc,00381h		;0d62
	call sub_04b1h		;0d65
	ld hl,(015eeh)		;0d68
	ld a,l			;0d6b
	rra			;0d6c
	jp nc,l0d78h		;0d6d
	ld c,04fh		;0d70
	call sub_0490h		;0d72
	jp l0d7dh		;0d75
l0d78h:
	ld c,057h		;0d78
	call sub_0490h		;0d7a
l0d7dh:
	ld bc,l0384h		;0d7d
	call sub_04b1h		;0d80
	call sub_0d1bh		;0d83
l0d86h:
	ld c,001h		;0d86
	ld hl,015ech		;0d88
	call sub_14b4h		;0d8b
	ex de,hl			;0d8e
	dec hl			;0d8f
	ld (hl),e			;0d90
	inc hl			;0d91
	ld (hl),d			;0d92
	ld c,001h		;0d93
	ld hl,015eeh		;0d95
	call sub_14b4h		;0d98
	ex de,hl			;0d9b
	dec hl			;0d9c
	ld (hl),e			;0d9d
	inc hl			;0d9e
	ld (hl),d			;0d9f
	inc hl			;0da0
	inc (hl)			;0da1
	jp l0d44h		;0da2
l0da5h:
	call sub_04a0h		;0da5
	ret			;0da8
sub_0da9h:
	ld a,(0005ch)		;0da9
	cp 000h		;0dac
	jp z,l0db9h		;0dae
	ld a,(0005ch)		;0db1
	dec a			;0db4
	ld c,a			;0db5
	call sub_05cch		;0db6
l0db9h:
	ret			;0db9
sub_0dbah:
	call sub_05b3h		;0dba
	call sub_0da9h		;0dbd
	ld hl,01556h		;0dc0
	ld (hl),000h		;0dc3
	ld hl,02a05h		;0dc5
	ld (hl),0ffh		;0dc8
	call sub_136ch		;0dca
	rra			;0dcd
	jp nc,l0de4h		;0dce
	ld a,(02a05h)		;0dd1
	cp 000h		;0dd4
	jp nz,l0ddah		;0dd6
	ret			;0dd9
l0ddah:
	ld a,(02a05h)		;0dda
	dec a			;0ddd
	ld (02a05h),a		;0dde
	jp l0df0h		;0de1
l0de4h:
	ld a,(0005dh)		;0de4
	cp 020h		;0de7
	jp nz,l0df0h		;0de9
	call sub_0d26h		;0dec
	ret			;0def
l0df0h:
	ld hl,00000h		;0df0
	ld (015f1h),hl		;0df3
	ld a,l			;0df6
	ld (0005ch),a		;0df7
	ld hl,00068h		;0dfa
	ld (hl),03fh		;0dfd
	ld hl,0006ah		;0dff
	ld (hl),03fh		;0e02
	ld bc,0005ch		;0e04
	call sub_0519h		;0e07
l0e0ah:
	ld a,(01560h)		;0e0a
	cp 0ffh		;0e0d
	jp z,l1043h		;0e0f
	ld a,(01560h)		;0e12
	and 003h		;0e15
	add a,a			;0e17
	add a,a			;0e18
	add a,a			;0e19
	add a,a			;0e1a
	add a,a			;0e1b
	add a,080h		;0e1c
	ld l,a			;0e1e
	ld h,000h		;0e1f
	ld (029f3h),hl		;0e21
	ld hl,02a04h		;0e24
	ld (hl),000h		;0e27
	ld hl,00000h		;0e29
	ld (029f7h),hl		;0e2c
l0e2fh:
	ld a,(02a04h)		;0e2f
	cpl			;0e32
	ld bc,015f1h		;0e33
	ld de,029f7h		;0e36
	push af			;0e39
	call sub_14ceh		;0e3a
	sbc a,a			;0e3d
	pop bc			;0e3e
	ld c,b			;0e3f
	and c			;0e40
	rra			;0e41
	jp nc,l0e98h		;0e42
	call sub_135dh		;0e45
	ld hl,029ffh		;0e48
	ld (hl),001h		;0e4b
l0e4dh:
	ld a,00bh		;0e4d
	ld hl,029ffh		;0e4f
	cp (hl)			;0e52
	jp c,l0e8eh		;0e53
	ld hl,(029ffh)		;0e56
	ld h,000h		;0e59
	ex de,hl			;0e5b
	ld hl,(029f3h)		;0e5c
	add hl,de			;0e5f
	push hl			;0e60
	ld hl,(029ffh)		;0e61
	ld h,000h		;0e64
	ex de,hl			;0e66
	ld hl,(029f5h)		;0e67
	add hl,de			;0e6a
	pop bc			;0e6b
	ld a,(bc)			;0e6c
	cp (hl)			;0e6d
	jp z,l0e79h		;0e6e
	ld hl,029ffh		;0e71
	ld (hl),00bh		;0e74
	jp l0e84h		;0e76
l0e79h:
	ld a,(029ffh)		;0e79
	sub 00bh		;0e7c
	sub 001h		;0e7e
	sbc a,a			;0e80
	ld (02a04h),a		;0e81
l0e84h:
	ld a,(029ffh)		;0e84
	inc a			;0e87
	ld (029ffh),a		;0e88
	jp nz,l0e4dh		;0e8b
l0e8eh:
	ld hl,(029f7h)		;0e8e
	inc hl			;0e91
	ld (029f7h),hl		;0e92
	jp l0e2fh		;0e95
l0e98h:
	ld a,(02a04h)		;0e98
	rra			;0e9b
	jp nc,l0ea9h		;0e9c
	ld hl,(029f7h)		;0e9f
	dec hl			;0ea2
	ld (029f7h),hl		;0ea3
	jp l0f60h		;0ea6
l0ea9h:
	ld hl,(015f1h)		;0ea9
	ld (029f7h),hl		;0eac
	inc hl			;0eaf
	ld (015f1h),hl		;0eb0
	call sub_135dh		;0eb3
	ld de,l0200h		;0eb6
	ld hl,015f1h		;0eb9
	call sub_14e9h		;0ebc
	sbc a,a			;0ebf
	ld de,00010h		;0ec0
	ld hl,(029f5h)		;0ec3
	add hl,de			;0ec6
	ex de,hl			;0ec7
	ld hl,00006h		;0ec8
	push af			;0ecb
	call sub_14e9h		;0ecc
	sbc a,a			;0ecf
	cpl			;0ed0
	pop bc			;0ed1
	ld c,b			;0ed2
	or c			;0ed3
	rra			;0ed4
	jp nc,l0eedh		;0ed5
	ld bc,003a5h		;0ed8
	call sub_04d2h		;0edb
	ld hl,00000h		;0ede
	ld (029f7h),hl		;0ee1
	ld hl,00001h		;0ee4
	ld (015f1h),hl		;0ee7
	call sub_135dh		;0eea
l0eedh:
	ld hl,(029f7h)		;0eed
	ld bc,015f3h		;0ef0
	add hl,hl			;0ef3
	add hl,bc			;0ef4
	push hl			;0ef5
	ld hl,(029f7h)		;0ef6
	ex de,hl			;0ef9
	pop hl			;0efa
	ld (hl),e			;0efb
	inc hl			;0efc
	ld (hl),d			;0efd
	ld hl,029ffh		;0efe
	ld (hl),000h		;0f01
l0f03h:
	ld a,00bh		;0f03
	ld hl,029ffh		;0f05
	cp (hl)			;0f08
	jp c,l0f2eh		;0f09
	ld hl,(029ffh)		;0f0c
	ld h,000h		;0f0f
	ex de,hl			;0f11
	ld hl,(029f3h)		;0f12
	add hl,de			;0f15
	push hl			;0f16
	ld hl,(029ffh)		;0f17
	ld h,000h		;0f1a
	ex de,hl			;0f1c
	ld hl,(029f5h)		;0f1d
	add hl,de			;0f20
	pop bc			;0f21
	ld a,(bc)			;0f22
	ld (hl),a			;0f23
	ld a,(029ffh)		;0f24
	inc a			;0f27
	ld (029ffh),a		;0f28
	jp nz,l0f03h		;0f2b
l0f2eh:
	ld hl,(029f7h)		;0f2e
	ld bc,019f3h		;0f31
	add hl,hl			;0f34
	add hl,bc			;0f35
	ld a,000h		;0f36
	ld (hl),a			;0f38
	inc hl			;0f39
	ld (hl),000h		;0f3a
	ld hl,(029f7h)		;0f3c
	ld bc,01df3h		;0f3f
	add hl,hl			;0f42
	add hl,bc			;0f43
	ld (hl),a			;0f44
	inc hl			;0f45
	ld (hl),000h		;0f46
	ld hl,(029f7h)		;0f48
	ld bc,021f3h		;0f4b
	add hl,hl			;0f4e
	add hl,bc			;0f4f
	ld (hl),a			;0f50
	inc hl			;0f51
	ld (hl),000h		;0f52
	ld hl,(029f7h)		;0f54
	ld bc,025f3h		;0f57
	add hl,hl			;0f5a
	add hl,bc			;0f5b
	ld (hl),a			;0f5c
	inc hl			;0f5d
	ld (hl),000h		;0f5e
l0f60h:
	ld hl,(029f7h)		;0f60
	ld bc,019f3h		;0f63
	add hl,hl			;0f66
	add hl,bc			;0f67
	ld c,(hl)			;0f68
	inc hl			;0f69
	ld b,(hl)			;0f6a
	inc bc			;0f6b
	ld hl,(029f7h)		;0f6c
	push bc			;0f6f
	ld bc,019f3h		;0f70
	add hl,hl			;0f73
	add hl,bc			;0f74
	pop bc			;0f75
	ld (hl),c			;0f76
	inc hl			;0f77
	ld (hl),b			;0f78
	ld hl,(029f7h)		;0f79
	ld bc,025f3h		;0f7c
	add hl,hl			;0f7f
	add hl,bc			;0f80
	ld bc,0000fh		;0f81
	push hl			;0f84
	ld hl,(029f3h)		;0f85
	add hl,bc			;0f88
	ld a,(hl)			;0f89
	pop de			;0f8a
	call sub_1456h		;0f8b
	ld bc,0000ch		;0f8e
	push hl			;0f91
	ld hl,(029f3h)		;0f92
	add hl,bc			;0f95
	ld bc,00004h		;0f96
	push hl			;0f99
	ld hl,(01557h)		;0f9a
	add hl,bc			;0f9d
	ld a,(hl)			;0f9e
	pop hl			;0f9f
	and (hl)			;0fa0
	ld l,a			;0fa1
	ld h,000h		;0fa2
	ld de,00080h		;0fa4
	call sub_148fh		;0fa7
	pop bc			;0faa
	add hl,bc			;0fab
	push hl			;0fac
	ld hl,(029f7h)		;0fad
	ld bc,025f3h		;0fb0
	add hl,hl			;0fb3
	add hl,bc			;0fb4
	pop bc			;0fb5
	ld (hl),c			;0fb6
	inc hl			;0fb7
	ld (hl),b			;0fb8
	ld hl,02a00h		;0fb9
	ld (hl),001h		;0fbc
	ld bc,00005h		;0fbe
	ld hl,(01557h)		;0fc1
	add hl,bc			;0fc4
	ld a,0ffh		;0fc5
	call sub_14e6h		;0fc7
	jp nc,l0fd2h		;0fca
	ld hl,02a00h		;0fcd
	ld (hl),002h		;0fd0
l0fd2h:
	ld hl,029ffh		;0fd2
	ld (hl),010h		;0fd5
l0fd7h:
	ld a,01fh		;0fd7
	ld hl,029ffh		;0fd9
	cp (hl)			;0fdc
	jp c,l103dh		;0fdd
	jp l0ff1h		;0fe0
l0fe3h:
	ld a,(02a00h)		;0fe3
	ld hl,029ffh		;0fe6
	add a,(hl)			;0fe9
	ld (hl),a			;0fea
	jp nc,l0fd7h		;0feb
	jp l103dh		;0fee
l0ff1h:
	ld hl,(029ffh)		;0ff1
	ld h,000h		;0ff4
	ex de,hl			;0ff6
	ld hl,(029f3h)		;0ff7
	add hl,de			;0ffa
	ld a,(hl)			;0ffb
	ld (02a01h),a		;0ffc
	ld a,(02a00h)		;0fff
	cp 002h		;1002
	jp nz,l101ch		;1004
	ld hl,(029ffh)		;1007
	ld h,000h		;100a
	ld bc,00001h		;100c
	add hl,bc			;100f
	ex de,hl			;1010
	ld hl,(029f3h)		;1011
	add hl,de			;1014
	ld a,(02a01h)		;1015
	or (hl)			;1018
	ld (02a01h),a		;1019
l101ch:
	ld a,(02a01h)		;101c
	cp 000h		;101f
	jp z,l103ah		;1021
	ld hl,(029f7h)		;1024
	ld bc,021f3h		;1027
	add hl,hl			;102a
	add hl,bc			;102b
	push hl			;102c
	ld hl,(029f7h)		;102d
	ld bc,01df3h		;1030
	add hl,hl			;1033
	add hl,bc			;1034
	ex de,hl			;1035
	pop bc			;1036
	call sub_0773h		;1037
l103ah:
	jp l0fe3h		;103a
l103dh:
	call sub_052ch		;103d
	jp l0e0ah		;1040
l1043h:
	ld a,000h		;1043
	ld de,015f1h		;1045
	call sub_14dbh		;1048
	or l			;104b
	jp nz,l1058h		;104c
	ld bc,l03bah		;104f
	call sub_04d2h		;1052
	jp l135ch		;1055
l1058h:
	ld a,(02a05h)		;1058
	cp 0ffh		;105b
	jp nz,l1293h		;105d
	ld a,001h		;1060
	ld hl,015f1h		;1062
	call sub_14e6h		;1065
	jp nc,l1164h		;1068
	ld hl,00001h		;106b
	ld (029f9h),hl		;106e
l1071h:
	ld a,000h		;1071
	ld hl,029f9h		;1073
	call sub_14e6h		;1076
	jp nc,l1164h		;1079
	ld hl,00000h		;107c
	ld (029f9h),hl		;107f
	ld hl,00000h		;1082
	ld (029fdh),hl		;1085
l1088h:
	ld hl,(015f1h)		;1088
	dec hl			;108b
	dec hl			;108c
	ex de,hl			;108d
	ld hl,029fdh		;108e
	call sub_14e9h		;1091
	jp c,l1161h		;1094
	ld hl,(029fdh)		;1097
	ld bc,015f5h		;109a
	add hl,hl			;109d
	add hl,bc			;109e
	ld e,(hl)			;109f
	inc hl			;10a0
	ld d,(hl)			;10a1
	ex de,hl			;10a2
	ld (029f7h),hl		;10a3
	call sub_135dh		;10a6
	ld hl,(029f5h)		;10a9
	ld (029f3h),hl		;10ac
	ld hl,(029fdh)		;10af
	ld bc,015f3h		;10b2
	add hl,hl			;10b5
	add hl,bc			;10b6
	ld e,(hl)			;10b7
	inc hl			;10b8
	ld d,(hl)			;10b9
	ex de,hl			;10ba
	ld (029f7h),hl		;10bb
	call sub_135dh		;10be
	ld hl,029ffh		;10c1
	ld (hl),001h		;10c4
l10c6h:
	ld a,00bh		;10c6
	ld hl,029ffh		;10c8
	cp (hl)			;10cb
	jp c,l1154h		;10cc
	ld hl,(029ffh)		;10cf
	ld h,000h		;10d2
	ex de,hl			;10d4
	ld hl,(029f3h)		;10d5
	add hl,de			;10d8
	ld a,(hl)			;10d9
	ld (02a02h),a		;10da
	push hl			;10dd
	ld hl,(029ffh)		;10de
	ld h,000h		;10e1
	ex de,hl			;10e3
	ld hl,(029f5h)		;10e4
	add hl,de			;10e7
	ld a,(hl)			;10e8
	ld (02a03h),a		;10e9
	ld c,a			;10ec
	pop de			;10ed
	ld a,(de)			;10ee
	cp c			;10ef
	jp nc,l113bh		;10f0
	ld hl,(029fdh)		;10f3
	ld bc,015f3h		;10f6
	add hl,hl			;10f9
	add hl,bc			;10fa
	ld e,(hl)			;10fb
	inc hl			;10fc
	ld d,(hl)			;10fd
	ex de,hl			;10fe
	ld (029fbh),hl		;10ff
	ld hl,(029fdh)		;1102
	ld bc,015f5h		;1105
	add hl,hl			;1108
	add hl,bc			;1109
	push hl			;110a
	ld hl,(029fdh)		;110b
	ld bc,015f3h		;110e
	add hl,hl			;1111
	add hl,bc			;1112
	ex (sp),hl			;1113
	ld c,(hl)			;1114
	inc hl			;1115
	ld b,(hl)			;1116
	pop hl			;1117
	ld (hl),c			;1118
	inc hl			;1119
	ld (hl),b			;111a
	ld hl,(029fdh)		;111b
	ld bc,015f5h		;111e
	add hl,hl			;1121
	add hl,bc			;1122
	push hl			;1123
	ld hl,(029fbh)		;1124
	ex de,hl			;1127
	pop hl			;1128
	ld (hl),e			;1129
	inc hl			;112a
	ld (hl),d			;112b
	ld hl,(029f9h)		;112c
	inc hl			;112f
	ld (029f9h),hl		;1130
	ld hl,029ffh		;1133
	ld (hl),00bh		;1136
	jp l114ah		;1138
l113bh:
	ld a,(02a03h)		;113b
	ld hl,02a02h		;113e
	cp (hl)			;1141
	jp nc,l114ah		;1142
	ld hl,029ffh		;1145
	ld (hl),00bh		;1148
l114ah:
	ld a,(029ffh)		;114a
	inc a			;114d
	ld (029ffh),a		;114e
	jp nz,l10c6h		;1151
l1154h:
	ld de,00001h		;1154
	ld hl,(029fdh)		;1157
	add hl,de			;115a
	ld (029fdh),hl		;115b
	jp nc,l1088h		;115e
l1161h:
	jp l1071h		;1161
l1164h:
	ld a,(01556h)		;1164
	rra			;1167
	jp nc,l1174h		;1168
	ld bc,003c9h		;116b
	call sub_04d2h		;116e
	jp l1177h		;1171
l1174h:
	call sub_04a0h		;1174
l1177h:
	ld bc,l03d0h		;1177
	call sub_04b1h		;117a
	ld hl,00000h		;117d
	ld (029f9h),hl		;1180
l1183h:
	ld bc,015f1h		;1183
	ld de,029f9h		;1186
	call sub_14ceh		;1189
	jp nc,l128dh		;118c
	ld hl,(029f9h)		;118f
	ld bc,015f3h		;1192
	add hl,hl			;1195
	add hl,bc			;1196
	ld e,(hl)			;1197
	inc hl			;1198
	ld d,(hl)			;1199
	ex de,hl			;119a
	ld (029f7h),hl		;119b
	call sub_135dh		;119e
	call sub_04a0h		;11a1
	ld l,010h		;11a4
	push hl			;11a6
	ld hl,(029f5h)		;11a7
	ld b,h			;11aa
	ld c,l			;11ab
	ld de,0005ch		;11ac
	pop hl			;11af
l11b0h:
	ld a,(bc)			;11b0
	ld (de),a			;11b1
	inc bc			;11b2
	inc de			;11b3
	dec l			;11b4
	jp nz,l11b0h		;11b5
	ld hl,0005ch		;11b8
	ld (hl),000h		;11bb
	ld a,(01556h)		;11bd
	rra			;11c0
	jp nc,l11e9h		;11c1
	ld bc,0005ch		;11c4
	call sub_05a3h		;11c7
	ld a,(0007fh)		;11ca
	cp 000h		;11cd
	jp z,l11dbh		;11cf
	ld bc,l03e6h		;11d2
	call sub_04b1h		;11d5
	jp l11e6h		;11d8
l11dbh:
	ld hl,(0007dh)		;11db
	ld b,h			;11de
	ld c,l			;11df
	ld de,02710h		;11e0
	call sub_0701h		;11e3
l11e6h:
	call sub_04abh		;11e6
l11e9h:
	ld hl,(029f7h)		;11e9
	ld bc,025f3h		;11ec
	add hl,hl			;11ef
	add hl,bc			;11f0
	ld c,(hl)			;11f1
	inc hl			;11f2
	ld b,(hl)			;11f3
	ld de,02710h		;11f4
	call sub_0701h		;11f7
	call sub_04abh		;11fa
	ld hl,(029f7h)		;11fd
	ld bc,021f3h		;1200
	add hl,hl			;1203
	add hl,bc			;1204
	ld c,(hl)			;1205
	inc hl			;1206
	ld b,(hl)			;1207
	ld de,02710h		;1208
	call sub_0701h		;120b
	ld c,06bh		;120e
	call sub_0490h		;1210
	call sub_04abh		;1213
	ld hl,(029f7h)		;1216
	ld bc,019f3h		;1219
	add hl,hl			;121c
	add hl,bc			;121d
	ld c,(hl)			;121e
	inc hl			;121f
	ld b,(hl)			;1220
	ld de,l03e8h		;1221
	call sub_0701h		;1224
	call sub_04abh		;1227
	ld c,052h		;122a
	call sub_0490h		;122c
	ld c,02fh		;122f
	call sub_0490h		;1231
	ld bc,00009h		;1234
	ld hl,(029f5h)		;1237
	add hl,bc			;123a
	ld a,(hl)			;123b
	rlca			;123c
	rra			;123d
	jp nc,l1249h		;123e
	ld c,04fh		;1241
	call sub_0490h		;1243
	jp l124eh		;1246
l1249h:
	ld c,057h		;1249
	call sub_0490h		;124b
l124eh:
	call sub_04abh		;124e
	call sub_0538h		;1251
	add a,041h		;1254
	ld c,a			;1256
	call sub_0490h		;1257
	ld c,03ah		;125a
	call sub_0490h		;125c
	ld bc,0000ah		;125f
	ld hl,(029f5h)		;1262
	add hl,bc			;1265
	ld a,(hl)			;1266
	rlca			;1267
	ld (02a01h),a		;1268
	rra			;126b
	jp nc,l1274h		;126c
	ld c,028h		;126f
	call sub_0490h		;1271
l1274h:
	call sub_13c0h		;1274
	ld a,(02a01h)		;1277
	rra			;127a
	jp nc,l1283h		;127b
	ld c,029h		;127e
	call sub_0490h		;1280
l1283h:
	ld hl,(029f9h)		;1283
	inc hl			;1286
	ld (029f9h),hl		;1287
	jp l1183h		;128a
l128dh:
	call sub_0d26h		;128d
	jp l135ch		;1290
l1293h:
	ld hl,00000h		;1293
	ld (029f9h),hl		;1296
l1299h:
	ld bc,015f1h		;1299
	ld de,029f9h		;129c
	call sub_14ceh		;129f
	jp nc,l135ch		;12a2
	call sub_04e4h		;12a5
	rra			;12a8
	jp nc,l12b0h		;12a9
	call sub_0812h		;12ac
	ret			;12af
l12b0h:
	ld hl,(029f9h)		;12b0
	ld (029f7h),hl		;12b3
	call sub_135dh		;12b6
	call sub_04a0h		;12b9
	call sub_13c0h		;12bc
	ld hl,(02a05h)		;12bf
	ld c,l			;12c2
	ld b,000h		;12c3
	ld hl,l1317h		;12c5
	add hl,bc			;12c8
	add hl,bc			;12c9
	ld e,(hl)			;12ca
	inc hl			;12cb
	ld d,(hl)			;12cc
	ex de,hl			;12cd
	jp (hl)			;12ce
	ld bc,00009h		;12cf
	ld hl,(029f5h)		;12d2
	add hl,bc			;12d5
	ld a,080h		;12d6
	or (hl)			;12d8
	ld hl,(029f5h)		;12d9
	add hl,bc			;12dc
	ld (hl),a			;12dd
	jp l131fh		;12de
	ld bc,00009h		;12e1
	ld hl,(029f5h)		;12e4
	add hl,bc			;12e7
	ld a,07fh		;12e8
	and (hl)			;12ea
	ld hl,(029f5h)		;12eb
	add hl,bc			;12ee
	ld (hl),a			;12ef
	jp l131fh		;12f0
	ld bc,0000ah		;12f3
	ld hl,(029f5h)		;12f6
	add hl,bc			;12f9
	ld a,080h		;12fa
	or (hl)			;12fc
	ld hl,(029f5h)		;12fd
code1_end:
	add hl,bc			;1300
	ld (hl),a			;1301
	jp l131fh		;1302
	ld bc,0000ah		;1305
	ld hl,(029f5h)		;1308
	add hl,bc			;130b
	ld a,07fh		;130c
	and (hl)			;130e
	ld hl,(029f5h)		;130f
	add hl,bc			;1312
	ld (hl),a			;1313
	jp l131fh		;1314
l1317h:
	rst 8			;1317
	ld (de),a			;1318
	pop hl			;1319
	ld (de),a			;131a
	di			;131b
	ld (de),a			;131c
	dec b			;131d
	inc de			;131e
l131fh:
	ld l,010h		;131f
	push hl			;1321
	ld hl,(029f5h)		;1322
	ld b,h			;1325
	ld c,l			;1326
	ld de,0005ch		;1327
	pop hl			;132a
l132bh:
	ld a,(bc)			;132b
	ld (de),a			;132c
	inc bc			;132d
	inc de			;132e
	dec l			;132f
	jp nz,l132bh		;1330
	ld hl,0005ch		;1333
	ld (hl),000h		;1336
	call sub_0575h		;1338
	ld bc,l03ech		;133b
	call sub_04b1h		;133e
	ld a,(02a05h)		;1341
	add a,a			;1344
	add a,a			;1345
	ld c,a			;1346
	ld b,000h		;1347
	ld hl,00199h		;1349
	add hl,bc			;134c
	ld b,h			;134d
	ld c,l			;134e
	call sub_04b1h		;134f
	ld hl,(029f9h)		;1352
	inc hl			;1355
	ld (029f9h),hl		;1356
	jp l1299h		;1359
l135ch:
	ret			;135c
sub_135dh:
	ld hl,(029f7h)		;135d
	add hl,hl			;1360
	add hl,hl			;1361
	add hl,hl			;1362
	add hl,hl			;1363
	ld de,02a08h		;1364
	add hl,de			;1367
	ld (029f5h),hl		;1368
	ret			;136b
sub_136ch:
	ld a,(0006dh)		;136c
	cp 020h		;136f
	jp nz,l1377h		;1371
	ld a,000h		;1374
	ret			;1376
l1377h:
	ld l,004h		;1377
	ld de,01592h		;1379
	ld bc,0006eh		;137c
l137fh:
	ld a,(bc)			;137f
	ld (de),a			;1380
	inc bc			;1381
	inc de			;1382
	dec l			;1383
	jp nz,l137fh		;1384
	ld a,(01592h)		;1387
	sub 053h		;138a
	sub 001h		;138c
	sbc a,a			;138e
	push af			;138f
	ld a,(01593h)		;1390
	sub 020h		;1393
	sub 001h		;1395
	sbc a,a			;1397
	pop bc			;1398
	ld c,b			;1399
	and c			;139a
	rra			;139b
	jp nc,l13a7h		;139c
	ld hl,01556h		;139f
	ld (hl),001h		;13a2
	ld a,0feh		;13a4
	ret			;13a6
l13a7h:
	ld e,004h		;13a7
	ld bc,001a9h		;13a9
	call sub_0a19h		;13ac
	ld (02a05h),a		;13af
	cp 000h		;13b2
	jp nz,l13bdh		;13b4
	ld bc,l038eh		;13b7
	call sub_04d2h		;13ba
l13bdh:
	ld a,001h		;13bd
	ret			;13bf
sub_13c0h:
	ld hl,02a06h		;13c0
	ld (hl),001h		;13c3
l13c5h:
	ld a,00bh		;13c5
	ld hl,02a06h		;13c7
	cp (hl)			;13ca
	jp c,l1401h		;13cb
	ld hl,(02a06h)		;13ce
	ld h,000h		;13d1
	ex de,hl			;13d3
	ld hl,(029f5h)		;13d4
	add hl,de			;13d7
	ld a,07fh		;13d8
	and (hl)			;13da
	ld (02a07h),a		;13db
	cp 020h		;13de
	jp z,l13f7h		;13e0
	ld a,(02a06h)		;13e3
	cp 009h		;13e6
	jp nz,l13f0h		;13e8
	ld c,02eh		;13eb
	call sub_0490h		;13ed
l13f0h:
	ld hl,(02a07h)		;13f0
	ld c,l			;13f3
	call sub_0490h		;13f4
l13f7h:
	ld a,(02a06h)		;13f7
	inc a			;13fa
	ld (02a06h),a		;13fb
	jp nz,l13c5h		;13fe
l1401h:
	ret			;1401
sub_1402h:
	call l0639h		;1402
	call l0639h		;1405
	ld a,(01592h)		;1408
	cp 03dh		;140b
	jp nz,l142fh		;140d
	call l0639h		;1410
	ld bc,003f5h		;1413
	call sub_05feh		;1416
	rra			;1419
	jp nc,l1426h		;141a
	call sub_0da9h		;141d
	call sub_0563h		;1420
	jp l142ch		;1423
l1426h:
	ld bc,003f9h		;1426
	call sub_04d2h		;1429
l142ch:
	jp l1448h		;142c
l142fh:
	call sub_0da9h		;142f
	ld e,008h		;1432
	ld bc,00139h		;1434
	call sub_0a19h		;1437
	cp 008h		;143a
	jp nz,l1445h		;143c
	call sub_08c3h		;143f
	jp l1448h		;1442
l1445h:
	call sub_0dbah		;1445
l1448h:
	ret			;1448
	ld l,c			;1449
	ld h,b			;144a
sub_144bh:
	ld c,(hl)			;144b
	inc hl			;144c
	ld b,(hl)			;144d
	ld a,(de)			;144e
	add a,c			;144f
	ld l,a			;1450
	inc de			;1451
	ld a,(de)			;1452
	adc a,b			;1453
	ld h,a			;1454
	ret			;1455
sub_1456h:
	ex de,hl			;1456
	ld e,a			;1457
	ld d,000h		;1458
	ex de,hl			;145a
	ld a,(de)			;145b
	add a,l			;145c
	ld l,a			;145d
	inc de			;145e
	ld a,(de)			;145f
	adc a,h			;1460
	ld h,a			;1461
	ret			;1462
sub_1463h:
	ex de,hl			;1463
	ld e,a			;1464
	ld d,000h		;1465
	ex de,hl			;1467
	ld a,(de)			;1468
	and l			;1469
	ld l,a			;146a
	inc de			;146b
	ld a,(de)			;146c
	and h			;146d
	ld h,a			;146e
	ret			;146f
sub_1470h:
	ld b,h			;1470
	ld c,l			;1471
sub_1472h:
	ld hl,00000h		;1472
	ld a,010h		;1475
l1477h:
	push af			;1477
	add hl,hl			;1478
	ex de,hl			;1479
	sub a			;147a
	add hl,hl			;147b
	ex de,hl			;147c
	adc a,l			;147d
	sub c			;147e
	ld l,a			;147f
	ld a,h			;1480
	sbc a,b			;1481
	ld h,a			;1482
	inc de			;1483
	jp nc,l1489h		;1484
	add hl,bc			;1487
	dec de			;1488
l1489h:
	pop af			;1489
	dec a			;148a
	jp nz,l1477h		;148b
	ret			;148e
sub_148fh:
	ld b,h			;148f
	ld c,l			;1490
sub_1491h:
	ld hl,00000h		;1491
	ld a,010h		;1494
l1496h:
	add hl,hl			;1496
	ex de,hl			;1497
	add hl,hl			;1498
	ex de,hl			;1499
	jp nc,l149eh		;149a
	add hl,bc			;149d
l149eh:
	dec a			;149e
	jp nz,l1496h		;149f
	ret			;14a2
sub_14a3h:
	ld a,(hl)			;14a3
l14a4h:
	rlca			;14a4
	dec c			;14a5
	jp nz,l14a4h		;14a6
	ret			;14a9
	ld e,(hl)			;14aa
	inc hl			;14ab
	ld d,(hl)			;14ac
	ex de,hl			;14ad
sub_14aeh:
	add hl,hl			;14ae
	dec c			;14af
	jp nz,sub_14aeh		;14b0
	ret			;14b3
sub_14b4h:
	ld e,(hl)			;14b4
	inc hl			;14b5
	ld d,(hl)			;14b6
	ex de,hl			;14b7
l14b8h:
	ld a,h			;14b8
	or a			;14b9
	rra			;14ba
	ld h,a			;14bb
	ld a,l			;14bc
	rra			;14bd
	ld l,a			;14be
	dec c			;14bf
	jp nz,l14b8h		;14c0
	ret			;14c3
sub_14c4h:
	ld e,a			;14c4
	ld d,000h		;14c5
	ld a,e			;14c7
	sub l			;14c8
	ld l,a			;14c9
	ld a,d			;14ca
	sbc a,h			;14cb
	ld h,a			;14cc
	ret			;14cd
sub_14ceh:
	ld l,c			;14ce
	ld h,b			;14cf
	ld c,(hl)			;14d0
	inc hl			;14d1
	ld b,(hl)			;14d2
sub_14d3h:
	ld a,(de)			;14d3
	sub c			;14d4
	ld l,a			;14d5
	inc de			;14d6
	ld a,(de)			;14d7
	sbc a,b			;14d8
	ld h,a			;14d9
	ret			;14da
sub_14dbh:
	ld l,a			;14db
	ld h,000h		;14dc
	ld a,(de)			;14de
	sub l			;14df
	ld l,a			;14e0
	inc de			;14e1
	ld a,(de)			;14e2
	sbc a,h			;14e3
	ld h,a			;14e4
	ret			;14e5
sub_14e6h:
	ld e,a			;14e6
	ld d,000h		;14e7
sub_14e9h:
	ld a,e			;14e9
	sub (hl)			;14ea
	ld e,a			;14eb
	ld a,d			;14ec
	inc hl			;14ed
	sbc a,(hl)			;14ee
	ld d,a			;14ef
	ex de,hl			;14f0
	ret			;14f1
	ld a,(de)			;14f2
	ld a,(de)			;14f3
	ld a,(de)			;14f4
	ld a,(de)			;14f5
	ld a,(de)			;14f6
	ld a,(de)			;14f7
	ld a,(de)			;14f8
	ld a,(de)			;14f9
	ld a,(de)			;14fa
	ld a,(de)			;14fb
	ld a,(de)			;14fc
	ld a,(de)			;14fd
	ld a,(de)			;14fe
	ld a,(de)			;14ff
