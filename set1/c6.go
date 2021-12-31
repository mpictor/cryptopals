package main

import (
	"bufio"
	"cryptopals/lib"
	"encoding/base64"
	"fmt"
	"os"
)

// https://cryptopals.com/sets/1/challenges/6
// stopped with good hamming code - need to work on the file now
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
	ks := keysizeSearch(2, 40, data)
	type keyscore struct {
		k []byte
		s int
	}
	var keys [3]keyscore
	for idx, s := range ks {
		tb := transposeBlocks(s, data)
		scores := make([]lib.Scoremap, len(tb))
		key := make([]byte, len(tb))
		for i, b := range tb {
			scores[i] = lib.GetScores(b)
			s := lib.Top(scores[i], nil)
			key[i] = s.X
		}
		keys[idx].k = key
		keys[idx].s = lib.ScoreSeq(lib.EncryptXor(data[:128], key))
	}
	var best keyscore
	for _, k := range keys {
		if k.s > best.s {
			best.s = k.s
			best.k = k.k
		}
	}
	fmt.Println(string(best.k), best.s)
	fmt.Println(string(lib.EncryptXor(data, best.k)))
}

//hamming distance for key of size s. goes through entire file, giving a more accurate number
func longHam(s int, data []byte) (h int) {
	nb := len(data) / s
	for i := 0; i < nb; i++ {
		a := data[i*s : (i+1)*s]
		b := data[(i+1)*s : (i+2)*s]
		h += lib.Hamdist(a, b)
	}
	return
}

//returns best 3 key sizes in [min,max] for data
func keysizeSearch(min, max int, data []byte) (bestSizes [3]int) {
	var tops = []int{66666, 66666, 66666}
	for keysize := min; keysize <= max; keysize++ {
		s := longHam(keysize, data)
		if s < tops[2] {
			if s < tops[1] {
				if s < tops[0] {
					tops[2] = tops[1]
					tops[1] = tops[0]
					tops[0] = s
					bestSizes[2] = bestSizes[1]
					bestSizes[1] = bestSizes[0]
					bestSizes[0] = keysize
				} else {
					tops[2] = tops[1]
					tops[1] = s
					bestSizes[2] = bestSizes[1]
					bestSizes[1] = keysize
				}
			} else {
				tops[2] = s
				bestSizes[2] = keysize
			}
		}
	}
	return
}

//count off by n, for the non-sentient
func transposeBlocks(ks int, data []byte) (blocks [][]byte) {
	for b := 0; b < ks; b++ {
		var blk []byte
		for i := b; i < len(data); i += ks {
			blk = append(blk, data[i])
		}
		blocks = append(blocks, blk)
	}
	return
}
