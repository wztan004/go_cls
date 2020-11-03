package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"errors"
	// "golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	IC			string
	Email		string
	Firstname	string
	Lastname	string
	Username	string
	Password	[]byte
}

func main() {
	err := WriteCSV(`assignment4_cp2\sandbox\posts.csv`,[]string{"one","two","three","four","five","six"})
	fmt.Println(err)
}


// WriteCSV returns error if len(input) doesn't match csv columns
// https://asciinema.org/a/138540
// https://gobyexample.com/variadic-functions
func WriteCSV(path string, input []string) (error) {
	file, err := os.OpenFile(path, os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.Read()
	if err != nil {
		panic(err)
	}
	numCols := len(record)
	if numCols != len(input) {
		return errors.New("Columns doesn't match")
	}

	var data [][]string
	data = append(data, input)

	w := csv.NewWriter(file)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	return nil
}