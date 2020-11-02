package ds

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