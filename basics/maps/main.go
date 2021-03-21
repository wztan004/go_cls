// maps
// think of it as a hash table

package main

import "fmt"


func main() {
	// instantiate
	fmt.Println("===Instantiate===")
	m := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	// delete a key
	fmt.Println("===Delete===")
	delete(m, "two")

	// test existence of a key
	fmt.Println("===Test existence===")
	i, ok := m["one"]
	fmt.Println(i, ok)

	i, ok = m["two"]
	fmt.Println(i, ok)

	// iterate
	fmt.Println("===Iterate===")
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
