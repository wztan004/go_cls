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
	res := utils.ReadFile("assignment4_cp3/confidential/venues_202010.csv")
	fmt.Println(res)

}