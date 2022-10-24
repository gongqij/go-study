"".main STEXT size=503 args=0x0 locals=0xf0 funcid=0x0 align=0x0
	0x0000 00000 (demo.go:5)	TEXT	"".main(SB), ABIInternal, $240-0
	0x0000 00000 (demo.go:5)	LEAQ	-112(SP), R12
	0x0005 00005 (demo.go:5)	CMPQ	R12, 16(R14)
	0x0009 00009 (demo.go:5)	PCDATA	$0, $-2
	0x0009 00009 (demo.go:5)	JLS	493
	0x000f 00015 (demo.go:5)	PCDATA	$0, $-1
	0x000f 00015 (demo.go:5)	SUBQ	$240, SP
	0x0016 00022 (demo.go:5)	MOVQ	BP, 232(SP)
	0x001e 00030 (demo.go:5)	LEAQ	232(SP), BP
	0x0026 00038 (demo.go:5)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0026 00038 (demo.go:5)	FUNCDATA	$1, gclocals·28835bfeb0e793947754a03b87c8a295(SB)
	0x0026 00038 (demo.go:5)	FUNCDATA	$2, "".main.stkobj(SB)
	0x0026 00038 (demo.go:6)	MOVQ	$200, "".x+24(SP)
	0x002f 00047 (demo.go:7)	MOVL	$200, AX
	0x0034 00052 (demo.go:7)	PCDATA	$1, $0
	0x0034 00052 (demo.go:7)	CALL	runtime.convT64(SB)
	0x0039 00057 (demo.go:7)	MOVQ	AX, ""..autotmp_5+40(SP)
	0x003e 00062 (demo.go:7)	LEAQ	type.int(SB), CX
	0x0045 00069 (demo.go:7)	MOVQ	CX, "".any+104(SP)
	0x004a 00074 (demo.go:7)	MOVQ	AX, "".any+112(SP)
	0x004f 00079 (demo.go:8)	MOVUPS	X15, ""..autotmp_4+152(SP)
	0x0058 00088 (demo.go:8)	LEAQ	""..autotmp_4+152(SP), AX
	0x0060 00096 (demo.go:8)	MOVQ	AX, ""..autotmp_7+32(SP)
	0x0065 00101 (demo.go:8)	TESTB	AL, (AX)
	0x0067 00103 (demo.go:8)	MOVQ	"".any+104(SP), CX
	0x006c 00108 (demo.go:8)	MOVQ	"".any+112(SP), DX
	0x0071 00113 (demo.go:8)	MOVQ	CX, ""..autotmp_4+152(SP)
	0x0079 00121 (demo.go:8)	MOVQ	DX, ""..autotmp_4+160(SP)
	0x0081 00129 (demo.go:8)	TESTB	AL, (AX)
	0x0083 00131 (demo.go:8)	JMP	133
	0x0085 00133 (demo.go:8)	MOVQ	AX, ""..autotmp_6+184(SP)
	0x008d 00141 (demo.go:8)	MOVQ	$1, ""..autotmp_6+192(SP)
	0x0099 00153 (demo.go:8)	MOVQ	$1, ""..autotmp_6+200(SP)
	0x00a5 00165 (demo.go:8)	MOVL	$1, BX
	0x00aa 00170 (demo.go:8)	MOVQ	BX, CX
	0x00ad 00173 (demo.go:8)	CALL	fmt.Println(SB)
	0x00b2 00178 (demo.go:10)	MOVUPS	X15, "".g+72(SP)
	0x00b8 00184 (demo.go:10)	LEAQ	go.string."Go"(SB), DX
	0x00bf 00191 (demo.go:10)	MOVQ	DX, "".g+72(SP)
	0x00c4 00196 (demo.go:10)	MOVQ	$2, "".g+80(SP)
	0x00cd 00205 (demo.go:11)	MOVQ	DX, ""..autotmp_8+136(SP)
	0x00d5 00213 (demo.go:11)	MOVQ	$2, ""..autotmp_8+144(SP)
	0x00e1 00225 (demo.go:11)	LEAQ	""..autotmp_8+136(SP), DX
	0x00e9 00233 (demo.go:11)	TESTB	AL, (DX)
	0x00eb 00235 (demo.go:11)	MOVQ	""..autotmp_8+136(SP), AX
	0x00f3 00243 (demo.go:11)	MOVQ	AX, ""..autotmp_9+120(SP)
	0x00f8 00248 (demo.go:11)	MOVQ	$2, ""..autotmp_9+128(SP)
	0x0104 00260 (demo.go:11)	MOVL	$2, BX
	0x0109 00265 (demo.go:11)	CALL	runtime.convTstring(SB)
	0x010e 00270 (demo.go:11)	MOVQ	AX, ""..autotmp_10+64(SP)
	0x0113 00275 (demo.go:11)	LEAQ	go.itab."".Gopher,"".coder(SB), DX
	0x011a 00282 (demo.go:11)	MOVQ	DX, "".c+88(SP)
	0x011f 00287 (demo.go:11)	MOVQ	AX, "".c+96(SP)
	0x0124 00292 (demo.go:12)	MOVUPS	X15, ""..autotmp_4+152(SP)
	0x012d 00301 (demo.go:12)	LEAQ	""..autotmp_4+152(SP), DX
	0x0135 00309 (demo.go:12)	MOVQ	DX, ""..autotmp_12+56(SP)
	0x013a 00314 (demo.go:12)	MOVQ	"".c+88(SP), DX
	0x013f 00319 (demo.go:12)	MOVQ	"".c+96(SP), SI
	0x0144 00324 (demo.go:12)	MOVQ	DX, ""..autotmp_13+168(SP)
	0x014c 00332 (demo.go:12)	MOVQ	SI, ""..autotmp_13+176(SP)
	0x0154 00340 (demo.go:12)	MOVQ	DX, ""..autotmp_14+48(SP)
	0x0159 00345 (demo.go:12)	CMPQ	""..autotmp_14+48(SP), $0
	0x015f 00351 (demo.go:12)	NOP
	0x0160 00352 (demo.go:12)	JNE	356
	0x0162 00354 (demo.go:12)	JMP	367
	0x0164 00356 (demo.go:12)	MOVQ	8(DX), DX
	0x0168 00360 (demo.go:12)	MOVQ	DX, ""..autotmp_14+48(SP)
	0x016d 00365 (demo.go:12)	JMP	369
	0x016f 00367 (demo.go:12)	PCDATA	$1, $-1
	0x016f 00367 (demo.go:12)	JMP	369
	0x0171 00369 (demo.go:12)	MOVQ	""..autotmp_12+56(SP), DX
	0x0176 00374 (demo.go:12)	TESTB	AL, (DX)
	0x0178 00376 (demo.go:12)	MOVQ	""..autotmp_13+176(SP), SI
	0x0180 00384 (demo.go:12)	MOVQ	""..autotmp_14+48(SP), R8
	0x0185 00389 (demo.go:12)	MOVQ	R8, (DX)
	0x0188 00392 (demo.go:12)	LEAQ	8(DX), DI
	0x018c 00396 (demo.go:12)	PCDATA	$0, $-2
	0x018c 00396 (demo.go:12)	CMPL	runtime.writeBarrier(SB), $0
	0x0193 00403 (demo.go:12)	JEQ	407
	0x0195 00405 (demo.go:12)	JMP	416
	0x0197 00407 (demo.go:12)	MOVQ	SI, 8(DX)
	0x019b 00411 (demo.go:12)	JMP	423
	0x019d 00413 (demo.go:12)	NOP
	0x01a0 00416 (demo.go:12)	CALL	runtime.gcWriteBarrierSI(SB)
	0x01a5 00421 (demo.go:12)	JMP	423
	0x01a7 00423 (demo.go:12)	PCDATA	$0, $-1
	0x01a7 00423 (demo.go:12)	MOVQ	""..autotmp_12+56(SP), AX
	0x01ac 00428 (demo.go:12)	TESTB	AL, (AX)
	0x01ae 00430 (demo.go:12)	JMP	432
	0x01b0 00432 (demo.go:12)	MOVQ	AX, ""..autotmp_11+208(SP)
	0x01b8 00440 (demo.go:12)	MOVQ	$1, ""..autotmp_11+216(SP)
	0x01c4 00452 (demo.go:12)	MOVQ	$1, ""..autotmp_11+224(SP)
	0x01d0 00464 (demo.go:12)	MOVL	$1, BX
	0x01d5 00469 (demo.go:12)	MOVQ	BX, CX
	0x01d8 00472 (demo.go:12)	PCDATA	$1, $0
	0x01d8 00472 (demo.go:12)	CALL	fmt.Println(SB)
	0x01dd 00477 (demo.go:13)	MOVQ	232(SP), BP
	0x01e5 00485 (demo.go:13)	ADDQ	$240, SP
	0x01ec 00492 (demo.go:13)	RET
	0x01ed 00493 (demo.go:13)	NOP
	0x01ed 00493 (demo.go:5)	PCDATA	$1, $-1
	0x01ed 00493 (demo.go:5)	PCDATA	$0, $-2
	0x01ed 00493 (demo.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x01f2 00498 (demo.go:5)	PCDATA	$0, $-1
	0x01f2 00498 (demo.go:5)	JMP	0
	0x0000 4c 8d 64 24 90 4d 3b 66 10 0f 86 de 01 00 00 48  L.d$.M;f.......H
	0x0010 81 ec f0 00 00 00 48 89 ac 24 e8 00 00 00 48 8d  ......H..$....H.
	0x0020 ac 24 e8 00 00 00 48 c7 44 24 18 c8 00 00 00 b8  .$....H.D$......
	0x0030 c8 00 00 00 e8 00 00 00 00 48 89 44 24 28 48 8d  .........H.D$(H.
	0x0040 0d 00 00 00 00 48 89 4c 24 68 48 89 44 24 70 44  .....H.L$hH.D$pD
	0x0050 0f 11 bc 24 98 00 00 00 48 8d 84 24 98 00 00 00  ...$....H..$....
	0x0060 48 89 44 24 20 84 00 48 8b 4c 24 68 48 8b 54 24  H.D$ ..H.L$hH.T$
	0x0070 70 48 89 8c 24 98 00 00 00 48 89 94 24 a0 00 00  pH..$....H..$...
	0x0080 00 84 00 eb 00 48 89 84 24 b8 00 00 00 48 c7 84  .....H..$....H..
	0x0090 24 c0 00 00 00 01 00 00 00 48 c7 84 24 c8 00 00  $........H..$...
	0x00a0 00 01 00 00 00 bb 01 00 00 00 48 89 d9 e8 00 00  ..........H.....
	0x00b0 00 00 44 0f 11 7c 24 48 48 8d 15 00 00 00 00 48  ..D..|$HH......H
	0x00c0 89 54 24 48 48 c7 44 24 50 02 00 00 00 48 89 94  .T$HH.D$P....H..
	0x00d0 24 88 00 00 00 48 c7 84 24 90 00 00 00 02 00 00  $....H..$.......
	0x00e0 00 48 8d 94 24 88 00 00 00 84 02 48 8b 84 24 88  .H..$......H..$.
	0x00f0 00 00 00 48 89 44 24 78 48 c7 84 24 80 00 00 00  ...H.D$xH..$....
	0x0100 02 00 00 00 bb 02 00 00 00 e8 00 00 00 00 48 89  ..............H.
	0x0110 44 24 40 48 8d 15 00 00 00 00 48 89 54 24 58 48  D$@H......H.T$XH
	0x0120 89 44 24 60 44 0f 11 bc 24 98 00 00 00 48 8d 94  .D$`D...$....H..
	0x0130 24 98 00 00 00 48 89 54 24 38 48 8b 54 24 58 48  $....H.T$8H.T$XH
	0x0140 8b 74 24 60 48 89 94 24 a8 00 00 00 48 89 b4 24  .t$`H..$....H..$
	0x0150 b0 00 00 00 48 89 54 24 30 48 83 7c 24 30 00 90  ....H.T$0H.|$0..
	0x0160 75 02 eb 0b 48 8b 52 08 48 89 54 24 30 eb 02 eb  u...H.R.H.T$0...
	0x0170 00 48 8b 54 24 38 84 02 48 8b b4 24 b0 00 00 00  .H.T$8..H..$....
	0x0180 4c 8b 44 24 30 4c 89 02 48 8d 7a 08 83 3d 00 00  L.D$0L..H.z..=..
	0x0190 00 00 00 74 02 eb 09 48 89 72 08 eb 0a 0f 1f 00  ...t...H.r......
	0x01a0 e8 00 00 00 00 eb 00 48 8b 44 24 38 84 00 eb 00  .......H.D$8....
	0x01b0 48 89 84 24 d0 00 00 00 48 c7 84 24 d8 00 00 00  H..$....H..$....
	0x01c0 01 00 00 00 48 c7 84 24 e0 00 00 00 01 00 00 00  ....H..$........
	0x01d0 bb 01 00 00 00 48 89 d9 e8 00 00 00 00 48 8b ac  .....H.......H..
	0x01e0 24 e8 00 00 00 48 81 c4 f0 00 00 00 c3 e8 00 00  $....H..........
	0x01f0 00 00 e9 09 fe ff ff                             .......
	rel 3+0 t=23 type.int+0
	rel 3+0 t=23 type."".Gopher+0
	rel 53+4 t=7 runtime.convT64+0
	rel 65+4 t=14 type.int+0
	rel 174+4 t=7 fmt.Println+0
	rel 187+4 t=14 go.string."Go"+0
	rel 266+4 t=7 runtime.convTstring+0
	rel 278+4 t=14 go.itab."".Gopher,"".coder+0
	rel 398+4 t=14 runtime.writeBarrier+-1
	rel 417+4 t=7 runtime.gcWriteBarrierSI+0
	rel 473+4 t=7 fmt.Println+0
	rel 494+4 t=7 runtime.morestack_noctxt+0
"".Gopher.code STEXT size=218 args=0x10 locals=0x68 funcid=0x0 align=0x0
	0x0000 00000 (demo.go:24)	TEXT	"".Gopher.code(SB), ABIInternal, $104-16
	0x0000 00000 (demo.go:24)	CMPQ	SP, 16(R14)
	0x0004 00004 (demo.go:24)	PCDATA	$0, $-2
	0x0004 00004 (demo.go:24)	JLS	188
	0x000a 00010 (demo.go:24)	PCDATA	$0, $-1
	0x000a 00010 (demo.go:24)	SUBQ	$104, SP
	0x000e 00014 (demo.go:24)	MOVQ	BP, 96(SP)
	0x0013 00019 (demo.go:24)	LEAQ	96(SP), BP
	0x0018 00024 (demo.go:24)	FUNCDATA	$0, gclocals·2d7c1615616d4cf40d01b3385155ed6e(SB)
	0x0018 00024 (demo.go:24)	FUNCDATA	$1, gclocals·151237c263d0ac053c215b44512d9922(SB)
	0x0018 00024 (demo.go:24)	FUNCDATA	$2, "".Gopher.code.stkobj(SB)
	0x0018 00024 (demo.go:24)	FUNCDATA	$5, "".Gopher.code.arginfo1(SB)
	0x0018 00024 (demo.go:24)	MOVQ	AX, "".p+112(SP)
	0x001d 00029 (demo.go:24)	MOVQ	BX, "".p+120(SP)
	0x0022 00034 (demo.go:25)	MOVUPS	X15, ""..autotmp_1+56(SP)
	0x0028 00040 (demo.go:25)	LEAQ	""..autotmp_1+56(SP), CX
	0x002d 00045 (demo.go:25)	MOVQ	CX, ""..autotmp_3+48(SP)
	0x0032 00050 (demo.go:25)	MOVQ	"".p+112(SP), AX
	0x0037 00055 (demo.go:25)	MOVQ	"".p+120(SP), BX
	0x003c 00060 (demo.go:25)	PCDATA	$1, $1
	0x003c 00060 (demo.go:25)	NOP
	0x0040 00064 (demo.go:25)	CALL	runtime.convTstring(SB)
	0x0045 00069 (demo.go:25)	MOVQ	AX, ""..autotmp_4+40(SP)
	0x004a 00074 (demo.go:25)	MOVQ	""..autotmp_3+48(SP), CX
	0x004f 00079 (demo.go:25)	TESTB	AL, (CX)
	0x0051 00081 (demo.go:25)	LEAQ	type.string(SB), DX
	0x0058 00088 (demo.go:25)	MOVQ	DX, (CX)
	0x005b 00091 (demo.go:25)	LEAQ	8(CX), DI
	0x005f 00095 (demo.go:25)	PCDATA	$0, $-2
	0x005f 00095 (demo.go:25)	CMPL	runtime.writeBarrier(SB), $0
	0x0066 00102 (demo.go:25)	JEQ	106
	0x0068 00104 (demo.go:25)	JMP	112
	0x006a 00106 (demo.go:25)	MOVQ	AX, 8(CX)
	0x006e 00110 (demo.go:25)	JMP	119
	0x0070 00112 (demo.go:25)	CALL	runtime.gcWriteBarrier(SB)
	0x0075 00117 (demo.go:25)	JMP	119
	0x0077 00119 (demo.go:25)	PCDATA	$0, $-1
	0x0077 00119 (demo.go:25)	MOVQ	""..autotmp_3+48(SP), CX
	0x007c 00124 (demo.go:25)	TESTB	AL, (CX)
	0x007e 00126 (demo.go:25)	NOP
	0x0080 00128 (demo.go:25)	JMP	130
	0x0082 00130 (demo.go:25)	MOVQ	CX, ""..autotmp_2+72(SP)
	0x0087 00135 (demo.go:25)	MOVQ	$1, ""..autotmp_2+80(SP)
	0x0090 00144 (demo.go:25)	MOVQ	$1, ""..autotmp_2+88(SP)
	0x0099 00153 (demo.go:25)	LEAQ	go.string."I am coding %s language\n"(SB), AX
	0x00a0 00160 (demo.go:25)	MOVL	$24, BX
	0x00a5 00165 (demo.go:25)	MOVL	$1, DI
	0x00aa 00170 (demo.go:25)	MOVQ	DI, SI
	0x00ad 00173 (demo.go:25)	PCDATA	$1, $2
	0x00ad 00173 (demo.go:25)	CALL	fmt.Printf(SB)
	0x00b2 00178 (demo.go:26)	MOVQ	96(SP), BP
	0x00b7 00183 (demo.go:26)	ADDQ	$104, SP
	0x00bb 00187 (demo.go:26)	RET
	0x00bc 00188 (demo.go:26)	NOP
	0x00bc 00188 (demo.go:24)	PCDATA	$1, $-1
	0x00bc 00188 (demo.go:24)	PCDATA	$0, $-2
	0x00bc 00188 (demo.go:24)	MOVQ	AX, 8(SP)
	0x00c1 00193 (demo.go:24)	MOVQ	BX, 16(SP)
	0x00c6 00198 (demo.go:24)	CALL	runtime.morestack_noctxt(SB)
	0x00cb 00203 (demo.go:24)	MOVQ	8(SP), AX
	0x00d0 00208 (demo.go:24)	MOVQ	16(SP), BX
	0x00d5 00213 (demo.go:24)	PCDATA	$0, $-1
	0x00d5 00213 (demo.go:24)	JMP	0
	0x0000 49 3b 66 10 0f 86 b2 00 00 00 48 83 ec 68 48 89  I;f.......H..hH.
	0x0010 6c 24 60 48 8d 6c 24 60 48 89 44 24 70 48 89 5c  l$`H.l$`H.D$pH.\
	0x0020 24 78 44 0f 11 7c 24 38 48 8d 4c 24 38 48 89 4c  $xD..|$8H.L$8H.L
	0x0030 24 30 48 8b 44 24 70 48 8b 5c 24 78 0f 1f 40 00  $0H.D$pH.\$x..@.
	0x0040 e8 00 00 00 00 48 89 44 24 28 48 8b 4c 24 30 84  .....H.D$(H.L$0.
	0x0050 01 48 8d 15 00 00 00 00 48 89 11 48 8d 79 08 83  .H......H..H.y..
	0x0060 3d 00 00 00 00 00 74 02 eb 06 48 89 41 08 eb 07  =.....t...H.A...
	0x0070 e8 00 00 00 00 eb 00 48 8b 4c 24 30 84 01 66 90  .......H.L$0..f.
	0x0080 eb 00 48 89 4c 24 48 48 c7 44 24 50 01 00 00 00  ..H.L$HH.D$P....
	0x0090 48 c7 44 24 58 01 00 00 00 48 8d 05 00 00 00 00  H.D$X....H......
	0x00a0 bb 18 00 00 00 bf 01 00 00 00 48 89 fe e8 00 00  ..........H.....
	0x00b0 00 00 48 8b 6c 24 60 48 83 c4 68 c3 48 89 44 24  ..H.l$`H..h.H.D$
	0x00c0 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08  .H.\$......H.D$.
	0x00d0 48 8b 5c 24 10 e9 26 ff ff ff                    H.\$..&...
	rel 3+0 t=23 type.string+0
	rel 65+4 t=7 runtime.convTstring+0
	rel 84+4 t=14 type.string+0
	rel 97+4 t=14 runtime.writeBarrier+-1
	rel 113+4 t=7 runtime.gcWriteBarrier+0
	rel 156+4 t=14 go.string."I am coding %s language\n"+0
	rel 174+4 t=7 fmt.Printf+0
	rel 199+4 t=7 runtime.morestack_noctxt+0
"".Gopher.debug STEXT size=218 args=0x10 locals=0x68 funcid=0x0 align=0x0
	0x0000 00000 (demo.go:28)	TEXT	"".Gopher.debug(SB), ABIInternal, $104-16
	0x0000 00000 (demo.go:28)	CMPQ	SP, 16(R14)
	0x0004 00004 (demo.go:28)	PCDATA	$0, $-2
	0x0004 00004 (demo.go:28)	JLS	188
	0x000a 00010 (demo.go:28)	PCDATA	$0, $-1
	0x000a 00010 (demo.go:28)	SUBQ	$104, SP
	0x000e 00014 (demo.go:28)	MOVQ	BP, 96(SP)
	0x0013 00019 (demo.go:28)	LEAQ	96(SP), BP
	0x0018 00024 (demo.go:28)	FUNCDATA	$0, gclocals·2d7c1615616d4cf40d01b3385155ed6e(SB)
	0x0018 00024 (demo.go:28)	FUNCDATA	$1, gclocals·151237c263d0ac053c215b44512d9922(SB)
	0x0018 00024 (demo.go:28)	FUNCDATA	$2, "".Gopher.debug.stkobj(SB)
	0x0018 00024 (demo.go:28)	FUNCDATA	$5, "".Gopher.debug.arginfo1(SB)
	0x0018 00024 (demo.go:28)	MOVQ	AX, "".p+112(SP)
	0x001d 00029 (demo.go:28)	MOVQ	BX, "".p+120(SP)
	0x0022 00034 (demo.go:29)	MOVUPS	X15, ""..autotmp_1+56(SP)
	0x0028 00040 (demo.go:29)	LEAQ	""..autotmp_1+56(SP), CX
	0x002d 00045 (demo.go:29)	MOVQ	CX, ""..autotmp_3+48(SP)
	0x0032 00050 (demo.go:29)	MOVQ	"".p+112(SP), AX
	0x0037 00055 (demo.go:29)	MOVQ	"".p+120(SP), BX
	0x003c 00060 (demo.go:29)	PCDATA	$1, $1
	0x003c 00060 (demo.go:29)	NOP
	0x0040 00064 (demo.go:29)	CALL	runtime.convTstring(SB)
	0x0045 00069 (demo.go:29)	MOVQ	AX, ""..autotmp_4+40(SP)
	0x004a 00074 (demo.go:29)	MOVQ	""..autotmp_3+48(SP), CX
	0x004f 00079 (demo.go:29)	TESTB	AL, (CX)
	0x0051 00081 (demo.go:29)	LEAQ	type.string(SB), DX
	0x0058 00088 (demo.go:29)	MOVQ	DX, (CX)
	0x005b 00091 (demo.go:29)	LEAQ	8(CX), DI
	0x005f 00095 (demo.go:29)	PCDATA	$0, $-2
	0x005f 00095 (demo.go:29)	CMPL	runtime.writeBarrier(SB), $0
	0x0066 00102 (demo.go:29)	JEQ	106
	0x0068 00104 (demo.go:29)	JMP	112
	0x006a 00106 (demo.go:29)	MOVQ	AX, 8(CX)
	0x006e 00110 (demo.go:29)	JMP	119
	0x0070 00112 (demo.go:29)	CALL	runtime.gcWriteBarrier(SB)
	0x0075 00117 (demo.go:29)	JMP	119
	0x0077 00119 (demo.go:29)	PCDATA	$0, $-1
	0x0077 00119 (demo.go:29)	MOVQ	""..autotmp_3+48(SP), CX
	0x007c 00124 (demo.go:29)	TESTB	AL, (CX)
	0x007e 00126 (demo.go:29)	NOP
	0x0080 00128 (demo.go:29)	JMP	130
	0x0082 00130 (demo.go:29)	MOVQ	CX, ""..autotmp_2+72(SP)
	0x0087 00135 (demo.go:29)	MOVQ	$1, ""..autotmp_2+80(SP)
	0x0090 00144 (demo.go:29)	MOVQ	$1, ""..autotmp_2+88(SP)
	0x0099 00153 (demo.go:29)	LEAQ	go.string."I am debuging %s language\n"(SB), AX
	0x00a0 00160 (demo.go:29)	MOVL	$26, BX
	0x00a5 00165 (demo.go:29)	MOVL	$1, DI
	0x00aa 00170 (demo.go:29)	MOVQ	DI, SI
	0x00ad 00173 (demo.go:29)	PCDATA	$1, $2
	0x00ad 00173 (demo.go:29)	CALL	fmt.Printf(SB)
	0x00b2 00178 (demo.go:30)	MOVQ	96(SP), BP
	0x00b7 00183 (demo.go:30)	ADDQ	$104, SP
	0x00bb 00187 (demo.go:30)	RET
	0x00bc 00188 (demo.go:30)	NOP
	0x00bc 00188 (demo.go:28)	PCDATA	$1, $-1
	0x00bc 00188 (demo.go:28)	PCDATA	$0, $-2
	0x00bc 00188 (demo.go:28)	MOVQ	AX, 8(SP)
	0x00c1 00193 (demo.go:28)	MOVQ	BX, 16(SP)
	0x00c6 00198 (demo.go:28)	CALL	runtime.morestack_noctxt(SB)
	0x00cb 00203 (demo.go:28)	MOVQ	8(SP), AX
	0x00d0 00208 (demo.go:28)	MOVQ	16(SP), BX
	0x00d5 00213 (demo.go:28)	PCDATA	$0, $-1
	0x00d5 00213 (demo.go:28)	JMP	0
	0x0000 49 3b 66 10 0f 86 b2 00 00 00 48 83 ec 68 48 89  I;f.......H..hH.
	0x0010 6c 24 60 48 8d 6c 24 60 48 89 44 24 70 48 89 5c  l$`H.l$`H.D$pH.\
	0x0020 24 78 44 0f 11 7c 24 38 48 8d 4c 24 38 48 89 4c  $xD..|$8H.L$8H.L
	0x0030 24 30 48 8b 44 24 70 48 8b 5c 24 78 0f 1f 40 00  $0H.D$pH.\$x..@.
	0x0040 e8 00 00 00 00 48 89 44 24 28 48 8b 4c 24 30 84  .....H.D$(H.L$0.
	0x0050 01 48 8d 15 00 00 00 00 48 89 11 48 8d 79 08 83  .H......H..H.y..
	0x0060 3d 00 00 00 00 00 74 02 eb 06 48 89 41 08 eb 07  =.....t...H.A...
	0x0070 e8 00 00 00 00 eb 00 48 8b 4c 24 30 84 01 66 90  .......H.L$0..f.
	0x0080 eb 00 48 89 4c 24 48 48 c7 44 24 50 01 00 00 00  ..H.L$HH.D$P....
	0x0090 48 c7 44 24 58 01 00 00 00 48 8d 05 00 00 00 00  H.D$X....H......
	0x00a0 bb 1a 00 00 00 bf 01 00 00 00 48 89 fe e8 00 00  ..........H.....
	0x00b0 00 00 48 8b 6c 24 60 48 83 c4 68 c3 48 89 44 24  ..H.l$`H..h.H.D$
	0x00c0 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08  .H.\$......H.D$.
	0x00d0 48 8b 5c 24 10 e9 26 ff ff ff                    H.\$..&...
	rel 3+0 t=23 type.string+0
	rel 65+4 t=7 runtime.convTstring+0
	rel 84+4 t=14 type.string+0
	rel 97+4 t=14 runtime.writeBarrier+-1
	rel 113+4 t=7 runtime.gcWriteBarrier+0
	rel 156+4 t=14 go.string."I am debuging %s language\n"+0
	rel 174+4 t=7 fmt.Printf+0
	rel 199+4 t=7 runtime.morestack_noctxt+0
"".(*Gopher).code STEXT dupok size=119 args=0x8 locals=0x28 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*Gopher).code(SB), DUPOK|WRAPPER|ABIInternal, $40-8
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	85
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$40, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 32(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	32(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	102
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·2589ca35330fc0fce83503f4569854a0(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, "".(*Gopher).code.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+48(SP)
	0x0022 00034 (<autogenerated>:1)	TESTQ	AX, AX
	0x0025 00037 (<autogenerated>:1)	JNE	41
	0x0027 00039 (<autogenerated>:1)	JMP	79
	0x0029 00041 (<autogenerated>:1)	TESTB	AL, (AX)
	0x002b 00043 (<autogenerated>:1)	MOVQ	(AX), CX
	0x002e 00046 (<autogenerated>:1)	MOVQ	8(AX), BX
	0x0032 00050 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_1+16(SP)
	0x0037 00055 (<autogenerated>:1)	MOVQ	BX, ""..autotmp_1+24(SP)
	0x003c 00060 (<autogenerated>:1)	MOVQ	CX, AX
	0x003f 00063 (<autogenerated>:1)	PCDATA	$1, $1
	0x003f 00063 (<autogenerated>:1)	NOP
	0x0040 00064 (<autogenerated>:1)	CALL	"".Gopher.code(SB)
	0x0045 00069 (<autogenerated>:1)	MOVQ	32(SP), BP
	0x004a 00074 (<autogenerated>:1)	ADDQ	$40, SP
	0x004e 00078 (<autogenerated>:1)	RET
	0x004f 00079 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x0054 00084 (<autogenerated>:1)	XCHGL	AX, AX
	0x0055 00085 (<autogenerated>:1)	NOP
	0x0055 00085 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0055 00085 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x005a 00090 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005f 00095 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0064 00100 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0064 00100 (<autogenerated>:1)	JMP	0
	0x0066 00102 (<autogenerated>:1)	LEAQ	48(SP), R13
	0x006b 00107 (<autogenerated>:1)	CMPQ	(R12), R13
	0x006f 00111 (<autogenerated>:1)	JNE	29
	0x0071 00113 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x0075 00117 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 4f 48 83 ec 28 48 89 6c 24 20 48  I;f.vOH..(H.l$ H
	0x0010 8d 6c 24 20 4d 8b 66 20 4d 85 e4 75 49 48 89 44  .l$ M.f M..uIH.D
	0x0020 24 30 48 85 c0 75 02 eb 26 84 00 48 8b 08 48 8b  $0H..u..&..H..H.
	0x0030 58 08 48 89 4c 24 10 48 89 5c 24 18 48 89 c8 90  X.H.L$.H.\$.H...
	0x0040 e8 00 00 00 00 48 8b 6c 24 20 48 83 c4 28 c3 e8  .....H.l$ H..(..
	0x0050 00 00 00 00 90 48 89 44 24 08 e8 00 00 00 00 48  .....H.D$......H
	0x0060 8b 44 24 08 eb 9a 4c 8d 6c 24 30 4d 39 2c 24 75  .D$...L.l$0M9,$u
	0x0070 ac 49 89 24 24 eb a6                             .I.$$..
	rel 65+4 t=7 "".Gopher.code+0
	rel 80+4 t=7 runtime.panicwrap+0
	rel 91+4 t=7 runtime.morestack_noctxt+0
"".(*Gopher).debug STEXT dupok size=119 args=0x8 locals=0x28 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*Gopher).debug(SB), DUPOK|WRAPPER|ABIInternal, $40-8
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	85
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$40, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 32(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	32(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	102
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·2589ca35330fc0fce83503f4569854a0(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, "".(*Gopher).debug.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+48(SP)
	0x0022 00034 (<autogenerated>:1)	TESTQ	AX, AX
	0x0025 00037 (<autogenerated>:1)	JNE	41
	0x0027 00039 (<autogenerated>:1)	JMP	79
	0x0029 00041 (<autogenerated>:1)	TESTB	AL, (AX)
	0x002b 00043 (<autogenerated>:1)	MOVQ	(AX), CX
	0x002e 00046 (<autogenerated>:1)	MOVQ	8(AX), BX
	0x0032 00050 (<autogenerated>:1)	MOVQ	CX, ""..autotmp_1+16(SP)
	0x0037 00055 (<autogenerated>:1)	MOVQ	BX, ""..autotmp_1+24(SP)
	0x003c 00060 (<autogenerated>:1)	MOVQ	CX, AX
	0x003f 00063 (<autogenerated>:1)	PCDATA	$1, $1
	0x003f 00063 (<autogenerated>:1)	NOP
	0x0040 00064 (<autogenerated>:1)	CALL	"".Gopher.debug(SB)
	0x0045 00069 (<autogenerated>:1)	MOVQ	32(SP), BP
	0x004a 00074 (<autogenerated>:1)	ADDQ	$40, SP
	0x004e 00078 (<autogenerated>:1)	RET
	0x004f 00079 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x0054 00084 (<autogenerated>:1)	XCHGL	AX, AX
	0x0055 00085 (<autogenerated>:1)	NOP
	0x0055 00085 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0055 00085 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x005a 00090 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005f 00095 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0064 00100 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0064 00100 (<autogenerated>:1)	JMP	0
	0x0066 00102 (<autogenerated>:1)	LEAQ	48(SP), R13
	0x006b 00107 (<autogenerated>:1)	CMPQ	(R12), R13
	0x006f 00111 (<autogenerated>:1)	JNE	29
	0x0071 00113 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x0075 00117 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 4f 48 83 ec 28 48 89 6c 24 20 48  I;f.vOH..(H.l$ H
	0x0010 8d 6c 24 20 4d 8b 66 20 4d 85 e4 75 49 48 89 44  .l$ M.f M..uIH.D
	0x0020 24 30 48 85 c0 75 02 eb 26 84 00 48 8b 08 48 8b  $0H..u..&..H..H.
	0x0030 58 08 48 89 4c 24 10 48 89 5c 24 18 48 89 c8 90  X.H.L$.H.\$.H...
	0x0040 e8 00 00 00 00 48 8b 6c 24 20 48 83 c4 28 c3 e8  .....H.l$ H..(..
	0x0050 00 00 00 00 90 48 89 44 24 08 e8 00 00 00 00 48  .....H.D$......H
	0x0060 8b 44 24 08 eb 9a 4c 8d 6c 24 30 4d 39 2c 24 75  .D$...L.l$0M9,$u
	0x0070 ac 49 89 24 24 eb a6                             .I.$$..
	rel 65+4 t=7 "".Gopher.debug+0
	rel 80+4 t=7 runtime.panicwrap+0
	rel 91+4 t=7 runtime.morestack_noctxt+0
"".coder.code STEXT dupok size=108 args=0x10 locals=0x10 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".coder.code(SB), DUPOK|WRAPPER|ABIInternal, $16-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	60
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	87
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, "".coder.code.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(SP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	BX, ""..this+32(SP)
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	MOVQ	24(AX), CX
	0x002d 00045 (<autogenerated>:1)	MOVQ	BX, AX
	0x0030 00048 (<autogenerated>:1)	PCDATA	$1, $1
	0x0030 00048 (<autogenerated>:1)	CALL	CX
	0x0032 00050 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x0037 00055 (<autogenerated>:1)	ADDQ	$16, SP
	0x003b 00059 (<autogenerated>:1)	RET
	0x003c 00060 (<autogenerated>:1)	NOP
	0x003c 00060 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003c 00060 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0041 00065 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x0046 00070 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0050 00080 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	JMP	0
	0x0057 00087 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x005c 00092 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0064 00100 (<autogenerated>:1)	JNE	29
	0x0066 00102 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x006a 00106 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 36 48 83 ec 10 48 89 6c 24 08 48  I;f.v6H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 3a 48 89 44  .l$.M.f M..u:H.D
	0x0020 24 18 48 89 5c 24 20 84 00 48 8b 48 18 48 89 d8  $.H.\$ ..H.H.H..
	0x0030 ff d1 48 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24  ..H.l$.H....H.D$
	0x0040 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08  .H.\$......H.D$.
	0x0050 48 8b 5c 24 10 eb a9 4c 8d 6c 24 18 0f 1f 40 00  H.\$...L.l$...@.
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 2+0 t=24 type."".coder+96
	rel 48+0 t=10 +0
	rel 71+4 t=7 runtime.morestack_noctxt+0
"".coder.debug STEXT dupok size=108 args=0x10 locals=0x10 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".coder.debug(SB), DUPOK|WRAPPER|ABIInternal, $16-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	60
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	87
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, "".coder.debug.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(SP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	BX, ""..this+32(SP)
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	MOVQ	32(AX), CX
	0x002d 00045 (<autogenerated>:1)	MOVQ	BX, AX
	0x0030 00048 (<autogenerated>:1)	PCDATA	$1, $1
	0x0030 00048 (<autogenerated>:1)	CALL	CX
	0x0032 00050 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x0037 00055 (<autogenerated>:1)	ADDQ	$16, SP
	0x003b 00059 (<autogenerated>:1)	RET
	0x003c 00060 (<autogenerated>:1)	NOP
	0x003c 00060 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003c 00060 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0041 00065 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x0046 00070 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0050 00080 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	JMP	0
	0x0057 00087 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x005c 00092 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0064 00100 (<autogenerated>:1)	JNE	29
	0x0066 00102 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x006a 00106 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 36 48 83 ec 10 48 89 6c 24 08 48  I;f.v6H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 3a 48 89 44  .l$.M.f M..u:H.D
	0x0020 24 18 48 89 5c 24 20 84 00 48 8b 48 20 48 89 d8  $.H.\$ ..H.H H..
	0x0030 ff d1 48 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24  ..H.l$.H....H.D$
	0x0040 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08  .H.\$......H.D$.
	0x0050 48 8b 5c 24 10 eb a9 4c 8d 6c 24 18 0f 1f 40 00  H.\$...L.l$...@.
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 2+0 t=24 type."".coder+104
	rel 48+0 t=10 +0
	rel 71+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.string."Go" SRODATA dupok size=2
	0x0000 47 6f                                            Go
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*func()- SRODATA dupok size=9
	0x0000 00 07 2a 66 75 6e 63 28 29                       ..*func()
type.*func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9b 90 75 1b 08 08 08 36 00 00 00 00 00 00 00 00  ..u....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 48+8 t=1 type.func()+0
type.func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f6 bc 82 f6 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00                                      ....
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 44+4 t=-32763 type.*func()+0
runtime.interequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.interequal+0
type..namedata.*main.coder- SRODATA dupok size=13
	0x0000 00 0b 2a 6d 61 69 6e 2e 63 6f 64 65 72           ..*main.coder
type.*"".coder SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 35 a8 c6 33 08 08 08 36 00 00 00 00 00 00 00 00  5..3...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.coder-+0
	rel 48+8 t=1 type."".coder+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type..namedata.code- SRODATA dupok size=6
	0x0000 00 04 63 6f 64 65                                ..code
type..namedata.debug- SRODATA dupok size=7
	0x0000 00 05 64 65 62 75 67                             ..debug
type."".coder SRODATA dupok size=112
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 d0 2f 69 91 07 08 08 14 00 00 00 00 00 00 00 00  ./i.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.interequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*main.coder-+0
	rel 44+4 t=5 type.*"".coder+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".coder+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+4 t=5 type..namedata.code-+0
	rel 100+4 t=5 type.func()+0
	rel 104+4 t=5 type..namedata.debug-+0
	rel 108+4 t=5 type.func()+0
runtime.strequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.strequal+0
type..namedata.*main.Gopher. SRODATA dupok size=14
	0x0000 01 0c 2a 6d 61 69 6e 2e 47 6f 70 68 65 72        ..*main.Gopher
type..namedata.*func(*main.Gopher)- SRODATA dupok size=21
	0x0000 00 13 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 47 6f  ..*func(*main.Go
	0x0010 70 68 65 72 29                                   pher)
type.*func(*"".Gopher) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a3 c4 a7 0a 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Gopher)-+0
	rel 48+8 t=1 type.func(*"".Gopher)+0
type.func(*"".Gopher) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 e1 b0 70 ba 02 08 08 33 00 00 00 00 00 00 00 00  ..p....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Gopher)-+0
	rel 44+4 t=-32763 type.*func(*"".Gopher)+0
	rel 56+8 t=1 type.*"".Gopher+0
type.*"".Gopher SRODATA dupok size=104
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ed 07 27 8c 09 08 08 36 00 00 00 00 00 00 00 00  ..'....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 02 00 00 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Gopher.+0
	rel 48+8 t=1 type."".Gopher+0
	rel 56+4 t=5 type..importpath."".+0
	rel 72+4 t=5 type..namedata.code-+0
	rel 76+4 t=26 type.func()+0
	rel 80+4 t=26 "".(*Gopher).code+0
	rel 84+4 t=26 "".(*Gopher).code+0
	rel 88+4 t=5 type..namedata.debug-+0
	rel 92+4 t=26 type.func()+0
	rel 96+4 t=26 "".(*Gopher).debug+0
	rel 100+4 t=26 "".(*Gopher).debug+0
type..namedata.*func(main.Gopher)- SRODATA dupok size=20
	0x0000 00 12 2a 66 75 6e 63 28 6d 61 69 6e 2e 47 6f 70  ..*func(main.Gop
	0x0010 68 65 72 29                                      her)
type.*func("".Gopher) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 fe b8 05 90 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Gopher)-+0
	rel 48+8 t=1 type.func("".Gopher)+0
type.func("".Gopher) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 40 58 c1 7c 02 08 08 33 00 00 00 00 00 00 00 00  @X.|...3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Gopher)-+0
	rel 44+4 t=-32763 type.*func("".Gopher)+0
	rel 56+8 t=1 type."".Gopher+0
type..namedata.language- SRODATA dupok size=10
	0x0000 00 08 6c 61 6e 67 75 61 67 65                    ..language
type."".Gopher SRODATA dupok size=152
	0x0000 10 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 02 2c fd 86 07 08 08 19 00 00 00 00 00 00 00 00  .,..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 02 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.strequal·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Gopher.+0
	rel 44+4 t=5 type.*"".Gopher+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".Gopher+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.language-+0
	rel 104+8 t=1 type.string+0
	rel 120+4 t=5 type..namedata.code-+0
	rel 124+4 t=26 type.func()+0
	rel 128+4 t=26 "".(*Gopher).code+0
	rel 132+4 t=26 "".Gopher.code+0
	rel 136+4 t=5 type..namedata.debug-+0
	rel 140+4 t=26 type.func()+0
	rel 144+4 t=26 "".(*Gopher).debug+0
	rel 148+4 t=26 "".Gopher.debug+0
go.itab."".Gopher,"".coder SRODATA dupok size=40
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 02 2c fd 86 00 00 00 00 00 00 00 00 00 00 00 00  .,..............
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type."".coder+0
	rel 8+8 t=1 type."".Gopher+0
	rel 24+8 t=-32767 "".(*Gopher).code+0
	rel 32+8 t=-32767 "".(*Gopher).debug+0
go.string."I am coding %s language\n" SRODATA dupok size=24
	0x0000 49 20 61 6d 20 63 6f 64 69 6e 67 20 25 73 20 6c  I am coding %s l
	0x0010 61 6e 67 75 61 67 65 0a                          anguage.
go.string."I am debuging %s language\n" SRODATA dupok size=26
	0x0000 49 20 61 6d 20 64 65 62 75 67 69 6e 67 20 25 73  I am debuging %s
	0x0010 20 6c 61 6e 67 75 61 67 65 0a                     language.
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·28835bfeb0e793947754a03b87c8a295 SRODATA dupok size=12
	0x0000 01 00 00 00 19 00 00 00 00 00 00 00              ............
"".main.stkobj SRODATA static size=40
	0x0000 02 00 00 00 00 00 00 00 a0 ff ff ff 10 00 00 00  ................
	0x0010 08 00 00 00 00 00 00 00 b0 ff ff ff 10 00 00 00  ................
	0x0020 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.01+0
	rel 36+4 t=5 runtime.gcbits.02+0
gclocals·2d7c1615616d4cf40d01b3385155ed6e SRODATA dupok size=11
	0x0000 03 00 00 00 01 00 00 00 01 00 00                 ...........
gclocals·151237c263d0ac053c215b44512d9922 SRODATA dupok size=11
	0x0000 03 00 00 00 07 00 00 00 00 02 00                 ...........
"".Gopher.code.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
"".Gopher.code.arginfo1 SRODATA static dupok size=9
	0x0000 fe fe 00 08 08 08 fd fd ff                       .........
"".Gopher.debug.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
"".Gopher.debug.arginfo1 SRODATA static dupok size=9
	0x0000 fe fe 00 08 08 08 fd fd ff                       .........
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·2589ca35330fc0fce83503f4569854a0 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 00                    ..........
"".(*Gopher).code.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".(*Gopher).debug.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
gclocals·09cf9819fc716118c209c2d2155a3632 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 02 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
"".coder.code.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
"".coder.debug.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
"".Foo STEXT size=135 args=0x10 locals=0x10 funcid=0x0 align=0x0
	0x0000 00000 (demo.go:10)	TEXT	"".Foo(SB), ABIInternal, $16-16
	0x0000 00000 (demo.go:10)	CMPQ	SP, 16(R14)
	0x0004 00004 (demo.go:10)	PCDATA	$0, $-2
	0x0004 00004 (demo.go:10)	JLS	105
	0x0006 00006 (demo.go:10)	PCDATA	$0, $-1
	0x0006 00006 (demo.go:10)	SUBQ	$16, SP
	0x000a 00010 (demo.go:10)	MOVQ	BP, 8(SP)
	0x000f 00015 (demo.go:10)	LEAQ	8(SP), BP
	0x0014 00020 (demo.go:10)	FUNCDATA	$0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x0014 00020 (demo.go:10)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0014 00020 (demo.go:10)	FUNCDATA	$5, "".Foo.arginfo1(SB)
	0x0014 00020 (demo.go:10)	MOVQ	AX, "".me+24(SP)
	0x0019 00025 (demo.go:10)	MOVQ	BX, "".me+32(SP)
	0x001e 00030 (demo.go:11)	TESTB	AL, (AX)
	0x0020 00032 (demo.go:11)	MOVQ	40(AX), CX
	0x0024 00036 (demo.go:11)	MOVQ	BX, AX
	0x0027 00039 (demo.go:11)	PCDATA	$1, $0
	0x0027 00039 (demo.go:11)	CALL	CX
	0x0029 00041 (demo.go:12)	MOVQ	"".me+24(SP), CX
	0x002e 00046 (demo.go:12)	TESTB	AL, (CX)
	0x0030 00048 (demo.go:12)	MOVQ	32(CX), CX
	0x0034 00052 (demo.go:12)	MOVQ	"".me+32(SP), AX
	0x0039 00057 (demo.go:12)	CALL	CX
	0x003b 00059 (demo.go:13)	MOVQ	"".me+24(SP), CX
	0x0040 00064 (demo.go:13)	TESTB	AL, (CX)
	0x0042 00066 (demo.go:13)	MOVQ	48(CX), CX
	0x0046 00070 (demo.go:13)	MOVQ	"".me+32(SP), AX
	0x004b 00075 (demo.go:13)	CALL	CX
	0x004d 00077 (demo.go:14)	MOVQ	"".me+24(SP), CX
	0x0052 00082 (demo.go:14)	TESTB	AL, (CX)
	0x0054 00084 (demo.go:14)	MOVQ	24(CX), CX
	0x0058 00088 (demo.go:14)	MOVQ	"".me+32(SP), AX
	0x005d 00093 (demo.go:14)	PCDATA	$1, $1
	0x005d 00093 (demo.go:14)	CALL	CX
	0x005f 00095 (demo.go:15)	MOVQ	8(SP), BP
	0x0064 00100 (demo.go:15)	ADDQ	$16, SP
	0x0068 00104 (demo.go:15)	RET
	0x0069 00105 (demo.go:15)	NOP
	0x0069 00105 (demo.go:10)	PCDATA	$1, $-1
	0x0069 00105 (demo.go:10)	PCDATA	$0, $-2
	0x0069 00105 (demo.go:10)	MOVQ	AX, 8(SP)
	0x006e 00110 (demo.go:10)	MOVQ	BX, 16(SP)
	0x0073 00115 (demo.go:10)	CALL	runtime.morestack_noctxt(SB)
	0x0078 00120 (demo.go:10)	MOVQ	8(SP), AX
	0x007d 00125 (demo.go:10)	MOVQ	16(SP), BX
	0x0082 00130 (demo.go:10)	PCDATA	$0, $-1
	0x0082 00130 (demo.go:10)	JMP	0
	0x0000 49 3b 66 10 76 63 48 83 ec 10 48 89 6c 24 08 48  I;f.vcH...H.l$.H
	0x0010 8d 6c 24 08 48 89 44 24 18 48 89 5c 24 20 84 00  .l$.H.D$.H.\$ ..
	0x0020 48 8b 48 28 48 89 d8 ff d1 48 8b 4c 24 18 84 01  H.H(H....H.L$...
	0x0030 48 8b 49 20 48 8b 44 24 20 ff d1 48 8b 4c 24 18  H.I H.D$ ..H.L$.
	0x0040 84 01 48 8b 49 30 48 8b 44 24 20 ff d1 48 8b 4c  ..H.I0H.D$ ..H.L
	0x0050 24 18 84 01 48 8b 49 18 48 8b 44 24 20 ff d1 48  $...H.I.H.D$ ..H
	0x0060 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24 08 48 89  .l$.H....H.D$.H.
	0x0070 5c 24 10 e8 00 00 00 00 48 8b 44 24 08 48 8b 5c  \$......H.D$.H.\
	0x0080 24 10 e9 79 ff ff ff                             $..y...
	rel 2+0 t=24 type."".MyInterface+112
	rel 2+0 t=24 type."".MyInterface+104
	rel 2+0 t=24 type."".MyInterface+120
	rel 2+0 t=24 type."".MyInterface+96
	rel 39+0 t=10 +0
	rel 57+0 t=10 +0
	rel 75+0 t=10 +0
	rel 93+0 t=10 +0
	rel 116+4 t=7 runtime.morestack_noctxt+0
"".MyStruct.Print STEXT nosplit size=1 args=0x0 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (demo.go:19)	TEXT	"".MyStruct.Print(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (demo.go:19)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (demo.go:19)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (demo.go:19)	FUNCDATA	$5, "".MyStruct.Print.arginfo1(SB)
	0x0000 00000 (demo.go:19)	RET
	0x0000 c3                                               .
"".MyStruct.Hello STEXT nosplit size=1 args=0x0 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (demo.go:20)	TEXT	"".MyStruct.Hello(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (demo.go:20)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (demo.go:20)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (demo.go:20)	FUNCDATA	$5, "".MyStruct.Hello.arginfo1(SB)
	0x0000 00000 (demo.go:20)	RET
	0x0000 c3                                               .
"".MyStruct.World STEXT nosplit size=1 args=0x0 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (demo.go:21)	TEXT	"".MyStruct.World(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (demo.go:21)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (demo.go:21)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (demo.go:21)	FUNCDATA	$5, "".MyStruct.World.arginfo1(SB)
	0x0000 00000 (demo.go:21)	RET
	0x0000 c3                                               .
"".MyStruct.AWK STEXT nosplit size=1 args=0x0 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (demo.go:22)	TEXT	"".MyStruct.AWK(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (demo.go:22)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (demo.go:22)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (demo.go:22)	FUNCDATA	$5, "".MyStruct.AWK.arginfo1(SB)
	0x0000 00000 (demo.go:22)	RET
	0x0000 c3                                               .
"".main STEXT size=56 args=0x0 locals=0x18 funcid=0x0 align=0x0
	0x0000 00000 (demo.go:24)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (demo.go:24)	CMPQ	SP, 16(R14)
	0x0004 00004 (demo.go:24)	PCDATA	$0, $-2
	0x0004 00004 (demo.go:24)	JLS	49
	0x0006 00006 (demo.go:24)	PCDATA	$0, $-1
	0x0006 00006 (demo.go:24)	SUBQ	$24, SP
	0x000a 00010 (demo.go:24)	MOVQ	BP, 16(SP)
	0x000f 00015 (demo.go:24)	LEAQ	16(SP), BP
	0x0014 00020 (demo.go:24)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (demo.go:24)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (demo.go:26)	LEAQ	go.itab."".MyStruct,"".MyInterface(SB), AX
	0x001b 00027 (demo.go:26)	LEAQ	runtime.zerobase(SB), BX
	0x0022 00034 (demo.go:26)	PCDATA	$1, $0
	0x0022 00034 (demo.go:26)	CALL	"".Foo(SB)
	0x0027 00039 (demo.go:27)	MOVQ	16(SP), BP
	0x002c 00044 (demo.go:27)	ADDQ	$24, SP
	0x0030 00048 (demo.go:27)	RET
	0x0031 00049 (demo.go:27)	NOP
	0x0031 00049 (demo.go:24)	PCDATA	$1, $-1
	0x0031 00049 (demo.go:24)	PCDATA	$0, $-2
	0x0031 00049 (demo.go:24)	CALL	runtime.morestack_noctxt(SB)
	0x0036 00054 (demo.go:24)	PCDATA	$0, $-1
	0x0036 00054 (demo.go:24)	JMP	0
	0x0000 49 3b 66 10 76 2b 48 83 ec 18 48 89 6c 24 10 48  I;f.v+H...H.l$.H
	0x0010 8d 6c 24 10 48 8d 05 00 00 00 00 48 8d 1d 00 00  .l$.H......H....
	0x0020 00 00 e8 00 00 00 00 48 8b 6c 24 10 48 83 c4 18  .......H.l$.H...
	0x0030 c3 e8 00 00 00 00 eb c8                          ........
	rel 2+0 t=23 type."".MyStruct+0
	rel 23+4 t=14 go.itab."".MyStruct,"".MyInterface+0
	rel 30+4 t=14 runtime.zerobase+0
	rel 35+4 t=7 "".Foo+0
	rel 50+4 t=7 runtime.morestack_noctxt+0
"".(*MyStruct).AWK STEXT dupok size=95 args=0x8 locals=0x8 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*MyStruct).AWK(SB), DUPOK|WRAPPER|ABIInternal, $8-8
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	61
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$8, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x000e 00014 (<autogenerated>:1)	LEAQ	(SP), BP
	0x0012 00018 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0016 00022 (<autogenerated>:1)	TESTQ	R12, R12
	0x0019 00025 (<autogenerated>:1)	JNE	78
	0x001b 00027 (<autogenerated>:1)	NOP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$5, "".(*MyStruct).AWK.arginfo1(SB)
	0x001b 00027 (<autogenerated>:1)	MOVQ	AX, ""..this+16(SP)
	0x0020 00032 (<autogenerated>:1)	TESTQ	AX, AX
	0x0023 00035 (<autogenerated>:1)	JNE	39
	0x0025 00037 (<autogenerated>:1)	JMP	55
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	PCDATA	$1, $1
	0x0029 00041 (<autogenerated>:1)	CALL	"".MyStruct.AWK(SB)
	0x002e 00046 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0032 00050 (<autogenerated>:1)	ADDQ	$8, SP
	0x0036 00054 (<autogenerated>:1)	RET
	0x0037 00055 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x003c 00060 (<autogenerated>:1)	XCHGL	AX, AX
	0x003d 00061 (<autogenerated>:1)	NOP
	0x003d 00061 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003d 00061 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003d 00061 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0042 00066 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0047 00071 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x004c 00076 (<autogenerated>:1)	PCDATA	$0, $-1
	0x004c 00076 (<autogenerated>:1)	JMP	0
	0x004e 00078 (<autogenerated>:1)	LEAQ	16(SP), R13
	0x0053 00083 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0057 00087 (<autogenerated>:1)	JNE	27
	0x0059 00089 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x005d 00093 (<autogenerated>:1)	JMP	27
	0x0000 49 3b 66 10 76 37 48 83 ec 08 48 89 2c 24 48 8d  I;f.v7H...H.,$H.
	0x0010 2c 24 4d 8b 66 20 4d 85 e4 75 33 48 89 44 24 10  ,$M.f M..u3H.D$.
	0x0020 48 85 c0 75 02 eb 10 84 00 e8 00 00 00 00 48 8b  H..u..........H.
	0x0030 2c 24 48 83 c4 08 c3 e8 00 00 00 00 90 48 89 44  ,$H..........H.D
	0x0040 24 08 e8 00 00 00 00 48 8b 44 24 08 eb b2 4c 8d  $......H.D$...L.
	0x0050 6c 24 10 4d 39 2c 24 75 c2 49 89 24 24 eb bc     l$.M9,$u.I.$$..
	rel 42+4 t=7 "".MyStruct.AWK+0
	rel 56+4 t=7 runtime.panicwrap+0
	rel 67+4 t=7 runtime.morestack_noctxt+0
"".(*MyStruct).Hello STEXT dupok size=95 args=0x8 locals=0x8 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*MyStruct).Hello(SB), DUPOK|WRAPPER|ABIInternal, $8-8
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	61
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$8, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x000e 00014 (<autogenerated>:1)	LEAQ	(SP), BP
	0x0012 00018 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0016 00022 (<autogenerated>:1)	TESTQ	R12, R12
	0x0019 00025 (<autogenerated>:1)	JNE	78
	0x001b 00027 (<autogenerated>:1)	NOP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$5, "".(*MyStruct).Hello.arginfo1(SB)
	0x001b 00027 (<autogenerated>:1)	MOVQ	AX, ""..this+16(SP)
	0x0020 00032 (<autogenerated>:1)	TESTQ	AX, AX
	0x0023 00035 (<autogenerated>:1)	JNE	39
	0x0025 00037 (<autogenerated>:1)	JMP	55
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	PCDATA	$1, $1
	0x0029 00041 (<autogenerated>:1)	CALL	"".MyStruct.Hello(SB)
	0x002e 00046 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0032 00050 (<autogenerated>:1)	ADDQ	$8, SP
	0x0036 00054 (<autogenerated>:1)	RET
	0x0037 00055 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x003c 00060 (<autogenerated>:1)	XCHGL	AX, AX
	0x003d 00061 (<autogenerated>:1)	NOP
	0x003d 00061 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003d 00061 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003d 00061 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0042 00066 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0047 00071 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x004c 00076 (<autogenerated>:1)	PCDATA	$0, $-1
	0x004c 00076 (<autogenerated>:1)	JMP	0
	0x004e 00078 (<autogenerated>:1)	LEAQ	16(SP), R13
	0x0053 00083 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0057 00087 (<autogenerated>:1)	JNE	27
	0x0059 00089 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x005d 00093 (<autogenerated>:1)	JMP	27
	0x0000 49 3b 66 10 76 37 48 83 ec 08 48 89 2c 24 48 8d  I;f.v7H...H.,$H.
	0x0010 2c 24 4d 8b 66 20 4d 85 e4 75 33 48 89 44 24 10  ,$M.f M..u3H.D$.
	0x0020 48 85 c0 75 02 eb 10 84 00 e8 00 00 00 00 48 8b  H..u..........H.
	0x0030 2c 24 48 83 c4 08 c3 e8 00 00 00 00 90 48 89 44  ,$H..........H.D
	0x0040 24 08 e8 00 00 00 00 48 8b 44 24 08 eb b2 4c 8d  $......H.D$...L.
	0x0050 6c 24 10 4d 39 2c 24 75 c2 49 89 24 24 eb bc     l$.M9,$u.I.$$..
	rel 42+4 t=7 "".MyStruct.Hello+0
	rel 56+4 t=7 runtime.panicwrap+0
	rel 67+4 t=7 runtime.morestack_noctxt+0
"".(*MyStruct).Print STEXT dupok size=95 args=0x8 locals=0x8 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*MyStruct).Print(SB), DUPOK|WRAPPER|ABIInternal, $8-8
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	61
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$8, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x000e 00014 (<autogenerated>:1)	LEAQ	(SP), BP
	0x0012 00018 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0016 00022 (<autogenerated>:1)	TESTQ	R12, R12
	0x0019 00025 (<autogenerated>:1)	JNE	78
	0x001b 00027 (<autogenerated>:1)	NOP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$5, "".(*MyStruct).Print.arginfo1(SB)
	0x001b 00027 (<autogenerated>:1)	MOVQ	AX, ""..this+16(SP)
	0x0020 00032 (<autogenerated>:1)	TESTQ	AX, AX
	0x0023 00035 (<autogenerated>:1)	JNE	39
	0x0025 00037 (<autogenerated>:1)	JMP	55
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	PCDATA	$1, $1
	0x0029 00041 (<autogenerated>:1)	CALL	"".MyStruct.Print(SB)
	0x002e 00046 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0032 00050 (<autogenerated>:1)	ADDQ	$8, SP
	0x0036 00054 (<autogenerated>:1)	RET
	0x0037 00055 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x003c 00060 (<autogenerated>:1)	XCHGL	AX, AX
	0x003d 00061 (<autogenerated>:1)	NOP
	0x003d 00061 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003d 00061 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003d 00061 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0042 00066 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0047 00071 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x004c 00076 (<autogenerated>:1)	PCDATA	$0, $-1
	0x004c 00076 (<autogenerated>:1)	JMP	0
	0x004e 00078 (<autogenerated>:1)	LEAQ	16(SP), R13
	0x0053 00083 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0057 00087 (<autogenerated>:1)	JNE	27
	0x0059 00089 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x005d 00093 (<autogenerated>:1)	JMP	27
	0x0000 49 3b 66 10 76 37 48 83 ec 08 48 89 2c 24 48 8d  I;f.v7H...H.,$H.
	0x0010 2c 24 4d 8b 66 20 4d 85 e4 75 33 48 89 44 24 10  ,$M.f M..u3H.D$.
	0x0020 48 85 c0 75 02 eb 10 84 00 e8 00 00 00 00 48 8b  H..u..........H.
	0x0030 2c 24 48 83 c4 08 c3 e8 00 00 00 00 90 48 89 44  ,$H..........H.D
	0x0040 24 08 e8 00 00 00 00 48 8b 44 24 08 eb b2 4c 8d  $......H.D$...L.
	0x0050 6c 24 10 4d 39 2c 24 75 c2 49 89 24 24 eb bc     l$.M9,$u.I.$$..
	rel 42+4 t=7 "".MyStruct.Print+0
	rel 56+4 t=7 runtime.panicwrap+0
	rel 67+4 t=7 runtime.morestack_noctxt+0
"".(*MyStruct).World STEXT dupok size=95 args=0x8 locals=0x8 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*MyStruct).World(SB), DUPOK|WRAPPER|ABIInternal, $8-8
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	61
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$8, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x000e 00014 (<autogenerated>:1)	LEAQ	(SP), BP
	0x0012 00018 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0016 00022 (<autogenerated>:1)	TESTQ	R12, R12
	0x0019 00025 (<autogenerated>:1)	JNE	78
	0x001b 00027 (<autogenerated>:1)	NOP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$5, "".(*MyStruct).World.arginfo1(SB)
	0x001b 00027 (<autogenerated>:1)	MOVQ	AX, ""..this+16(SP)
	0x0020 00032 (<autogenerated>:1)	TESTQ	AX, AX
	0x0023 00035 (<autogenerated>:1)	JNE	39
	0x0025 00037 (<autogenerated>:1)	JMP	55
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	PCDATA	$1, $1
	0x0029 00041 (<autogenerated>:1)	CALL	"".MyStruct.World(SB)
	0x002e 00046 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0032 00050 (<autogenerated>:1)	ADDQ	$8, SP
	0x0036 00054 (<autogenerated>:1)	RET
	0x0037 00055 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x003c 00060 (<autogenerated>:1)	XCHGL	AX, AX
	0x003d 00061 (<autogenerated>:1)	NOP
	0x003d 00061 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003d 00061 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003d 00061 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0042 00066 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0047 00071 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x004c 00076 (<autogenerated>:1)	PCDATA	$0, $-1
	0x004c 00076 (<autogenerated>:1)	JMP	0
	0x004e 00078 (<autogenerated>:1)	LEAQ	16(SP), R13
	0x0053 00083 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0057 00087 (<autogenerated>:1)	JNE	27
	0x0059 00089 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x005d 00093 (<autogenerated>:1)	JMP	27
	0x0000 49 3b 66 10 76 37 48 83 ec 08 48 89 2c 24 48 8d  I;f.v7H...H.,$H.
	0x0010 2c 24 4d 8b 66 20 4d 85 e4 75 33 48 89 44 24 10  ,$M.f M..u3H.D$.
	0x0020 48 85 c0 75 02 eb 10 84 00 e8 00 00 00 00 48 8b  H..u..........H.
	0x0030 2c 24 48 83 c4 08 c3 e8 00 00 00 00 90 48 89 44  ,$H..........H.D
	0x0040 24 08 e8 00 00 00 00 48 8b 44 24 08 eb b2 4c 8d  $......H.D$...L.
	0x0050 6c 24 10 4d 39 2c 24 75 c2 49 89 24 24 eb bc     l$.M9,$u.I.$$..
	rel 42+4 t=7 "".MyStruct.World+0
	rel 56+4 t=7 runtime.panicwrap+0
	rel 67+4 t=7 runtime.morestack_noctxt+0
"".MyInterface.AWK STEXT dupok size=108 args=0x10 locals=0x10 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".MyInterface.AWK(SB), DUPOK|WRAPPER|ABIInternal, $16-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	60
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	87
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, "".MyInterface.AWK.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(SP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	BX, ""..this+32(SP)
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	MOVQ	24(AX), CX
	0x002d 00045 (<autogenerated>:1)	MOVQ	BX, AX
	0x0030 00048 (<autogenerated>:1)	PCDATA	$1, $1
	0x0030 00048 (<autogenerated>:1)	CALL	CX
	0x0032 00050 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x0037 00055 (<autogenerated>:1)	ADDQ	$16, SP
	0x003b 00059 (<autogenerated>:1)	RET
	0x003c 00060 (<autogenerated>:1)	NOP
	0x003c 00060 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003c 00060 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0041 00065 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x0046 00070 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0050 00080 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	JMP	0
	0x0057 00087 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x005c 00092 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0064 00100 (<autogenerated>:1)	JNE	29
	0x0066 00102 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x006a 00106 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 36 48 83 ec 10 48 89 6c 24 08 48  I;f.v6H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 3a 48 89 44  .l$.M.f M..u:H.D
	0x0020 24 18 48 89 5c 24 20 84 00 48 8b 48 18 48 89 d8  $.H.\$ ..H.H.H..
	0x0030 ff d1 48 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24  ..H.l$.H....H.D$
	0x0040 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08  .H.\$......H.D$.
	0x0050 48 8b 5c 24 10 eb a9 4c 8d 6c 24 18 0f 1f 40 00  H.\$...L.l$...@.
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 2+0 t=24 type."".MyInterface+96
	rel 48+0 t=10 +0
	rel 71+4 t=7 runtime.morestack_noctxt+0
"".MyInterface.Hello STEXT dupok size=108 args=0x10 locals=0x10 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".MyInterface.Hello(SB), DUPOK|WRAPPER|ABIInternal, $16-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	60
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	87
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, "".MyInterface.Hello.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(SP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	BX, ""..this+32(SP)
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	MOVQ	32(AX), CX
	0x002d 00045 (<autogenerated>:1)	MOVQ	BX, AX
	0x0030 00048 (<autogenerated>:1)	PCDATA	$1, $1
	0x0030 00048 (<autogenerated>:1)	CALL	CX
	0x0032 00050 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x0037 00055 (<autogenerated>:1)	ADDQ	$16, SP
	0x003b 00059 (<autogenerated>:1)	RET
	0x003c 00060 (<autogenerated>:1)	NOP
	0x003c 00060 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003c 00060 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0041 00065 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x0046 00070 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0050 00080 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	JMP	0
	0x0057 00087 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x005c 00092 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0064 00100 (<autogenerated>:1)	JNE	29
	0x0066 00102 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x006a 00106 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 36 48 83 ec 10 48 89 6c 24 08 48  I;f.v6H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 3a 48 89 44  .l$.M.f M..u:H.D
	0x0020 24 18 48 89 5c 24 20 84 00 48 8b 48 20 48 89 d8  $.H.\$ ..H.H H..
	0x0030 ff d1 48 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24  ..H.l$.H....H.D$
	0x0040 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08  .H.\$......H.D$.
	0x0050 48 8b 5c 24 10 eb a9 4c 8d 6c 24 18 0f 1f 40 00  H.\$...L.l$...@.
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 2+0 t=24 type."".MyInterface+104
	rel 48+0 t=10 +0
	rel 71+4 t=7 runtime.morestack_noctxt+0
"".MyInterface.Print STEXT dupok size=108 args=0x10 locals=0x10 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".MyInterface.Print(SB), DUPOK|WRAPPER|ABIInternal, $16-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	60
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	87
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, "".MyInterface.Print.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(SP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	BX, ""..this+32(SP)
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	MOVQ	40(AX), CX
	0x002d 00045 (<autogenerated>:1)	MOVQ	BX, AX
	0x0030 00048 (<autogenerated>:1)	PCDATA	$1, $1
	0x0030 00048 (<autogenerated>:1)	CALL	CX
	0x0032 00050 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x0037 00055 (<autogenerated>:1)	ADDQ	$16, SP
	0x003b 00059 (<autogenerated>:1)	RET
	0x003c 00060 (<autogenerated>:1)	NOP
	0x003c 00060 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003c 00060 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0041 00065 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x0046 00070 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0050 00080 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	JMP	0
	0x0057 00087 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x005c 00092 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0064 00100 (<autogenerated>:1)	JNE	29
	0x0066 00102 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x006a 00106 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 36 48 83 ec 10 48 89 6c 24 08 48  I;f.v6H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 3a 48 89 44  .l$.M.f M..u:H.D
	0x0020 24 18 48 89 5c 24 20 84 00 48 8b 48 28 48 89 d8  $.H.\$ ..H.H(H..
	0x0030 ff d1 48 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24  ..H.l$.H....H.D$
	0x0040 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08  .H.\$......H.D$.
	0x0050 48 8b 5c 24 10 eb a9 4c 8d 6c 24 18 0f 1f 40 00  H.\$...L.l$...@.
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 2+0 t=24 type."".MyInterface+112
	rel 48+0 t=10 +0
	rel 71+4 t=7 runtime.morestack_noctxt+0
"".MyInterface.World STEXT dupok size=108 args=0x10 locals=0x10 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".MyInterface.World(SB), DUPOK|WRAPPER|ABIInternal, $16-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	60
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	87
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, "".MyInterface.World.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, ""..this+24(SP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	BX, ""..this+32(SP)
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	MOVQ	48(AX), CX
	0x002d 00045 (<autogenerated>:1)	MOVQ	BX, AX
	0x0030 00048 (<autogenerated>:1)	PCDATA	$1, $1
	0x0030 00048 (<autogenerated>:1)	CALL	CX
	0x0032 00050 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x0037 00055 (<autogenerated>:1)	ADDQ	$16, SP
	0x003b 00059 (<autogenerated>:1)	RET
	0x003c 00060 (<autogenerated>:1)	NOP
	0x003c 00060 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003c 00060 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0041 00065 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x0046 00070 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0050 00080 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	JMP	0
	0x0057 00087 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x005c 00092 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0064 00100 (<autogenerated>:1)	JNE	29
	0x0066 00102 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x006a 00106 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 36 48 83 ec 10 48 89 6c 24 08 48  I;f.v6H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 3a 48 89 44  .l$.M.f M..u:H.D
	0x0020 24 18 48 89 5c 24 20 84 00 48 8b 48 30 48 89 d8  $.H.\$ ..H.H0H..
	0x0030 ff d1 48 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24  ..H.l$.H....H.D$
	0x0040 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08  .H.\$......H.D$.
	0x0050 48 8b 5c 24 10 eb a9 4c 8d 6c 24 18 0f 1f 40 00  H.\$...L.l$...@.
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 2+0 t=24 type."".MyInterface+120
	rel 48+0 t=10 +0
	rel 71+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*func()- SRODATA dupok size=9
	0x0000 00 07 2a 66 75 6e 63 28 29                       ..*func()
type.*func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9b 90 75 1b 08 08 08 36 00 00 00 00 00 00 00 00  ..u....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 48+8 t=1 type.func()+0
type.func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f6 bc 82 f6 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00                                      ....
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 44+4 t=-32763 type.*func()+0
runtime.interequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.interequal+0
type..namedata.*main.MyInterface. SRODATA dupok size=19
	0x0000 01 11 2a 6d 61 69 6e 2e 4d 79 49 6e 74 65 72 66  ..*main.MyInterf
	0x0010 61 63 65                                         ace
type.*"".MyInterface SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9c ec de 3f 08 08 08 36 00 00 00 00 00 00 00 00  ...?...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.MyInterface.+0
	rel 48+8 t=1 type."".MyInterface+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type..namedata.AWK. SRODATA dupok size=5
	0x0000 01 03 41 57 4b                                   ..AWK
type..namedata.Hello. SRODATA dupok size=7
	0x0000 01 05 48 65 6c 6c 6f                             ..Hello
type..namedata.Print. SRODATA dupok size=7
	0x0000 01 05 50 72 69 6e 74                             ..Print
type..namedata.World. SRODATA dupok size=7
	0x0000 01 05 57 6f 72 6c 64                             ..World
type."".MyInterface SRODATA dupok size=128
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 b0 8e 5c db 07 08 08 14 00 00 00 00 00 00 00 00  ..\.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 04 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  ........0.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.interequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*main.MyInterface.+0
	rel 44+4 t=5 type.*"".MyInterface+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".MyInterface+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+4 t=5 type..namedata.AWK.+0
	rel 100+4 t=5 type.func()+0
	rel 104+4 t=5 type..namedata.Hello.+0
	rel 108+4 t=5 type.func()+0
	rel 112+4 t=5 type..namedata.Print.+0
	rel 116+4 t=5 type.func()+0
	rel 120+4 t=5 type..namedata.World.+0
	rel 124+4 t=5 type.func()+0
runtime.memequal0·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal0+0
type..namedata.*main.MyStruct. SRODATA dupok size=16
	0x0000 01 0e 2a 6d 61 69 6e 2e 4d 79 53 74 72 75 63 74  ..*main.MyStruct
type..namedata.*func(*main.MyStruct)- SRODATA dupok size=23
	0x0000 00 15 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 4d 79  ..*func(*main.My
	0x0010 53 74 72 75 63 74 29                             Struct)
type.*func(*"".MyStruct) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 56 ef 87 87 08 08 08 36 00 00 00 00 00 00 00 00  V......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.MyStruct)-+0
	rel 48+8 t=1 type.func(*"".MyStruct)+0
type.func(*"".MyStruct) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 db ae 78 91 02 08 08 33 00 00 00 00 00 00 00 00  ..x....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.MyStruct)-+0
	rel 44+4 t=-32763 type.*func(*"".MyStruct)+0
	rel 56+8 t=1 type.*"".MyStruct+0
type.*"".MyStruct SRODATA dupok size=136
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 58 6c 00 8b 09 08 08 36 00 00 00 00 00 00 00 00  Xl.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 04 00 04 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.MyStruct.+0
	rel 48+8 t=1 type."".MyStruct+0
	rel 56+4 t=5 type..importpath."".+0
	rel 72+4 t=5 type..namedata.AWK.+0
	rel 76+4 t=26 type.func()+0
	rel 80+4 t=26 "".(*MyStruct).AWK+0
	rel 84+4 t=26 "".(*MyStruct).AWK+0
	rel 88+4 t=5 type..namedata.Hello.+0
	rel 92+4 t=26 type.func()+0
	rel 96+4 t=26 "".(*MyStruct).Hello+0
	rel 100+4 t=26 "".(*MyStruct).Hello+0
	rel 104+4 t=5 type..namedata.Print.+0
	rel 108+4 t=26 type.func()+0
	rel 112+4 t=26 "".(*MyStruct).Print+0
	rel 116+4 t=26 "".(*MyStruct).Print+0
	rel 120+4 t=5 type..namedata.World.+0
	rel 124+4 t=26 type.func()+0
	rel 128+4 t=26 "".(*MyStruct).World+0
	rel 132+4 t=26 "".(*MyStruct).World+0
runtime.gcbits. SRODATA dupok size=0
type..namedata.*func(main.MyStruct)- SRODATA dupok size=22
	0x0000 00 14 2a 66 75 6e 63 28 6d 61 69 6e 2e 4d 79 53  ..*func(main.MyS
	0x0010 74 72 75 63 74 29                                truct)
type.*func("".MyStruct) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ed af 89 2c 08 08 08 36 00 00 00 00 00 00 00 00  ...,...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.MyStruct)-+0
	rel 48+8 t=1 type.func("".MyStruct)+0
type.func("".MyStruct) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 32 f8 b5 89 02 08 08 33 00 00 00 00 00 00 00 00  2......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.MyStruct)-+0
	rel 44+4 t=-32763 type.*func("".MyStruct)+0
	rel 56+8 t=1 type."".MyStruct+0
type."".MyStruct SRODATA dupok size=160
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 dc a7 2c a3 0f 01 01 19 00 00 00 00 00 00 00 00  ..,.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 04 00 04 00 10 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.memequal0·f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.MyStruct.+0
	rel 44+4 t=5 type.*"".MyStruct+0
	rel 56+8 t=1 type."".MyStruct+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+4 t=5 type..namedata.AWK.+0
	rel 100+4 t=26 type.func()+0
	rel 104+4 t=26 "".(*MyStruct).AWK+0
	rel 108+4 t=26 "".MyStruct.AWK+0
	rel 112+4 t=5 type..namedata.Hello.+0
	rel 116+4 t=26 type.func()+0
	rel 120+4 t=26 "".(*MyStruct).Hello+0
	rel 124+4 t=26 "".MyStruct.Hello+0
	rel 128+4 t=5 type..namedata.Print.+0
	rel 132+4 t=26 type.func()+0
	rel 136+4 t=26 "".(*MyStruct).Print+0
	rel 140+4 t=26 "".MyStruct.Print+0
	rel 144+4 t=5 type..namedata.World.+0
	rel 148+4 t=26 type.func()+0
	rel 152+4 t=26 "".(*MyStruct).World+0
	rel 156+4 t=26 "".MyStruct.World+0
go.itab."".MyStruct,"".MyInterface SRODATA dupok size=56
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 dc a7 2c a3 00 00 00 00 00 00 00 00 00 00 00 00  ..,.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type."".MyInterface+0
	rel 8+8 t=1 type."".MyStruct+0
	rel 24+8 t=-32767 "".(*MyStruct).AWK+0
	rel 32+8 t=-32767 "".(*MyStruct).Hello+0
	rel 40+8 t=-32767 "".(*MyStruct).Print+0
	rel 48+8 t=-32767 "".(*MyStruct).World+0
gclocals·09cf9819fc716118c209c2d2155a3632 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 02 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
"".Foo.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".MyStruct.Print.arginfo1 SRODATA static dupok size=3
	0x0000 fe fd ff                                         ...
"".MyStruct.Hello.arginfo1 SRODATA static dupok size=3
	0x0000 fe fd ff                                         ...
"".MyStruct.World.arginfo1 SRODATA static dupok size=3
	0x0000 fe fd ff                                         ...
"".MyStruct.AWK.arginfo1 SRODATA static dupok size=3
	0x0000 fe fd ff                                         ...
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
"".(*MyStruct).AWK.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".(*MyStruct).Hello.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".(*MyStruct).Print.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".(*MyStruct).World.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".MyInterface.AWK.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
"".MyInterface.Hello.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
"".MyInterface.Print.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
"".MyInterface.World.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
