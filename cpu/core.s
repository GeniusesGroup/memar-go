// func ActiveCoreID() uint64
TEXT ·ActiveCoreID(SB),7,$0
	MOVQ $0xB,AX
	XORQ CX,CX
	CPUID
	MOVQ DX,ret+0(FP)
	RET
