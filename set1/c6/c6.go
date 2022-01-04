package main

import (
	"bufio"
	"cryptopals/lib/block"
	"cryptopals/lib/key"
	"cryptopals/lib/xor"
	"encoding/base64"
	"fmt"
	"os"
)

// https://cryptopals.com/sets/1/challenges/6

func main() {
	var data []byte
	f, err := os.Open("c6.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l, err := base64.StdEncoding.DecodeString(scanner.Text())
		if err != nil {
			panic(err)
		}
		data = append(data, l...)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	ks := key.SizeSearch(2, 40, data)
	fmt.Println("best key sizes:", ks)
	type keyscore struct {
		k []byte
		s int
	}
	var keys [3]keyscore
	for idx, s := range ks {
		tb := block.Transpose(s, data)
		scores := make([]key.Scoremap, len(tb))
		keybytes := make([]byte, len(tb))
		for i, b := range tb {
			scores[i] = key.GetScores(b, key.PTAsciiText)
			s := key.Top(scores[i], nil)
			keybytes[i] = s.X
		}
		keys[idx].k = keybytes
		keys[idx].s = key.ScoreSeq(xor.EncryptXor(data[:128], keybytes), key.PTAsciiText)
	}
	var best keyscore
	for _, k := range keys {
		if k.s > best.s {
			best.s = k.s
			best.k = k.k
		}
	}
	fmt.Printf("best key: score=%d key=%q\n", best.s, string(best.k))
	fmt.Println(string(xor.EncryptXor(data, best.k))[:60], "...")
}
