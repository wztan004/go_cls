package utils

import (
	"encoding/csv"
	"log"
	"os"
	"errors"
	"assignment4_cp3/datastruct"
	"assignment4_cp3/constants"
	"sync"
)

// WriteCSV receives a file path and paste the input into the given
// file path.
// Returns error if len(input) doesn't match CSV columns.
// Reference: https://asciinema.org/a/138540
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

// ReadFile returns the content of a file, given a file path.
func ReadFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return record
}

// readFileConcurrently is a helper function for ReadMultipleFilesConcurrently.
func readFileConcurrently(path string, ch chan <-[][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	ch <- record
}

// ReadMultipleFilesConcurrently reads the two venue files concurrently.
// Note that this function is not generalized and only applied to venue
// files.
func ReadMultipleFilesConcurrently() [][]string{
	var wg sync.WaitGroup
	wg.Add(2)
	firstchan := make(chan [][]string)
	secondchan := make(chan [][]string)
	go readFileConcurrently(constants.LatestMthLess1,firstchan, &wg)
	go readFileConcurrently(constants.LatestMth,secondchan, &wg)
	x1 := <- firstchan
	x2 := <- secondchan
	wg.Wait()
	x1 = append(x1, x2...)
	return x1
}

// GetUserCSV returns UserServer of the user. Returns error if the user does
// not exist.
func GetUserCSV(username string) (datastruct.UserServer, error) {
	// reading a CSV file
	file, err := os.Open(constants.UserFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	user := datastruct.UserServer{}
	for _, item := range record {
		if (item[4] == username) {
			user.IC = item[0]
			user.Email = item[1]
			user.Firstname = item[2]
			user.Lastname = item[3]
			user.Username = item[4]
			user.Password = item[5]
			return user, nil
		}
	}
	return user, errors.New("Cannot find user")
}

