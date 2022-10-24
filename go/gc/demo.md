"".main STEXT size=40 args=0x0 locals=0x8 funcid=0x0 align=0x0
	0x0000 00000 (gc.go:5)	TEXT	"".main(SB), ABIInternal, $8-0
	0x0000 00000 (gc.go:5)	CMPQ	SP, 16(R14)
	0x0004 00004 (gc.go:5)	PCDATA	$0, $-2
	0x0004 00004 (gc.go:5)	JLS	33
	0x0006 00006 (gc.go:5)	PCDATA	$0, $-1
	0x0006 00006 (gc.go:5)	SUBQ	$8, SP
	0x000a 00010 (gc.go:5)	MOVQ	BP, (SP)
	0x000e 00014 (gc.go:5)	LEAQ	(SP), BP
	0x0012 00018 (gc.go:5)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (gc.go:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (gc.go:6)	PCDATA	$1, $0
	0x0012 00018 (gc.go:6)	CALL	runtime.GC(SB)
	0x0017 00023 (gc.go:7)	MOVQ	(SP), BP
	0x001b 00027 (gc.go:7)	ADDQ	$8, SP
	0x001f 00031 (gc.go:7)	NOP
	0x0020 00032 (gc.go:7)	RET
	0x0021 00033 (gc.go:7)	NOP
	0x0021 00033 (gc.go:5)	PCDATA	$1, $-1
	0x0021 00033 (gc.go:5)	PCDATA	$0, $-2
	0x0021 00033 (gc.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x0026 00038 (gc.go:5)	PCDATA	$0, $-1
	0x0026 00038 (gc.go:5)	JMP	0
	0x0000 49 3b 66 10 76 1b 48 83 ec 08 48 89 2c 24 48 8d  I;f.v.H...H.,$H.
	0x0010 2c 24 e8 00 00 00 00 48 8b 2c 24 48 83 c4 08 90  ,$.....H.,$H....
	0x0020 c3 e8 00 00 00 00 eb d8                          ........
	rel 19+4 t=7 runtime.GC+0
	rel 34+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime..inittask+0
type..importpath.runtime. SRODATA dupok size=9
	0x0000 00 07 72 75 6e 74 69 6d 65                       ..runtime
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
