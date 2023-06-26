package main

import (
	"fmt"
	"log"
	"os"

	"emmecorelli/config.yaml"

	"github.com/akamensky/argparse"
	"gopkg.in/yaml.v2"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("config-cli", "Manipulate config file and keys")
	// Create file flag
	f := parser.File("f", "file", os.O_RDWR, 0600, &argparse.Options{Required: true, Help: "Config file to manipulate"})

	encryptCmd := parser.NewCommand("encrypt", "encrypt Config file, out is stdout")
	decryptCmd := parser.NewCommand("decrypt", "decrypt Config file, out is stdout")

	p := parser.String("p", "password", &argparse.Options{Required: true, Help: "Encryption password"})
	k := parser.String("k", "key", &argparse.Options{Help: "key value to encrypt/decrypt password"})

	// Parse input
	if err := parser.Parse(os.Args); err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		log.Print(parser.Usage(err))
		return
	}

	data, err := getData(f.Name(), *k)
	if err != nil {
		panic(err)
	}

	if encryptCmd.Happened() {
		if err := config.Encrypt(&data, []byte(*p)); err != nil {
			panic(err)
		}
		log.Printf("Encrypt process\n%s\n", string(data))
	} else if decryptCmd.Happened() {
		if err := config.Decrypt(&data, []byte(*p)); err != nil {
			panic(err)
		}
		log.Printf("Decrypt process\n%v\n", string(data))
	} else {
		// In fact we will never hit this one
		// because commands and sub-commands are considered as required
		err := fmt.Errorf("bad arguments, please check usage")
		log.Print(parser.Usage(err))
	}

}

func getData(filename string, key string) ([]byte, error) {
	log.Printf("Read '%q'\n", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Read '%q' error: %v!", filename, err)
		return nil, err
	}

	if key == "" {
		log.Println("Return data")
		return data, nil
	}

	log.Println("Unmarshal yalm data")
	var m map[string]interface{}
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("Unmarshal %v\n", err)
		return nil, err
	}

	log.Printf("Check[%q]\n", key)
	if data, ok := m[key]; !ok {
		log.Fatalf("Key not found [%v]\n", key)
		return nil, err
	} else {
		return []byte(data.(string)), err
	}
}
