// struct

package main

import "fmt"

type customer struct {
	name string
	age int
}

func main() {
	c := customer{"wz", 30}
	fmt.Println(c)
	fmt.Println(c.age)
}