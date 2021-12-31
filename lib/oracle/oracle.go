package oracle

import (
	"cryptopals/lib/ecb"
	"cryptopals/lib/pkcs7"
	"encoding/base64"
)

func Encryption_oracle_with_unknown(key, plaintext, unknown []byte) (ciphertext []byte) {
	ud, _ := base64.StdEncoding.DecodeString(string(unknown))
	pt := append(plaintext, ud...)
	pt = pkcs7.Pkcs7pad(pt, 16)
	ciphertext = ecb.Encrypt_aes_ecb(pt, key)
	return
}
