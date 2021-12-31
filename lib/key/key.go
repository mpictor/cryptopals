package key

import (
	"math/rand"
)

func RandKey() []byte {
	return RandBytes16()
}
func RandBytes16() []byte {
	return append(RandBytes8(), RandBytes8()...)
}

//splits a uint64 into bytes
func RandBytes8() (out []byte) {
	r := rand.Uint64()
	for i := 0; i < 8; i++ {
		v := r & 0xff
		r = r >> 8
		out = append(out, byte(v))
	}
	return
}
