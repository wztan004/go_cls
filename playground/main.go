package main

import "fmt"

func main() {
    c := make(chan int)    
    go func() {
        fmt.Println("received:", <-c)
    }()
	c <- 1
}