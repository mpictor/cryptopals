package key

import (
	"cryptopals/lib"
	"cryptopals/lib/block"
	"cryptopals/lib/hamming"
	"cryptopals/lib/xor"
	"fmt"
	"math"
	"strings"
)

//returns best 3 key sizes in [min,max] for data
func SizeSearch(min, max int, data []byte) []int {
	ksd := KeySearchData{Data: data}
	ksd.SizeSearch(min, max)
	return ksd.Sizes
}

//search for plausible keys
func KeySearch(data []byte, ks []int) KeyScore {
	ksd := KeySearchData{
		Data:   data,
		Sizes:  ks,
		PTType: PTAsciiText,
	}
	ksd.KeySearch()
	return ksd.BestKey
}

//what will the plaintext look like once decrypted - ascii text, arm binary, ...
type PlainTextType int

const (
	PTUnknown PlainTextType = iota
	PTAsciiText
	PTArm32le
	//TODO
)

var strPTmap = map[string]PlainTextType{
	"ascii":   PTAsciiText,
	"arm32le": PTArm32le,
}

func PTFromStr(s string) (PlainTextType, error) {
	s = strings.ToLower(s)
	ptt, ok := strPTmap[s]
	if ok {
		return ptt, nil
	}
	var ptn []string
	for k := range strPTmap {
		ptn = append(ptn, k)
	}
	return 0, fmt.Errorf("unrecognized type %s; allowed are %v", s, ptn)
}

type KeySearchData struct {
	Data    []byte
	Sizes   []int
	PTType  PlainTextType
	BestKey KeyScore
}

//returns best 3 key sizes in [min,max] for data
func (ksd *KeySearchData) SizeSearch(min, max int) {
	var bestSizes = make([]int, 3)
	var tops = []int{math.MaxInt64, math.MaxInt64, math.MaxInt64}
	for keysize := min; keysize <= max; keysize++ {
		s := hamming.LongHam(keysize, ksd.Data)
		if lib.VV {
			fmt.Println("score", keysize, s)
		}
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
	ksd.Sizes = bestSizes
}

type KeyScore struct {
	K []byte
	S int
}

func (ks KeyScore) String() string {
	return fmt.Sprintf("score=%d len=%d key(hex)=%x, key(bytes)=%#v", ks.S, len(ks.K), ks.K, ks.K)
}

//search for plausible keys
func (ksd *KeySearchData) KeySearch() {
	var keys = make([]KeyScore, len(ksd.Sizes))
	for idx, s := range ksd.Sizes {
		tb := block.Transpose(s, ksd.Data)
		scores := make([]Scoremap, len(tb))
		key := make([]byte, len(tb))
		for i, b := range tb {
			scores[i] = GetScores(b, ksd.PTType)
			s := Top(scores[i], nil)
			key[i] = s.X
		}
		if len(key) == 0 {
			continue
		}
		keys[idx].K = key
		dlen := min(len(ksd.Data), len(key)*16)
		keys[idx].S = ScoreSeq(xor.EncryptXor(ksd.Data[:dlen], key), ksd.PTType)
		if lib.VV {
			fmt.Println("key #", idx, keys[idx].String())
		}
	}
	for _, k := range keys {
		if k.S > ksd.BestKey.S {
			ksd.BestKey.S = k.S
			ksd.BestKey.K = k.K
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
