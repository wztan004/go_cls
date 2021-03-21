// https://leetcode.com/problems/top-k-frequent-elements/

package main

import "fmt"

func main() {
	k := 2
	res := []int{}
	Map := map[int]int {
		1: 3,
		2: 2,
		3: 1,
	}

	for ; k > 0; k-- {
		ress := 0 // to be updated with the most frequent key
		index := -1 // -1 is just a placeholder
		for i, e := range Map {
			if index < e {
				index = e
				ress = i
			}
			fmt.Println("q",i, index, ress)
		}
		delete(Map, ress)
		res = append(res, ress) // most frequent number of will be appended here, k times.
	}
	fmt.Println(res)
}