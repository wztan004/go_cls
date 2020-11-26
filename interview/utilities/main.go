package main
import "fmt"


/*
Hash maps / Dictionaries
	- Create a frequency table
	- Get key-value pair of each entry
*/


// // Create a frequency table
// func main() {
	
// 	nums := []int{1,1,1,2,2,3}
// 	Map := make(map[int]int)
// 	for _, e := range nums {
// 		Map[e]++
// 	}
// 	fmt.Println(Map)
// }


// // Iterate through a hash map
// func main() {
// 	import "fmt"
// 	Map := map[int]int {
// 		1: 3,
// 		2: 2,
// 		3: 1,
// 	}

// 	for i, j := range Map {
// 		fmt.Println(i, j)
// 	}
// }


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