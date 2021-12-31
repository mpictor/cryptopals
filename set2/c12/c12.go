package main

import (
	// "bytes"
	"cryptopals/lib/block"
	"cryptopals/lib/enctype"
	"cryptopals/lib/key"
	"cryptopals/lib/oracle"
	"fmt"
)

func main() {
	blk := []byte("AAAAAAAAAAAAAAAA")
	key := key.RandKey()
	e := func(b []byte) []byte { return oracle.Encryption_oracle_with_unknown(key, b, unknown) }
	//#1
	bs := block.Find_blksize(e)
	fmt.Println("blocksize", bs)
	//#2
	et := enctype.Detect_ecb_cbc_16(e(append(blk, blk...)))
	fmt.Println("blocktype", et.String())

	//#3-6
	for i := 15; i >= 0; i-- {
		m := make_dict2(e, blk[1:], i)
		out := e(blk[:i])
		var b [16]byte
		copy(b[:], out)
		v := m[b]
		blk = append(blk[1:], v)
		fmt.Println("==", string(blk), i, v)
	}
	//too lazy to continue for additional bytes...
	fmt.Printf("blk %x\n", blk)
	fmt.Printf("key %x\n", key)
}

var unknown = []byte(`Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`)

func make_dict(enc func(pt []byte) []byte, pat []byte, pos int) map[byte][16]byte {
	m := make(map[byte][16]byte)
	var i byte
	p := make([]byte, 16)
	copy(p[:], pat) //unexpected: if we don't copy, pat's modifications are seen in calling function
	for ; i < 255; i++ {
		p[15] = i
		out := enc(p)
		var v [16]byte
		copy(v[:], out[:16])
		m[i] = v
	}
	return m
}

func make_dict2(enc func(pt []byte) []byte, pat []byte, pos int) map[[16]byte]byte {
	m := make(map[[16]byte]byte)
	var i byte
	p := make([]byte, 16)
	copy(p[:], pat) //unexpected: if we don't copy, pat's modifications are seen in calling function
	for ; i < 255; i++ {
		p[15] = i
		out := enc(p)
		var v [16]byte
		copy(v[:], out[:16])
		m[v] = i
	}
	return m
}
