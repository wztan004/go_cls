//TODO compare checksum and log file

package main

import (
	"fmt"
	"crypto/sha256"
	"io/ioutil"
	"encoding/hex"
)

// func main () {
// 	// Get our known Log checksum from checksum file.
// 	logBytes, err := ioutil.ReadFile("errors.txt")
// 	if err !=nil {
// 		fmt.Println("Error")
// 	}

// 	logStr := string(logBytes) // convert content to a 'string'
// 	logHash32 := sha256.Sum256(logBytes) // // Compute our current log's SHA256 hash
// 	logHash := logHash32[:] // convert from 32bytes to byte https://stackoverflow.com/questions/28046949/convert-fixed-size-array-to-variable-sized-array-in-go
// 	hash := hex.EncodeToString(logHash)

// 	fmt.Println("logBytes", logBytes)
// 	fmt.Println("logStr", logStr)
// 	fmt.Println("logHash32", logHash32)
// 	fmt.Println("logHash", logHash)
// 	fmt.Println("hash", hash)

// 	// Compare our calculated hash with our stored hash
// 	if logStr == hash {
// 		// Ok the checksums match.
// 		fmt.Println("Log integrity OK.")
// 	} else {
// 		// The file integrity has been compromised...
// 		fmt.Println("File Tampering detected.")
// 	}
	
// }

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main () {
	// Get our known Log checksum from checksum file.
	logBytes, err := ioutil.ReadFile("errors.txt")
	if err !=nil {
		fmt.Println("Error")
	}
	logHash32 := sha256.Sum256(logBytes) // // Compute our current log's SHA256 hash

	fmt.Println(logHash32)

	logHash := logHash32[:]
	hashStr := hex.EncodeToString(logHash)

	fmt.Println(hashStr)

    err = ioutil.WriteFile("checksum.txt", []byte(hashStr), 0644)
    check(err)

}