package main

import (
	"bufio"
	"cryptopals/lib"
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
		scores := make([]lib.Scoremap, len(tb))
		key := make([]byte, len(tb))
		for i, b := range tb {
			scores[i] = lib.GetScores(b)
			s := lib.Top(scores[i], nil)
			key[i] = s.X
		}
		keys[idx].k = key
		keys[idx].s = lib.ScoreSeq(xor.EncryptXor(data[:128], key))
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
