// Package datastruct is used to describe the main structs and functions
// used across the projects.
package datastruct

import (
	"time"
)

// A Venue is used when listing out venue booking/unbooking process.
// Usually used with VenueAvailability.
type Venue struct {
	Date     string
	Type     string
	Capacity string
	BookedBy string
	Username string
}

// A VenueAvailability is used as a list when listing out venue 
// booking or unbooking process.
// Usually used with Venue.
type VenueAvailability struct {
	MyUser        UserClient
	VenueUser     []Venue
	VenueUnbook   []Venue
	VenueAll      []Venue
}

// UserServer is used in signing up account and authentication.
type UserServer struct {
	IC			string
	Email		string
	Firstname	string
	Lastname	string
	Username	string
	Password	string
}

// UserClient is used to display a user's profile after logging-in.
type UserClient struct {
	Username	string
	Firstname	string
	Lastname	string
	Email		string
}

// A Session is used in client browser cookie. Includes a timestamp of when a
// user last visited a page.
type Session struct {
	SessionUUID	string
	Username	string
	CreatedAt	time.Time
}

