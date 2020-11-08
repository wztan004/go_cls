// Package utils offers shared functions that could potentially be used
// across multiple packages. 
package utils

import (
	"fmt"
	"golang.org/x/crypto/blake2b"
)

// CreateChecksum returns a checksum based on a string using BLAKE hash
// function. Reference: https://godoc.org/golang.org/x/crypto/blake2b
func CreateChecksum(plain string) (checksum string) {
	b := []byte(plain)
	h := blake2b.Sum256(b)
	return fmt.Sprintf("%x", h)
}