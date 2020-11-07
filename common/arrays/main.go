package main

import "fmt"

func main() {
	a := [][]string{{"a", "b"}, {"c", "d"}}
	b := [][]string{{"1", "2"}, {"3", "4"}}
	a = append(a, b...)
	fmt.Println(a)
}
