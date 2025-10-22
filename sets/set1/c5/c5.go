package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	in := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	key := []byte("ICE")
	xwant := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	//want,_ := hex.DecodeString(xwant)
	got := encryptXor(in, key)
	xgot := hex.EncodeToString(got)
	if xgot == xwant {
		fmt.Println("match")
	}
	fmt.Println(xwant)
	fmt.Println(xgot)
}

func encryptXor(in, key []byte) (out []byte) {
	out = make([]byte, len(in))
	l := len(key)
	for i := range in {
		out[i] = in[i] ^ key[i%l]
	}
	return
}
