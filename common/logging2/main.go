// custom loggers
// Go in Action 2 Day 2

package main

import (
	"io"
	"log"
	"os"
)

var (
	wlog *log.Logger // Be concerned
	elog *log.Logger // Error problem
	clog *log.Logger // Critical problem
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	wlog = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	clog = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	wlog.Println("There is something you need to know about")
	clog.Println("Something has failed")
}