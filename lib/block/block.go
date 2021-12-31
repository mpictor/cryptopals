package block

import (
	"bytes"
)

func Find_blksize(enc func(pt []byte) []byte) (bs int) {
	//increase repeating input length until output repeats; divide by two, that's the blocksize
	var pt []byte
	for il := 2; il < 512; il += 2 {
		pt = append(pt, []byte("aa")...)
		out := enc(pt)
		if bytes.Equal(out[:il/2], out[il/2:il]) {
			bs = il / 2
			return
		}
	}
	return 0
}
