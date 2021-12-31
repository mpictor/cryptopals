package ecb

import (
	"bytes"
	"cryptopals/lib/vectors"
	"testing"
)

//func Encrypt_aes_ecb(plaintext, key []byte) []byte
func TestEncAesEcb(t *testing.T) {

	key := []byte("YELLOW SUBMARINE")
	ct := Encrypt_aes_ecb(vectors.TestTextAlign16, key)
	out := Decrypt_aes_ecb(ct, key)
	if !bytes.Equal(vectors.TestTextAlign16, out) {
		t.Errorf("in/out mismatch")
	}
}

//func Decrypt_aes_ecb(ciphertext, key []byte) []byte
//func TestDecAesEcb(t*testing.T){}
