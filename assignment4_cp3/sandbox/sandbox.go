package main

import (
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	// "encoding/csv"
	// "encoding/hex"
	// "errors"
	"fmt"
	// "log"
	// "net/http"
	// "os"
	// "time"
)

type data struct {
	MyUser datastruct.UserClient
	VenueUser []datastruct.Venue
	VenueUnbooked []datastruct.Venue
	VenueAll []datastruct.Venue
}

func main() {
	res := utils.ReadFile(`assignment4_cp3\confidential\venues_202009.csv`)

	var returnRes [][]string

	date := "20200901"
	venueType := "Bar"
	capacity := "40"
	email := "email"
	username := "username"

	toReturn := false
	for _, v := range res {
		if (date == v[0] && venueType == v[1] && capacity == v[2]) {
			v[3] = email
			v[4] = username
			toReturn = true
			returnRes = append(returnRes, v)
		} else {
			returnRes = append(returnRes, v)
		}
	}

	fmt.Println(toReturn)
	fmt.Println(returnRes)
}



// Progress

// AuthenticateUser (ON HOLD, TO RESOLVE A SUB FUNCTION)
// Create New Session (DONE)
// Putting New Session To Linked List (DONE)
// COMPARE TIME (DONE)
// ADDING USER TO CURRENT SESSION (IN PROGRESS)
