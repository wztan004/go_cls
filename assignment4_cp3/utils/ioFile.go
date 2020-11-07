package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"errors"
	"assignment4_cp3/datastruct"
)



func InitializeUsers() {
	bPassword := CreateChecksum("password")

	allUsers := []datastruct.UserServer{
		datastruct.UserServer{
			IC: "S1111111A", 
			Email: "admin@admin.com",
			Firstname: "Firstname",
			Lastname: "Lastname",
			Username: "admin", 
			Password: bPassword,
		},
		datastruct.UserServer{
			IC: "S1111111B", 
			Email: "user@user.com",
			Firstname: "Firstname",
			Lastname: "Lastname",
			Username: "user", 
			Password: bPassword,
		},
	}
	
	// creating a CSV file
	csvFile, err := os.Create(`confidential/users.csv`)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for _, user := range allUsers {
		line := []string{user.IC, user.Email, user.Firstname, user.Lastname, user.Username, string(user.Password)}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	fmt.Println("users initialized")
}



func CreateNewBookingCSV(path string, res [][]string) {
	csvFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for _, user := range res {
		line := []string{user[0], user[1], user[2], user[3], user[4]}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}

func CreateNewUserCSV(path string, res [][]string) {
	csvFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for _, user := range res {
		line := []string{user[0], user[1], user[2], user[3], user[4], user[5]}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
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

func GetUserCSV(username string) (datastruct.UserServer, error) {
	// reading a CSV file
	file, err := os.Open(`confidential/users.csv`)
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
