// +build !noasm

// Generated by PeachPy 0.2.0 from stadtx.py


// func coreShort(v0 uint64, v1 uint64, key_base uintptr, key_len int64, key_cap int64) uint64
TEXT ·coreShort(SB),4,$0-48
	MOVQ v0+0(FP), AX
	MOVQ v1+8(FP), BX
	MOVQ key_base+16(FP), CX
	MOVQ key_len+24(FP), DX
	MOVQ DX, DI
	SHRQ $3, DI
	CMPQ DI, $0
	JEQ shortCore0
	CMPQ DI, $1
	JEQ shortCore1
	CMPQ DI, $2
	JEQ shortCore2
	CMPQ DI, $3
	JEQ shortCore3
shortCore3:
	MOVQ 0(CX), DI
	MOVQ $11248740756307325503, SI
	IMULQ SI, DI
	ADDQ DI, AX
	RORQ $17, AX
	XORQ BX, AX
	RORQ $53, BX
	ADDQ AX, BX
	ADDQ $8, CX
	SUBQ $8, DX
shortCore2:
	MOVQ 0(CX), DI
	MOVQ $11248740756307325503, SI
	IMULQ SI, DI
	ADDQ DI, AX
	RORQ $17, AX
	XORQ BX, AX
	RORQ $53, BX
	ADDQ AX, BX
	ADDQ $8, CX
	SUBQ $8, DX
shortCore1:
	MOVQ 0(CX), DI
	MOVQ $11248740756307325503, SI
	IMULQ SI, DI
	ADDQ DI, AX
	RORQ $17, AX
	XORQ BX, AX
	RORQ $53, BX
	ADDQ AX, BX
	ADDQ $8, CX
	SUBQ $8, DX
shortCore0:
	CMPQ DX, $0
	JEQ shortTail0
	CMPQ DX, $1
	JEQ shortTail1
	CMPQ DX, $2
	JEQ shortTail2
	CMPQ DX, $3
	JEQ shortTail3
	CMPQ DX, $4
	JEQ shortTail4
	CMPQ DX, $5
	JEQ shortTail5
	CMPQ DX, $6
	JEQ shortTail6
	MOVBQZX 6(CX), DX
	SHLQ $32, DX
	ADDQ DX, AX
shortTail6:
	MOVBQZX 5(CX), DX
	SHLQ $48, DX
	ADDQ DX, BX
shortTail5:
	MOVBQZX 4(CX), DX
	SHLQ $16, DX
	ADDQ DX, AX
shortTail4:
	XORQ DX, DX
	MOVL 0(CX), DX
	ADDQ DX, BX
	JMP shortAfter
shortTail3:
	MOVBQZX 2(CX), DX
	SHLQ $48, DX
	ADDQ DX, AX
shortTail2:
	XORQ DX, DX
	MOVW 0(CX), DX
	ADDQ DX, BX
	JMP shortAfter
shortTail1:
	MOVBQZX 0(CX), DX
	ADDQ DX, AX
shortTail0:
	RORQ $32, BX
	XORQ $255, BX
shortAfter:
	XORQ AX, BX
	RORQ $33, AX
	ADDQ BX, AX
	ROLQ $17, BX
	XORQ AX, BX
	ROLQ $43, AX
	ADDQ BX, AX
	ROLQ $31, BX
	SUBQ AX, BX
	ROLQ $13, AX
	XORQ BX, AX
	SUBQ AX, BX
	ROLQ $41, AX
	ADDQ BX, AX
	ROLQ $37, BX
	XORQ AX, BX
	RORQ $39, AX
	ADDQ BX, AX
	RORQ $15, BX
	ADDQ AX, BX
	ROLQ $15, AX
	XORQ BX, AX
	RORQ $5, BX
	XORQ BX, AX
	MOVQ AX, ret+40(FP)
	RET

// func coreLong(v0 uint64, v1 uint64, v2 uint64, v3 uint64, key_base uintptr, key_len int64, key_cap int64) uint64
TEXT ·coreLong(SB),4,$0-64
	MOVQ v0+0(FP), AX
	MOVQ v1+8(FP), BX
	MOVQ v2+16(FP), CX
	MOVQ v3+24(FP), DX
	MOVQ key_base+32(FP), DI
	MOVQ key_len+40(FP), SI
