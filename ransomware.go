package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func checkError(e error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error Recovery:", r)
		}
	}()
	if e != nil {
		panic(e)
	}
}

func generateKey(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(path, key string) {
	// initialize AES in GCM mode
	block, err := aes.NewCipher([]byte(generateKey(key)))
	checkError(err)
	gcm, err := cipher.NewGCM(block)
	checkError(err)

	// looping through the target files
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Println("Encrypting " + path + "...") // encrypt the file

			// read file contents
			original, err := os.ReadFile(path)
			checkError(err)

			// encrypt bytes
			nonce := make([]byte, gcm.NonceSize())
			_, err = io.ReadFull(rand.Reader, nonce)
			checkError(err)
			encrypted := gcm.Seal(nonce, nonce, original, nil)

			// write the encrypted contents to file with a .enc extension
			err = os.WriteFile(path + ".enc", encrypted, 0666)
			checkError(err)
			os.Remove(path) // delete the original file
		}
		return nil
	})
}

func decrypt(path, key string) {
	// Initialize AES in GCM mode
	block, err := aes.NewCipher([]byte(generateKey(key)))
	checkError(err)
	gcm, err := cipher.NewGCM(block)
	checkError(err)

	// looping through target files
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		// skip if directory
		if !info.IsDir() && filepath.Ext(path) == ".enc" {
			fmt.Println("Decrypting " + path + "...")

			// read file contents
			encrypted, err := os.ReadFile(path)
			checkError(err)
			if len(encrypted) < gcm.NonceSize() {
				panic("Invalid Data")
			}

			// Decrypt bytes
			nonce := encrypted[:gcm.NonceSize()]
			encrypted = encrypted[gcm.NonceSize():]

			original, err := gcm.Open(nil, nonce, encrypted, nil)
			checkError(err)

			// write decrypted contents
			filename := strings.TrimSuffix(path, ".enc")
			err = os.WriteFile(filename, original, 0666)
			checkError(err)
			os.Remove(path) // delete the encrypted file
		}
		return nil
	})
}

func main() {
	encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	encryptPath := encryptCmd.String("path", "", "target directory to encrypt")
	encryptKey := encryptCmd.String("key", "", "key for encryption")

	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decryptPath := decryptCmd.String("path", "", "target directory to decrypt")
	decryptKey := decryptCmd.String("key", "", "key for decryption")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encrypt":
		encryptCmd.Parse(os.Args[2:])
		encrypt(*encryptPath, *encryptKey)

	case "decrypt":
		decryptCmd.Parse(os.Args[2:])
		decrypt(*decryptPath, *decryptKey)

	default:
		fmt.Println("Expected 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}
}
