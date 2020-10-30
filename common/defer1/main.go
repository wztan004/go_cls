// defers in inner functions will execute finish first
// source: myself


package main

import (
	"fmt"
)

func main() {
	callB()
	fmt.Println("A")
}

func callB() {
	defer fmt.Println("B")
}