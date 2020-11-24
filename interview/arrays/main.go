package main

import (
    "fmt"
)

func RemoveIndex(s []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}


func main() {
    all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("all: ", all) //[0 1 2 3 4 5 6 7 8 9]
	
	removeIndex := RemoveIndex(all, 91)
	fmt.Println("removeIndex: ", removeIndex) //[0 1 2 3 4 6 7 8 9]
	fmt.Println("all: ", all) //[0 1 2 3 4 6 7 8 9]	
}