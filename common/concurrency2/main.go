// Wait groups with for loops
// Note: Compare between the two main()
// Concurrency in Go: Tools and Techniques for Developers, p42

package main

import (
	"fmt"
	"sync"
)


// Given a slice "hello", "greetings", "good day", start a goroutine for each of them
// Ensure that "good day doesn't appear thrice"
func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}

	wg.Wait()
}