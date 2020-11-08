// Used for testing purposes. Ignore this package/file.
package main

import (
	"fmt"
)

func main() {
	heq("1")
}


func heq(qq string) {
	if qq == "2" {
		fmt.Println("a")
		return
	}
	if qq == "1" {
		fmt.Println("b")
		return
	}
	fmt.Println("c")
}