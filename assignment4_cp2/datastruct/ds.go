package datastruct

import (
	"time"
)

type Venue struct {
	ID       string
	Date     string
	Capacity string
	IsBooked string
}

type User struct {
	Username string
	Password []byte
	First    string
	Last     string
}

type Data struct {
	MyUser     User
	VenueNames map[string][]Venue
	VenueUser  []Venue
	VenueAll   []Venue
}

type UUID []byte
type Session struct {
	SessionUUID	string
	Username	string
	CreatedAt	time.Time
}