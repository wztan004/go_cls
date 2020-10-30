// log.Fatal
// Go in Action 2 Day 2

package main

import (
	"fmt"
	"log"
)

func initFunc(i int) {
	if i < 2 {
		fmt.Printf("Var %d - initialized\n", i)
	} else {
		log.Fatal("Init failure - Terminating.")
	}
}

func main() {
	i := 1
	for i < 3 {
		initFunc(i)
		i++
	}
	fmt.Println("Initialized all variables successfully")
}