// Wait groups
// Concurrency in Go: Tools and Techniques for Developers, p41

package main

import (
	"fmt"
	"sync"
)

// create variables wg, salutation
// update salutation (hello -> welcome)
func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}