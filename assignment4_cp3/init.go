package main

import (
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	"assignment4_cp3/constants"
	"encoding/csv"
	"os"
)

// initializeUsers create "admin" and "user" accounts.
func initializeUsers() {
	bPassword := utils.CreateChecksum("password")

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
	csvFile, err := os.Create(constants.UserFile)
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

