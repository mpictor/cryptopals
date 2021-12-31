package cbc

import (
	"crypto/aes"
	"cryptopals/lib/xor"
	"fmt"
)

func Encrypt_aes_cbc(plaintext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(plaintext)%bs != 0 {
		panic(fmt.Sprintf("Need a multiple of the blocksize %d, got %d", bs, len(plaintext)))
	}
	if len(iv) != bs {
		panic(fmt.Sprintf("iv is not same size as block - got %d, need %d", len(iv), bs))
	}
	ciphertext := make([]byte, len(plaintext))
	offs := 0
	for offs < len(plaintext) {
		var xin []byte
		if offs == 0 {
			xin = iv
		} else {
			xin = ciphertext[offs-bs : offs]
		}
		x := xor.EncryptXor(plaintext[offs:offs+bs], xin)

		block.Encrypt(ciphertext[offs:], x)
		offs += bs
	}
	return ciphertext
}

func Decrypt_aes_cbc(ciphertext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(ciphertext)%bs != 0 {
		panic(fmt.Sprintf("Need a multiple of the blocksize %d, got %d", bs, len(ciphertext)))
	}
	if len(iv) != bs {
		panic(fmt.Sprintf("iv is not same size as block - got %d, need %d", len(iv), bs))
	}
	plaintext := make([]byte, len(ciphertext))
	offs := 0
	tmp := make([]byte, bs)
	for offs < len(ciphertext) {
		var xin []byte
		if offs == 0 {
			xin = iv
		} else {
			xin = ciphertext[offs-bs : offs]
		}

		block.Decrypt(tmp, ciphertext[offs:])
		x := xor.EncryptXor(tmp, xin)
		copy(plaintext[offs:offs+bs], x)
		offs += bs
	}
	return plaintext
}
