package config

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

func Load(ptr interface{}, filename string, createIfNotExists bool) error {
	if err := defaults.Set(ptr); err != nil {
		return err
	}

	file, err := open(ptr, filename, createIfNotExists)
	if err != nil {
		return err
	}
	defer file.Close()

	return yaml.NewDecoder(file).Decode(ptr)
}

func open(ptr interface{}, filename string, createIfNotExists bool) (*os.File, error) {

	if !fileExists(filename) && createIfNotExists {

		if filename == "" {
			filename = "./config.yaml"
		}

		file, err := os.Create(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		if err := yaml.NewEncoder(file).Encode(ptr); err != nil {
			return nil, err
		}
	}

	return os.Open(filename)
}

// function to check if file exists
func fileExists(filename string) bool {
	_, error := os.Stat(filename)
	return !os.IsNotExist(error)
}
