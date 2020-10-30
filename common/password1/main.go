// Storing passwords
// Go in Action 2 Day 2

package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {
	password := []byte("47;u5:B(95m72;Xq")

	// Hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	fmt.Println(hashedPassword)


	// Compares hashed password and byte password
	fmt.Println(bcrypt.CompareHashAndPassword(hashedPassword, []byte("47;u5:B(95m72;Xq")))
}