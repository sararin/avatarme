package avatarme

import (
	"testing"
	"io/ioutil"
)

func TestEncrypt(t *testing.T) {
	g, _ := Encrypt([]byte("test msg"))
	if g != 16644892400343469032 {
		t.Errorf("Encrypt([]byte(\"test msg\")) = %d, want 16644892400343469032", g)
	}
	
	f, _ := ioutil.ReadFile("testfile.txt")
	g, _ = Encrypt(f)
	if g != 16644892400343469032 {
		t.Errorf("Encrypt([]byte(\"test msg\")) = %d, want 16644892400343469032", g)
	}

	g, _ = Encrypt([]byte(""))
	if g != 14695981039346656037 {
		t.Errorf("Encrypt([]byte(\"\")) = %d, want 14695981039346656037", g)
	}

	g, _ = Encrypt([]byte{0})
	if g != 12638153115695167455 {
		t.Errorf("Encrypt([]byte{0}) = %d, want 12638153115695167455", g)
	}

	g, err := Encrypt([]byte{})
	if err == nil {
		t.Errorf("Encrypt([]byte{}) = %d, want Error about empty array, g)
	}
}