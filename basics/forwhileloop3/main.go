// for / while loop
// GoSchool: Basics

package main

import "fmt"

func main(){
	// Convert this Python code into Go.
	//
	// strings = ["hello", "world"]
	// for i, j in enumerate(strings):
	// 	print(i, j)

	strings := []string{"hello", "world"}
	for i, j := range strings {
		fmt.Println(i, j)
	}
}