loop_begin:
		MOVQ 0(DI), BP
		MOVQ $2150174947, R8
		IMULQ R8, BP
		ADDQ BP, AX
		ROLQ $57, AX
		XORQ DX, AX
		MOVQ 8(DI), BP
		MOVQ $2174423983, R8
		IMULQ R8, BP
		ADDQ BP, BX
		ROLQ $63, BX
		XORQ CX, BX
		MOVQ 16(DI), BP
		MOVQ $2446010341, R8
		IMULQ R8, BP
		ADDQ BP, CX
		RORQ $47, CX
		ADDQ AX, CX
		MOVQ 24(DI), BP
		MOVQ $3248646593, R8
		IMULQ R8, BP
		ADDQ BP, DX
		RORQ $11, DX
		SUBQ BX, DX
		ADDQ $32, DI
		SUBQ $32, SI
		CMPQ SI, $32
		JGE loop_begin
	MOVQ SI, R8
	MOVQ SI, BP
	SHRQ $3, BP
	CMPQ BP, $0
	JEQ longCore0
	CMPQ BP, $1
	JEQ longCore1
	CMPQ BP, $2
	JEQ longCore2
	CMPQ BP, $3
	JEQ longCore3
longCore3:
	MOVQ 0(DI), BP
	MOVQ $2150174947, R9
	IMULQ R9, BP
	ADDQ BP, AX
	ROLQ $57, AX
	XORQ DX, AX
	ADDQ $8, DI
	SUBQ $8, SI
longCore2:
	MOVQ 0(DI), BP
	MOVQ $2174423983, R9
	IMULQ R9, BP
	ADDQ BP, BX
	ROLQ $63, BX
	XORQ CX, BX
	ADDQ $8, DI
	SUBQ $8, SI
longCore1:
	MOVQ 0(DI), BP
	MOVQ $2446010341, R9
	IMULQ R9, BP
	ADDQ BP, CX
	RORQ $47, CX
	ADDQ AX, CX
	ADDQ $8, DI
	SUBQ $8, SI
longCore0:
	RORQ $11, DX
	SUBQ BX, DX
	ADDQ $1, R8
	MOVQ $11248740756307325503, BP
	IMULQ BP, R8
	XORQ R8, AX
	CMPQ SI, $0
	JEQ longTail0
	CMPQ SI, $1
	JEQ longTail1
	CMPQ SI, $2
	JEQ longTail2
	CMPQ SI, $3
	JEQ longTail3
	CMPQ SI, $4
	JEQ longTail4
	CMPQ SI, $5
	JEQ longTail5
	CMPQ SI, $6
	JEQ longTail6
	CMPQ SI, $7
	JEQ longTail7
longTail7:
	MOVBQZX 6(DI), SI
	ADDQ SI, BX
longTail6:
	XORQ SI, SI
	MOVW 4(DI), SI
	ADDQ SI, CX
	MOVL 0(DI), SI
	ADDQ SI, DX
	JMP longAfter
longTail5:
	MOVBQZX 4(DI), SI
	ADDQ SI, BX
longTail4:
	XORQ SI, SI
	MOVL 0(DI), SI
	ADDQ SI, CX
	JMP longAfter
longTail3:
	MOVBQZX 2(DI), SI
	ADDQ SI, DX
longTail2:
	XORQ SI, SI
	MOVW 0(DI), SI
	ADDQ SI, BX
	JMP longAfter
longTail1:
	MOVBQZX 0(DI), SI
	ADDQ SI, CX
longTail0:
	ROLQ $32, DX
	XORQ $255, DX
longAfter:
	SUBQ CX, BX
	RORQ $19, AX
	SUBQ AX, BX
	RORQ $53, BX
	XORQ BX, DX
	SUBQ DX, AX
	ROLQ $43, DX
	ADDQ DX, AX
	RORQ $3, AX
	SUBQ AX, DX
	RORQ $43, CX
	SUBQ DX, CX
	ROLQ $55, CX
	XORQ AX, CX
	SUBQ CX, BX
	RORQ $7, DX
	SUBQ CX, DX
	RORQ $31, CX
	ADDQ CX, DX
	SUBQ BX, CX
	RORQ $39, DX
	XORQ DX, CX
	RORQ $17, DX
	XORQ CX, DX
	ADDQ DX, BX
	RORQ $9, BX
	XORQ BX, CX
	ROLQ $24, CX
	XORQ CX, DX
	RORQ $59, DX
	RORQ $1, AX
	SUBQ BX, AX
	XORQ BX, AX
	XORQ DX, CX
	XORQ CX, AX
	MOVQ AX, ret+56(FP)
	RET