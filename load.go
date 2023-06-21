package config

import (
	"io/fs"
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

func Load(ptr interface{}, filename string, createIfNotExists bool, key []byte) error {
	if err := defaults.Set(ptr); err != nil {
		return err
	}

	out, err := read(ptr, filename, createIfNotExists, key)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(out, ptr)
}

func read(ptr interface{}, filename string, createIfNotExists bool, key []byte) ([]byte, error) {

	if filename == "" {
		filename = "./config.yaml"
	}

	if !fileExists(filename) && createIfNotExists {
		out, err := yaml.Marshal(ptr)
		if err != nil {
			return nil, err
		}

		if err := encrypt(&out, key); err != nil {
			return nil, err
		}

		if err := os.WriteFile(filename, out, fs.ModePerm); err != nil {
			return nil, err
		}
	}

	out, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := decrypt(&out, key); err != nil {
		return nil, err
	}

	return out, nil
}

func fileExists(filename string) bool {
	_, error := os.Stat(filename)
	return !os.IsNotExist(error)
}
