package config_test

import (
	"errors"
	"io/fs"
	"os"
	"reflect"
	"testing"

	"emmecorelli/config.yaml"
)

func TestConfigFile(t *testing.T) {

	type Conf struct {
		Server string `yaml:"Server" default:"myServer"`
	}

	fileTest := []struct {
		name     string
		filename string
		create   bool
		server   string
		err      error
	}{
		{name: "empty  ", filename: "./empty.yaml", create: false, server: "myServer", err: nil},
		{name: "setted ", filename: "./setted.yaml", create: false, server: "NewServer", err: nil},
		{
			name: "err    ", filename: "./config.yaml", create: false, server: "Undef", err: fs.ErrNotExist,
		},
		{name: "default", filename: "", create: true, server: "myServer", err: nil},
	}

	if err := os.Chdir("test"); err != nil {
		t.Errorf("Test directory doesn't exists: %q!", err)
	}
	if err := os.Remove("./config.yaml"); err != nil {
		t.Logf("Unable to del config.yaml: %q!", err)
	}

	for _, tt := range fileTest {
		// using tt.name from the case to use it as the 't.Run' test name
		t.Run(tt.name, func(t *testing.T) {
			var got = Conf{}
			if err := config.Load(&got, tt.filename, tt.create, nil); err != tt.err {
				if errors.Is(err, tt.err) {
					return
				}
				t.Errorf("Unexpected error %q!", reflect.TypeOf(err))
			}

			if got.Server != tt.server {
				t.Errorf("got %q want %q", got.Server, tt.server)
			}
		})

	}

	if err := os.Chdir(".."); err != nil {
		t.Errorf("Fail to come back to home: %q!", err)
	}
}

func TestCryptFile(t *testing.T) {

	const filename = "crypto.yaml"
	key128 := []byte("1234567890abcdef")

	type Conf struct {
		Server string `yaml:"Server" default:"myServer"`
	}

	if err := os.Chdir("test"); err != nil {
		t.Errorf("Test directory doesn't exists: %q!", err)
	}
	if err := os.Remove(filename); err != nil {
		t.Logf("Unable to del %s: %q!", filename, err)
	}

	var got = Conf{}
	want := "myServer"

	if err := config.Load(&got, filename, true, nil); err != nil {
		t.Errorf("Unexpected error %q!", err)
	}
	if err := config.EncryptFile(filename, key128); err != nil {
		t.Errorf("Unable to encrypt %s: %q!", filename, err)
	}
	if err := config.Load(&got, filename, false, key128); err != nil {
		t.Errorf("Unexpected error %q!", err)
	}

	if got.Server != want {
		t.Errorf("got %q want %q", got.Server, want)
	}

}
