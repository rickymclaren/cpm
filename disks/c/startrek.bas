10 REM Copyright 1982 by The Wizard of OsZ
20 REM 'Super StarTrek' adapted from Lyn Cochran 'STAR TREK'
30 REM for Osborne I computer by Microcosm Associates.
40 REM
50 REM
60 CLEAR 500
70 DIM D(7),K(2,2),S(7,7),Q(7,7)
80 PRINT CHR$(26);TAB(10)"The Wizard of OsZ - 'Super StarTrek'"
90 PRINT TAB(10)"  (C) 1982, by The Wizard of OsZ"
100 PRINT:PRINT:RANDOMIZE:PRINT:T0=INT(RND(1)*20+20)*100:T=T0:E=3000:P=10:S=0:S9=20
110 DEF FND(D)=SQR((K(I,0)-S1)^2+(K(I,1)-S2)^2)
120 GOSUB 2310:Q2=Y2:Q1=X2:GOSUB 2310:S2=Y2:S1=X2:FOR I=0 TO 7:D(I)=0:NEXT I
130 Q$="   <*>+++>!< * ":A$="NASRLRPHTOSHDACORE":H$="I  II IIIIV "
140 D$="Warp EnginesS.R. SensorsL.R. SensorsPhaser Cntrl"
150 D$=D$+"Photon TubesDamage CntrlShield CntrlComputer"
160 G$="Antares Sirius Rigel Deneb Procyon Capella Vega "
170 G$=G$+"Betelgeuse Canopus Aldebaran Altair Regulus "
180 G$=G$+"Sagittarius Arcturus Pollux Spica "
190 B9=0:K9=0:FOR I=0 TO 7:FOR J=0 TO 7:N=RND(1):K3=0-(N>.8)-(N>.95)-(N>.98)
200 K9=K9+K3:B3=0:N=RND(1):IF N>.96 THEN B3=1:B9=B9+1
210 S3=INT(RND(1)*8+1):Q(I,J)=-(K3*100+B3*10+S3)
220 NEXT J,I:IF B9=0 THEN Q(5,2)=-114:B9=1:K9=K9+1
230 K0=K9:PRINT"Your orders are as follows:":PRINT
240 PRINT
250 PRINT"Destroy the"K9"Klingon warships (marked '+++') which" 
260 PRINT"have  invaded the Galaxy before they can attack  the" 
270 PRINT"Federation Headquarters by StarDate  -"T0+30"....this" 
280 PRINT"gives you 30 days to complete your assigned mission." 
290 PRINT"There   are"B9"  StarBases  located  throughout  the" 
300 PRINT"Galaxy for resupplying your ship..."
310 PRINT:PRINT:INPUT"Are you ready to assume command";S$
320 IF LEFT$(S$,1)="N" THEN 10
330 PRINT:PRINT:INPUT"Do you require operational instructions Captain";S$
340 IF LEFT$(S$,1)="Y" THEN GOSUB 2580
350 K3=0:B3=0:S3=0:GOSUB 2360:IF T<>T0 THEN 380
360 PRINT CHR$(26):PRINT"Your mission begins with your StarShip located"
370 PRINT"In the Galactic Quadrant, "S$:GOTO 390
380 PRINT"Now entering "S$" Quadrant..."
390 Q(Q1,Q2)=ABS(Q(Q1,Q2)):N=Q(Q1,Q2):X=N*.01
400 K3=INT(X):B3=INT((X-K3)*10):S3=N-INT(N*.1)*10
410 IF S<200 THEN PRINT"...Watch-out, your shields are dangerously low!
420 FOR I=0 TO 7:FOR J=0 TO 7:S(I,J)=1:K(INT(I/3),INT(J/3))=0:NEXT J,I
430 GOSUB 2310:S1=Y2:S2=X2:S(Y2,X2)=2:IF K3=0 THEN 460
440 FOR I=0 TO K3-1:GOSUB 2290:S(Y2,X2)=3:K(I,0)=Y2:K(I,1)=X2:K(I,2)=200
450 NEXT
460 IF B3<>0 THEN FOR I=1 TO B3:GOSUB 2290:S(Y2,X2)=4:BY=Y2:BX=X2:NEXT I
470 FOR I=1 TO S3:GOSUB 2290:S(Y2,X2)=5:NEXT I
480 GOSUB 1770
490 IF S+E>10 AND(E>10 OR D(6)=0)THEN 550
500 PRINT"Oops...Fatal Error!  You've just stranded your  ship" 
510 PRINT"in space!  You have insufficient maneuvering energy," 
520 PRINT"and  Shield Control is presently incapable of cross-"
530 PRINT"circuiting to Engine Room..."
540 GOTO 1660
550 INPUT"Command:";S$:FOR I=1 TO 9
560 IF LEFT$(S$,2)=MID$(A$,(I-1)*2+1,2)THEN 590
570 NEXT I:PRINT"Science Officer Spock reports:"
580 PRINT"'That's TOTALLY ILLOGICAL, Captain!'":GOTO 490
590 ON I GOTO 600,480,980,1050,1220,1380,1450,1980,1660
600 INPUT"Course (1 to 8.9):";C:IF C>=1 AND C<9 THEN 620
610 PRINT"Lt. Sulu reports, 'Incorrect course data, Sir!'":GOTO 490
620 INPUT"Warp (0 to 12):";W:IF W>0 AND W<=12 THEN 650
630 PRINT"Chief Engineer Scott reports, 'The engines won't";
640 PRINT" Take Warp"W"!'":GOTO 600
650 IF D(0)=0 OR W<=.2 THEN 670
660 PRINT"Warp Engines are damaged, maximum speed = Warp 0.2!":GOTO 600
670 N=INT(W*8+.5):IF E-N>=0 THEN 730
680 PRINT"Engineering reports,":PRINT"'Insufficient energy available "
690 PRINT"for maneuvering at Warp"W"!'":IF S<N-E OR D(6)>0 THEN 490
700 PRINT"Deflector Control Room acknowledges,"
710 PRINT"'"S"units of energy presently deployed"
720 PRINT"to shields at this time...'":GOTO 490
730 IF K3<=0 THEN 770
740 FOR I=0 TO 2:IF K(I,2)>0 THEN S(K(I,0),K(I,1))=1:GOSUB 2290:S(Y2,X2)=3
750 K(I,0)=Y2:K(I,1)=X2:NEXT I
760 GOSUB 1560
770 J=-1:GOSUB 2420:IF RND(1)>.2 THEN 800
780 IF RND(1)<.6 THEN J=1
790 GOSUB 2430
800 S(S1,S2)=1:Y1=S1+.5:X1=S2+.5:N=INT(W*8)
810 FOR I=1 TO N:GOSUB 2320:IF J=0 THEN 870
820 IF S(Y2,X2)=1 THEN NEXT I:S1=Y2:S2=X2:GOTO 860
830 S1=INT(Y1-Y):S2=INT(X1-X):PRINT"Warp Engines shut down at Sector";
840 PRINT S1+1"-"S2+1"due to bad"
850 PRINT"navigation! Keep your eyes open!!!"
860 S(S1,S2)=2:GOSUB 2230:GOTO 480
870 X1=Q1:X2=Q2:Q1=INT(Q1+W*Y+(S1+.5)/8):Q2=INT(Q2+W*X+(S2+.5)/8):N=0
880 IF Q1<0 OR Q1>7 THEN N=1:Q1=Q1-(Q1<0)+(Q1>7):S1=Q1:GOTO 880
890 IF Q2<0 OR Q2>7 THEN N=1:Q2=Q2-(Q2<0)+(Q2>7):S2=Q2:GOTO 890
900 IF N=0 THEN 970
910 PRINT"Lt. Uhura reports message from StarFleet Command:"
920 PRINT"'Permission  to   attempt  crossing   of   Galactic" 
930 PRINT"Perimeter is hereby DENIED!  Shut down your engines" 
940 PRINT"at once!!!  You are trespassing beyond the Galactic"
950 PRINT"Perimeter into non-Federation territory...'"
960 IF X1=Q1 AND X2=Q2 THEN 860
970 GOSUB 2230:GOTO 350
980 IF D(2)>0 THEN PRINT"Long Range Sensors are INOPERABLE":GOTO 490
990 PRINT"Long Range Sensor Scan for Quadrant"Q1+1"-"Q2+1
1000 FOR I=Q1-1 TO Q1+1:FOR J=Q2-1 TO Q2+1:PRINT"   ";
1010 IF I<0 OR I>7 OR J<0 OR J>7 THEN PRINT"***";:GOTO 1040
1020 S$=STR$(Q(I,J)):S$="00"+MID$(S$,2):PRINT RIGHT$(S$,3);
1030 IF D(7)=0 THEN Q(I,J)=ABS(Q(I,J))
1040 NEXT J:PRINT:NEXT I:GOTO 490
1050 IF K3>0 THEN 1080
1060 PRINT"Science Officer Spock reports:"
1070 PRINT"'Sensors show no Klingon Ships in this Quadrant.'":GOTO 490
1080 IF D(3)>0 THEN PRINT"Phasers are INOPERATIVE":GOTO 490
1090 IF D(7)>0 THEN PRINT"Computer failure hampers ACCURACY!"
1100 PRINT"Phasers locked on target, energy available ="E
1110 INPUT"How many unit of energy to fire";W:IF W<=0 THEN 490
1120 IF E-W<0 THEN 1100
1130 E=E-W:GOSUB 1560:IF D(7)>0 THEN W=W*RND(1)
1140 FOR I=0 TO 2:IF K(I,2)<=0 THEN 1210
1150 H=INT(W/K3):H=INT((W/FND(0))*(RND(1)+2)):IF H>.15*K(I,2)THEN 1170
1160 PRINT"Sensors show no damage to enemy at"K(I,0)+1"-"K(I,1)+1:GOTO 120
1170 K(I,2)=K(I,2)-H
1180 PRINT H"Energy Unit hit on Klingon at Sector"K(I,0)+1"-"K(I,1)+1
1190 IF K(I,2)>0 THEN PRINT"  ("K(I,2)"Energy Units remaining)":GOTO 1210
1200 GOSUB 2530
1210 NEXT I:GOTO 490
1220 IF D(4)>0 THEN PRINT"Photon Tubes are NOT OPERATIONAL":GOTO 490
1230 IF P=0 THEN PRINT"All Photon Torpedoes EXPENDED":GOTO 490
1240 INPUT"Torpedo course (1 to 8.9)";C:IF C>=1 AND C<9 THEN 1260
1250 PRINT"Ensign Chekov reports: 'Incorrect course data, Captain!'":GOTO 490
1260 P=P-1:PRINT"Torpedo tracking:":Y1=S1+.5:X1=S2+.5
1270 GOSUB 2320:IF J=0 THEN PRINT"Oops...Torpedo missed!":GOTO 1370
1280 PRINT TAB(15)Y2+1"-"X2+1:ON S(Y2,X2)GOTO 1270,0,1290,1310,1360
1290 FOR I=0 TO 2:IF INT(Y2)<>K(I,0)OR INT(X2)<>K(I,1)THEN NEXT I
1300 GOSUB 2530:GOTO 1370
1310 FOR QQ=1 TO 5000:NEXT QQ:PRINT"What the...? You destroyed a StarBase!!!":B3=B3-1:B9=B9-1:GOSUB 2550
1320 IF B3>0 THEN 1370
1330 PRINT CHR$(26):PRINT"That does it Captain...  You are hereby relieved of"
1340 PRINT"command of the StarShip Enterprise!!!"
1350 PRINT:GOTO 1660
1360 IF S(Y2,X2)=5 THEN PRINT"Star at"Y2+1"-"X2+1"absorbed Torpedo energy!"
1370 GOSUB 1560:GOTO 490
1380 IF D(6)>0 THEN PRINT"Shield Control is non-operational at this time!":GOTO 490
1390 PRINT"Energy available ="E+S;
1400 PRINT:INPUT" Number of Energy Units to Shields:";N:IF N>0 AND S<>N THEN 1420
1410 PRINT"Shield Energy level is UNCHANGED":GOTO 490
1420 IF E+S-N<0 THEN 1390
1430 E=E+S-N:S=N:PRINT"Deflector Control Room Report:"
1440 PRINT"'Shields now at"S"per your command, Sir'":GOTO 490
1450 IF D(5)>0 THEN PRINT"Damage Control report is not available":GOTO 490
1460 N=0:FOR I=0 TO 7:IF D(I)>0 THEN N=N+.1
1470 NEXT I:IF N=0 OR C$<>"Docked"THEN 1540
1480 N=N+.5*RND(1):IF N>=1 THEN N=.9
1490 PRINT"Technicians are standing by to effect repairs to the"
1500 PRINT"Enterprise; estimated time to repair: "N"StarDates."
1510 PRINT:INPUT"Will you authorize the Repair Order, Captain";S$
1520 IF LEFT$(S$,1)<>"Y"THEN 490
1530 T=T+N+.1:J=-100:GOSUB 2420:GOTO 490
1540 PRINT:PRINT"Device"TAB(14)"state of repair"
1550 J=0:GOSUB 2420:PRINT:GOTO 490
1560 IF K3<=0 THEN RETURN
1570 IF C$="Docked"THEN PRINT"StarBase shields the Enterprise":RETURN
1580 FOR I=0 TO 2:IF K(I,2)<=0 THEN 1650
1590 H=FND(0):H=INT(K(I,2)/H):H=H*(2+RND(1)):S=S-H
1600 PRINT H"Unit hit on Enterprise from Sector"K(I,0)+1"-"K(I,1)+1
1610 IF S<0 THEN 1670
1620 PRINT"  (Shields down to"S"Units.)"
1630 IF H<20 OR RND(1)>.6 OR H/S<.02 THEN 1650
1640 J=H/S+.5*RND(1):GOSUB 2430
1650 NEXT I:RETURN
1660 PRINT"It is StarDate"T:GOTO 1700
1670 PRINT"The Enterprise has been DESTROYED!!!  The Federation" 
1680 PRINT"will  now  be  summarily conquered  by  the  Klingon" 
1690 PRINT"Empire, you TWIT! You would do better as Ships Cook!"
1700 PRINT:PRINT"There  were"K9"Battle Cruisers, left at the  end  of"
1710 PRINT"your mission...":PRINT:PRINT 
1720 INPUT"Want another chance to save the Federation";S$:IF LEFT$(S$,1)="Y"THEN 80
1730 END
1740 PRINT"Congratulations,  Captain!  The last Klingon Battle" 
1750 PRINT"Cruiser menacing the Federation has been destroyed."
1760 PRINT"Your StarFleet Rating is"INT(K0/(T-T0)*1000):GOTO 1720
1770 FOR I=S1-1 TO S1+1:FOR J=S2-1 TO S2+1:IF I<0 OR I>7 OR J<0 OR J>7 THEN 1790
1780 IF S(I,J)=4 THEN 1800
1790 NEXT J,I:GOTO 1820
1800 C$="Docked":E=3000:P=10
1810 PRINT"Shields dropped for docking purposes":S=0:GOTO 1850
1820 IF K3>0 THEN C$="Red":GOTO 1850
1830 IF E<300 THEN C$="Yellow":GOTO 1850
1840 C$="Green"
1850 IF D(1)>0 THEN PRINT"Oops, Short Range Sensors are OUT!":RETURN
1860 S$="---------------------------------":PRINT S$
1870 FOR I=0 TO 7:FOR J=0 TO 7
1880 PRINT MID$(Q$,(S(I,J)-1)*3+1,3);" ";:NEXT J:PRINT TAB(35)
1890 IF I=0 THEN PRINT" "
1900 IF I=1 THEN PRINT"StarDate"TAB(45)T
1910 IF I=2 THEN PRINT"Condition"TAB(46)C$
1920 IF I=3 THEN PRINT"Quadrant"TAB(45)Q1+1"-"Q2+1
1930 IF I=4 THEN PRINT"Sector"TAB(45)S1+1"-"S2+1
1940 IF I=5 THEN PRINT"Energy"TAB(45)E
1950 IF I=6 THEN PRINT"Torpedoes"TAB(45)P
1960 IF I=7 THEN PRINT"Shields"TAB(45)S
1970 NEXT I:PRINT S$:RETURN
1980 J=0:IF D(7)>0 THEN PRINT"Computer disabled!":GOTO 490
1990 INPUT"Computer active and awaiting command:";N
2000 IF N>0 AND N<6 THEN ON N GOTO 2030,2070,2090,2110,2140
2010 PRINT"Ships computer reports:"
2020 PRINT"'I don't understand that command, Sir!'":GOTO 490
2030 PRINT"Computer Record of the Galaxy for Quadrant "Q1+1"-"Q2+1
2040 FOR I=0 TO 7:FOR J=0 TO 7:PRINT"  ";:IF Q(I,J)<0 THEN PRINT"***";:GOTO 2060
2050 S$=STR$(ABS(Q(I,J))):S$="00"+MID$(S$,2):PRINT RIGHT$(S$,3);
2060 NEXT J:PRINT:NEXT I:GOTO 490
2070 PRINT"  Status Report:":PRINT K9"Klingons left"
2080 PRINT T0+30-T"StarDates left:":PRINT B9"StarBases left:":GOTO 1450
2090 PRINT:J=1:FOR I=0 TO 2:IF K(I,2)<=0 THEN 2220
2100 Y1=S1:X1=S2:Y2=K(I,0):X2=K(I,1):GOTO 2150
2110 Y1=S1:X1=S2:Y2=BY:X2=BX:IF B3>0 THEN 2150
2120 PRINT"Mr. Spock reports, 'Sensors show no StarBases in this Quadrant'"
2130 GOTO 490
2140 INPUT"Initial coordinates";Y1,X1:INPUT"Final coordinates";Y2,X2
2150 Y2=Y1-Y2:X2=X2-X1:N=0:C=7:IF X2<0 THEN C=5:GOTO 2190
2160 IF X2>0 AND Y2>=0 THEN C=1:GOTO 2190
2170 IF X2>0 THEN C=9:GOTO 2190
2180 IF Y2>0 THEN C=3
2190 IF Y2<>0 AND X2<>0 THEN N=1.27324*ATN(Y2/X2)
2200 PRINT"Direction ="C+N
2210 PRINT"Distance ="SQR(Y2^2+X2^2):IF J=0 THEN 490
2220 NEXT I:GOTO 490
2230 E=E-INT(W*8+.5)-10:J=1:IF E>=0 THEN 2260
2240 PRINT"Shield Control supplied energy to complete the maneuver."
2250 S=S+E:E=0:IF S<=0 THEN S=0
2260 IF W<1 THEN J=.1*INT(10*W)
2270 T=T+J:IF T>T0+30 THEN 1660
2280 RETURN
2290 GOSUB 2310:IF S(Y2,X2)<>1 THEN 2290
2300 RETURN
2310 Y2=INT(RND(1)*8):X2=INT(RND(1)*8):RETURN
2320 J=1:Y=(C-1)*.785398:X=COS(Y):Y=-SIN(Y)
2330 Y1=Y1+Y:X1=X1+X:Y2=INT(Y1):X2=INT(X1)
2340 IF X2<0 OR X2>7 OR Y2<0 OR Y2>7 THEN J=0
2350 RETURN
2360 J=0:N=1:FOR I=1 TO LEN(G$):IF MID$(G$,I,1)<>" "THEN 2390
2370 IF J=Q1*2-(Q2>3)THEN 2400
2380 N=I+1:J=J+1
2390 NEXT I
2400 S$=MID$(G$,N,I-N)
2410 N=(Q2+4*(Q2>3))*3+1:S$=S$+" "+MID$(H$,N,3):RETURN
2420 FOR X=0 TO 7:N=X:GOTO 2440
2430 N=INT(RND(1)*8):X=7
2440 IF J=0 THEN PRINT MID$(D$,N*12+1,11),D(N):GOTO 2510
2450 Y=D(N):D(N)=Y+J:IF D(N)<0 THEN D(N)=0
2460 IF(J>0 AND Y>0)OR(J<0 AND Y=0)THEN 2510
2470 PRINT"Damage Control report:";:PRINT"  "MID$(D$,N*12+1,11)" ";
2480 IF J>0 THEN PRINT"DAMAGED":GOTO 2510
2490 IF D(N)=0 THEN PRINT"Repair COMPLETED":GOTO 2510
2500 PRINT"repair IMPROVED"
2510 IF X<>7 THEN NEXT
2520 RETURN
2530 PRINT CHR$(7)"> Klingon destroyed <":K3=K3-1:K9=K9-1:Y2=K(I,0):X2=K(I,1)
2540 K(I,2)=0:IF K9<1 THEN 1740
2550 N=SGN(Q(Q1,Q2)):Q(Q1,Q2)=N*(K3*100+B3*10+S3):S(Y2,X2)=1
2560 RETURN
2570 END
2580 PRINT CHR$(26):PRINT
2590 PRINT "     At  the  'Command:' prompt,  you may enter  the"
2600 PRINT "following  directives to the on-board ship computer,"
2610 PRINT "navigational  control,  and  weapon system  for  the"
2620 PRINT "StarShip Enterprise:"
2630 GOSUB 3490
2640 PRINT "CO  - Computer,  select the following  commands  for"
2650 PRINT "computational evaluation of Records, Status, Klingon"
2660 PRINT "and StarBase locations, and Hyper-Warp Navigation."
2670 PRINT
2680 PRINT "     1 - Computer Record of the Galaxy (ref: 'LR')."
2690 PRINT "     2 - Status Report (ref: 'SR')."
2700 PRINT "     3 - Klingon location (direction and distance)."
2710 PRINT "     4 - StarBase location     -      -     -     ."
2720 PRINT "     5 - Hyper-Warp (start and end coordinates)."
2730 GOSUB 3490
2740 PRINT "NA  - Navigate,  allows directional control and warp"
2750 PRINT "factor in the range of 1 to 8.9 for direction, and 1"
2760 PRINT "to 12 for warp factor. Direction is indicated by the"
2770 PRINT "following diagramatic representation:"
2780 PRINT
2790 PRINT "                        3.0"
2800 PRINT "                        --- "
2810 PRINT "           4.0           |         2.0"
2820 PRINT "              \          |        /"
2830 PRINT "                         |"
2840 PRINT "                         |"
2850 PRINT "                         |"
2860 PRINT                   
2870 PRINT "      5.0 |------------ <*> ------------| 1.0"
2880 PRINT "                                         (8.9)"
2890 PRINT "                         |"
2900 PRINT "                         |"
2910 PRINT "                         |"
2920 PRINT "              /          |        \"
2930 PRINT "           6.0           |         8.0"
2940 PRINT "                        ---"    
2950 PRINT "                        7.0"
2960 GOSUB 3490
2970 PRINT "PH  - Phasor  control,  allows diversion  of  Energy" 
2980 PRINT "Units  from Phasor Bank,  maneuvering power for Warp" 
2990 PRINT "Engines,  or Shield Control storage allocation.  Use" 
3000 PRINT "only as much energy as can be safely allocated  from" 
3010 PRINT "Warp Engines,  or you will be stranded with no means" 
3020 PRINT "of returning to a StarBase for re-supply."
3030 GOSUB 3490
3040 PRINT "TO  - Photon  Torpedos may be used when you  have  a" 
3050 PRINT "'clear  shot'  (no stars in the way!) at  a  Klingon" 
3060 PRINT "Battle Cruiser."
3070 GOSUB 3490
3080 PRINT "SH - Shield Control,  allocates energy (and at times" 
3090 PRINT "diverted  automatically) from available Energy  Unit" 
3100 PRINT "resources.  Warning:  It  is  recommended  that  you" 
3110 PRINT "allocate some minimum amount of Energy Units   PRIOR" 
3120 PRINT "to encountering a Klingon Battle Cruiser..."
3130 GOSUB 3490
3140 PRINT "LR - Long Range Sensor Scan, depicts the surrounding" 
3150 PRINT "eight Quadrants,  with the Enterprise located in the" 
3160 PRINT "center.  Significant digits represent (from left  to" 
3170 PRINT "right),  (1)  the number of Klingon Battle Cruisers," 
3180 PRINT "(2) the number of StarBases,  and (3) the number  of" 
3190 PRINT "Stars.  For example:"
3200 PRINT
3210 PRINT"           1 Star         1 StarBase"
3220 PRINT"                 \      /"
3230 PRINT" 3 Klingons --> 301   018   *** <-- Edge of Galaxy"
3240 PRINT"                003   008   ***"
3250 PRINT" 1 Klingon  --> 108   001   ***"
3260 PRINT"                 /       \"
3270 PRINT"          8 Stars          1 Star"
3280 PRINT
3290 PRINT"The  StarShip Enterprise is in the middle,  with '0'" 
3300 PRINT"Klingons,  '0' Starbases, and '8' Stars. The edge of" 
3310 PRINT"the Galxay is off to the right.  When displaying the" 
3320 PRINT"Galactic Map, '***' represents unexplored Quadrants."
3330 GOSUB 3490
3340 PRINT "SR  - Short Range Sensor Scan,  depicts the  current" 
3350 PRINT "Quadrant and status in which you have the Enterprise" 
3360 PRINT "located." 
3370 GOSUB 3490
3380 PRINT "     Klingon   Battle  Cruisers  are  identified  as" 
3390 PRINT "'+++',  StarBases are '>!<,  Stars are '*',  and the" 
3400 PRINT "StarShip Enterprise as '<*>'."  
3410 PRINT
3420 PRINT "     Your aids Mr. Spock, Lt. Uharu, Engineer Scott," 
3430 PRINT "and Navigator Chekov, are along on this mission, for" 
3440 PRINT "your help in battling the Klingon Empire hoard..."
3450 GOSUB 3490
3460 PRINT "     The  Federation  hopes that you can  save  them" 
3470 PRINT "from the terrible wrath of the Klingon Empire...good" 
3480 PRINT "hunting Captain!"
3490 PRINT:INPUT"Press RETURN to continue:",QQ$:PRINT CHR$(26):RETURN
Klingon Empire...good" 
3480 PRINT "hunting Captain!"
3