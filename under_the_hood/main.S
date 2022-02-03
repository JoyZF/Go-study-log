TEXT main.main(SB) /Users/joy/workspace/gopath/src/study/under_the_hood/main.go
  main.go:5		0x10a31c0		65488b0c2530000000	MOVQ GS:0x30, CX			
  main.go:5		0x10a31c9		488d4424f8		LEAQ -0x8(SP), AX			
  main.go:5		0x10a31ce		483b4110		CMPQ 0x10(CX), AX			
  main.go:5		0x10a31d2		0f86d9000000		JBE 0x10a32b1				
  main.go:5		0x10a31d8		4881ec88000000		SUBQ $0x88, SP				
  main.go:5		0x10a31df		4889ac2480000000	MOVQ BP, 0x80(SP)			
  main.go:5		0x10a31e7		488dac2480000000	LEAQ 0x80(SP), BP			
  main.go:6		0x10a31ef		48c744244001000000	MOVQ $0x1, 0x40(SP)			
  main.go:7		0x10a31f8		48c744243802000000	MOVQ $0x2, 0x38(SP)			
  main.go:8		0x10a3201		488b442440		MOVQ 0x40(SP), AX			
  main.go:8		0x10a3206		4883c002		ADDQ $0x2, AX				
  main.go:8		0x10a320a		4889442430		MOVQ AX, 0x30(SP)			
  main.go:9		0x10a320f		0f57c0			XORPS X0, X0				
  main.go:9		0x10a3212		0f11442458		MOVUPS X0, 0x58(SP)			
  main.go:9		0x10a3217		488d442458		LEAQ 0x58(SP), AX			
  main.go:9		0x10a321c		4889442450		MOVQ AX, 0x50(SP)			
  main.go:9		0x10a3221		488b442430		MOVQ 0x30(SP), AX			
  main.go:9		0x10a3226		48890424		MOVQ AX, 0(SP)				
  main.go:9		0x10a322a		e8b16ef6ff		CALL runtime.convT64(SB)		
  main.go:9		0x10a322f		488b442408		MOVQ 0x8(SP), AX			
  main.go:9		0x10a3234		4889442448		MOVQ AX, 0x48(SP)			
  main.go:9		0x10a3239		488b4c2450		MOVQ 0x50(SP), CX			
  main.go:9		0x10a323e		8401			TESTB AL, 0(CX)				
  main.go:9		0x10a3240		488d15d9a80000		LEAQ runtime.types+42816(SB), DX	
  main.go:9		0x10a3247		488911			MOVQ DX, 0(CX)				
  main.go:9		0x10a324a		488d7908		LEAQ 0x8(CX), DI			
  main.go:9		0x10a324e		833d9b910e0000		CMPL $0x0, runtime.writeBarrier(SB)	
  main.go:9		0x10a3255		7402			JE 0x10a3259				
  main.go:9		0x10a3257		eb51			JMP 0x10a32aa				
  main.go:9		0x10a3259		48894108		MOVQ AX, 0x8(CX)			
  main.go:9		0x10a325d		eb00			JMP 0x10a325f				
  main.go:9		0x10a325f		488b442450		MOVQ 0x50(SP), AX			
  main.go:9		0x10a3264		8400			TESTB AL, 0(AX)				
  main.go:9		0x10a3266		eb00			JMP 0x10a3268				
  main.go:9		0x10a3268		4889442468		MOVQ AX, 0x68(SP)			
  main.go:9		0x10a326d		48c744247001000000	MOVQ $0x1, 0x70(SP)			
  main.go:9		0x10a3276		48c744247801000000	MOVQ $0x1, 0x78(SP)			
  main.go:9		0x10a327f		48890424		MOVQ AX, 0(SP)				
  main.go:9		0x10a3283		48c744240801000000	MOVQ $0x1, 0x8(SP)			
  main.go:9		0x10a328c		48c744241001000000	MOVQ $0x1, 0x10(SP)			
  main.go:9		0x10a3295		e8869affff		CALL fmt.Println(SB)			
  main.go:10		0x10a329a		488bac2480000000	MOVQ 0x80(SP), BP			
  main.go:10		0x10a32a2		4881c488000000		ADDQ $0x88, SP				
  main.go:10		0x10a32a9		c3			RET					
  main.go:9		0x10a32aa		e8510efcff		CALL runtime.gcWriteBarrier(SB)		
  main.go:9		0x10a32af		ebae			JMP 0x10a325f				
  main.go:5		0x10a32b1		e8caf2fbff		CALL runtime.morestack_noctxt(SB)	
  main.go:5		0x10a32b6		e905ffffff		JMP main.main(SB)			
