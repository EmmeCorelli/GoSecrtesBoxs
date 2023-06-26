package config_test

import (
	"emmecorelli/config.yaml"
	"testing"
)

func TestCrypto(t *testing.T) {
	key128 := []byte("1234567890abcdef")
	got := []byte("Server: myServer")
	want := []byte("Server: myServer")

	if err := config.Encrypt(&got, key128); err != nil {
		t.Logf("Unexpected error %q!", err)
	}
	if err := config.Decrypt(&got, key128); err != nil {
		t.Logf("Unexpected error %q!", err)
	}

	if string(got) != string(want) {
		t.Errorf("got %q want %q", got, want)
	}

}
