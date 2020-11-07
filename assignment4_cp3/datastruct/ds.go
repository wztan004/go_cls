package datastruct

import (
	"time"
)

type Venue struct {
	Date     string
	Type     string
	Capacity string
	BookedBy string
	Username string
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
	Username	string
	Firstname	string
	Lastname	string
	Email		string
}

type Data struct {
	MyUser     UserClient
	VenueUser  []Venue
	VenueAll   []Venue
}

type UUID []byte
type Session struct {
	SessionUUID	string
	Username	string
	CreatedAt	time.Time
}