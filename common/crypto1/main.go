// AES encrypt and decrypt
// Note1: The cryptographic keys should also not be hardcoded in the source code (as it is in the example).
// Note2: It is recommended to use more modern approach than AES, such as chacha20poly1305 or NaCl 
// Go in Action 2 Day 2

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)
func main() {
	key := []byte("Encryption Key should be 32 char")
	data := []byte("Welcome to Go in Action 2")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	encryptedData := aesgcm.Seal(nil, nonce, data, nil)
	fmt.Printf("Encrypted: %x\n", encryptedData)
	decryptedData, err := aesgcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Decrypted: %s\n", decryptedData)
}