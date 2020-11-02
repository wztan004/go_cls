// automatically break the loop when a channel is closed
// Concurrency in Go: Tools and Techniques for Developers

package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}