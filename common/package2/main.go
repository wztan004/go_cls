// global variables

package main

import (
	ds "common/package2/datastruct"
	"fmt"
)

func main() {
	type A = ds.Item
	p := A{"3", 3}
	fmt.Println(p)
}