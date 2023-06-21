package main

import (
	"emmecorelli/config.yaml"
	"fmt"
)

type Config struct {
	Server   string `yaml:"Server"     default:"myServer"`
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
)

func main() {

	var cnf = Config{}
	if err := config.Load(&cnf, filename, false, nil); err != nil {
		fmt.Printf("Unable to read %q: %q!\n", filename, err)
	}

	fmt.Println("Server: ", cnf.Server)

	if err := config.Load(&cnf, confname, false, nil); err != nil {
		fmt.Printf("Unable to read %q: %q!\n", confname, err)
	}

	fmt.Println("Server: ", cnf.Server)

	if err := config.Load(&cnf, cryptoname, false, []byte(key128)); err != nil {
		fmt.Printf("Unable to read %q: %q!\n", cryptoname, err)
	}

	fmt.Println("Server: ", cnf.Server)

}
