package config

import (
	"io/fs"
	"os"
	"reflect"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

func Load(ptr interface{}, filename string, createIfNotExists bool, key []byte, onlyFields bool) error {
	if err := defaults.Set(ptr); err != nil {
		return err
	}

	keyFile := key
	if onlyFields {
		keyFile = nil
	}

	out, err := read(ptr, filename, createIfNotExists, keyFile)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(out, ptr); err != nil {
		return err
	}

	if key != nil && onlyFields {
		val := reflect.Indirect(reflect.ValueOf(ptr))
		for i := 0; i < reflect.ValueOf(ptr).Elem().NumField(); i++ {

			structField := val.Type().Field(i)
			dataField := val.Field(i)

			// structField, found := reflect.TypeOf(ptr).Elem().FieldByName(name)
			// if !found {
			// 	continue
			// }
			if structField.Tag.Get("encrypted") == "true" {
				if dataField.Kind() == reflect.String {
					value := []byte(dataField.String())
					if err := Decrypt(&value, key); err != nil {
						return err
					}
					dataField.SetString(string(value))
				}

			}

		}

	}

	return nil
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

		if err := Encrypt(&out, key); err != nil {
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

	if err := Decrypt(&out, key); err != nil {
		return nil, err
	}

	return out, nil
}

func fileExists(filename string) bool {
	_, error := os.Stat(filename)
	return !os.IsNotExist(error)
}
