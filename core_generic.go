package stadtx

import "encoding/binary"

func coreShort(v0, v1 uint64, key []byte) uint64 {

	switch len(key) >> 3 {
	case 3:
		v0 += binary.LittleEndian.Uint64(key) * k3U64
		v0 = rotr64(v0, 17) ^ v1
		v1 = rotr64(v1, 53) + v0
		key = key[8:]
		fallthrough
	case 2:
		v0 += binary.LittleEndian.Uint64(key) * k3U64
		v0 = rotr64(v0, 17) ^ v1
		v1 = rotr64(v1, 53) + v0
		key = key[8:]
		fallthrough
	case 1:
		v0 += binary.LittleEndian.Uint64(key) * k3U64
		v0 = rotr64(v0, 17) ^ v1
		v1 = rotr64(v1, 53) + v0
		key = key[8:]
		fallthrough
	case 0:
	default:
		break
	}

	switch len(key) & 0x7 {
	case 7:
		v0 += uint64(key[6]) << 32
		fallthrough
	case 6:
		v1 += uint64(key[5]) << 48
		fallthrough
	case 5:
		v0 += uint64(key[4]) << 16
		fallthrough
	case 4:
		v1 += uint64(binary.LittleEndian.Uint32(key))
		break
	case 3:
		v0 += uint64(key[2]) << 48
		fallthrough
	case 2:
		v1 += uint64(binary.LittleEndian.Uint16(key))
		break
	case 1:
		v0 += uint64(key[0])
		fallthrough
	case 0:
		v1 = rotl64(v1, 32) ^ 0xFF
		break
	}
	v1 ^= v0
	v0 = rotr64(v0, 33) + v1
	v1 = rotl64(v1, 17) ^ v0
	v0 = rotl64(v0, 43) + v1
	v1 = rotl64(v1, 31) - v0
	v0 = rotl64(v0, 13) ^ v1
	v1 -= v0
	v0 = rotl64(v0, 41) + v1
	v1 = rotl64(v1, 37) ^ v0
	v0 = rotr64(v0, 39) + v1
	v1 = rotr64(v1, 15) + v0
	v0 = rotl64(v0, 15) ^ v1
	v1 = rotr64(v1, 5)
	return v0 ^ v1
}

func coreLong(v0, v1, v2, v3 uint64, key []byte) uint64 {
	for len(key) >= 32 {
		v0 += binary.LittleEndian.Uint64(key[0:]) * k2U32
		v0 = rotl64(v0, 57) ^ v3
		v1 += binary.LittleEndian.Uint64(key[8:]) * k3U32
		v1 = rotl64(v1, 63) ^ v2
		v2 += binary.LittleEndian.Uint64(key[16:]) * k4U32
		v2 = rotr64(v2, 47) + v0
		v3 += binary.LittleEndian.Uint64(key[24:]) * k5U32
		v3 = rotr64(v3, 11) - v1
		key = key[32:]
	}

	l := len(key) // needed for after the switch

	switch len(key) >> 3 {
	case 3:
		v0 += binary.LittleEndian.Uint64(key) * k2U32
		key = key[8:]
		v0 = rotl64(v0, 57) ^ v3
		fallthrough
	case 2:
		v1 += binary.LittleEndian.Uint64(key) * k3U32
		key = key[8:]
		v1 = rotl64(v1, 63) ^ v2
		fallthrough
	case 1:
		v2 += binary.LittleEndian.Uint64(key) * k4U32
		key = key[8:]
		v2 = rotr64(v2, 47) + v0
		fallthrough

	case 0:
		v3 = rotr64(v3, 11) - v1
	}

	v0 ^= uint64(l+1) * k3U64
	switch len(key) & 0x7 {
	case 7:
		v1 += uint64(key[6])
		fallthrough
	case 6:
		v2 += uint64(binary.LittleEndian.Uint16(key[4:]))
		v3 += uint64(binary.LittleEndian.Uint32(key))
		break
	case 5:
		v1 += uint64(key[4])
		fallthrough
	case 4:
		v2 += uint64(binary.LittleEndian.Uint32(key))
		break
	case 3:
		v3 += uint64(key[2])
		fallthrough
	case 2:
		v1 += uint64(binary.LittleEndian.Uint16(key))
		break
	case 1:
		v2 += uint64(key[0])
		fallthrough
	case 0:
		v3 = rotl64(v3, 32) ^ 0xFF
		break
	}

	v1 -= v2
	v0 = rotr64(v0, 19)
	v1 -= v0
	v1 = rotr64(v1, 53)
	v3 ^= v1
	v0 -= v3
	v3 = rotl64(v3, 43)
	v0 += v3
	v0 = rotr64(v0, 3)
	v3 -= v0
	v2 = rotr64(v2, 43) - v3
	v2 = rotl64(v2, 55) ^ v0
	v1 -= v2
	v3 = rotr64(v3, 7) - v2
	v2 = rotr64(v2, 31)
	v3 += v2
	v2 -= v1
	v3 = rotr64(v3, 39)
	v2 ^= v3
	v3 = rotr64(v3, 17) ^ v2
	v1 += v3
	v1 = rotr64(v1, 9)
	v2 ^= v1
	v2 = rotl64(v2, 24)
	v3 ^= v2
	v3 = rotr64(v3, 59)
	v0 = rotr64(v0, 1) - v1

	return v0 ^ v1 ^ v2 ^ v3
}
