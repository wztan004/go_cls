// Wait groups with two goroutines
// Concurrency in Go: Tools and Techniques for Developers, p47

package main

import (
	"fmt"
	"sync"
	"time"
)

// Run two goroutines.
// The first goroutine should print "1st goroutine sleeping..."
// The second goroutine should print "2nd goroutine sleeping..."
// The execution order will be done through time.Sleep()
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("finis")

}