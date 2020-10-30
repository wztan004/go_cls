// maps and structs
// https://tour.golang.org/moretypes/2

package main

import "fmt"

type venue struct {
	ID string
	Date string
	Capacity string
	IsBooked string
}

func main() {
	mVenue := venue{"mID", "mDate", "mCapacity", "mBooked"}
	mVenue1 := venue{"mID1", "mDate1", "mCapacity1", "mBooked1"}
	
	m := make(map[string][]venue) // create an empty map instance
	s := []venue{} // create an empty slicee
	s = append(s,mVenue) // append a map instance to a slice of a map
	s = append(s,mVenue1)
	
	m["ee"] = s
	m["ff"] = s
	fmt.Println(m)
	
}
