import peachpy.x86_64


k0U64 = 0xb89b0f8e1655514f
k1U64 = 0x8c6f736011bd5127
k2U64 = 0x8f29bd94edce7b39
k3U64 = 0x9c1b8e1e9628323f

k2U32 = 0x802910e3
k3U32 = 0x819b13af
k4U32 = 0x91cb27e5
k5U32 = 0xc1a269c1

def imul(r,k):
    t = GeneralPurposeRegister64()
    MOV(t, k)
    IMUL(r, t)

v0 = Argument(uint64_t)
v1 = Argument(uint64_t)
v2 = Argument(uint64_t)
v3 = Argument(uint64_t)
key_base = Argument(ptr())
key_len = Argument(int64_t)
key_cap = Argument(int64_t)

with Function("coreShort", (v0, v1, key_base, key_len, key_cap), uint64_t) as function:

    reg_v0 = GeneralPurposeRegister64()
    reg_v1 = GeneralPurposeRegister64()

    LOAD.ARGUMENT(reg_v0, v0)
    LOAD.ARGUMENT(reg_v1, v1)

    reg_ptr = GeneralPurposeRegister64()
    reg_ptr_len = GeneralPurposeRegister64()

    LOAD.ARGUMENT(reg_ptr, key_base)
    LOAD.ARGUMENT(reg_ptr_len, key_len)

    reg_u64s = GeneralPurposeRegister64()
    MOV(reg_u64s, reg_ptr_len)
    SHR(reg_u64s, 3)

    labels = [Label("shortCore%d" % i) for i in range(4)]

    for i in range(4):
        CMP(reg_u64s, i)
        JE(labels[i])

    for i in range(3, 0, -1):
        LABEL(labels[i])
        r = GeneralPurposeRegister64()
        MOV(r, [reg_ptr])
        imul(r, k3U64)
        ADD(reg_v0, r)
        ROR(reg_v0, 17)
        XOR(reg_v0, reg_v1)
        ROR(reg_v1, 53)
        ADD(reg_v1, reg_v0)
        ADD(reg_ptr,8)
        SUB(reg_ptr_len,8)

    LABEL(labels[0])

    labels = [Label("shortTail%d" % i) for i in range(8)]

    for i in range(7):
        CMP(reg_ptr_len, i)
        JE(labels[i])

    after = Label("shortAfter")

    reg_ch = GeneralPurposeRegister64()

    LABEL(labels[7])
    MOVZX(reg_ch, byte[reg_ptr+6])
    SHL(reg_ch, 32)
    ADD(reg_v0, reg_ch)

    LABEL(labels[6])
    MOVZX(reg_ch, byte[reg_ptr+5])
    SHL(reg_ch, 48)
    ADD(reg_v1, reg_ch)

    LABEL(labels[5])
    MOVZX(reg_ch, byte[reg_ptr+4])
    SHL(reg_ch, 16)
    ADD(reg_v0, reg_ch)

    LABEL(labels[4])
    XOR(reg_ch, reg_ch)
    MOV(reg_ch.as_dword, dword[reg_ptr])
    ADD(reg_v1, reg_ch)

    JMP(after)

    LABEL(labels[3])
    MOVZX(reg_ch, byte[reg_ptr+2])
    SHL(reg_ch, 48)
    ADD(reg_v0, reg_ch)

    LABEL(labels[2])
    XOR(reg_ch, reg_ch)
    MOV(reg_ch.as_word, word[reg_ptr])
    ADD(reg_v1, reg_ch)

    JMP(after)

    LABEL(labels[1])
    MOVZX(reg_ch, byte[reg_ptr])
    ADD(reg_v0, reg_ch)

    LABEL(labels[0])
    ROR(reg_v1, 32)
    XOR(reg_v1, 0xFF)

    LABEL(after)

    XOR(reg_v1, reg_v0)

    ROR(reg_v0, 33)
    ADD(reg_v0, reg_v1)

    ROL(reg_v1, 17)
    XOR(reg_v1, reg_v0)

    ROL(reg_v0, 43)
    ADD(reg_v0, reg_v1)

    ROL(reg_v1, 31)
    SUB(reg_v1, reg_v0)

    ROL(reg_v0, 13)
    XOR(reg_v0, reg_v1)

    SUB(reg_v1, reg_v0)

    ROL(reg_v0, 41)
    ADD(reg_v0, reg_v1)

    ROL(reg_v1, 37)
    XOR(reg_v1, reg_v0)

    ROR(reg_v0, 39)
    ADD(reg_v0, reg_v1)

    ROR(reg_v1, 15)
    ADD(reg_v1, reg_v0)

    ROL(reg_v0, 15)
    XOR(reg_v0, reg_v1)

    ROR(reg_v1, 5)

    XOR(reg_v0, reg_v1)

    RETURN(reg_v0)


