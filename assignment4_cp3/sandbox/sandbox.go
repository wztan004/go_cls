package main

import (
	// "assignment4_cp3/datastruct"
	// "assignment4_cp3/utils"
	"crypto/subtle"
	// "encoding/csv"
	// "encoding/hex"
	// "errors"
	"fmt"
	// "log"
	// "net/http"
	// "os"
	// "sync"
	// "time"
)

func main() {
	x := []byte("123")
	y := []byte("123")
	fmt.Println(subtle.ConstantTimeCompare(x,y))
}

// func ReadFileConcurrently(path string, ch chan <-[][]string) {
// 	defer wg.Done()
// 	file, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	record, err := reader.ReadAll()
// 	if err != nil {
// 		panic(err)
// 	}
// 	ch <- record
// }
