// +build amd64
// +build !noasm

package stadtx

//go:generate python -m peachpy.x86_64 stadtx.py -S -o stadtx_amd64.s -mabi=goasm
//go:noescape

func coreShort(v0, v1 uint64, key []byte) uint64
func coreLong(v0, v1, v2, v3 uint64, key []byte) uint64
