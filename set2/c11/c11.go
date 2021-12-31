package main

import (
	"cryptopals/lib/cbc"
	"cryptopals/lib/ecb"
	"cryptopals/lib/pkcs7"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	sub := []byte(`YELLAR SUBBY`)
	var pt []byte
	for i := 0; i < 256; i++ {
		pt = append(pt, sub...)
	}
	ct := encryption_oracle(pt)
	typ := detect_ecb_cbc(ct)
	fmt.Printf("pt=%s\nct=%q\ntype=%s\n", pt, ct, typ)
}
func encryption_oracle(plaintext []byte) (ciphertext []byte) {
	rand.Seed(int64(os.Getpid()))
	typ := encType(rand.Intn(2) + 1)
	prePadl := rand.Intn(5) + 5
	postPadl := rand.Intn(5) + 5

	key := randBytes16()
	pt := append(fill('j', prePadl), plaintext...)
	pt = append(pt, fill('k', postPadl)...)
	pt = pkcs7.Pkcs7pad(pt, 16)
	if typ == CBC {
		iv := randBytes16()
		ciphertext = cbc.Encrypt_aes_cbc(pt, key, iv)
	} else {
		ciphertext = ecb.Encrypt_aes_ecb(pt, key)
	}
	fmt.Printf("type=%s\n", typ)
	return
}

func fill(b byte, n int) (out []byte) {
	for i := 0; i < n; i++ {
		out = append(out, b)
	}
	return
}

func randBytes16() []byte {
	return append(randBytes8(), randBytes8()...)
}

//splits a uint64 into bytes
func randBytes8() (out []byte) {
	r := rand.Uint64()
	for i := 0; i < 8; i++ {
		v := r & 0xff
		r = r >> 8
		out = append(out, byte(v))
	}
	return
}

type encType int

const (
	unknown encType = iota
	ECB
	CBC
)
const blockLen = 16

func (e encType) String() string {
	switch e {
	case ECB:
		return "ECB"
	case CBC:
		return "CBC"
	case unknown:
		return "unknown"
	default:
		return "(undefined)"
	}
}

type block [blockLen]byte

func detect_ecb_cbc(data []byte) encType {
	blocks := make(map[block]interface{})
	for len(data) > 0 {
		var blk block
		copy(blk[:], data[:16])
		data = data[16:]
		_, ok := blocks[blk]
		if ok {
			return ECB
		}
		blocks[blk] = nil
	}

	return CBC
}
