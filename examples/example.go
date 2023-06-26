package main

import (
	"emmecorelli/config.yaml"
	"fmt"
)

type Config struct {
	Server   string `yaml:"Server"     default:"myServer" encrypted:"true"`
	Username string `yaml:"User"       default:"Admin"`
	Password string `yaml:"Password"   default:"Admin"`
	Env      string `yaml:"Enviroment" default:"Test"`
	Skip     bool   `yaml:"Skip"       default:"false"`
}

const (
	filename   = "./config.yaml"
	confname   = "./cnf/config.yaml"
	cryptoname = "../test/crypto.yaml"
	key128     = "1234567890abcdef"
	plain      = "ServerChipred"
)

func main() {

	var got = Config{}

	if err := config.Load(&got, "../test/field.yaml", true, []byte(key128), true); err != nil {
		fmt.Printf("Unexpected error %q!\n", err)
	}
	fmt.Println("Server: ", got.Server)

	var cnf = Config{}
	if err := config.Load(&cnf, filename, false, nil, false); err != nil {
		fmt.Printf("Unable to read %q: %q!\n", filename, err)
	}

	fmt.Println("Server: ", cnf.Server)

	if err := config.Load(&cnf, confname, false, nil, false); err != nil {
		fmt.Printf("Unable to read %q: %q!\n", confname, err)
	}

	fmt.Println("Server: ", cnf.Server)

	if err := config.Load(&cnf, cryptoname, false, []byte(key128), false); err != nil {
		fmt.Printf("Unable to read %q: %q!\n", cryptoname, err)
	}

	fmt.Println("Server: ", cnf.Server)
	toEncrypt := []byte(plain)
	if err := config.Encrypt(&toEncrypt, []byte(key128)); err != nil {
		fmt.Printf("Encoding error: %q!\n", err)
	}
	fmt.Printf("Encoding [%s]: %s!\n", plain, string(toEncrypt))

}
