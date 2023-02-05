### Simulating a Ransomware Program
---
This code was adapted and improved upon after watching William Moody's [Youtube](https://www.youtube.com/watch?v=9B3xas3McQU).


- Set up a 'home' target to encrypt the files therein  
`./setup`

##### **Running v0.1**
- Encrypt the files  
  `go run encryption.go`

- decrypt the files  
  `go run decryption.go`

##### **Running v0.2**
This version weaponises this as a cli tool
```shell
# Encrypt the files 
go run ransomware.go encrypt -path ./home -key usethiskey

# Decrypt the files
go run ransomware.go decrypt -path ./home -key usethiskey

# Compile the Program for Linux and Windows Os
GOOS=linux go build ransomware.go
GOOS=windows go build ransomware.go
```  

## Disclaimer
---
This program was developed only for informational purpose and **must** not be used illegally.

## Issues
Although this simulates encrypting and decrypting files in a given directory. The files are rendered unrecoverable if the user passes a wrong key. [This](https://crypto.stackexchange.com/questions/84355/can-aes-gcm-mode-detect-an-incorrect-key-and-refuse-to-decrypt) stackexchange resource somewhat explains how aes-gcm works.

## References
- https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
- https://www.programming-books.io/essential/go/encryption-and-decryption-with-aes-gcm-474ffe54eb92473b908b5ef162789cad
- https://go.dev/src/crypto/aes/aes_gcm.go
- https://go.dev/src/crypto/aes/aes_gcm.go