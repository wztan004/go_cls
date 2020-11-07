// Package route the app to navigate around
package route

import (
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	"assignment4_cp3/constants"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var tpl *template.Template = template.Must(template.ParseGlob("templates/*"))
var mapUsers = map[string]datastruct.UserClient{}
var mapSessions = map[string]string{}
var mData datastruct.Data
var mLinkedList datastruct.LinkedList
var wlog *log.Logger // Be concerned
var elog *log.Logger // Critical problem

type data struct {
	MyUser        datastruct.UserClient
	VenueUser     []datastruct.Venue
	VenueUnbooked []datastruct.Venue
	VenueAll      []datastruct.Venue
}

func init() {
	utils.InitializeUsers()

	file, err := os.OpenFile(constants.LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	wlog = log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	elog = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Index k
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Start: Index()")
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		fmt.Println("Index(): User is not logged in")
		fmt.Println("Index() 1")
		http.Redirect(res, req, "/login", 302)
		return
	}

	fmt.Println("Index(): User is logged in")

	var mData data
	mData.MyUser = mUserClient
	fmt.Println("!", mUserClient)

	venueCSVStruct := func(i []string) datastruct.Venue {
		var mVenue datastruct.Venue
		mVenue.Date = i[0]
		mVenue.Type = i[1]
		mVenue.Capacity = i[2]
		mVenue.BookedBy = i[3]
		mVenue.Username = i[4]
		return mVenue
	}

	updateVenueCSV := func() {
		for _, i := range utils.ReadMultipleFilesConcurrently() {
			if i[4] == "not booked" {
				mData.VenueUnbooked = append(mData.VenueUnbooked, venueCSVStruct(i))
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
		// Breaks off the loop once it's found
		for _, k := range []string{
			`confidential\venues_202009.csv`,
			`confidential\venues_202010.csv`,
		} {
			hasBooked, err := EditVenue("book", k, date, venueType, capacity, mUserClient)
			fmt.Println("Index11", hasBooked, err)
			if err != nil {
				log.Fatalln(err)
			}
			if hasBooked == true {
				fmt.Println("Index121 breaking out")
				http.Redirect(res, req, "/", http.StatusSeeOther)
				return
			}
			fmt.Println("this shouldn't appear after 121")
		}

		http.Error(res, "Check your input again. You can only enter available venues", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(res, "index.gohtml", mData)
}

func EditVenue(action string, path string, date string, venueType string, capacity string, mUserClient datastruct.UserClient) (bool, error) {
	if action != "book" && action != "remove" {
		return false, errors.New("Choose either 'book' or 'remove'")
	}

	// 1. Read each CSV
	// TODO Variadic
	res := utils.ReadFile(path)

	// 2. If record matches, update the array
	var returnRes [][]string

	toReturn := false
	for _, v := range res {
		if date == v[0] && venueType == v[1] && capacity == v[2] {
			if action == "book" {
				v[3] = mUserClient.Email
				v[4] = mUserClient.Username
			} else if action == "remove" {
				v[3] = "not booked"
				v[4] = "not booked"
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

// Signup allows users to sign up (5Nov20 refactored)
func Signup(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Signup")

	// 1. If logged in, move to Index(). If not logged in, stay.
	isLoggedIn, _ := alreadyLoggedIn(req)
	if isLoggedIn {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	// 2. Perform form validation
	var errorString string

	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		ic := req.FormValue("icnumber")
		email := req.FormValue("email")

		if _, ok := mapUsers[username]; ok {
			errorString += "Username already taken"
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
			fmt.Println("navigate > Signup > Error signing up")
			tpl.ExecuteTemplate(res, "signup.gohtml", errorString)
			return
		}

		// 3. CSV User: Update
		bPassword := utils.CreateChecksum(password)
		utils.WriteCSV(`confidential\users.csv`, []string{ic, email, firstname, lastname, username, bPassword})

		startSession(res, req, username)

		fmt.Println("navigate > Signup > No error signing up")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

// Login : allows log in
func Login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Login()", time.Now())

	isLoggedIn, _ := alreadyLoggedIn(req)
	if isLoggedIn {
		fmt.Println("Login() 1")
		http.Redirect(res, req, "/", 302)
		return
	}

	fmt.Println("Login() 2")

	// process form submission
	if req.Method == http.MethodPost {
		fmt.Println("Login() -> Submit")
		username := req.FormValue("username")
		password := req.FormValue("password")
		ok := authenticateUser1(username, password)

		if ok {
			startSession(res, req, username)

			fmt.Println("Login() -> Submit -> Set cookie")
			fmt.Println("Login() -> Submit -> Redirect")
			http.Redirect(res, req, "/", 302)
			fmt.Println("End: Login()")
			return
		}
		fmt.Println("Login() -> Failed submission -> Redirect")
		http.Redirect(res, req, "/login", 302)
		fmt.Println("End: Login()")
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

// Logout : logs the user out of the app
func Logout(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Logout()")
	isLoggedIn, _ := alreadyLoggedIn(req)
	if !isLoggedIn {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	// TODO delete the session

	myCookie := &http.Cookie{
		Name:   "_cookie",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, myCookie)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}


// alreadyLoggedIn checks if user is part of active user list
func alreadyLoggedIn(req *http.Request) (bool, datastruct.UserClient) {
	fmt.Println("Start: alreadyLoggedIn()", time.Now())

	var userClient datastruct.UserClient
	myCookie, err := req.Cookie("_cookie")
	if err != nil {
		fmt.Println("alreadyLoggedIn(): Not logged in (client cookie not found)")
		fmt.Println("End: alreadyLoggedIn()")
		return false, userClient
	}

	res, userstring := mLinkedList.CheckSessionID(myCookie.Value)
	fmt.Println("!myCookie.Value", myCookie.Value)
	fmt.Println("!userstring", userstring)

	if res == true {
		fmt.Println("alreadyLoggedIn(): Logged in")
		userServer, err := utils.GetUserCSV(userstring)
		if err != nil {
			errors.New("Unable to get user")
		}
		userClient.Username = userServer.Username
		userClient.Firstname = userServer.Firstname
		userClient.Lastname = userServer.Lastname
		userClient.Email = userServer.Email
		fmt.Println("!userClient", userClient)
		return true, userClient
	}
	fmt.Println("alreadyLoggedIn(): Not logged in (session cookie not found)")
	return false, userClient
}

// authenticateUser1
func authenticateUser1(username string, password string) bool {
	fmt.Println("Start: authenticateUser1()")
	user, err := utils.GetUserCSV(username)
	if err != nil {
		// user does not exist
		fmt.Println("End: authenticateUser1() -> false")
		return false
	}
	bPassword := utils.CreateChecksum(password)
	if string(user.Password) == bPassword {
		fmt.Println("End: authenticateUser1() -> true")
		return true
	}
	fmt.Println("End: authenticateUser1() -> false")
	return false
}

func startSession(res http.ResponseWriter, req *http.Request, username string) {
	mSession := utils.CreateSessionStruct(username)
	mLinkedList.Remove(username)
	mLinkedList.EnqueueSession(mSession)

	// Creating session cookie: client side
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    mSession.SessionUUID,
		HttpOnly: true,
	}

	http.SetCookie(res, &cookie)
}
