package main

import (
	"encoding/hex"
	"fmt"
)

type scoremap [256]int
type score struct {
	x byte
	s int
}

func main() {
	in := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bin, _ := hex.DecodeString(in)
	//bout := make([]byte,len(bin))
	var scores scoremap //:= make(scoremap)
	//var x byte = 0x58
	var x byte
	for x = 0; x < 255; x++ {
		scores[x] = scoreSeq(myxor(bin, x))
	}
	//fmt.Println("scored")
	/*var max int
	var mk byte
	for k,s := range scores {
		if s >= max {
			max=s
			mk=k
		}
	}*/
	//fmt.Println(mk,max,string(myxor(bin,mk)))
	//fmt.Println(string(bout))
	showtopn(bin, scores, 5)
}

func scoreSeq(bout []byte) (s int) {
	for _, b := range bout {
		if b == 32 { //space
			s += 5
		} else if (b <= 90 && b >= 65) || (b <= 122 && b >= 97) { //A-Z a-z
			s++
		} else if b >= 123 || b <= 31 { //non-printing
			//FIXME \n \r \t etc shouldn't get bad scores
			s -= 5
		}
	}
	return
}

func myxor(bin []byte, x byte) (bout []byte) {
	//fmt.Println("xor",x)
	bout = make([]byte, len(bin))
	for i := range bin {
		bout[i] = bin[i] ^ x
	}
	return
}

func top(scores scoremap, not []byte) (t score) {
scoreloop:
	for x, s := range scores {
		for _, n := range not {
			if byte(x) == n {
				//fmt.Println("not",x)
				continue scoreloop
			}
		}
		if s >= t.s {
			//fmt.Println("top",x,s)
			t.s = s
			t.x = byte(x)
		}
	}
	return
}
func topn(scores scoremap, n int) (tops []score) {
	var not []byte
	for i := 0; i < n; i++ {
		s := top(scores, not)
		tops = append(tops, s)
		not = append(not, s.x)
		//fmt.Println("topn:",len(tops),len(not))
	}
	return
}

func showtopn(bin []byte, scores scoremap, n int) {
	tops := topn(scores, n)
	for _, t := range tops {
		fmt.Printf("0x%x %02d %s\n", t.x, t.s, string(myxor(bin, t.x)))
	}
}
