package main

import (
	"bufio"
	"fmt"
	"os"
)

const blockLen = 32

type block [blockLen]byte

func main() {
	f, err := os.Open("c8.txt")
	if err != nil {
		panic(err)
	}
	blocks := make(map[block]interface{})
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		l := len(line)
		for i := 0; i < l-31; i += 32 {
			var blk block
			copy(blk[:], line[i:i+32])
			_, present := blocks[blk]
			if present {
				fmt.Println("duplicate:", string(blk[:]))
			} else {
				blocks[blk] = nil
			}
		}
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
