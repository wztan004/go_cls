//how to access and write a log file without overriding
//https://stackoverflow.com/questions/19965795/how-to-write-log-to-file

package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("testlogfile.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	
	log.SetOutput(f)
	log.Println("This is a test log entry43242")
}
