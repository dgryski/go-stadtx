// Package stadtx implements Stadtx Hash
/*

   https://github.com/demerphq/BeagleHash

*/
package stadtx

import "encoding/binary"

func rotl64(x uint64, r uint64) uint64 {
	return (((x) << (r)) | ((x) >> (64 - r)))
}

func rotr64(x uint64, r uint64) uint64 {
	return (((x) >> (r)) | ((x) << (64 - r)))
}

func scramble64(v, prime uint64) uint64 {
	v ^= (v >> 13)
	v ^= (v << 35)
	v ^= (v >> 30)
	v *= prime
	v ^= (v >> 19)
	v ^= (v << 15)
	v ^= (v >> 46)
	return v
}

func SeedState(seed []uint64) State {

	var state State

	state[0] = seed[0] ^ 0x43f6a8885a308d31
	state[1] = seed[1] ^ 0x3198a2e03707344a
	state[2] = seed[0] ^ 0x4093822299f31d00
	state[3] = seed[1] ^ 0x82efa98ec4e6c894

	if state[0] == 0 {
		state[0] = 1
	}
	if state[1] == 0 {
		state[1] = 2
	}
	if state[2] == 0 {
		state[2] = 4
	}
	if state[3] == 0 {
		state[3] = 8
	}

	state[0] = scramble64(state[0], 0x801178846e899d17)
	state[0] = scramble64(state[0], 0xdd51e5d1c9a5a151)
	state[1] = scramble64(state[1], 0x93a7d6c8c62e4835)
	state[1] = scramble64(state[1], 0x803340f36895c2b5)
	state[2] = scramble64(state[2], 0xbea9344eb7565eeb)
	state[2] = scramble64(state[2], 0xcd95d1e509b995cd)
	state[3] = scramble64(state[3], 0x9999791977e30c13)
	state[3] = scramble64(state[3], 0xaab8b6b05abfc6cd)

	return state
}

const (
	k0U64 = 0xb89b0f8e1655514f
	k1U64 = 0x8c6f736011bd5127
	k2U64 = 0x8f29bd94edce7b39
	k3U64 = 0x9c1b8e1e9628323f

	k2U32 = 0x802910e3
	k3U32 = 0x819b13af
	k4U32 = 0x91cb27e5
	k5U32 = 0xc1a269c1
)

type State [4]uint64

func Hash(state *State, key []byte) uint64 {

	v0 := state[0] ^ (uint64(len(key)+1) * k0U64)
	v1 := state[1] ^ (uint64(len(key)+2) * k1U64)

	if len(key) < 32 {

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

	// len >= 32

	v2 := state[2] ^ (uint64(len(key)+3) * k2U64)
	v3 := state[3] ^ (uint64(len(key)+4) * k3U64)

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
