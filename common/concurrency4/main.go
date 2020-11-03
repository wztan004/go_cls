// Add multiple goroutines at once
// Concurrency in Go: Tools and Techniques for Developers, p48

package main

import (
	"fmt"
	"sync"
)

// Run a goroutine hello() 5 times.
// hello() is an inner function and runs fmt.Printf("Hello from %v!\n", id)
// Call wg.Add() once only, instead of once per goroutine


func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}