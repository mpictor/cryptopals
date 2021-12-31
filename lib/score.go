package lib

import (
	"fmt"
)

//note - some of this only works for single-byte keys

type Scoremap [256]int
type Score struct {
	X byte
	S int
}

func GetScores(bin []byte) (scores Scoremap) {
	var x byte
	for x = 0; x < 255; x++ {
		scores[x] = ScoreSeq(EncryptXor(bin, []byte{x}))
	}
	//showtopn(bin,scores,5)
	return
}

//score plaintext candidate, higher scores better
func ScoreSeq(bout []byte) (s int) {
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

func Top(scores Scoremap, not []byte) (t Score) {
scoreloop:
	for x, s := range scores {
		for _, n := range not {
			if byte(x) == n {
				continue scoreloop
			}
		}
		if s >= t.S {
			t.S = s
			t.X = byte(x)
		}
	}
	return
}
func Topn(scores Scoremap, n int) (tops []Score) {
	var not []byte
	for i := 0; i < n; i++ {
		s := Top(scores, not)
		tops = append(tops, s)
		not = append(not, s.X)
	}
	return
}

func Showtopn(bin []byte, scores Scoremap, n int) {
	tops := Topn(scores, n)
	for _, t := range tops {
		if t.S > len(bin) {
			fmt.Printf("0x%x %d %s\n", t.X, t.S, string(EncryptXor(bin, []byte{t.X})))
		}
	}
}
