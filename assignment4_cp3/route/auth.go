// Package route includes routing details, including authentication and key
// options (e.g. book venues) within the web app.
package route

import (
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	"assignment4_cp3/constants"
	"crypto/subtle"
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template = template.Must(template.ParseGlob("templates/*"))
var mapUsers = map[string]datastruct.UserClient{}
var mLinkedList datastruct.SessionLinkedList
var wlog *log.Logger // Be concerned
var elog *log.Logger // Error problem
var clog *log.Logger // Critical problem

// init (initializes) "admin" and "user" accounts, and logging objects.
func init() { 
	utils.InitializeUsers()

	file, err := os.OpenFile(constants.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	wlog = log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	elog = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	clog = log.New(io.MultiWriter(file, os.Stderr), "CRITICAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Index is the home page. If the user is logged in, it will show the user's
// current booked venues and main actions to take in the system.
func Index(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		http.Redirect(res, req, "/login", 302)
		return
	}

	var mData datastruct.VenueAvailability
	mData.MyUser = mUserClient

	venueCSVStruct := func(i []string) datastruct.Venue {
		var mVenue datastruct.Venue
		mVenue.Date = i[0]
		mVenue.Type = i[1]
		mVenue.Capacity = i[2]
		mVenue.BookedBy = i[3]
		mVenue.Username = i[4]
		return mVenue
	}

	// Reads multiple CSV venue files and update mData struct 
	updateVenueCSV := func() {
		for _, i := range utils.ReadMultipleFilesConcurrently() {
			if i[4] == "unbook" {
				mData.VenueUnbook = append(mData.VenueUnbook, venueCSVStruct(i))
			} else if i[4] == mUserClient.Username {
				mData.VenueUser = append(mData.VenueUser, venueCSVStruct(i))
			}
			mData.VenueAll = append(mData.VenueAll, venueCSVStruct(i))
		}
	}
	updateVenueCSV()

	if req.Method == http.MethodPost {
		date := strings.TrimSpace(req.FormValue("date"))
		venueType := strings.TrimSpace(req.FormValue("venueType"))
		capacity := strings.TrimSpace(req.FormValue("capacity"))

		// Goes through each CSV to see if the requested venue is in each CSV
		// Once found, update the venue info and break off the loop.
		for _, k := range []string{
			`confidential\venues_202009.csv`,
			`confidential\venues_202010.csv`,
		} {
			hasBooked, err := EditVenue("book", k, date, venueType, capacity, mUserClient)
			if err != nil {
				log.Fatalln(err)
			}
			if hasBooked == true {
				http.Redirect(res, req, "/", http.StatusSeeOther)
				return
			}
		}
		http.Error(res, "Check your input again. You can only enter available venues", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(res, "index.gohtml", mData)
}

// EditVenue allows the user to either add or unbook a venue. Adding a venue can
// only be applied to unbooked venues, removing a venue can only be applied to
// booked venues. It edits the last two fields of the CSV file: email and
// username. Adding a venue would set the two fields to user-specific info,
// while removing a venue would set them to "unbook".
//
// Action should be either "book" or "unbook".
// Path is the relative path to the CSV venue file.
// Date follows the format of "YYYYMMDD".
// Capacity is an integer but formatted in string.
// mUserClient is the user's info.
func EditVenue(action string, path string, date string, venueType string, capacity string, mUserClient datastruct.UserClient) (bool, error) {
	if action != "book" && action != "unbook" {
		return false, errors.New("Choose either 'book' or 'unbook'")
	}

	// 1. Read a CSV file
	res := utils.ReadFile(path)

	// 2. If record matches, update the array
	var returnRes [][]string

	toReturn := false
	for _, v := range res {
		if date == v[0] && venueType == v[1] && capacity == v[2] {
			if action == "book" {
				v[3] = mUserClient.Email
				v[4] = mUserClient.Username
			} else if action == "unbook" {
				v[3] = "unbook"
				v[4] = "unbook"
			}
			toReturn = true
			returnRes = append(returnRes, v)
		} else {
			returnRes = append(returnRes, v)
		}
	}

	// 3. Save the array into the same CSV
	if toReturn {
		utils.CreateNewBookingCSV(path, returnRes)
	}

	// 4. Return function
	return toReturn, nil
}

// Signup starts account registration, validates their fields, saves the
// profile and then redirect to the home page.
func Signup(res http.ResponseWriter, req *http.Request) {
	// 1. If logged in, redirect to Index().
	isLoggedIn, _ := alreadyLoggedIn(req)
	if isLoggedIn {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	// 2. Perform form validation with regular expressions.
	var errorString string
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		ic := req.FormValue("icnumber")
		email := req.FormValue("email")
		if _, err := utils.GetUserCSV(username); err == nil {
			errorString += "Username is already taken!"
		}
		if res, err := utils.VerifyIC(ic); res == false {
			errorString += err.Error()
		}
		if res, err := utils.VerifyEmail(email); res == false {
			errorString += err.Error()
		}
		if res, err := utils.VerifyPassword(password); res == false {
			errorString += err.Error()
		}
		if len(errorString) > 0 {
			tpl.ExecuteTemplate(res, "signup.gohtml", errorString)
			return
		}

		// 3. CSV User: Update
		bPassword := utils.CreateChecksum(password)
		utils.WriteCSV(`confidential\users.csv`, []string{ic, email, firstname, lastname, username, bPassword})
		startSession(res, req, username)
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

// Login allows user to authenticate. If successful, start a session and
// redirect to home page.
func Login(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, _ := alreadyLoggedIn(req)
	if isLoggedIn {
		http.Redirect(res, req, "/", 302)
		return
	}
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		ok := authenticateUser(username, password)
		if ok {
			startSession(res, req, username)
			http.Redirect(res, req, "/", 302)
			return
		}
		http.Redirect(res, req, "/login", 302)
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

// Logout removes the session in the server and expires the client cookie.
func Logout(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	mLinkedList.RemoveSession(mUserClient.Username)

	myCookie := &http.Cookie{
		Name:   "_cookie",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, myCookie)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}


// alreadyLoggedIn checks if user is part of active session list and returns
// the client's profile.
func alreadyLoggedIn(req *http.Request) (bool, datastruct.UserClient) {
	var userClient datastruct.UserClient
	myCookie, err := req.Cookie("_cookie")
	if err != nil {
		return false, userClient
	}

	res, userstring := mLinkedList.CheckSessionID(myCookie.Value)

	if res == true {
		userServer, err := utils.GetUserCSV(userstring)
		if err != nil {
			errors.New("Unable to get user")
		}
		userClient.Username = userServer.Username
		userClient.Firstname = userServer.Firstname
		userClient.Lastname = userServer.Lastname
		userClient.Email = userServer.Email
		return true, userClient
	}
	return false, userClient
}

// authenticateUser checks if the username and password is correct.
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

// startSession updates the user session in the server and sets client cookie
func startSession(res http.ResponseWriter, req *http.Request, username string) {
	mSession := utils.CreateSessionStruct(username)
	mLinkedList.RemoveSession(username)
	mLinkedList.EnqueueSession(mSession)

	// Create session cookie: client side
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    mSession.SessionUUID,
		HttpOnly: true,
	}

	http.SetCookie(res, &cookie)
}
