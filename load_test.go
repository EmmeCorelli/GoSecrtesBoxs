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
			name: "err ", filename: "./config.yaml", create: false, server: "Undef", err: fs.ErrNotExist,
		},
		{name: "default", filename: "", create: true, server: "myServer", err: nil},
	}

	if err := os.Chdir("test"); err != nil {
		t.Errorf("Test directory doesn't exists: %q!", err)
	}
	if err := os.Remove("./config.yaml"); err != nil {
		t.Errorf("Unable to del config.yaml: %q!", err)
	}

	for _, tt := range fileTest {
		// using tt.name from the case to use it as the 't.Run' test name
		t.Run(tt.name, func(t *testing.T) {
			var got = Conf{}
			if err := config.Load(&got, tt.filename, tt.create); err != tt.err {
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
}

func TestStruct(t *testing.T) {

	type Conf struct {
		Server string `yaml:"Server"`
	}

}
