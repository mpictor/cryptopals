package key

import (
	"cryptopals/lib/xor"
	"encoding/binary"
	"fmt"
	"os"
)

//note - some of this only works for single-byte keys

type Scoremap [256]int
type Score struct {
	X byte
	S int
}

//TODO methods for scoring executable code

func GetScores(bin []byte, ptype PlainTextType) (scores Scoremap) {
	var x byte
	for x = 0; x < 255; x++ {
		scores[x] = ScoreSeq(xor.EncryptXor(bin, []byte{x}), ptype)
	}
	//showtopn(bin,scores,5)
	return
}

//score plaintext candidate, higher scores better
func ScoreSeq(bout []byte, ptype PlainTextType) int {
	switch ptype {
	case PTAsciiText:
		return ScoreASCIISeq(bout)
	case PTArm32le:
		return ScoreArm32leSeq(bout)
	default:
		fmt.Fprintf(os.Stderr, "unhandled plaintext type %d\n", ptype)
		os.Exit(1)
	}
	return 0 //unreachable
}

//score plaintext candidate, higher scores better
func ScoreASCIISeq(bout []byte) (s int) {
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

//score plaintext candidate, higher scores better
func ScoreArm32leSeq(bout []byte) (s int) {
	//arm 32 le hist (sample size 1, go binary)
	//first is most prevalent
	//0xe58d0004, 0xe3500000, 0xe3a00000, 0xe58d0008, 0xe59b0000, 0x00000000, 0xe58d1004, 0xe58d000c, 0xe58d1008, 0xe1a0300e
	insts := []uint32{0xe58d0004, 0xe3500000, 0xe3a00000, 0xe58d0008, 0xe59b0000, 0x00000000, 0xe58d1004, 0xe58d000c, 0xe58d1008, 0xe1a0300e}
	for len(bout) >= 4 {
		inst := binary.LittleEndian.Uint32(bout[:4])
		bout = bout[4:]
		for n, i := range insts {
			if i == inst {
				s++
				if n < len(insts)/2 {
					//the first ones are more prevalent, so give them a bit of a boost
					s++
				}
				break
			}
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
			fmt.Printf("0x%x %d %s\n", t.X, t.S, string(xor.EncryptXor(bin, []byte{t.X})))
		}
	}
}
