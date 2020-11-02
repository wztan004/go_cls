// The receiving form of the <- operator can also optionally return two values
// Concurrency in Go: Tools and Techniques for Developers

package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutation)
}