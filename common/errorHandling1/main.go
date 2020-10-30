// error messages
// https://gobyexample.com/errors

package main

import (
	"fmt"
	"errors"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg, nil
}

func main() {
	if result, err := f1(42); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("No error. Result is", result)
	}
}