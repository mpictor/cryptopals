package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

type scoremap [256]int
type score struct {
	x byte
	s int
}

func main() {
	f, err := os.Open("c4.txt")
	if err != nil {
		panic(err)
	}
	var lines [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bin, _ := hex.DecodeString(scanner.Text())
		lines = append(lines, bin)
	}
	//var scores scoremap //:= make(scoremap)
	scores := make([]scoremap, len(lines))
	for i, l := range lines {
		scores[i] = getScores(l)
	}
	for i := range scores {
		showtopn(lines[i], scores[i], 1)
	}
}

func getScores(bin []byte) (scores scoremap) {
	var x byte
	for x = 0; x < 255; x++ {
		scores[x] = scoreSeq(myxor(bin, x))
	}
	//showtopn(bin,scores,5)
	return
}

func scoreSeq(bout []byte) (s int) {
	for _, b := range bout {
		if b == 32 {
			//space
			s += 5
			continue
		}
		if (b <= 90 && b >= 65) || (b <= 122 && b >= 97) {
			//A-Z a-z
			s++
			continue
		}
		if b == 9 || b == 10 || b == 13 {
			//no penalty for \t \n \r
			continue
		}
		if b >= 123 || b <= 31 {
			//non-printing
			s -= 5
			continue
		}
	}
	return
}

func myxor(bin []byte, x byte) (bout []byte) {
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
				continue scoreloop
			}
		}
		if s >= t.s {
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
	}
	return
}

func showtopn(bin []byte, scores scoremap, n int) {
	tops := topn(scores, n)
	for _, t := range tops {
		if t.s > len(bin) {
			fmt.Printf("0x%x %d %s\n", t.x, t.s, string(myxor(bin, t.x)))
		}
	}
}
