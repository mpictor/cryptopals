package main

import (
	"fmt"
	"time"
)

func main() {
	var b byte = 0
	for ; b <= 255; b++ {
		time.Sleep(time.Millisecond)
		if b < 5 || b > 250 {
			fmt.Println(b)
		}
	}
}
