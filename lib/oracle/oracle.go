package oracle

import (
	"cryptopals/lib"
	"encoding/base64"
)

func Encryption_oracle_with_unknown(key, plaintext, unknown []byte) (ciphertext []byte) {
	ud, _ := base64.StdEncoding.DecodeString(string(unknown))
	pt := append(plaintext, ud...)
	pt = lib.Pkcs7pad(pt, 16)
	ciphertext = lib.Encrypt_aes_ecb(pt, key)
	return
}
