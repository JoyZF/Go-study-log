TEXT main.main(SB) /Users/joy/workspace/gopath/src/study/plan9/main.go
  main.go:10		0x10a31c0		65488b0c2530000000		MOVQ GS:0x30, CX			
  main.go:10		0x10a31c9		488d4424b8			LEAQ -0x48(SP), AX			
  main.go:10		0x10a31ce		483b4110			CMPQ 0x10(CX), AX			
  main.go:10		0x10a31d2		0f86f9010000			JBE 0x10a33d1				
  main.go:10		0x10a31d8		4881ecc8000000			SUBQ $0xc8, SP				
  main.go:10		0x10a31df		4889ac24c0000000		MOVQ BP, 0xc0(SP)			
  main.go:10		0x10a31e7		488dac24c0000000		LEAQ 0xc0(SP), BP			
  main.go:11		0x10a31ef		48c744243801000000		MOVQ $0x1, 0x38(SP)			
  main.go:12		0x10a31f8		48c744243802000000		MOVQ $0x2, 0x38(SP)			
  main.go:13		0x10a3201		48c744243002000000		MOVQ $0x2, 0x30(SP)			
  main.go:14		0x10a320a		488b442438			MOVQ 0x38(SP), AX			
  main.go:14		0x10a320f		4889442430			MOVQ AX, 0x30(SP)			
  main.go:15		0x10a3214		0f57c0				XORPS X0, X0				
  main.go:15		0x10a3217		0f11442468			MOVUPS X0, 0x68(SP)			
  main.go:15		0x10a321c		488d442468			LEAQ 0x68(SP), AX			
  main.go:15		0x10a3221		4889442458			MOVQ AX, 0x58(SP)			
  main.go:15		0x10a3226		488b442430			MOVQ 0x30(SP), AX			
  main.go:15		0x10a322b		48890424			MOVQ AX, 0(SP)				
  main.go:15		0x10a322f		e8ac6ef6ff			CALL runtime.convT64(SB)		
  main.go:15		0x10a3234		488b442408			MOVQ 0x8(SP), AX			
  main.go:15		0x10a3239		4889442450			MOVQ AX, 0x50(SP)			
  main.go:15		0x10a323e		488b4c2458			MOVQ 0x58(SP), CX			
  main.go:15		0x10a3243		8401				TESTB AL, 0(CX)				
  main.go:15		0x10a3245		488d15f4a90000			LEAQ runtime.types+42816(SB), DX	
  main.go:15		0x10a324c		488911				MOVQ DX, 0(CX)				
  main.go:15		0x10a324f		488d7908			LEAQ 0x8(CX), DI			
  main.go:15		0x10a3253		833db6910e0000			CMPL $0x0, runtime.writeBarrier(SB)	
  main.go:15		0x10a325a		7409				JE 0x10a3265				
  main.go:15		0x10a325c		0f1f4000			NOPL 0(AX)				
  main.go:15		0x10a3260		e962010000			JMP 0x10a33c7				
  main.go:15		0x10a3265		48894108			MOVQ AX, 0x8(CX)			
  main.go:15		0x10a3269		eb00				JMP 0x10a326b				
  main.go:15		0x10a326b		488b442458			MOVQ 0x58(SP), AX			
  main.go:15		0x10a3270		8400				TESTB AL, 0(AX)				
  main.go:15		0x10a3272		eb00				JMP 0x10a3274				
  main.go:15		0x10a3274		48898424a8000000		MOVQ AX, 0xa8(SP)			
  main.go:15		0x10a327c		48c78424b000000001000000	MOVQ $0x1, 0xb0(SP)			
  main.go:15		0x10a3288		48c78424b800000001000000	MOVQ $0x1, 0xb8(SP)			
  main.go:15		0x10a3294		48890424			MOVQ AX, 0(SP)				
  main.go:15		0x10a3298		48c744240801000000		MOVQ $0x1, 0x8(SP)			
  main.go:15		0x10a32a1		48c744241001000000		MOVQ $0x1, 0x10(SP)			
  main.go:15		0x10a32aa		e8719affff			CALL fmt.Println(SB)			
  main.go:16		0x10a32af		0f57c0				XORPS X0, X0				
  main.go:16		0x10a32b2		0f11442468			MOVUPS X0, 0x68(SP)			
  main.go:16		0x10a32b7		488d442468			LEAQ 0x68(SP), AX			
  main.go:16		0x10a32bc		4889442448			MOVQ AX, 0x48(SP)			
  main.go:16		0x10a32c1		8400				TESTB AL, 0(AX)				
  main.go:16		0x10a32c3		488d0d76a90000			LEAQ runtime.types+42816(SB), CX	
  main.go:16		0x10a32ca		48894c2468			MOVQ CX, 0x68(SP)			
  main.go:16		0x10a32cf		488d0d1a340400			LEAQ $f64.bfe62e42fefa39ef+8(SB), CX	
  main.go:16		0x10a32d6		48894c2470			MOVQ CX, 0x70(SP)			
  main.go:16		0x10a32db		8400				TESTB AL, 0(AX)				
  main.go:16		0x10a32dd		eb00				JMP 0x10a32df				
  main.go:16		0x10a32df		4889842490000000		MOVQ AX, 0x90(SP)			
  main.go:16		0x10a32e7		48c784249800000001000000	MOVQ $0x1, 0x98(SP)			
  main.go:16		0x10a32f3		48c78424a000000001000000	MOVQ $0x1, 0xa0(SP)			
  main.go:16		0x10a32ff		48890424			MOVQ AX, 0(SP)				
  main.go:16		0x10a3303		48c744240801000000		MOVQ $0x1, 0x8(SP)			
  main.go:16		0x10a330c		48c744241001000000		MOVQ $0x1, 0x10(SP)			
  main.go:16		0x10a3315		e8069affff			CALL fmt.Println(SB)			
  main.go:17		0x10a331a		0f57c0				XORPS X0, X0				
  main.go:17		0x10a331d		0f11442468			MOVUPS X0, 0x68(SP)			
  main.go:17		0x10a3322		488d442468			LEAQ 0x68(SP), AX			
  main.go:17		0x10a3327		4889442440			MOVQ AX, 0x40(SP)			
  main.go:17		0x10a332c		488b05355f0a00			MOVQ main.d(SB), AX			
  main.go:17		0x10a3333		48890424			MOVQ AX, 0(SP)				
  main.go:17		0x10a3337		e8a46df6ff			CALL runtime.convT64(SB)		
  main.go:17		0x10a333c		488b442408			MOVQ 0x8(SP), AX			
  main.go:17		0x10a3341		4889442460			MOVQ AX, 0x60(SP)			
  main.go:17		0x10a3346		488b4c2440			MOVQ 0x40(SP), CX			
  main.go:17		0x10a334b		8401				TESTB AL, 0(CX)				
  main.go:17		0x10a334d		488d15eca80000			LEAQ runtime.types+42816(SB), DX	
  main.go:17		0x10a3354		488911				MOVQ DX, 0(CX)				
  main.go:17		0x10a3357		488d7908			LEAQ 0x8(CX), DI			
  main.go:17		0x10a335b		833dae900e0000			CMPL $0x0, runtime.writeBarrier(SB)	
  main.go:17		0x10a3362		7402				JE 0x10a3366				
  main.go:17		0x10a3364		eb57				JMP 0x10a33bd				
  main.go:17		0x10a3366		48894108			MOVQ AX, 0x8(CX)			
  main.go:17		0x10a336a		eb00				JMP 0x10a336c				
  main.go:17		0x10a336c		488b442440			MOVQ 0x40(SP), AX			
  main.go:17		0x10a3371		8400				TESTB AL, 0(AX)				
  main.go:17		0x10a3373		eb00				JMP 0x10a3375				
  main.go:17		0x10a3375		4889442478			MOVQ AX, 0x78(SP)			
  main.go:17		0x10a337a		48c784248000000001000000	MOVQ $0x1, 0x80(SP)			
  main.go:17		0x10a3386		48c784248800000001000000	MOVQ $0x1, 0x88(SP)			
  main.go:17		0x10a3392		48890424			MOVQ AX, 0(SP)				
  main.go:17		0x10a3396		48c744240801000000		MOVQ $0x1, 0x8(SP)			
  main.go:17		0x10a339f		48c744241001000000		MOVQ $0x1, 0x10(SP)			
  main.go:17		0x10a33a8		e87399ffff			CALL fmt.Println(SB)			
  main.go:18		0x10a33ad		488bac24c0000000		MOVQ 0xc0(SP), BP			
  main.go:18		0x10a33b5		4881c4c8000000			ADDQ $0xc8, SP				
  main.go:18		0x10a33bc		c3				RET					
  main.go:17		0x10a33bd		0f1f00				NOPL 0(AX)				
  main.go:17		0x10a33c0		e83b0dfcff			CALL runtime.gcWriteBarrier(SB)		
  main.go:17		0x10a33c5		eba5				JMP 0x10a336c				
  main.go:15		0x10a33c7		e8340dfcff			CALL runtime.gcWriteBarrier(SB)		
  main.go:15		0x10a33cc		e99afeffff			JMP 0x10a326b				
  main.go:10		0x10a33d1		e8aaf1fbff			CALL runtime.morestack_noctxt(SB)	
  main.go:10		0x10a33d6		e9e5fdffff			JMP main.main(SB)			
