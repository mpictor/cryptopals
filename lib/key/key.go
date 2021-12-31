package key

import (
	"cryptopals/lib/hamming"
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

//returns best 3 key sizes in [min,max] for data
func SizeSearch(min, max int, data []byte) (bestSizes [3]int) {
	var tops = []int{66666, 66666, 66666}
	for keysize := min; keysize <= max; keysize++ {
		s := hamming.LongHam(keysize, data)
		if s < tops[2] {
			if s < tops[1] {
				if s < tops[0] {
					tops[2] = tops[1]
					tops[1] = tops[0]
					tops[0] = s
					bestSizes[2] = bestSizes[1]
					bestSizes[1] = bestSizes[0]
					bestSizes[0] = keysize
				} else {
					tops[2] = tops[1]
					tops[1] = s
					bestSizes[2] = bestSizes[1]
					bestSizes[1] = keysize
				}
			} else {
				tops[2] = s
				bestSizes[2] = keysize
			}
		}
	}
	return
}
