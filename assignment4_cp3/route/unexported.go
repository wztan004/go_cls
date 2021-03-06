// Package route is For internal functions
package route

import (
	"assignment4_cp3/constants"
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	"crypto/subtle"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"strings"
	"time"
)


// authenticateUser returns true if the username and password is correct.
func authenticateUser(username string, password string) bool {
	user, err := utils.GetUserCSV(username)
	if err != nil {
		return false
	}
	bPassword := utils.CreateChecksum(password)

	x := []byte(user.Password)
	y := []byte(bPassword)
	result := subtle.ConstantTimeCompare(x,y)

	if result == 1 {
		return true
	}
	// This log should not shown to the user.
	wlog.Println("Log in failed with wrong password. Username:", username)
	return false
}

// createNewUserCSV creates a new CSV that stores user data, given a 2D array.
func createNewUserCSV(path string, res [][]string) {
	csvFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for _, user := range res {
		line := []string{user[0], user[1], user[2], user[3], user[4], user[5]}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}



// toIndexPage returns to the Index page
func toIndexPage(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/", http.StatusSeeOther)
	return
}

// alreadyLoggedIn checks if user is part of active session list and returns
// the client's profile.
func alreadyLoggedIn(req *http.Request) (bool, datastruct.UserClient) {
	var userClient datastruct.UserClient
	myCookie, err := req.Cookie(constants.CookieName)
	if err != nil {
		return false, userClient
	}

	res, userstring := mLinkedList.CheckSessionID(myCookie.Value)

	if res == true {
		userServer, err := utils.GetUserCSV(userstring)
		if err != nil {
			wlog.Println("Unable to get user")
		}
		userClient.Username = userServer.Username
		userClient.Firstname = userServer.Firstname
		userClient.Lastname = userServer.Lastname
		userClient.Email = userServer.Email
		return true, userClient
	}
	return false, userClient
}



// startSession updates the user session in the server and sets client cookie
func startSession(res http.ResponseWriter, req *http.Request, username string) {
	mSession := createSessionStruct(username)
	mLinkedList.RemoveSession(username)
	mLinkedList.EnqueueSession(mSession)

	// Create session cookie: client side
	cookie := http.Cookie{
		Name:     constants.CookieName,
		Value:    mSession.SessionUUID,
		HttpOnly: true,
	}

	http.SetCookie(res, &cookie)
}

// createUUID creates a V4 UUID. Reference: https://github.com/satori/go.uuid
func createUUID() (id string) {
	mUUID, err := uuid.NewV4()
	if err != nil {
		wlog.Printf("Something went wrong: %s", err)
		return
	}
	return fmt.Sprintf("%s", mUUID)
}

// createSessionStruct returns a new Session.
func createSessionStruct(username string) (datastruct.Session) {
	mUUID := createUUID()
	mSession := datastruct.Session{
		SessionUUID:	mUUID, 
		Username:		username, 
		CreatedAt: 		time.Now(),
	}
	return mSession
}

// createNewUserCSV creates a new CSV that stores user data, given a 2D array.
func createNewBookingCSV(path string, res [][]string) error {
	if !strings.Contains(path, constants.VenueRegex) {
		return errors.New("Check if the file path is correct")
	}
	csvFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for _, venue := range res {
		if len(venue) != 5 {
			return errors.New("Check if the file path is correct")
		}
		line := []string{venue[0], venue[1], venue[2], venue[3], venue[4]}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	return nil
}


// editNames allows the user to change first name and last name.
func editNames(username string, firstname string, lastname string) {
	var returnRes [][]string
	result := utils.ReadFile(constants.UserFile)
	for _, v := range result {
		if (username == v[4]) {
			v[2] = firstname
			v[3] = lastname
			returnRes = append(returnRes, v)
		} else {
			returnRes = append(returnRes, v)
		}
	}
	createNewUserCSV(constants.UserFile, returnRes)
}

// updateTemplateData updates the struct that's passed to the HTML template.
// The struct is filled with venue data (booked, unbooked, all)
func updateTemplateData(mData *datastruct.VenueAvailability, mUserClient *datastruct.UserClient) {
	venueCSVStruct := func(i []string) datastruct.Venue {
		var mVenue datastruct.Venue
		mVenue.Date = i[0]
		mVenue.Type = i[1]
		mVenue.Capacity = i[2]
		mVenue.BookedBy = i[3]
		mVenue.Username = i[4]
		return mVenue
	}

	for _, i := range utils.ReadMultipleFilesConcurrently() {
		if i[4] == "unbook" {
			mData.VenueUnbook = append(mData.VenueUnbook, venueCSVStruct(i))
		} else if i[4] == mUserClient.Username {
			mData.VenueUser = append(mData.VenueUser, venueCSVStruct(i))
		}
		mData.VenueAll = append(mData.VenueAll, venueCSVStruct(i))
	}
}
