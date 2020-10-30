// changing the pointer value affects the original variable
// go programming language ebook

package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x // p, of type *int, points to x
	fmt.Println(p) // "0xc0000120a0"
	fmt.Println(*p) // "1"
	*p = 2 // equivalent to x = 2
	fmt.Println(x) // "2"
}