with Function("coreLong", (v0, v1, v2, v3, key_base, key_len, key_cap), uint64_t) as function:

    reg_v0 = GeneralPurposeRegister64()
    reg_v1 = GeneralPurposeRegister64()
    reg_v2 = GeneralPurposeRegister64()
    reg_v3 = GeneralPurposeRegister64()

    LOAD.ARGUMENT(reg_v0, v0)
    LOAD.ARGUMENT(reg_v1, v1)
    LOAD.ARGUMENT(reg_v2, v2)
    LOAD.ARGUMENT(reg_v3, v3)

    reg_ptr = GeneralPurposeRegister64()
    reg_ptr_len = GeneralPurposeRegister64()

    LOAD.ARGUMENT(reg_ptr, key_base)
    LOAD.ARGUMENT(reg_ptr_len, key_len)

    # we know len is >= 32 at this point

    r = GeneralPurposeRegister64()
    with Loop() as loop:
        MOV(r, [reg_ptr])
        imul(r, k2U32)
        ADD(reg_v0, r)
        ROL(reg_v0, 57)
        XOR(reg_v0, reg_v3)

        MOV(r, [reg_ptr + 8])
        imul(r, k3U32)
        ADD(reg_v1, r)
        ROL(reg_v1, 63)
        XOR(reg_v1, reg_v2)

        MOV(r, [reg_ptr + 16])
        imul(r, k4U32)
        ADD(reg_v2, r)
        ROR(reg_v2, 47)
        ADD(reg_v2, reg_v0)

        MOV(r, [reg_ptr + 24])
        imul(r, k5U32)
        ADD(reg_v3, r)
        ROR(reg_v3, 11)
        SUB(reg_v3, reg_v1)

        ADD(reg_ptr, 32)
        SUB(reg_ptr_len, 32)

        CMP(reg_ptr_len, 32)
        JGE(loop.begin)


    reg_ptr_len_saved = GeneralPurposeRegister64()
    MOV(reg_ptr_len_saved, reg_ptr_len)

    reg_u64s = GeneralPurposeRegister64()
    MOV(reg_u64s, reg_ptr_len)
    SHR(reg_u64s, 3)

    labels = [Label("longCore%d" % i) for i in range(4)]

    for i in range(4):
        CMP(reg_u64s, i)
        JE(labels[i])

    LABEL(labels[3])

    MOV(r, [reg_ptr])
    imul(r, k2U32)
    ADD(reg_v0, r)
    ROL(reg_v0, 57)
    XOR(reg_v0, reg_v3)
    ADD(reg_ptr, 8)
    SUB(reg_ptr_len, 8)

    LABEL(labels[2])

    MOV(r, [reg_ptr])
    imul(r, k3U32)
    ADD(reg_v1, r)
    ROL(reg_v1, 63)
    XOR(reg_v1, reg_v2)
    ADD(reg_ptr, 8)
    SUB(reg_ptr_len, 8)

    LABEL(labels[1])

    MOV(r, [reg_ptr])
    imul(r, k4U32)
    ADD(reg_v2, r)
    ROR(reg_v2, 47)
    ADD(reg_v2, reg_v0)
    ADD(reg_ptr, 8)
    SUB(reg_ptr_len, 8)

    LABEL(labels[0])

    ROR(reg_v3, 11)
    SUB(reg_v3, reg_v1)

    ADD(reg_ptr_len_saved, 1)
    imul(reg_ptr_len_saved, k3U64)
    XOR(reg_v0, reg_ptr_len_saved)

    labels = [Label("longTail%d" % i) for i in range(8)]

    for i in range(8):
        CMP(reg_ptr_len, i)
        JE(labels[i])

    after = Label("longAfter")

    reg_ch = GeneralPurposeRegister64()

    LABEL(labels[7])
    MOVZX(reg_ch, byte[reg_ptr+6])
    ADD(reg_v1, reg_ch)

    LABEL(labels[6])
    XOR(reg_ch, reg_ch)
    MOV(reg_ch.as_word, word[reg_ptr + 4])
    ADD(reg_v2, reg_ch)
    MOV(reg_ch.as_dword, dword[reg_ptr])
    ADD(reg_v3, reg_ch)
    JMP(after)

    LABEL(labels[5])
    MOVZX(reg_ch, byte[reg_ptr+4])
    ADD(reg_v1, reg_ch)

    LABEL(labels[4])
    XOR(reg_ch, reg_ch)
    MOV(reg_ch.as_dword, dword[reg_ptr])
    ADD(reg_v2, reg_ch)

    JMP(after)

    LABEL(labels[3])
    MOVZX(reg_ch, byte[reg_ptr+2])
    ADD(reg_v3, reg_ch)

    LABEL(labels[2])
    XOR(reg_ch, reg_ch)
    MOV(reg_ch.as_word, word[reg_ptr])
    ADD(reg_v1, reg_ch)

    JMP(after)

    LABEL(labels[1])
    MOVZX(reg_ch, byte[reg_ptr])
    ADD(reg_v2, reg_ch)

    LABEL(labels[0])
    ROL(reg_v3, 32)
    XOR(reg_v3, 0xFF)

    LABEL(after)

    ## finalize

    SUB(reg_v1, reg_v2)
    ROR(reg_v0, 19)
    SUB(reg_v1, reg_v0)
    ROR(reg_v1, 53)
    XOR(reg_v3, reg_v1)
    SUB(reg_v0, reg_v3)
    ROL(reg_v3, 43)
    ADD(reg_v0, reg_v3)
    ROR(reg_v0, 3)
    SUB(reg_v3, reg_v0)
    ROR(reg_v2, 43)
    SUB(reg_v2, reg_v3)
    ROL(reg_v2, 55)
    XOR(reg_v2, reg_v0)
    SUB(reg_v1, reg_v2)
    ROR(reg_v3, 7)
    SUB(reg_v3, reg_v2)
    ROR(reg_v2, 31)
    ADD(reg_v3, reg_v2)
    SUB(reg_v2, reg_v1)
    ROR(reg_v3, 39)
    XOR(reg_v2, reg_v3)
    ROR(reg_v3, 17)
    XOR(reg_v3, reg_v2)
    ADD(reg_v1, reg_v3)
    ROR(reg_v1, 9)
    XOR(reg_v2, reg_v1)
    ROL(reg_v2, 24)
    XOR(reg_v3, reg_v2)
    ROR(reg_v3, 59)
    ROR(reg_v0, 1)
    SUB(reg_v0, reg_v1)

    XOR(reg_v0, reg_v1)
    XOR(reg_v2, reg_v3)

    XOR(reg_v0, reg_v2)

    RETURN(reg_v0)
