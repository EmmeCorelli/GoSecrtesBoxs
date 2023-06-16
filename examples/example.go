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
	filename = "./config.yaml"
	confname = "./cnf/config.yaml"
)

func main() {

	var cnf = Config{}
	if err := config.Load(&cnf, filename, false); err != nil {
		fmt.Printf("Unable to read %q: %q!", filename, err)
	}

	fmt.Println("Server: ", cnf.Server)

	if err := config.Load(&cnf, confname, false); err != nil {
		fmt.Printf("Unable to read %q: %q!", confname, err)
	}

	fmt.Println("Server: ", cnf.Server)

}
