package cbc

import (
	"bytes"
	"cryptopals/lib/vectors"
	"testing"
)

//func Decrypt_aes_cbc(ciphertext, key, iv []byte) []byte
//func Encrypt_aes_cbc(plaintext, key, iv []byte) []byte
func TestAesCbc(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	ct := Encrypt_aes_cbc(vectors.TestTextAlign16, key, iv)
	out := Decrypt_aes_cbc(ct, key, iv)
	if !bytes.Equal(vectors.TestTextAlign16, out) {
		t.Errorf("mismatch:\n%q\n%q", vectors.TestTextAlign16, out)
	}
}
