package lib

import (
	"bytes"
	"testing"
)

//func Pkcs7pad(in []byte, blksiz int) (out []byte)
func TestPkcsPadding(t *testing.T) {
	in := []byte("YELLOW SUBMARINE")
	want := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	got := Pkcs7pad(in, 20)
	if !bytes.Equal(want, got) {
		t.Errorf("want \n%q, got\n%q\n", want, got)
	}
	want2 := []byte("YELLOW SUBMARINE\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e")
	got = Pkcs7pad(in, 15)
	if !bytes.Equal(want2, got) {
		t.Errorf("want \n%q (%d), got\n%q (%d)\n", want2, len(want2), got, len(got))
	}
}

//func Pkcs7strip(in []byte,blksiz int)(out []byte)
func TestPkcsStrip(t *testing.T) {
	want := []byte("YELLOW SUBMARINE")
	in := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	got := Pkcs7strip(in, 20)
	if !bytes.Equal(want, got) {
		t.Errorf("want \n%q, got\n%q\n", want, got)
	}
	in2 := []byte("YELLOW SUBMARINE\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e")
	got = Pkcs7strip(in2, 15)
	if !bytes.Equal(want, got) {
		t.Errorf("want \n%q (%d), got\n%q (%d)\n", want, len(want), got, len(got))
	}
	//not valid padding - pass through verbatim
	in3 := []byte("YELLOW SUBMARINE\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0d\x0e")
	got = Pkcs7strip(in3, 15)
	if !bytes.Equal(in3, got) {
		t.Errorf("want \n%q (%d), got\n%q (%d)\n", in3, len(in3), got, len(got))
	}
	in3 = []byte("YELLOW SUBMARINE\x0d\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e")
	got = Pkcs7strip(in3, 15)
	if !bytes.Equal(in3, got) {
		t.Errorf("want \n%q (%d), got\n%q (%d)\n", in3, len(in3), got, len(got))
	}
	in3 = []byte("YELLOW SUBMARINE\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f")
	got = Pkcs7strip(in3, 15)
	if !bytes.Equal(in3, got) {
		t.Errorf("want \n%q (%d), got\n%q (%d)\n", in3, len(in3), got, len(got))
	}
	//input not a multiple of blocksize
	in3 = []byte("YELLOW SUBMARINE\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f")
	got = Pkcs7strip(in3, 15)
	if len(got) != 0 {
		t.Errorf("want '', got\n%q (%d)\n", got, len(got))
	}
}
