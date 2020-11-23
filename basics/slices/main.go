// slices

package main

import "fmt"

func main(){
	slices := []string{"one","two","three"}
	fmt.Println(slices)

	slices = append(slices,"4")
	fmt.Println(slices)

	slices[3] = "four"
	fmt.Println(slices)
}
