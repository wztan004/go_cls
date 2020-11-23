// switch and case

package main

import "fmt"

func main() {
	// What do you think is the output?

	fmt.Println("Switch 1")
	switch 1 {
		case 1:
			fmt.Println("20 > 40")
		case 2:
			fallthrough
		case 3:
			fmt.Println("1 > 10")
		default:
			fmt.Println("None 1")
	}
	
	fmt.Println("Switch 2")
	switch 2 {
		case 1:
			fmt.Println("1")
			fallthrough
		case 2:
			fmt.Println("2")
			fallthrough
		case 3:
			fmt.Println("3")
	}

	fmt.Println("Switch 3")
	switch 3 {
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
		default:
			fmt.Println("3")
	}
}