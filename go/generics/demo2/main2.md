TEXT main.main(SB) /Users/gongqi/code/go-study/go/generics/demo2/main.go
  main.go:8		0x10899a0		493b6610		CMPQ 0x10(R14), SP			
  main.go:8		0x10899a4		0f8605020000		JBE 0x1089baf				
  main.go:8		0x10899aa		4883ec58		SUBQ $0x58, SP				
  main.go:8		0x10899ae		48896c2450		MOVQ BP, 0x50(SP)			
  main.go:8		0x10899b3		488d6c2450		LEAQ 0x50(SP), BP			
  main.go:9		0x10899b8		440f117c2428		MOVUPS X15, 0x28(SP)			
  main.go:19		0x10899be		31c0			XORL AX, AX				
  main.go:19		0x10899c0		e87bfaf7ff		CALL runtime.convT64(SB)		
  main.go:19		0x10899c5		488d0dc4660300		LEAQ main..dict.echo[int](SB), CX	
  main.go:19		0x10899cc		488b09			MOVQ 0(CX), CX				
  main.go:19		0x10899cf		48894c2428		MOVQ CX, 0x28(SP)			
  main.go:19		0x10899d4		4889442430		MOVQ AX, 0x30(SP)			
  main.go:19		0x10899d9		488d052d7d0100		LEAQ go.string.*+37(SB), AX		
  main.go:19		0x10899e0		bb02000000		MOVL $0x2, BX				
  main.go:19		0x10899e5		488d4c2428		LEAQ 0x28(SP), CX			
  main.go:19		0x10899ea		bf01000000		MOVL $0x1, DI				
  main.go:19		0x10899ef		4889fe			MOVQ DI, SI				
  main.go:19		0x10899f2		e82998ffff		CALL fmt.Sprintf(SB)			
  main.go:10		0x10899f7		90			NOPL					
  main.go:19		0x10899f8		440f117c2428		MOVUPS X15, 0x28(SP)			
  main.go:19		0x10899fe		31c0			XORL AX, AX				
  main.go:19		0x1089a00		e8bbf9f7ff		CALL runtime.convT32(SB)		
  main.go:19		0x1089a05		488d0d7c660300		LEAQ main..dict.echo[int32](SB), CX	
  main.go:19		0x1089a0c		488b09			MOVQ 0(CX), CX				
  main.go:19		0x1089a0f		48894c2428		MOVQ CX, 0x28(SP)			
  main.go:19		0x1089a14		4889442430		MOVQ AX, 0x30(SP)			
  main.go:19		0x1089a19		488d05ed7c0100		LEAQ go.string.*+37(SB), AX		
  main.go:19		0x1089a20		bb02000000		MOVL $0x2, BX				
  main.go:19		0x1089a25		488d4c2428		LEAQ 0x28(SP), CX			
  main.go:19		0x1089a2a		bf01000000		MOVL $0x1, DI				
  main.go:19		0x1089a2f		4889fe			MOVQ DI, SI				
  main.go:19		0x1089a32		e8e997ffff		CALL fmt.Sprintf(SB)			
  main.go:11		0x1089a37		90			NOPL					
  main.go:19		0x1089a38		440f117c2428		MOVUPS X15, 0x28(SP)			
  main.go:19		0x1089a3e		31c0			XORL AX, AX				
  main.go:19		0x1089a40		e87bf9f7ff		CALL runtime.convT32(SB)		
  main.go:19		0x1089a45		488d0d64660300		LEAQ main..dict.echo[uint32](SB), CX	
  main.go:19		0x1089a4c		488b09			MOVQ 0(CX), CX				
  main.go:19		0x1089a4f		48894c2428		MOVQ CX, 0x28(SP)			
  main.go:19		0x1089a54		4889442430		MOVQ AX, 0x30(SP)			
  main.go:19		0x1089a59		488d05ad7c0100		LEAQ go.string.*+37(SB), AX		
  main.go:19		0x1089a60		bb02000000		MOVL $0x2, BX				
  main.go:19		0x1089a65		488d4c2428		LEAQ 0x28(SP), CX			
  main.go:19		0x1089a6a		bf01000000		MOVL $0x1, DI				
  main.go:19		0x1089a6f		4889fe			MOVQ DI, SI				
  main.go:19		0x1089a72		e8a997ffff		CALL fmt.Sprintf(SB)			
  main.go:12		0x1089a77		90			NOPL					
  main.go:19		0x1089a78		440f117c2428		MOVUPS X15, 0x28(SP)			
  main.go:19		0x1089a7e		31c0			XORL AX, AX				
  main.go:19		0x1089a80		e8bbf9f7ff		CALL runtime.convT64(SB)		
  main.go:19		0x1089a85		488d0d2c660300		LEAQ main..dict.echo[uint64](SB), CX	
  main.go:19		0x1089a8c		488b09			MOVQ 0(CX), CX				
  main.go:19		0x1089a8f		48894c2428		MOVQ CX, 0x28(SP)			
  main.go:19		0x1089a94		4889442430		MOVQ AX, 0x30(SP)			
  main.go:19		0x1089a99		488d056d7c0100		LEAQ go.string.*+37(SB), AX		
  main.go:19		0x1089aa0		bb02000000		MOVL $0x2, BX				
  main.go:19		0x1089aa5		488d4c2428		LEAQ 0x28(SP), CX			
  main.go:19		0x1089aaa		bf01000000		MOVL $0x1, DI				
  main.go:19		0x1089aaf		4889fe			MOVQ DI, SI				
  main.go:19		0x1089ab2		e86997ffff		CALL fmt.Sprintf(SB)			
  main.go:13		0x1089ab7		90			NOPL					
  main.go:19		0x1089ab8		440f117c2428		MOVUPS X15, 0x28(SP)			
  main.go:19		0x1089abe		488d05107f0100		LEAQ go.string.*+749(SB), AX		
  main.go:19		0x1089ac5		bb05000000		MOVL $0x5, BX				
  main.go:19		0x1089aca		e8f1f9f7ff		CALL runtime.convTstring(SB)		
  main.go:19		0x1089acf		488d0dc2650300		LEAQ main..dict.echo[string](SB), CX	
  main.go:19		0x1089ad6		488b09			MOVQ 0(CX), CX				
  main.go:19		0x1089ad9		48894c2428		MOVQ CX, 0x28(SP)			
  main.go:19		0x1089ade		4889442430		MOVQ AX, 0x30(SP)			
  main.go:19		0x1089ae3		488d05237c0100		LEAQ go.string.*+37(SB), AX		
  main.go:19		0x1089aea		bb02000000		MOVL $0x2, BX				
  main.go:19		0x1089aef		488d4c2428		LEAQ 0x28(SP), CX			
  main.go:19		0x1089af4		bf01000000		MOVL $0x1, DI				
  main.go:19		0x1089af9		4889fe			MOVQ DI, SI				
  main.go:19		0x1089afc		0f1f4000		NOPL 0(AX)				
  main.go:19		0x1089b00		e81b97ffff		CALL fmt.Sprintf(SB)			
  main.go:14		0x1089b05		90			NOPL					
  main.go:19		0x1089b06		440f117c2428		MOVUPS X15, 0x28(SP)			
  main.go:19		0x1089b0c		488d0d8d650300		LEAQ main..dict.echo[struct {}](SB), CX	
  main.go:19		0x1089b13		488b09			MOVQ 0(CX), CX				
  main.go:19		0x1089b16		48894c2428		MOVQ CX, 0x28(SP)			
  main.go:19		0x1089b1b		488d0dbe940d00		LEAQ runtime.zerobase(SB), CX		
  main.go:19		0x1089b22		48894c2430		MOVQ CX, 0x30(SP)			
  main.go:19		0x1089b27		488d05df7b0100		LEAQ go.string.*+37(SB), AX		
  main.go:19		0x1089b2e		bb02000000		MOVL $0x2, BX				
  main.go:19		0x1089b33		488d4c2428		LEAQ 0x28(SP), CX			
  main.go:19		0x1089b38		bf01000000		MOVL $0x1, DI				
  main.go:19		0x1089b3d		4889fe			MOVQ DI, SI				
  main.go:19		0x1089b40		e8db96ffff		CALL fmt.Sprintf(SB)			
  main.go:15		0x1089b45		e8f6f8feff		CALL time.Now(SB)			
  main.go:19		0x1089b4a		4889442438		MOVQ AX, 0x38(SP)			
  main.go:19		0x1089b4f		48895c2440		MOVQ BX, 0x40(SP)			
  main.go:19		0x1089b54		48894c2448		MOVQ CX, 0x48(SP)			
  main.go:19		0x1089b59		440f117c2428		MOVUPS X15, 0x28(SP)			
  main.go:19		0x1089b5f		488d05bae70000		LEAQ runtime.types+58944(SB), AX	
  main.go:19		0x1089b66		488d5c2438		LEAQ 0x38(SP), BX			
  main.go:19		0x1089b6b		e850f7f7ff		CALL runtime.convT(SB)			
  main.go:19		0x1089b70		488d0d31650300		LEAQ main..dict.echo[time.Time](SB), CX	
  main.go:19		0x1089b77		488b09			MOVQ 0(CX), CX				
  main.go:19		0x1089b7a		48894c2428		MOVQ CX, 0x28(SP)			
  main.go:19		0x1089b7f		4889442430		MOVQ AX, 0x30(SP)			
  main.go:19		0x1089b84		488d05827b0100		LEAQ go.string.*+37(SB), AX		
  main.go:19		0x1089b8b		bb02000000		MOVL $0x2, BX				
  main.go:19		0x1089b90		488d4c2428		LEAQ 0x28(SP), CX			
  main.go:19		0x1089b95		bf01000000		MOVL $0x1, DI				
  main.go:19		0x1089b9a		4889fe			MOVQ DI, SI				
  main.go:19		0x1089b9d		0f1f00			NOPL 0(AX)				
  main.go:19		0x1089ba0		e87b96ffff		CALL fmt.Sprintf(SB)			
  main.go:16		0x1089ba5		488b6c2450		MOVQ 0x50(SP), BP			
  main.go:16		0x1089baa		4883c458		ADDQ $0x58, SP				
  main.go:16		0x1089bae		c3			RET					
  main.go:8		0x1089baf		e82ce6fcff		CALL runtime.morestack_noctxt.abi0(SB)	
  main.go:8		0x1089bb4		e9e7fdffff		JMP main.main(SB)			
