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

//count off by n, for the non-sentient
func Transpose(ks int, data []byte) (blocks [][]byte) {
	for b := 0; b < ks; b++ {
		var blk []byte
		for i := b; i < len(data); i += ks {
			blk = append(blk, data[i])
		}
		blocks = append(blocks, blk)
	}
	return
}
