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

type UserServer struct {
	IC			string
	Email		string
	Firstname	string
	Lastname	string
	Username	string
	Password	string
}

type UserClient struct {
	Username string
	Firstname    string
	Lastname     string
}

type Data struct {
	MyUser     UserClient
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