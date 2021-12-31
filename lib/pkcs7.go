package lib

import (
	"fmt"
)

//implement pkcs#7 padding

//len(out) must be a multiple of blksiz
func Pkcs7pad(in []byte, blksiz int) (out []byte) {
	lin := len(in)
	lpad := blksiz % lin
	if blksiz < lin {
		lpad = blksiz - lin%blksiz
	}
	out = make([]byte, lin+lpad)
	copy(out, in)
	for i := 0; i < lpad; i++ {
		out[lin+i] = byte(lpad)
	}
	return
}

func Pkcs7strip(in []byte, blksiz int) []byte {
	l := len(in)
	if l%blksiz != 0 {
		fmt.Println("bad size:", l, blksiz)
		return nil
	}
	padcount := in[l-1]
	if padcount >= byte(blksiz) {
		return in
	}
	//sanity
	for i := padcount; i > 0; i-- {
		if in[l-int(i)] != padcount {
			//fmt.Println("mismatch at", i, padcount, in)
			return in
		}
	}
	return in[:l-int(padcount)]
}
