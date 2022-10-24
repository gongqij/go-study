"".main STEXT size=185 args=0x0 locals=0x48 funcid=0x0 align=0x0
	0x0000 00000 (main.go:8)	TEXT	"".main(SB), ABIInternal, $72-0
	0x0000 00000 (main.go:8)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:8)	PCDATA	$0, $-2
	0x0004 00004 (main.go:8)	JLS	175
	0x000a 00010 (main.go:8)	PCDATA	$0, $-1
	0x000a 00010 (main.go:8)	SUBQ	$72, SP
	0x000e 00014 (main.go:8)	MOVQ	BP, 64(SP)
	0x0013 00019 (main.go:8)	LEAQ	64(SP), BP
	0x0018 00024 (main.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0018 00024 (main.go:8)	FUNCDATA	$1, gclocals·54241e171da8af6ae173d69da0236748(SB)
	0x0018 00024 (main.go:9)	LEAQ	""..dict.echo[int](SB), AX
	0x001f 00031 (main.go:9)	XORL	BX, BX
	0x0021 00033 (main.go:9)	PCDATA	$1, $0
	0x0021 00033 (main.go:9)	CALL	"".echo[go.shape.int_0](SB)
	0x0026 00038 (main.go:10)	LEAQ	""..dict.echo[int32](SB), AX
	0x002d 00045 (main.go:10)	XORL	BX, BX
	0x002f 00047 (main.go:10)	CALL	"".echo[go.shape.int32_0](SB)
	0x0034 00052 (main.go:11)	LEAQ	""..dict.echo[uint32](SB), AX
	0x003b 00059 (main.go:11)	XORL	BX, BX
	0x003d 00061 (main.go:11)	NOP
	0x0040 00064 (main.go:11)	CALL	"".echo[go.shape.uint32_0](SB)
	0x0045 00069 (main.go:12)	LEAQ	""..dict.echo[uint64](SB), AX
	0x004c 00076 (main.go:12)	XORL	BX, BX
	0x004e 00078 (main.go:12)	CALL	"".echo[go.shape.uint64_0](SB)
	0x0053 00083 (main.go:13)	LEAQ	""..dict.echo[string](SB), AX
	0x005a 00090 (main.go:13)	LEAQ	go.string."hello"(SB), BX
	0x0061 00097 (main.go:13)	MOVL	$5, CX
	0x0066 00102 (main.go:13)	CALL	"".echo[go.shape.string_0](SB)
	0x006b 00107 (main.go:14)	LEAQ	""..dict.echo[struct {}](SB), AX
	0x0072 00114 (main.go:14)	CALL	"".echo[go.shape.struct {}_0](SB)
	0x0077 00119 (main.go:15)	CALL	time.Now(SB)
	0x007c 00124 (main.go:15)	MOVQ	AX, ""..autotmp_0+40(SP)
	0x0081 00129 (main.go:15)	MOVQ	BX, ""..autotmp_0+48(SP)
	0x0086 00134 (main.go:15)	MOVQ	CX, ""..autotmp_0+56(SP)
	0x008b 00139 (main.go:15)	MOVQ	CX, DI
	0x008e 00142 (main.go:15)	MOVQ	BX, CX
	0x0091 00145 (main.go:15)	MOVQ	AX, BX
	0x0094 00148 (main.go:15)	LEAQ	""..dict.echo[time.Time](SB), AX
	0x009b 00155 (main.go:15)	NOP
	0x00a0 00160 (main.go:15)	CALL	"".echo[go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0](SB)
	0x00a5 00165 (main.go:16)	MOVQ	64(SP), BP
	0x00aa 00170 (main.go:16)	ADDQ	$72, SP
	0x00ae 00174 (main.go:16)	RET
	0x00af 00175 (main.go:16)	NOP
	0x00af 00175 (main.go:8)	PCDATA	$1, $-1
	0x00af 00175 (main.go:8)	PCDATA	$0, $-2
	0x00af 00175 (main.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x00b4 00180 (main.go:8)	PCDATA	$0, $-1
	0x00b4 00180 (main.go:8)	JMP	0
	0x0000 49 3b 66 10 0f 86 a5 00 00 00 48 83 ec 48 48 89  I;f.......H..HH.
	0x0010 6c 24 40 48 8d 6c 24 40 48 8d 05 00 00 00 00 31  l$@H.l$@H......1
	0x0020 db e8 00 00 00 00 48 8d 05 00 00 00 00 31 db e8  ......H......1..
	0x0030 00 00 00 00 48 8d 05 00 00 00 00 31 db 0f 1f 00  ....H......1....
	0x0040 e8 00 00 00 00 48 8d 05 00 00 00 00 31 db e8 00  .....H......1...
	0x0050 00 00 00 48 8d 05 00 00 00 00 48 8d 1d 00 00 00  ...H......H.....
	0x0060 00 b9 05 00 00 00 e8 00 00 00 00 48 8d 05 00 00  ...........H....
	0x0070 00 00 e8 00 00 00 00 e8 00 00 00 00 48 89 44 24  ............H.D$
	0x0080 28 48 89 5c 24 30 48 89 4c 24 38 48 89 cf 48 89  (H.\$0H.L$8H..H.
	0x0090 d9 48 89 c3 48 8d 05 00 00 00 00 0f 1f 44 00 00  .H..H........D..
	0x00a0 e8 00 00 00 00 48 8b 6c 24 40 48 83 c4 48 c3 e8  .....H.l$@H..H..
	0x00b0 00 00 00 00 e9 47 ff ff ff                       .....G...
	rel 27+4 t=14 ""..dict.echo[int]+0
	rel 34+4 t=7 "".echo[go.shape.int_0]+0
	rel 41+4 t=14 ""..dict.echo[int32]+0
	rel 48+4 t=7 "".echo[go.shape.int32_0]+0
	rel 55+4 t=14 ""..dict.echo[uint32]+0
	rel 65+4 t=7 "".echo[go.shape.uint32_0]+0
	rel 72+4 t=14 ""..dict.echo[uint64]+0
	rel 79+4 t=7 "".echo[go.shape.uint64_0]+0
	rel 86+4 t=14 ""..dict.echo[string]+0
	rel 93+4 t=14 go.string."hello"+0
	rel 103+4 t=7 "".echo[go.shape.string_0]+0
	rel 110+4 t=14 ""..dict.echo[struct {}]+0
	rel 115+4 t=7 "".echo[go.shape.struct {}_0]+0
	rel 120+4 t=7 time.Now+0
	rel 151+4 t=14 ""..dict.echo[time.Time]+0
	rel 161+4 t=7 "".echo[go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0]+0
	rel 176+4 t=7 runtime.morestack_noctxt+0
"".echo[go.shape.int_0] STEXT dupok size=268 args=0x10 locals=0x88 funcid=0x0 align=0x0
	0x0000 00000 (main.go:18)	TEXT	"".echo[go.shape.int_0](SB), DUPOK|ABIInternal, $136-16
	0x0000 00000 (main.go:18)	LEAQ	-8(SP), R12
	0x0005 00005 (main.go:18)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:18)	PCDATA	$0, $-2
	0x0009 00009 (main.go:18)	JLS	238
	0x000f 00015 (main.go:18)	PCDATA	$0, $-1
	0x000f 00015 (main.go:18)	SUBQ	$136, SP
	0x0016 00022 (main.go:18)	MOVQ	BP, 128(SP)
	0x001e 00030 (main.go:18)	LEAQ	128(SP), BP
	0x0026 00038 (main.go:18)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$1, gclocals·731dddc57432221a509e963fc20cf850(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$2, "".echo[go.shape.int_0].stkobj(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$5, "".echo[go.shape.int_0].arginfo1(SB)
	0x0026 00038 (main.go:18)	MOVQ	AX, ""..dict+144(SP)
	0x002e 00046 (main.go:18)	MOVQ	BX, "".t+152(SP)
	0x0036 00054 (main.go:18)	MOVUPS	X15, "".~r0+56(SP)
	0x003c 00060 (main.go:19)	MOVUPS	X15, ""..autotmp_3+88(SP)
	0x0042 00066 (main.go:19)	LEAQ	""..autotmp_3+88(SP), CX
	0x0047 00071 (main.go:19)	MOVQ	CX, ""..autotmp_6+48(SP)
	0x004c 00076 (main.go:19)	MOVQ	"".t+152(SP), AX
	0x0054 00084 (main.go:19)	PCDATA	$1, $1
	0x0054 00084 (main.go:19)	CALL	runtime.convT64(SB)
	0x0059 00089 (main.go:19)	PCDATA	$0, $-2
	0x0059 00089 (main.go:19)	MOVQ	AX, ""..autotmp_7+40(SP)
	0x005e 00094 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x0063 00099 (main.go:19)	TESTB	AL, (CX)
	0x0065 00101 (main.go:19)	MOVQ	""..dict+144(SP), DX
	0x006d 00109 (main.go:19)	PCDATA	$0, $-1
	0x006d 00109 (main.go:19)	TESTB	AL, (DX)
	0x006f 00111 (main.go:19)	MOVQ	(DX), DX
	0x0072 00114 (main.go:19)	MOVQ	DX, (CX)
	0x0075 00117 (main.go:19)	LEAQ	8(CX), DI
	0x0079 00121 (main.go:19)	PCDATA	$0, $-2
	0x0079 00121 (main.go:19)	CMPL	runtime.writeBarrier(SB), $0
	0x0080 00128 (main.go:19)	JEQ	132
	0x0082 00130 (main.go:19)	JMP	138
	0x0084 00132 (main.go:19)	MOVQ	AX, 8(CX)
	0x0088 00136 (main.go:19)	JMP	145
	0x008a 00138 (main.go:19)	CALL	runtime.gcWriteBarrier(SB)
	0x008f 00143 (main.go:19)	JMP	145
	0x0091 00145 (main.go:19)	PCDATA	$0, $-1
	0x0091 00145 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x0096 00150 (main.go:19)	TESTB	AL, (CX)
	0x0098 00152 (main.go:19)	JMP	154
	0x009a 00154 (main.go:19)	MOVQ	CX, ""..autotmp_5+104(SP)
	0x009f 00159 (main.go:19)	MOVQ	$1, ""..autotmp_5+112(SP)
	0x00a8 00168 (main.go:19)	MOVQ	$1, ""..autotmp_5+120(SP)
	0x00b1 00177 (main.go:19)	LEAQ	go.string."%v"(SB), AX
	0x00b8 00184 (main.go:19)	MOVL	$2, BX
	0x00bd 00189 (main.go:19)	MOVL	$1, DI
	0x00c2 00194 (main.go:19)	MOVQ	DI, SI
	0x00c5 00197 (main.go:19)	PCDATA	$1, $0
	0x00c5 00197 (main.go:19)	CALL	fmt.Sprintf(SB)
	0x00ca 00202 (main.go:19)	MOVQ	AX, ""..autotmp_4+72(SP)
	0x00cf 00207 (main.go:19)	MOVQ	BX, ""..autotmp_4+80(SP)
	0x00d4 00212 (main.go:19)	MOVQ	AX, "".~r0+56(SP)
	0x00d9 00217 (main.go:19)	MOVQ	BX, "".~r0+64(SP)
	0x00de 00222 (main.go:19)	MOVQ	128(SP), BP
	0x00e6 00230 (main.go:19)	ADDQ	$136, SP
	0x00ed 00237 (main.go:19)	RET
	0x00ee 00238 (main.go:19)	NOP
	0x00ee 00238 (main.go:18)	PCDATA	$1, $-1
	0x00ee 00238 (main.go:18)	PCDATA	$0, $-2
	0x00ee 00238 (main.go:18)	MOVQ	AX, 8(SP)
	0x00f3 00243 (main.go:18)	MOVQ	BX, 16(SP)
	0x00f8 00248 (main.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x00fd 00253 (main.go:18)	MOVQ	8(SP), AX
	0x0102 00258 (main.go:18)	MOVQ	16(SP), BX
	0x0107 00263 (main.go:18)	PCDATA	$0, $-1
	0x0107 00263 (main.go:18)	JMP	0
	0x0000 4c 8d 64 24 f8 4d 3b 66 10 0f 86 df 00 00 00 48  L.d$.M;f.......H
	0x0010 81 ec 88 00 00 00 48 89 ac 24 80 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 80 00 00 00 48 89 84 24 90 00 00 00 48 89  .$....H..$....H.
	0x0030 9c 24 98 00 00 00 44 0f 11 7c 24 38 44 0f 11 7c  .$....D..|$8D..|
	0x0040 24 58 48 8d 4c 24 58 48 89 4c 24 30 48 8b 84 24  $XH.L$XH.L$0H..$
	0x0050 98 00 00 00 e8 00 00 00 00 48 89 44 24 28 48 8b  .........H.D$(H.
	0x0060 4c 24 30 84 01 48 8b 94 24 90 00 00 00 84 02 48  L$0..H..$......H
	0x0070 8b 12 48 89 11 48 8d 79 08 83 3d 00 00 00 00 00  ..H..H.y..=.....
	0x0080 74 02 eb 06 48 89 41 08 eb 07 e8 00 00 00 00 eb  t...H.A.........
	0x0090 00 48 8b 4c 24 30 84 01 eb 00 48 89 4c 24 68 48  .H.L$0....H.L$hH
	0x00a0 c7 44 24 70 01 00 00 00 48 c7 44 24 78 01 00 00  .D$p....H.D$x...
	0x00b0 00 48 8d 05 00 00 00 00 bb 02 00 00 00 bf 01 00  .H..............
	0x00c0 00 00 48 89 fe e8 00 00 00 00 48 89 44 24 48 48  ..H.......H.D$HH
	0x00d0 89 5c 24 50 48 89 44 24 38 48 89 5c 24 40 48 8b  .\$PH.D$8H.\$@H.
	0x00e0 ac 24 80 00 00 00 48 81 c4 88 00 00 00 c3 48 89  .$....H.......H.
	0x00f0 44 24 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44  D$.H.\$......H.D
	0x0100 24 08 48 8b 5c 24 10 e9 f4 fe ff ff              $.H.\$......
	rel 85+4 t=7 runtime.convT64+0
	rel 123+4 t=14 runtime.writeBarrier+-1
	rel 139+4 t=7 runtime.gcWriteBarrier+0
	rel 180+4 t=14 go.string."%v"+0
	rel 198+4 t=7 fmt.Sprintf+0
	rel 249+4 t=7 runtime.morestack_noctxt+0
"".echo[go.shape.int32_0] STEXT dupok size=266 args=0x10 locals=0x88 funcid=0x0 align=0x0
	0x0000 00000 (main.go:18)	TEXT	"".echo[go.shape.int32_0](SB), DUPOK|ABIInternal, $136-16
	0x0000 00000 (main.go:18)	LEAQ	-8(SP), R12
	0x0005 00005 (main.go:18)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:18)	PCDATA	$0, $-2
	0x0009 00009 (main.go:18)	JLS	238
	0x000f 00015 (main.go:18)	PCDATA	$0, $-1
	0x000f 00015 (main.go:18)	SUBQ	$136, SP
	0x0016 00022 (main.go:18)	MOVQ	BP, 128(SP)
	0x001e 00030 (main.go:18)	LEAQ	128(SP), BP
	0x0026 00038 (main.go:18)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$1, gclocals·731dddc57432221a509e963fc20cf850(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$2, "".echo[go.shape.int32_0].stkobj(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$5, "".echo[go.shape.int32_0].arginfo1(SB)
	0x0026 00038 (main.go:18)	MOVQ	AX, ""..dict+144(SP)
	0x002e 00046 (main.go:18)	MOVL	BX, "".t+152(SP)
	0x0035 00053 (main.go:18)	MOVUPS	X15, "".~r0+56(SP)
	0x003b 00059 (main.go:19)	MOVUPS	X15, ""..autotmp_3+88(SP)
	0x0041 00065 (main.go:19)	LEAQ	""..autotmp_3+88(SP), CX
	0x0046 00070 (main.go:19)	MOVQ	CX, ""..autotmp_6+48(SP)
	0x004b 00075 (main.go:19)	MOVL	"".t+152(SP), AX
	0x0052 00082 (main.go:19)	PCDATA	$1, $1
	0x0052 00082 (main.go:19)	CALL	runtime.convT32(SB)
	0x0057 00087 (main.go:19)	PCDATA	$0, $-2
	0x0057 00087 (main.go:19)	MOVQ	AX, ""..autotmp_7+40(SP)
	0x005c 00092 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x0061 00097 (main.go:19)	TESTB	AL, (CX)
	0x0063 00099 (main.go:19)	MOVQ	""..dict+144(SP), DX
	0x006b 00107 (main.go:19)	PCDATA	$0, $-1
	0x006b 00107 (main.go:19)	TESTB	AL, (DX)
	0x006d 00109 (main.go:19)	MOVQ	(DX), DX
	0x0070 00112 (main.go:19)	MOVQ	DX, (CX)
	0x0073 00115 (main.go:19)	LEAQ	8(CX), DI
	0x0077 00119 (main.go:19)	PCDATA	$0, $-2
	0x0077 00119 (main.go:19)	CMPL	runtime.writeBarrier(SB), $0
	0x007e 00126 (main.go:19)	NOP
	0x0080 00128 (main.go:19)	JEQ	132
	0x0082 00130 (main.go:19)	JMP	138
	0x0084 00132 (main.go:19)	MOVQ	AX, 8(CX)
	0x0088 00136 (main.go:19)	JMP	145
	0x008a 00138 (main.go:19)	CALL	runtime.gcWriteBarrier(SB)
	0x008f 00143 (main.go:19)	JMP	145
	0x0091 00145 (main.go:19)	PCDATA	$0, $-1
	0x0091 00145 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x0096 00150 (main.go:19)	TESTB	AL, (CX)
	0x0098 00152 (main.go:19)	JMP	154
	0x009a 00154 (main.go:19)	MOVQ	CX, ""..autotmp_5+104(SP)
	0x009f 00159 (main.go:19)	MOVQ	$1, ""..autotmp_5+112(SP)
	0x00a8 00168 (main.go:19)	MOVQ	$1, ""..autotmp_5+120(SP)
	0x00b1 00177 (main.go:19)	LEAQ	go.string."%v"(SB), AX
	0x00b8 00184 (main.go:19)	MOVL	$2, BX
	0x00bd 00189 (main.go:19)	MOVL	$1, DI
	0x00c2 00194 (main.go:19)	MOVQ	DI, SI
	0x00c5 00197 (main.go:19)	PCDATA	$1, $0
	0x00c5 00197 (main.go:19)	CALL	fmt.Sprintf(SB)
	0x00ca 00202 (main.go:19)	MOVQ	AX, ""..autotmp_4+72(SP)
	0x00cf 00207 (main.go:19)	MOVQ	BX, ""..autotmp_4+80(SP)
	0x00d4 00212 (main.go:19)	MOVQ	AX, "".~r0+56(SP)
	0x00d9 00217 (main.go:19)	MOVQ	BX, "".~r0+64(SP)
	0x00de 00222 (main.go:19)	MOVQ	128(SP), BP
	0x00e6 00230 (main.go:19)	ADDQ	$136, SP
	0x00ed 00237 (main.go:19)	RET
	0x00ee 00238 (main.go:19)	NOP
	0x00ee 00238 (main.go:18)	PCDATA	$1, $-1
	0x00ee 00238 (main.go:18)	PCDATA	$0, $-2
	0x00ee 00238 (main.go:18)	MOVQ	AX, 8(SP)
	0x00f3 00243 (main.go:18)	MOVL	BX, 16(SP)
	0x00f7 00247 (main.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x00fc 00252 (main.go:18)	MOVQ	8(SP), AX
	0x0101 00257 (main.go:18)	MOVL	16(SP), BX
	0x0105 00261 (main.go:18)	PCDATA	$0, $-1
	0x0105 00261 (main.go:18)	JMP	0
	0x0000 4c 8d 64 24 f8 4d 3b 66 10 0f 86 df 00 00 00 48  L.d$.M;f.......H
	0x0010 81 ec 88 00 00 00 48 89 ac 24 80 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 80 00 00 00 48 89 84 24 90 00 00 00 89 9c  .$....H..$......
	0x0030 24 98 00 00 00 44 0f 11 7c 24 38 44 0f 11 7c 24  $....D..|$8D..|$
	0x0040 58 48 8d 4c 24 58 48 89 4c 24 30 8b 84 24 98 00  XH.L$XH.L$0..$..
	0x0050 00 00 e8 00 00 00 00 48 89 44 24 28 48 8b 4c 24  .......H.D$(H.L$
	0x0060 30 84 01 48 8b 94 24 90 00 00 00 84 02 48 8b 12  0..H..$......H..
	0x0070 48 89 11 48 8d 79 08 83 3d 00 00 00 00 00 66 90  H..H.y..=.....f.
	0x0080 74 02 eb 06 48 89 41 08 eb 07 e8 00 00 00 00 eb  t...H.A.........
	0x0090 00 48 8b 4c 24 30 84 01 eb 00 48 89 4c 24 68 48  .H.L$0....H.L$hH
	0x00a0 c7 44 24 70 01 00 00 00 48 c7 44 24 78 01 00 00  .D$p....H.D$x...
	0x00b0 00 48 8d 05 00 00 00 00 bb 02 00 00 00 bf 01 00  .H..............
	0x00c0 00 00 48 89 fe e8 00 00 00 00 48 89 44 24 48 48  ..H.......H.D$HH
	0x00d0 89 5c 24 50 48 89 44 24 38 48 89 5c 24 40 48 8b  .\$PH.D$8H.\$@H.
	0x00e0 ac 24 80 00 00 00 48 81 c4 88 00 00 00 c3 48 89  .$....H.......H.
	0x00f0 44 24 08 89 5c 24 10 e8 00 00 00 00 48 8b 44 24  D$..\$......H.D$
	0x0100 08 8b 5c 24 10 e9 f6 fe ff ff                    ..\$......
	rel 83+4 t=7 runtime.convT32+0
	rel 121+4 t=14 runtime.writeBarrier+-1
	rel 139+4 t=7 runtime.gcWriteBarrier+0
	rel 180+4 t=14 go.string."%v"+0
	rel 198+4 t=7 fmt.Sprintf+0
	rel 248+4 t=7 runtime.morestack_noctxt+0
"".echo[go.shape.uint32_0] STEXT dupok size=266 args=0x10 locals=0x88 funcid=0x0 align=0x0
	0x0000 00000 (main.go:18)	TEXT	"".echo[go.shape.uint32_0](SB), DUPOK|ABIInternal, $136-16
	0x0000 00000 (main.go:18)	LEAQ	-8(SP), R12
	0x0005 00005 (main.go:18)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:18)	PCDATA	$0, $-2
	0x0009 00009 (main.go:18)	JLS	238
	0x000f 00015 (main.go:18)	PCDATA	$0, $-1
	0x000f 00015 (main.go:18)	SUBQ	$136, SP
	0x0016 00022 (main.go:18)	MOVQ	BP, 128(SP)
	0x001e 00030 (main.go:18)	LEAQ	128(SP), BP
	0x0026 00038 (main.go:18)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$1, gclocals·731dddc57432221a509e963fc20cf850(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$2, "".echo[go.shape.uint32_0].stkobj(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$5, "".echo[go.shape.uint32_0].arginfo1(SB)
	0x0026 00038 (main.go:18)	MOVQ	AX, ""..dict+144(SP)
	0x002e 00046 (main.go:18)	MOVL	BX, "".t+152(SP)
	0x0035 00053 (main.go:18)	MOVUPS	X15, "".~r0+56(SP)
	0x003b 00059 (main.go:19)	MOVUPS	X15, ""..autotmp_3+88(SP)
	0x0041 00065 (main.go:19)	LEAQ	""..autotmp_3+88(SP), CX
	0x0046 00070 (main.go:19)	MOVQ	CX, ""..autotmp_6+48(SP)
	0x004b 00075 (main.go:19)	MOVL	"".t+152(SP), AX
	0x0052 00082 (main.go:19)	PCDATA	$1, $1
	0x0052 00082 (main.go:19)	CALL	runtime.convT32(SB)
	0x0057 00087 (main.go:19)	PCDATA	$0, $-2
	0x0057 00087 (main.go:19)	MOVQ	AX, ""..autotmp_7+40(SP)
	0x005c 00092 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x0061 00097 (main.go:19)	TESTB	AL, (CX)
	0x0063 00099 (main.go:19)	MOVQ	""..dict+144(SP), DX
	0x006b 00107 (main.go:19)	PCDATA	$0, $-1
	0x006b 00107 (main.go:19)	TESTB	AL, (DX)
	0x006d 00109 (main.go:19)	MOVQ	(DX), DX
	0x0070 00112 (main.go:19)	MOVQ	DX, (CX)
	0x0073 00115 (main.go:19)	LEAQ	8(CX), DI
	0x0077 00119 (main.go:19)	PCDATA	$0, $-2
	0x0077 00119 (main.go:19)	CMPL	runtime.writeBarrier(SB), $0
	0x007e 00126 (main.go:19)	NOP
	0x0080 00128 (main.go:19)	JEQ	132
	0x0082 00130 (main.go:19)	JMP	138
	0x0084 00132 (main.go:19)	MOVQ	AX, 8(CX)
	0x0088 00136 (main.go:19)	JMP	145
	0x008a 00138 (main.go:19)	CALL	runtime.gcWriteBarrier(SB)
	0x008f 00143 (main.go:19)	JMP	145
	0x0091 00145 (main.go:19)	PCDATA	$0, $-1
	0x0091 00145 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x0096 00150 (main.go:19)	TESTB	AL, (CX)
	0x0098 00152 (main.go:19)	JMP	154
	0x009a 00154 (main.go:19)	MOVQ	CX, ""..autotmp_5+104(SP)
	0x009f 00159 (main.go:19)	MOVQ	$1, ""..autotmp_5+112(SP)
	0x00a8 00168 (main.go:19)	MOVQ	$1, ""..autotmp_5+120(SP)
	0x00b1 00177 (main.go:19)	LEAQ	go.string."%v"(SB), AX
	0x00b8 00184 (main.go:19)	MOVL	$2, BX
	0x00bd 00189 (main.go:19)	MOVL	$1, DI
	0x00c2 00194 (main.go:19)	MOVQ	DI, SI
	0x00c5 00197 (main.go:19)	PCDATA	$1, $0
	0x00c5 00197 (main.go:19)	CALL	fmt.Sprintf(SB)
	0x00ca 00202 (main.go:19)	MOVQ	AX, ""..autotmp_4+72(SP)
	0x00cf 00207 (main.go:19)	MOVQ	BX, ""..autotmp_4+80(SP)
	0x00d4 00212 (main.go:19)	MOVQ	AX, "".~r0+56(SP)
	0x00d9 00217 (main.go:19)	MOVQ	BX, "".~r0+64(SP)
	0x00de 00222 (main.go:19)	MOVQ	128(SP), BP
	0x00e6 00230 (main.go:19)	ADDQ	$136, SP
	0x00ed 00237 (main.go:19)	RET
	0x00ee 00238 (main.go:19)	NOP
	0x00ee 00238 (main.go:18)	PCDATA	$1, $-1
	0x00ee 00238 (main.go:18)	PCDATA	$0, $-2
	0x00ee 00238 (main.go:18)	MOVQ	AX, 8(SP)
	0x00f3 00243 (main.go:18)	MOVL	BX, 16(SP)
	0x00f7 00247 (main.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x00fc 00252 (main.go:18)	MOVQ	8(SP), AX
	0x0101 00257 (main.go:18)	MOVL	16(SP), BX
	0x0105 00261 (main.go:18)	PCDATA	$0, $-1
	0x0105 00261 (main.go:18)	JMP	0
	0x0000 4c 8d 64 24 f8 4d 3b 66 10 0f 86 df 00 00 00 48  L.d$.M;f.......H
	0x0010 81 ec 88 00 00 00 48 89 ac 24 80 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 80 00 00 00 48 89 84 24 90 00 00 00 89 9c  .$....H..$......
	0x0030 24 98 00 00 00 44 0f 11 7c 24 38 44 0f 11 7c 24  $....D..|$8D..|$
	0x0040 58 48 8d 4c 24 58 48 89 4c 24 30 8b 84 24 98 00  XH.L$XH.L$0..$..
	0x0050 00 00 e8 00 00 00 00 48 89 44 24 28 48 8b 4c 24  .......H.D$(H.L$
	0x0060 30 84 01 48 8b 94 24 90 00 00 00 84 02 48 8b 12  0..H..$......H..
	0x0070 48 89 11 48 8d 79 08 83 3d 00 00 00 00 00 66 90  H..H.y..=.....f.
	0x0080 74 02 eb 06 48 89 41 08 eb 07 e8 00 00 00 00 eb  t...H.A.........
	0x0090 00 48 8b 4c 24 30 84 01 eb 00 48 89 4c 24 68 48  .H.L$0....H.L$hH
	0x00a0 c7 44 24 70 01 00 00 00 48 c7 44 24 78 01 00 00  .D$p....H.D$x...
	0x00b0 00 48 8d 05 00 00 00 00 bb 02 00 00 00 bf 01 00  .H..............
	0x00c0 00 00 48 89 fe e8 00 00 00 00 48 89 44 24 48 48  ..H.......H.D$HH
	0x00d0 89 5c 24 50 48 89 44 24 38 48 89 5c 24 40 48 8b  .\$PH.D$8H.\$@H.
	0x00e0 ac 24 80 00 00 00 48 81 c4 88 00 00 00 c3 48 89  .$....H.......H.
	0x00f0 44 24 08 89 5c 24 10 e8 00 00 00 00 48 8b 44 24  D$..\$......H.D$
	0x0100 08 8b 5c 24 10 e9 f6 fe ff ff                    ..\$......
	rel 83+4 t=7 runtime.convT32+0
	rel 121+4 t=14 runtime.writeBarrier+-1
	rel 139+4 t=7 runtime.gcWriteBarrier+0
	rel 180+4 t=14 go.string."%v"+0
	rel 198+4 t=7 fmt.Sprintf+0
	rel 248+4 t=7 runtime.morestack_noctxt+0
"".echo[go.shape.uint64_0] STEXT dupok size=268 args=0x10 locals=0x88 funcid=0x0 align=0x0
	0x0000 00000 (main.go:18)	TEXT	"".echo[go.shape.uint64_0](SB), DUPOK|ABIInternal, $136-16
	0x0000 00000 (main.go:18)	LEAQ	-8(SP), R12
	0x0005 00005 (main.go:18)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:18)	PCDATA	$0, $-2
	0x0009 00009 (main.go:18)	JLS	238
	0x000f 00015 (main.go:18)	PCDATA	$0, $-1
	0x000f 00015 (main.go:18)	SUBQ	$136, SP
	0x0016 00022 (main.go:18)	MOVQ	BP, 128(SP)
	0x001e 00030 (main.go:18)	LEAQ	128(SP), BP
	0x0026 00038 (main.go:18)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$1, gclocals·731dddc57432221a509e963fc20cf850(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$2, "".echo[go.shape.uint64_0].stkobj(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$5, "".echo[go.shape.uint64_0].arginfo1(SB)
	0x0026 00038 (main.go:18)	MOVQ	AX, ""..dict+144(SP)
	0x002e 00046 (main.go:18)	MOVQ	BX, "".t+152(SP)
	0x0036 00054 (main.go:18)	MOVUPS	X15, "".~r0+56(SP)
	0x003c 00060 (main.go:19)	MOVUPS	X15, ""..autotmp_3+88(SP)
	0x0042 00066 (main.go:19)	LEAQ	""..autotmp_3+88(SP), CX
	0x0047 00071 (main.go:19)	MOVQ	CX, ""..autotmp_6+48(SP)
	0x004c 00076 (main.go:19)	MOVQ	"".t+152(SP), AX
	0x0054 00084 (main.go:19)	PCDATA	$1, $1
	0x0054 00084 (main.go:19)	CALL	runtime.convT64(SB)
	0x0059 00089 (main.go:19)	PCDATA	$0, $-2
	0x0059 00089 (main.go:19)	MOVQ	AX, ""..autotmp_7+40(SP)
	0x005e 00094 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x0063 00099 (main.go:19)	TESTB	AL, (CX)
	0x0065 00101 (main.go:19)	MOVQ	""..dict+144(SP), DX
	0x006d 00109 (main.go:19)	PCDATA	$0, $-1
	0x006d 00109 (main.go:19)	TESTB	AL, (DX)
	0x006f 00111 (main.go:19)	MOVQ	(DX), DX
	0x0072 00114 (main.go:19)	MOVQ	DX, (CX)
	0x0075 00117 (main.go:19)	LEAQ	8(CX), DI
	0x0079 00121 (main.go:19)	PCDATA	$0, $-2
	0x0079 00121 (main.go:19)	CMPL	runtime.writeBarrier(SB), $0
	0x0080 00128 (main.go:19)	JEQ	132
	0x0082 00130 (main.go:19)	JMP	138
	0x0084 00132 (main.go:19)	MOVQ	AX, 8(CX)
	0x0088 00136 (main.go:19)	JMP	145
	0x008a 00138 (main.go:19)	CALL	runtime.gcWriteBarrier(SB)
	0x008f 00143 (main.go:19)	JMP	145
	0x0091 00145 (main.go:19)	PCDATA	$0, $-1
	0x0091 00145 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x0096 00150 (main.go:19)	TESTB	AL, (CX)
	0x0098 00152 (main.go:19)	JMP	154
	0x009a 00154 (main.go:19)	MOVQ	CX, ""..autotmp_5+104(SP)
	0x009f 00159 (main.go:19)	MOVQ	$1, ""..autotmp_5+112(SP)
	0x00a8 00168 (main.go:19)	MOVQ	$1, ""..autotmp_5+120(SP)
	0x00b1 00177 (main.go:19)	LEAQ	go.string."%v"(SB), AX
	0x00b8 00184 (main.go:19)	MOVL	$2, BX
	0x00bd 00189 (main.go:19)	MOVL	$1, DI
	0x00c2 00194 (main.go:19)	MOVQ	DI, SI
	0x00c5 00197 (main.go:19)	PCDATA	$1, $0
	0x00c5 00197 (main.go:19)	CALL	fmt.Sprintf(SB)
	0x00ca 00202 (main.go:19)	MOVQ	AX, ""..autotmp_4+72(SP)
	0x00cf 00207 (main.go:19)	MOVQ	BX, ""..autotmp_4+80(SP)
	0x00d4 00212 (main.go:19)	MOVQ	AX, "".~r0+56(SP)
	0x00d9 00217 (main.go:19)	MOVQ	BX, "".~r0+64(SP)
	0x00de 00222 (main.go:19)	MOVQ	128(SP), BP
	0x00e6 00230 (main.go:19)	ADDQ	$136, SP
	0x00ed 00237 (main.go:19)	RET
	0x00ee 00238 (main.go:19)	NOP
	0x00ee 00238 (main.go:18)	PCDATA	$1, $-1
	0x00ee 00238 (main.go:18)	PCDATA	$0, $-2
	0x00ee 00238 (main.go:18)	MOVQ	AX, 8(SP)
	0x00f3 00243 (main.go:18)	MOVQ	BX, 16(SP)
	0x00f8 00248 (main.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x00fd 00253 (main.go:18)	MOVQ	8(SP), AX
	0x0102 00258 (main.go:18)	MOVQ	16(SP), BX
	0x0107 00263 (main.go:18)	PCDATA	$0, $-1
	0x0107 00263 (main.go:18)	JMP	0
	0x0000 4c 8d 64 24 f8 4d 3b 66 10 0f 86 df 00 00 00 48  L.d$.M;f.......H
	0x0010 81 ec 88 00 00 00 48 89 ac 24 80 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 80 00 00 00 48 89 84 24 90 00 00 00 48 89  .$....H..$....H.
	0x0030 9c 24 98 00 00 00 44 0f 11 7c 24 38 44 0f 11 7c  .$....D..|$8D..|
	0x0040 24 58 48 8d 4c 24 58 48 89 4c 24 30 48 8b 84 24  $XH.L$XH.L$0H..$
	0x0050 98 00 00 00 e8 00 00 00 00 48 89 44 24 28 48 8b  .........H.D$(H.
	0x0060 4c 24 30 84 01 48 8b 94 24 90 00 00 00 84 02 48  L$0..H..$......H
	0x0070 8b 12 48 89 11 48 8d 79 08 83 3d 00 00 00 00 00  ..H..H.y..=.....
	0x0080 74 02 eb 06 48 89 41 08 eb 07 e8 00 00 00 00 eb  t...H.A.........
	0x0090 00 48 8b 4c 24 30 84 01 eb 00 48 89 4c 24 68 48  .H.L$0....H.L$hH
	0x00a0 c7 44 24 70 01 00 00 00 48 c7 44 24 78 01 00 00  .D$p....H.D$x...
	0x00b0 00 48 8d 05 00 00 00 00 bb 02 00 00 00 bf 01 00  .H..............
	0x00c0 00 00 48 89 fe e8 00 00 00 00 48 89 44 24 48 48  ..H.......H.D$HH
	0x00d0 89 5c 24 50 48 89 44 24 38 48 89 5c 24 40 48 8b  .\$PH.D$8H.\$@H.
	0x00e0 ac 24 80 00 00 00 48 81 c4 88 00 00 00 c3 48 89  .$....H.......H.
	0x00f0 44 24 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44  D$.H.\$......H.D
	0x0100 24 08 48 8b 5c 24 10 e9 f4 fe ff ff              $.H.\$......
	rel 85+4 t=7 runtime.convT64+0
	rel 123+4 t=14 runtime.writeBarrier+-1
	rel 139+4 t=7 runtime.gcWriteBarrier+0
	rel 180+4 t=14 go.string."%v"+0
	rel 198+4 t=7 fmt.Sprintf+0
	rel 249+4 t=7 runtime.morestack_noctxt+0
"".echo[go.shape.string_0] STEXT dupok size=295 args=0x18 locals=0x88 funcid=0x0 align=0x0
	0x0000 00000 (main.go:18)	TEXT	"".echo[go.shape.string_0](SB), DUPOK|ABIInternal, $136-24
	0x0000 00000 (main.go:18)	LEAQ	-8(SP), R12
	0x0005 00005 (main.go:18)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:18)	PCDATA	$0, $-2
	0x0009 00009 (main.go:18)	JLS	255
	0x000f 00015 (main.go:18)	PCDATA	$0, $-1
	0x000f 00015 (main.go:18)	SUBQ	$136, SP
	0x0016 00022 (main.go:18)	MOVQ	BP, 128(SP)
	0x001e 00030 (main.go:18)	LEAQ	128(SP), BP
	0x0026 00038 (main.go:18)	FUNCDATA	$0, gclocals·ba30782f8935b28ed1adaec603e72627(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$1, gclocals·f9b2f46f83a03191e4148409eebaa2a7(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$2, "".echo[go.shape.string_0].stkobj(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$5, "".echo[go.shape.string_0].arginfo1(SB)
	0x0026 00038 (main.go:18)	MOVQ	AX, ""..dict+144(SP)
	0x002e 00046 (main.go:18)	MOVQ	BX, "".t+152(SP)
	0x0036 00054 (main.go:18)	MOVQ	CX, "".t+160(SP)
	0x003e 00062 (main.go:18)	MOVUPS	X15, "".~r0+56(SP)
	0x0044 00068 (main.go:19)	MOVUPS	X15, ""..autotmp_3+88(SP)
	0x004a 00074 (main.go:19)	LEAQ	""..autotmp_3+88(SP), CX
	0x004f 00079 (main.go:19)	MOVQ	CX, ""..autotmp_6+48(SP)
	0x0054 00084 (main.go:19)	MOVQ	"".t+152(SP), AX
	0x005c 00092 (main.go:19)	MOVQ	"".t+160(SP), BX
	0x0064 00100 (main.go:19)	PCDATA	$1, $1
	0x0064 00100 (main.go:19)	CALL	runtime.convTstring(SB)
	0x0069 00105 (main.go:19)	PCDATA	$0, $-2
	0x0069 00105 (main.go:19)	MOVQ	AX, ""..autotmp_7+40(SP)
	0x006e 00110 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x0073 00115 (main.go:19)	TESTB	AL, (CX)
	0x0075 00117 (main.go:19)	MOVQ	""..dict+144(SP), DX
	0x007d 00125 (main.go:19)	PCDATA	$0, $-1
	0x007d 00125 (main.go:19)	TESTB	AL, (DX)
	0x007f 00127 (main.go:19)	MOVQ	(DX), DX
	0x0082 00130 (main.go:19)	MOVQ	DX, (CX)
	0x0085 00133 (main.go:19)	LEAQ	8(CX), DI
	0x0089 00137 (main.go:19)	PCDATA	$0, $-2
	0x0089 00137 (main.go:19)	CMPL	runtime.writeBarrier(SB), $0
	0x0090 00144 (main.go:19)	JEQ	148
	0x0092 00146 (main.go:19)	JMP	154
	0x0094 00148 (main.go:19)	MOVQ	AX, 8(CX)
	0x0098 00152 (main.go:19)	JMP	162
	0x009a 00154 (main.go:19)	CALL	runtime.gcWriteBarrier(SB)
	0x009f 00159 (main.go:19)	NOP
	0x00a0 00160 (main.go:19)	JMP	162
	0x00a2 00162 (main.go:19)	PCDATA	$0, $-1
	0x00a2 00162 (main.go:19)	MOVQ	""..autotmp_6+48(SP), CX
	0x00a7 00167 (main.go:19)	TESTB	AL, (CX)
	0x00a9 00169 (main.go:19)	JMP	171
	0x00ab 00171 (main.go:19)	MOVQ	CX, ""..autotmp_5+104(SP)
	0x00b0 00176 (main.go:19)	MOVQ	$1, ""..autotmp_5+112(SP)
	0x00b9 00185 (main.go:19)	MOVQ	$1, ""..autotmp_5+120(SP)
	0x00c2 00194 (main.go:19)	LEAQ	go.string."%v"(SB), AX
	0x00c9 00201 (main.go:19)	MOVL	$2, BX
	0x00ce 00206 (main.go:19)	MOVL	$1, DI
	0x00d3 00211 (main.go:19)	MOVQ	DI, SI
	0x00d6 00214 (main.go:19)	PCDATA	$1, $2
	0x00d6 00214 (main.go:19)	CALL	fmt.Sprintf(SB)
	0x00db 00219 (main.go:19)	MOVQ	AX, ""..autotmp_4+72(SP)
	0x00e0 00224 (main.go:19)	MOVQ	BX, ""..autotmp_4+80(SP)
	0x00e5 00229 (main.go:19)	MOVQ	AX, "".~r0+56(SP)
	0x00ea 00234 (main.go:19)	MOVQ	BX, "".~r0+64(SP)
	0x00ef 00239 (main.go:19)	MOVQ	128(SP), BP
	0x00f7 00247 (main.go:19)	ADDQ	$136, SP
	0x00fe 00254 (main.go:19)	RET
	0x00ff 00255 (main.go:19)	NOP
	0x00ff 00255 (main.go:18)	PCDATA	$1, $-1
	0x00ff 00255 (main.go:18)	PCDATA	$0, $-2
	0x00ff 00255 (main.go:18)	MOVQ	AX, 8(SP)
	0x0104 00260 (main.go:18)	MOVQ	BX, 16(SP)
	0x0109 00265 (main.go:18)	MOVQ	CX, 24(SP)
	0x010e 00270 (main.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x0113 00275 (main.go:18)	MOVQ	8(SP), AX
	0x0118 00280 (main.go:18)	MOVQ	16(SP), BX
	0x011d 00285 (main.go:18)	MOVQ	24(SP), CX
	0x0122 00290 (main.go:18)	PCDATA	$0, $-1
	0x0122 00290 (main.go:18)	JMP	0
	0x0000 4c 8d 64 24 f8 4d 3b 66 10 0f 86 f0 00 00 00 48  L.d$.M;f.......H
	0x0010 81 ec 88 00 00 00 48 89 ac 24 80 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 80 00 00 00 48 89 84 24 90 00 00 00 48 89  .$....H..$....H.
	0x0030 9c 24 98 00 00 00 48 89 8c 24 a0 00 00 00 44 0f  .$....H..$....D.
	0x0040 11 7c 24 38 44 0f 11 7c 24 58 48 8d 4c 24 58 48  .|$8D..|$XH.L$XH
	0x0050 89 4c 24 30 48 8b 84 24 98 00 00 00 48 8b 9c 24  .L$0H..$....H..$
	0x0060 a0 00 00 00 e8 00 00 00 00 48 89 44 24 28 48 8b  .........H.D$(H.
	0x0070 4c 24 30 84 01 48 8b 94 24 90 00 00 00 84 02 48  L$0..H..$......H
	0x0080 8b 12 48 89 11 48 8d 79 08 83 3d 00 00 00 00 00  ..H..H.y..=.....
	0x0090 74 02 eb 06 48 89 41 08 eb 08 e8 00 00 00 00 90  t...H.A.........
	0x00a0 eb 00 48 8b 4c 24 30 84 01 eb 00 48 89 4c 24 68  ..H.L$0....H.L$h
	0x00b0 48 c7 44 24 70 01 00 00 00 48 c7 44 24 78 01 00  H.D$p....H.D$x..
	0x00c0 00 00 48 8d 05 00 00 00 00 bb 02 00 00 00 bf 01  ..H.............
	0x00d0 00 00 00 48 89 fe e8 00 00 00 00 48 89 44 24 48  ...H.......H.D$H
	0x00e0 48 89 5c 24 50 48 89 44 24 38 48 89 5c 24 40 48  H.\$PH.D$8H.\$@H
	0x00f0 8b ac 24 80 00 00 00 48 81 c4 88 00 00 00 c3 48  ..$....H.......H
	0x0100 89 44 24 08 48 89 5c 24 10 48 89 4c 24 18 e8 00  .D$.H.\$.H.L$...
	0x0110 00 00 00 48 8b 44 24 08 48 8b 5c 24 10 48 8b 4c  ...H.D$.H.\$.H.L
	0x0120 24 18 e9 d9 fe ff ff                             $......
	rel 101+4 t=7 runtime.convTstring+0
	rel 139+4 t=14 runtime.writeBarrier+-1
	rel 155+4 t=7 runtime.gcWriteBarrier+0
	rel 197+4 t=14 go.string."%v"+0
	rel 215+4 t=7 fmt.Sprintf+0
	rel 271+4 t=7 runtime.morestack_noctxt+0
"".echo[go.shape.struct {}_0] STEXT dupok size=208 args=0x8 locals=0x88 funcid=0x0 align=0x0
	0x0000 00000 (main.go:18)	TEXT	"".echo[go.shape.struct {}_0](SB), DUPOK|ABIInternal, $136-8
	0x0000 00000 (main.go:18)	LEAQ	-8(SP), R12
	0x0005 00005 (main.go:18)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:18)	PCDATA	$0, $-2
	0x0009 00009 (main.go:18)	JLS	188
	0x000f 00015 (main.go:18)	PCDATA	$0, $-1
	0x000f 00015 (main.go:18)	SUBQ	$136, SP
	0x0016 00022 (main.go:18)	MOVQ	BP, 128(SP)
	0x001e 00030 (main.go:18)	LEAQ	128(SP), BP
	0x0026 00038 (main.go:18)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$1, gclocals·e1d295804c1c91c77a93c02d636362d3(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$2, "".echo[go.shape.struct {}_0].stkobj(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$5, "".echo[go.shape.struct {}_0].arginfo1(SB)
	0x0026 00038 (main.go:18)	PCDATA	$0, $-2
	0x0026 00038 (main.go:18)	MOVQ	AX, ""..dict+144(SP)
	0x002e 00046 (main.go:18)	MOVUPS	X15, "".~r0+56(SP)
	0x0034 00052 (main.go:19)	MOVUPS	X15, ""..autotmp_4+88(SP)
	0x003a 00058 (main.go:19)	LEAQ	""..autotmp_4+88(SP), CX
	0x003f 00063 (main.go:19)	MOVQ	CX, ""..autotmp_7+48(SP)
	0x0044 00068 (main.go:19)	TESTB	AL, (CX)
	0x0046 00070 (main.go:19)	MOVQ	""..dict+144(SP), DX
	0x004e 00078 (main.go:19)	PCDATA	$0, $-1
	0x004e 00078 (main.go:19)	TESTB	AL, (DX)
	0x0050 00080 (main.go:19)	MOVQ	(DX), DX
	0x0053 00083 (main.go:19)	MOVQ	DX, ""..autotmp_4+88(SP)
	0x0058 00088 (main.go:19)	LEAQ	runtime.zerobase(SB), DX
	0x005f 00095 (main.go:19)	MOVQ	DX, ""..autotmp_4+96(SP)
	0x0064 00100 (main.go:19)	TESTB	AL, (CX)
	0x0066 00102 (main.go:19)	JMP	104
	0x0068 00104 (main.go:19)	MOVQ	CX, ""..autotmp_6+104(SP)
	0x006d 00109 (main.go:19)	MOVQ	$1, ""..autotmp_6+112(SP)
	0x0076 00118 (main.go:19)	MOVQ	$1, ""..autotmp_6+120(SP)
	0x007f 00127 (main.go:19)	LEAQ	go.string."%v"(SB), AX
	0x0086 00134 (main.go:19)	MOVL	$2, BX
	0x008b 00139 (main.go:19)	MOVL	$1, DI
	0x0090 00144 (main.go:19)	MOVQ	DI, SI
	0x0093 00147 (main.go:19)	PCDATA	$1, $0
	0x0093 00147 (main.go:19)	CALL	fmt.Sprintf(SB)
	0x0098 00152 (main.go:19)	MOVQ	AX, ""..autotmp_5+72(SP)
	0x009d 00157 (main.go:19)	MOVQ	BX, ""..autotmp_5+80(SP)
	0x00a2 00162 (main.go:19)	MOVQ	AX, "".~r0+56(SP)
	0x00a7 00167 (main.go:19)	MOVQ	BX, "".~r0+64(SP)
	0x00ac 00172 (main.go:19)	MOVQ	128(SP), BP
	0x00b4 00180 (main.go:19)	ADDQ	$136, SP
	0x00bb 00187 (main.go:19)	RET
	0x00bc 00188 (main.go:19)	NOP
	0x00bc 00188 (main.go:18)	PCDATA	$1, $-1
	0x00bc 00188 (main.go:18)	PCDATA	$0, $-2
	0x00bc 00188 (main.go:18)	MOVQ	AX, 8(SP)
	0x00c1 00193 (main.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x00c6 00198 (main.go:18)	MOVQ	8(SP), AX
	0x00cb 00203 (main.go:18)	PCDATA	$0, $-1
	0x00cb 00203 (main.go:18)	JMP	0
	0x0000 4c 8d 64 24 f8 4d 3b 66 10 0f 86 ad 00 00 00 48  L.d$.M;f.......H
	0x0010 81 ec 88 00 00 00 48 89 ac 24 80 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 80 00 00 00 48 89 84 24 90 00 00 00 44 0f  .$....H..$....D.
	0x0030 11 7c 24 38 44 0f 11 7c 24 58 48 8d 4c 24 58 48  .|$8D..|$XH.L$XH
	0x0040 89 4c 24 30 84 01 48 8b 94 24 90 00 00 00 84 02  .L$0..H..$......
	0x0050 48 8b 12 48 89 54 24 58 48 8d 15 00 00 00 00 48  H..H.T$XH......H
	0x0060 89 54 24 60 84 01 eb 00 48 89 4c 24 68 48 c7 44  .T$`....H.L$hH.D
	0x0070 24 70 01 00 00 00 48 c7 44 24 78 01 00 00 00 48  $p....H.D$x....H
	0x0080 8d 05 00 00 00 00 bb 02 00 00 00 bf 01 00 00 00  ................
	0x0090 48 89 fe e8 00 00 00 00 48 89 44 24 48 48 89 5c  H.......H.D$HH.\
	0x00a0 24 50 48 89 44 24 38 48 89 5c 24 40 48 8b ac 24  $PH.D$8H.\$@H..$
	0x00b0 80 00 00 00 48 81 c4 88 00 00 00 c3 48 89 44 24  ....H.......H.D$
	0x00c0 08 e8 00 00 00 00 48 8b 44 24 08 e9 30 ff ff ff  ......H.D$..0...
	rel 91+4 t=14 runtime.zerobase+0
	rel 130+4 t=14 go.string."%v"+0
	rel 148+4 t=7 fmt.Sprintf+0
	rel 194+4 t=7 runtime.morestack_noctxt+0
"".echo[go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0] STEXT dupok size=364 args=0x20 locals=0xa0 funcid=0x0 align=0x0
	0x0000 00000 (main.go:18)	TEXT	"".echo[go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0](SB), DUPOK|ABIInternal, $160-32
	0x0000 00000 (main.go:18)	LEAQ	-32(SP), R12
	0x0005 00005 (main.go:18)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:18)	PCDATA	$0, $-2
	0x0009 00009 (main.go:18)	JLS	314
	0x000f 00015 (main.go:18)	PCDATA	$0, $-1
	0x000f 00015 (main.go:18)	SUBQ	$160, SP
	0x0016 00022 (main.go:18)	MOVQ	BP, 152(SP)
	0x001e 00030 (main.go:18)	LEAQ	152(SP), BP
	0x0026 00038 (main.go:18)	FUNCDATA	$0, gclocals·8e13821a52d7f75aaebe655ba3b03067(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$1, gclocals·b24549320fb5e410e9425e096449fc75(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$2, "".echo[go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0].stkobj(SB)
	0x0026 00038 (main.go:18)	FUNCDATA	$5, "".echo[go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0].arginfo1(SB)
	0x0026 00038 (main.go:18)	MOVQ	AX, ""..dict+168(SP)
	0x002e 00046 (main.go:18)	MOVQ	BX, "".t+176(SP)
	0x0036 00054 (main.go:18)	MOVQ	CX, "".t+184(SP)
	0x003e 00062 (main.go:18)	MOVQ	DI, "".t+192(SP)
	0x0046 00070 (main.go:18)	MOVUPS	X15, "".~r0+56(SP)
	0x004c 00076 (main.go:19)	MOVQ	"".t+176(SP), CX
	0x0054 00084 (main.go:19)	MOVQ	"".t+184(SP), DX
	0x005c 00092 (main.go:19)	MOVQ	"".t+192(SP), SI
	0x0064 00100 (main.go:19)	MOVQ	CX, ""..autotmp_3+128(SP)
	0x006c 00108 (main.go:19)	MOVQ	DX, ""..autotmp_3+136(SP)
	0x0074 00116 (main.go:19)	MOVQ	SI, ""..autotmp_3+144(SP)
	0x007c 00124 (main.go:19)	MOVUPS	X15, ""..autotmp_4+88(SP)
	0x0082 00130 (main.go:19)	LEAQ	""..autotmp_4+88(SP), CX
	0x0087 00135 (main.go:19)	MOVQ	CX, ""..autotmp_7+48(SP)
	0x008c 00140 (main.go:19)	LEAQ	type.go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0(SB), AX
	0x0093 00147 (main.go:19)	LEAQ	""..autotmp_3+128(SP), BX
	0x009b 00155 (main.go:19)	PCDATA	$1, $1
	0x009b 00155 (main.go:19)	NOP
	0x00a0 00160 (main.go:19)	CALL	runtime.convT(SB)
	0x00a5 00165 (main.go:19)	PCDATA	$0, $-2
	0x00a5 00165 (main.go:19)	MOVQ	AX, ""..autotmp_8+40(SP)
	0x00aa 00170 (main.go:19)	MOVQ	""..autotmp_7+48(SP), CX
	0x00af 00175 (main.go:19)	TESTB	AL, (CX)
	0x00b1 00177 (main.go:19)	MOVQ	""..dict+168(SP), DX
	0x00b9 00185 (main.go:19)	PCDATA	$0, $-1
	0x00b9 00185 (main.go:19)	TESTB	AL, (DX)
	0x00bb 00187 (main.go:19)	MOVQ	(DX), DX
	0x00be 00190 (main.go:19)	MOVQ	DX, (CX)
	0x00c1 00193 (main.go:19)	LEAQ	8(CX), DI
	0x00c5 00197 (main.go:19)	PCDATA	$0, $-2
	0x00c5 00197 (main.go:19)	CMPL	runtime.writeBarrier(SB), $0
	0x00cc 00204 (main.go:19)	JEQ	208
	0x00ce 00206 (main.go:19)	JMP	214
	0x00d0 00208 (main.go:19)	MOVQ	AX, 8(CX)
	0x00d4 00212 (main.go:19)	JMP	221
	0x00d6 00214 (main.go:19)	CALL	runtime.gcWriteBarrier(SB)
	0x00db 00219 (main.go:19)	JMP	221
	0x00dd 00221 (main.go:19)	PCDATA	$0, $-1
	0x00dd 00221 (main.go:19)	MOVQ	""..autotmp_7+48(SP), CX
	0x00e2 00226 (main.go:19)	TESTB	AL, (CX)
	0x00e4 00228 (main.go:19)	JMP	230
	0x00e6 00230 (main.go:19)	MOVQ	CX, ""..autotmp_6+104(SP)
	0x00eb 00235 (main.go:19)	MOVQ	$1, ""..autotmp_6+112(SP)
	0x00f4 00244 (main.go:19)	MOVQ	$1, ""..autotmp_6+120(SP)
	0x00fd 00253 (main.go:19)	LEAQ	go.string."%v"(SB), AX
	0x0104 00260 (main.go:19)	MOVL	$2, BX
	0x0109 00265 (main.go:19)	MOVL	$1, DI
	0x010e 00270 (main.go:19)	MOVQ	DI, SI
	0x0111 00273 (main.go:19)	PCDATA	$1, $2
	0x0111 00273 (main.go:19)	CALL	fmt.Sprintf(SB)
	0x0116 00278 (main.go:19)	MOVQ	AX, ""..autotmp_5+72(SP)
	0x011b 00283 (main.go:19)	MOVQ	BX, ""..autotmp_5+80(SP)
	0x0120 00288 (main.go:19)	MOVQ	AX, "".~r0+56(SP)
	0x0125 00293 (main.go:19)	MOVQ	BX, "".~r0+64(SP)
	0x012a 00298 (main.go:19)	MOVQ	152(SP), BP
	0x0132 00306 (main.go:19)	ADDQ	$160, SP
	0x0139 00313 (main.go:19)	RET
	0x013a 00314 (main.go:19)	NOP
	0x013a 00314 (main.go:18)	PCDATA	$1, $-1
	0x013a 00314 (main.go:18)	PCDATA	$0, $-2
	0x013a 00314 (main.go:18)	MOVQ	AX, 8(SP)
	0x013f 00319 (main.go:18)	MOVQ	BX, 16(SP)
	0x0144 00324 (main.go:18)	MOVQ	CX, 24(SP)
	0x0149 00329 (main.go:18)	MOVQ	DI, 32(SP)
	0x014e 00334 (main.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x0153 00339 (main.go:18)	MOVQ	8(SP), AX
	0x0158 00344 (main.go:18)	MOVQ	16(SP), BX
	0x015d 00349 (main.go:18)	MOVQ	24(SP), CX
	0x0162 00354 (main.go:18)	MOVQ	32(SP), DI
	0x0167 00359 (main.go:18)	PCDATA	$0, $-1
	0x0167 00359 (main.go:18)	JMP	0
	0x0000 4c 8d 64 24 e0 4d 3b 66 10 0f 86 2b 01 00 00 48  L.d$.M;f...+...H
	0x0010 81 ec a0 00 00 00 48 89 ac 24 98 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 98 00 00 00 48 89 84 24 a8 00 00 00 48 89  .$....H..$....H.
	0x0030 9c 24 b0 00 00 00 48 89 8c 24 b8 00 00 00 48 89  .$....H..$....H.
	0x0040 bc 24 c0 00 00 00 44 0f 11 7c 24 38 48 8b 8c 24  .$....D..|$8H..$
	0x0050 b0 00 00 00 48 8b 94 24 b8 00 00 00 48 8b b4 24  ....H..$....H..$
	0x0060 c0 00 00 00 48 89 8c 24 80 00 00 00 48 89 94 24  ....H..$....H..$
	0x0070 88 00 00 00 48 89 b4 24 90 00 00 00 44 0f 11 7c  ....H..$....D..|
	0x0080 24 58 48 8d 4c 24 58 48 89 4c 24 30 48 8d 05 00  $XH.L$XH.L$0H...
	0x0090 00 00 00 48 8d 9c 24 80 00 00 00 0f 1f 44 00 00  ...H..$......D..
	0x00a0 e8 00 00 00 00 48 89 44 24 28 48 8b 4c 24 30 84  .....H.D$(H.L$0.
	0x00b0 01 48 8b 94 24 a8 00 00 00 84 02 48 8b 12 48 89  .H..$......H..H.
	0x00c0 11 48 8d 79 08 83 3d 00 00 00 00 00 74 02 eb 06  .H.y..=.....t...
	0x00d0 48 89 41 08 eb 07 e8 00 00 00 00 eb 00 48 8b 4c  H.A..........H.L
	0x00e0 24 30 84 01 eb 00 48 89 4c 24 68 48 c7 44 24 70  $0....H.L$hH.D$p
	0x00f0 01 00 00 00 48 c7 44 24 78 01 00 00 00 48 8d 05  ....H.D$x....H..
	0x0100 00 00 00 00 bb 02 00 00 00 bf 01 00 00 00 48 89  ..............H.
	0x0110 fe e8 00 00 00 00 48 89 44 24 48 48 89 5c 24 50  ......H.D$HH.\$P
	0x0120 48 89 44 24 38 48 89 5c 24 40 48 8b ac 24 98 00  H.D$8H.\$@H..$..
	0x0130 00 00 48 81 c4 a0 00 00 00 c3 48 89 44 24 08 48  ..H.......H.D$.H
	0x0140 89 5c 24 10 48 89 4c 24 18 48 89 7c 24 20 e8 00  .\$.H.L$.H.|$ ..
	0x0150 00 00 00 48 8b 44 24 08 48 8b 5c 24 10 48 8b 4c  ...H.D$.H.\$.H.L
	0x0160 24 18 48 8b 7c 24 20 e9 94 fe ff ff              $.H.|$ .....
	rel 143+4 t=14 type.go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0+0
	rel 161+4 t=7 runtime.convT+0
	rel 199+4 t=14 runtime.writeBarrier+-1
	rel 215+4 t=7 runtime.gcWriteBarrier+0
	rel 256+4 t=14 go.string."%v"+0
	rel 274+4 t=7 fmt.Sprintf+0
	rel 335+4 t=7 runtime.morestack_noctxt+0
""..dict.echo[int] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.int+0
	rel 0+0 t=23 type.int+0
""..dict.echo[int32] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.int32+0
	rel 0+0 t=23 type.int32+0
""..dict.echo[uint32] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.uint32+0
	rel 0+0 t=23 type.uint32+0
""..dict.echo[uint64] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.uint64+0
	rel 0+0 t=23 type.uint64+0
""..dict.echo[string] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.string+0
	rel 0+0 t=23 type.string+0
""..dict.echo[struct {}] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.struct {}+0
	rel 0+0 t=23 type.struct {}+0
""..dict.echo[time.Time] SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type.time.Time+0
	rel 0+0 t=23 type.time.Time+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=40
	0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 fmt..inittask+0
	rel 32+8 t=1 time..inittask+0
go.string."hello" SRODATA dupok size=5
	0x0000 68 65 6c 6c 6f                                   hello
go.string."%v" SRODATA dupok size=2
	0x0000 25 76                                            %v
type..eqfunc24 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 runtime.memequal_varlen+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0- SRODATA dupok size=82
	0x0000 00 50 2a 67 6f 2e 73 68 61 70 65 2e 73 74 72 75  .P*go.shape.stru
	0x0010 63 74 20 7b 20 74 69 6d 65 2e 77 61 6c 6c 20 75  ct { time.wall u
	0x0020 69 6e 74 36 34 3b 20 74 69 6d 65 2e 65 78 74 20  int64; time.ext 
	0x0030 69 6e 74 36 34 3b 20 74 69 6d 65 2e 6c 6f 63 20  int64; time.loc 
	0x0040 2a 74 69 6d 65 2e 4c 6f 63 61 74 69 6f 6e 20 7d  *time.Location }
	0x0050 5f 30                                            _0
type.*go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 12 a7 74 e6 08 08 08 36 00 00 00 00 00 00 00 00  ..t....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0-+0
	rel 48+8 t=1 type.go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0+0
runtime.gcbits.04 SRODATA dupok size=1
	0x0000 04                                               .
type..importpath.time. SRODATA dupok size=6
	0x0000 00 04 74 69 6d 65                                ..time
type..importpath.go.shape. SRODATA dupok size=10
	0x0000 00 08 67 6f 2e 73 68 61 70 65                    ..go.shape
type..namedata.wall- SRODATA dupok size=6
	0x0000 00 04 77 61 6c 6c                                ..wall
type..namedata.ext- SRODATA dupok size=5
	0x0000 00 03 65 78 74                                   ..ext
type..namedata.loc- SRODATA dupok size=5
	0x0000 00 03 6c 6f 63                                   ..loc
type.go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0 SRODATA dupok size=168
	0x0000 18 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	0x0010 fc fb 2d 35 0f 08 08 19 00 00 00 00 00 00 00 00  ..-5............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 58 00 00 00 00 00 00 00  ........X.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 20 00 00 00 00 00 00 00                           .......
	rel 24+8 t=1 type..eqfunc24+0
	rel 32+8 t=1 runtime.gcbits.04+0
	rel 40+4 t=5 type..namedata.*go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0-+0
	rel 44+4 t=5 type.*go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0+0
	rel 48+8 t=1 type..importpath.time.+0
	rel 56+8 t=1 type.go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0+96
	rel 80+4 t=5 type..importpath.go.shape.+0
	rel 96+8 t=1 type..namedata.wall-+0
	rel 104+8 t=1 type.uint64+0
	rel 120+8 t=1 type..namedata.ext-+0
	rel 128+8 t=1 type.int64+0
	rel 144+8 t=1 type..namedata.loc-+0
	rel 152+8 t=1 type.*time.Location+0
runtime.memequal0·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal0+0
type..namedata.*struct {}- SRODATA dupok size=12
	0x0000 00 0a 2a 73 74 72 75 63 74 20 7b 7d              ..*struct {}
type.*struct {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4a 24 a9 e5 08 08 08 36 00 00 00 00 00 00 00 00  J$.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct {}-+0
	rel 48+8 t=1 type.struct {}+0
runtime.gcbits. SRODATA dupok size=0
type.struct {} SRODATA dupok size=80
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 1b ac f6 27 0a 01 01 19 00 00 00 00 00 00 00 00  ...'............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.memequal0·f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*struct {}-+0
	rel 44+4 t=-32763 type.*struct {}+0
	rel 56+8 t=1 type.struct {}+80
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·54241e171da8af6ae173d69da0236748 SRODATA dupok size=9
	0x0000 01 00 00 00 03 00 00 00 00                       .........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·731dddc57432221a509e963fc20cf850 SRODATA dupok size=12
	0x0000 02 00 00 00 0b 00 00 00 00 00 02 00              ............
"".echo[go.shape.int_0].stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
"".echo[go.shape.int_0].arginfo1 SRODATA static dupok size=3
	0x0000 08 08 ff                                         ...
"".echo[go.shape.int32_0].stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
"".echo[go.shape.int32_0].arginfo1 SRODATA static dupok size=3
	0x0000 08 04 ff                                         ...
"".echo[go.shape.uint32_0].stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
"".echo[go.shape.uint32_0].arginfo1 SRODATA static dupok size=3
	0x0000 08 04 ff                                         ...
"".echo[go.shape.uint64_0].stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
"".echo[go.shape.uint64_0].arginfo1 SRODATA static dupok size=3
	0x0000 08 08 ff                                         ...
gclocals·ba30782f8935b28ed1adaec603e72627 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 02 00 00                 ...........
gclocals·f9b2f46f83a03191e4148409eebaa2a7 SRODATA dupok size=14
	0x0000 03 00 00 00 0b 00 00 00 00 00 02 00 00 00        ..............
"".echo[go.shape.string_0].stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
"".echo[go.shape.string_0].arginfo1 SRODATA static dupok size=7
	0x0000 fe 08 08 10 08 fd ff                             .......
gclocals·e1d295804c1c91c77a93c02d636362d3 SRODATA dupok size=10
	0x0000 01 00 00 00 0a 00 00 00 00 00                    ..........
"".echo[go.shape.struct {}_0].stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
"".echo[go.shape.struct {}_0].arginfo1 SRODATA static dupok size=3
	0x0000 fe fd ff                                         ...
gclocals·8e13821a52d7f75aaebe655ba3b03067 SRODATA dupok size=11
	0x0000 03 00 00 00 04 00 00 00 08 00 00                 ...........
gclocals·b24549320fb5e410e9425e096449fc75 SRODATA dupok size=14
	0x0000 03 00 00 00 0e 00 00 00 00 00 02 00 00 00        ..............
"".echo[go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0].stkobj SRODATA static size=40
	0x0000 02 00 00 00 00 00 00 00 c0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00 e8 ff ff ff 18 00 00 00  ................
	0x0020 18 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
	rel 36+4 t=5 runtime.gcbits.04+0
"".echo[go.shape.struct { time.wall uint64; time.ext int64; time.loc *time.Location }_0].arginfo1 SRODATA static dupok size=9
	0x0000 fe 08 08 10 08 18 08 fd ff                       .........
