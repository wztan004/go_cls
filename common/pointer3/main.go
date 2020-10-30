// changing the pointer value affects the original variable
// go programming language ebook

package main

import (
	"fmt"
)

func main() {
	p := new(int) // p, of type *int, points to an unnamed int variable
	fmt.Println(p) // 0xc0000120a0
	fmt.Println(*p) // "0"
	*p = 2 // sets the unnamed int to 2
	fmt.Println(*p) // "2"
}
