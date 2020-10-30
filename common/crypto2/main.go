// recommend to use blake
// Go in Action 2 Day 2

package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"golang.org/x/crypto/blake2s"
)

func main() {
	hMd5 := md5.New()
	hSha := sha256.New()
	hBlake2s, _ := blake2s.New256(nil)
	io.WriteString(hMd5, "Welcome to Go Language Secure Coding Practices")
	io.WriteString(hSha, "Welcome to Go Language Secure Coding Practices")
	io.WriteString(hBlake2s, "Welcome to Go Language Secure Coding Practices")
	fmt.Printf("MD5 : %x\n", hMd5.Sum(nil))
	fmt.Printf("SHA256 : %x\n", hSha.Sum(nil))
	fmt.Printf("Blake2s-256: %x\n", hBlake2s.Sum(nil))
}
