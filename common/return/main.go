package main

import "fmt"

func main() {
	fmt.Println("A")

	if true {
		fmt.Println("B")
		return
	}

	fmt.Println("C")
}