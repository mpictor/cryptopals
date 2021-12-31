package lib

import (
	"testing"
)

func TestHamdist(t *testing.T) {
	a := []byte("this is a test")
	b := []byte("wokka wokka!!!")
	s := Hamdist(a, b)
	if s != 37 {
		t.Error(s, 37)
	}
	c := []byte("this is ` test")
	s = Hamdist(a, c)
	if s != 1 {
		t.Error(s, 1)
	}
	s = Hamdist(b, a)
	if s != 37 {
		t.Error(s, 37)
	}
	s = Hamdist(b, b)
	if s != 0 {
		t.Error(s, 0)
	}
}
