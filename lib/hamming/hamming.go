package hamming

func Hamdist(a, b []byte) (ham int) {
	if len(a) != len(b) {
		panic("differing lengths")
	}
	for i := range a {
		if a[i] == b[i] {
			continue
		}
		//count different bits
		m := byte(1)
		for s := 0; s < 7; s++ {
			if a[i]&m != b[i]&m {
				ham++
			}
			m <<= 1
		}
	}
	return
}

//hamming distance for key of size s. goes through entire file, giving a more accurate number
func LongHam(s int, data []byte) (h int) {
	nb := len(data) / s
	for i := 0; i < nb; i++ {
		a := data[i*s : (i+1)*s]
		b := data[(i+1)*s : (i+2)*s]
		h += Hamdist(a, b)
	}
	return
}
