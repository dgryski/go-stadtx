// +build amd64
// +build !noasm

package stadtx

//go:generate python -m peachpy.x86_64 stadtx.py -S -o stadtx_amd64.s -mabi=goasm
//go:noescape

func Hash(state *State, key []byte) uint64
