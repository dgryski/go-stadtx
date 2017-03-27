// Package stadtx implements Stadtx Hash
/*

   https://github.com/demerphq/BeagleHash

*/
package stadtx

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
		return coreShort(v0, v1, key)
	}

	// len >= 32

	v2 := state[2] ^ (uint64(len(key)+3) * k2U64)
	v3 := state[3] ^ (uint64(len(key)+4) * k3U64)

	return coreLong(v0, v1, v2, v3, key)

}
