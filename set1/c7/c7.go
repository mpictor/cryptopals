package main

import (
	"bufio"
	"cryptopals/lib/ecb"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("c7.txt")
	scanner := bufio.NewScanner(f)
	var buf []byte
	for scanner.Scan() {
		l, err := base64.StdEncoding.DecodeString(scanner.Text())
		if err != nil {
			panic(err)
		}
		buf = append(buf, l...)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	out := ecb.Decrypt_aes_ecb(buf, []byte("YELLOW SUBMARINE"))
	fmt.Println(string(out))
}
