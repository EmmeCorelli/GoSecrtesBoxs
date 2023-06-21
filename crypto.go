package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/fs"
	"os"
)

func EncryptFile(filename string, key []byte) error {
	out, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := encrypt(&out, key); err != nil {
		return err
	}

	if err := os.WriteFile(filename, out, fs.ModePerm); err != nil {
		return err
	}
	return nil
}

func encrypt(plain *[]byte, key []byte) error {
	if key == nil {
		return nil
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}
	*plain = gcm.Seal(nil, nonce, *plain, nil)
	return nil
}

func DecryptFile(filename string, key []byte) error {
	out, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := decrypt(&out, key); err != nil {
		return err
	}

	if err := os.WriteFile(filename, out, fs.ModePerm); err != nil {
		return err
	}
	return nil
}

func decrypt(encrypted *[]byte, key []byte) error {
	if key == nil {
		return nil
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}
	chiper := *encrypted
	nonce, chiper := chiper[:gcm.NonceSize()], chiper[gcm.NonceSize():]

	if *encrypted, err = gcm.Open(nil, nonce, chiper, nil); err != nil {
		return err
	}
	return nil
}
