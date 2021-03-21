// if-else
// GoSchool: Basics

package main

import "fmt"

func main(){
	conditionA := 10
	conditionB := 12
	if conditionA < conditionB {
		fmt.Println("A smaller than B")
	} else if conditionA > conditionB {
		fmt.Println("A bigger than B")
	} else {
		fmt.Println("A is same as C")
	}
}