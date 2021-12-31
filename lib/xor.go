package lib

func EncryptXor(in, key []byte) (out []byte) {
	out = make([]byte, len(in))
	l := len(key)
	for i := range in {
		out[i] = in[i] ^ key[i%l]
	}
	return
}
