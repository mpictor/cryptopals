package lib

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
