package utils

import (
	"encoding/csv"
	// "fmt"
	"log"
	"os"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type User struct {
	IC			string
	Email		string
	Firstname	string
	Lastname	string
	Username	string
	Password	[]byte
}

func InitializeUsers() {
	bPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)

	allUsers := []User{
		User{
			IC: "S1111111A", 
			Email: "admin@admin.com",
			Firstname: "Firstname",
			Lastname: "Lastname",
			Username: "admin", 
			Password: bPassword,
		},
		User{
			IC: "S1111111B", 
			Email: "user@user.com",
			Firstname: "Firstname",
			Lastname: "Lastname",
			Username: "user", 
			Password: bPassword,
		},
	}
	
	// creating a CSV file
	csvFile, err := os.Create(`security/users.csv`)
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

func ReadUserCSV() [][]string {
	file, err := os.Open(`security/users.csv`)
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


func GetUser(username string) (User, error) {
	// reading a CSV file
	file, err := os.Open(`security/users.csv`)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	user := User{}

	for _, item := range record {
		if (item[4] == username) {
			user.IC = item[0]
			user.Email = item[1]
			user.Firstname = item[2]
			user.Lastname = item[3]
			user.Username = item[4]
			user.Password = []byte(item[5])
			return user, nil
		}
	}
	return user, errors.New("Cannot find user")
}
