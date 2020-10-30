// changing the pointer value affects the original variable
// go programming language ebook

package main

import (
	"fmt"
)

func main() {
	v := 1
	incr(&v) // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}
	