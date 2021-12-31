package ecb

import (
	"crypto/aes"
	"fmt"
)

func Decrypt_aes_ecb(ciphertext, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(ciphertext)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	plaintext := make([]byte, len(ciphertext))
	offs := 0
	for offs < len(ciphertext) {
		block.Decrypt(plaintext[offs:], ciphertext[offs:])
		offs += bs
	}
	return plaintext
}

func Encrypt_aes_ecb(plaintext, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(plaintext)%bs != 0 {
		panic(fmt.Sprintf("Need a multiple of the blocksize %d, got %d", bs, len(plaintext)))
	}
	ciphertext := make([]byte, len(plaintext))
	offs := 0
	for offs < len(plaintext) {
		block.Encrypt(ciphertext[offs:], plaintext[offs:])
		offs += bs
	}
	return ciphertext
}
