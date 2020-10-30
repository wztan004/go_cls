// RNG
// Go in Action 2 Day 2

package main

import (
	cr "crypto/rand"
	"fmt"
	mr "math/rand"
	big "math/big"
)
	

func main() {
	// pseudo, will always return 1825
	fmt.Println("Random Number: ", mr.Intn(1984))

	// a more secure example
	rand, err := cr.Int(cr.Reader, big.NewInt(1984))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Random Number: %d\n", rand)
}