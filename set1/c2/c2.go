package main

import (
	"encoding/hex"

	"fmt"
)

func main() {
	in := "1c0111001f010100061a024b53535009181c"
	x := "686974207468652062756c6c277320657965"
	out := "746865206b696420646f6e277420706c6179"
	bin, _ := hex.DecodeString(in)
	bout := make([]byte, len(bin))
	bx, _ := hex.DecodeString(x)
	for i := range bin {
		bout[i] = bin[i] ^ bx[i]
	}
	xout := hex.EncodeToString(bout)
	if xout == out {
		fmt.Println("match")
	}
	fmt.Println(out)
	fmt.Println(xout)
	//fmt.Println("Hello World")
}
