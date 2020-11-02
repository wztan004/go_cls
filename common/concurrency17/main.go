// create a goroutine that clearly owns a channel, and a consumer that clearly 
// handles blocking and closing of a channel
// Concurrency in Go: Tools and Techniques for Developers

package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}