// validation
package regex

import (
	"fmt"
	"strings"
	"regexp"
	"unicode"
	"errors"
	"log"
)

var (
	Warning *log.Logger // Be concerned
)

// Verify IC format
func VerifyIC(input string) (bool, error) {
	strTrimmed := strings.TrimSpace(input)
	strChangeCase := strings.ToUpper(strTrimmed)
	res, err := regexp.MatchString(`^[STFG]\d{7}[A-Z]$`, strChangeCase)
	if err != nil {
		Warning.Println("Something went wrong in VerifyIC function")
	}
	if res == false {
		return false, errors.New("Make sure that your IC is entered correctly")
	}
	fmt.Println("1", err)
	return res, nil
}


func VerifyEmail(input string) (bool, error) {
	strTrimmed := strings.TrimSpace(input)
	strChangeCase := strings.ToLower(strTrimmed)
	res, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, strChangeCase)

	if err != nil {
		Warning.Println("Something went wrong in VerifyEmail function")
	}
	if res == false {
		return false, errors.New("Make sure that your email is entered correctly")
	}
	return res, nil
	
}


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

	finalErr = ""
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