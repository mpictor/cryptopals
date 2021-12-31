package lib

import (
	"bytes"
)

import (
	"testing"
)

var TestTextAlign16 = []byte(`You wanna battle me -- Anytime, anywhere

You thought that I was weak, Boy, you're dead wrong
So come on, everybody and sing this song

Say -- Play that funky music Say, go white boy, go white boy go
play that funky music Go white boy, go white boy, go
Lay down and boogie and play that funky music till you die.

Play that funky music Come on, Come on, let me hear
Play that funky music white boy you say it, say it
Play that funky music A little `)

//func Encrypt_aes_ecb(plaintext, key []byte) []byte
func TestEncAesEcb(t *testing.T) {

	key := []byte("YELLOW SUBMARINE")
	ct := Encrypt_aes_ecb(TestTextAlign16, key)
	out := Decrypt_aes_ecb(ct, key)
	if !bytes.Equal(TestTextAlign16, out) {
		t.Errorf("in/out mismatch")
	}
}

//func Decrypt_aes_ecb(ciphertext, key []byte) []byte
//func TestDecAesEcb(t*testing.T){}
