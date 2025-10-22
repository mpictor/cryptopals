package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	out := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	bytes, err := hex.DecodeString(in)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(bytes)
	//if err != nil {panic(err)}
	if out == enc {
		fmt.Println("equal")
	}
	fmt.Println("enc", enc)
	fmt.Println("out", out)
}
