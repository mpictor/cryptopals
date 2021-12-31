package lib

import (
	"bytes"
	"testing"
)

//func Decrypt_aes_cbc(ciphertext, key, iv []byte) []byte
//func Encrypt_aes_cbc(plaintext, key, iv []byte) []byte
func TestAesCbc(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	ct := Encrypt_aes_cbc(TestTextAlign16, key, iv)
	out := Decrypt_aes_cbc(ct, key, iv)
	if !bytes.Equal(TestTextAlign16, out) {
		t.Errorf("mismatch:\n%q\n%q", TestTextAlign16, out)
	}
}
