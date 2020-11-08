package utils

import (
	"fmt"
	"strings"
	"regexp"
	"unicode"
	"errors"
	"log"
)

var (
	w *log.Logger
)

// VerifyIC ensures that the string fulfills these conditions:
//
// 1. The first letter must be S, T, F, or G.
//
// 2. After the first letter, 7 digits must follow.
//
// 3. The element must be an alphabet.
//
// Whether the alphabet is upper case or lower case does not matter.
func VerifyIC(input string) (bool, error) {
	strTrimmed := strings.TrimSpace(input)
	strChangeCase := strings.ToUpper(strTrimmed)
	res, err := regexp.MatchString(`^[STFG]\d{7}[A-Z]$`, strChangeCase)
	if err != nil {
		w.Println("Something went wrong in VerifyIC function")
	}
	if res == false {
		return false, errors.New("Make sure that your IC is entered correctly")
	}
	return res, nil
}

// VerifyEmail ensures that the string fulfills these conditions:
//
// 1. There must be either an alphabet, period, or dash, followed by a single @.
//
// 2. The email domain must be between 2 to 4 letters.
func VerifyEmail(input string) (bool, error) {
	strTrimmed := strings.TrimSpace(input)
	strChangeCase := strings.ToLower(strTrimmed)
	res, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, strChangeCase)

	if err != nil {
		w.Println("Something went wrong in VerifyEmail function")
	}
	if res == false {
		return false, errors.New("Make sure that your email is entered correctly")
	}
	return res, nil
	
}

// VerifyPassword ensures that the string fulfills these conditions:
//
// There must be at least one lower case letter, upper case letter, a
// number, a special character, and has at least 8 minimum and 64 maximum
// characters.
func VerifyPassword(password string) (bool, error) {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var finalErr string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}

	appendError := func (checkErr string) {
		if len(finalErr) > 0 {
			finalErr = finalErr[:len(finalErr)-2]
			finalErr += ", " + checkErr + ". "
		} else {
			finalErr += "Include at least one " + checkErr + ". "
		}
	}

	if !lowercasePresent {
		appendError("lowercase character")
	}
	if !uppercasePresent {
		appendError("uppercase character")
	}
	if !numberPresent {
		appendError("one number")
	}
	if !specialCharPresent {
		appendError("one special character")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		finalErr += fmt.Sprintf("Password length must be between %d to %d characters long.", minPassLength, maxPassLength)
	}

	if len(finalErr) != 0 {
		finalErr = "Password error: " + finalErr
		return false, errors.New(finalErr)
	}
	return true, nil
}