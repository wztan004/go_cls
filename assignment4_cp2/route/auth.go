// Package route the app to navigate around
package route

import (
	"fmt"
	"html/template"
	"net/http"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	ds "assignment4_cp1/dataStruct"
	"assignment4_cp1/utils"
	"log"
	"os"
	"assignment4_cp1/constants"
	"io"
)

var tpl *template.Template = template.Must(template.ParseGlob("templates/*"))
var mapUsers = map[string]ds.User{}
var mapSessions = map[string]string{}
var mData ds.Data
var wlog *log.Logger // Be concerned
var elog *log.Logger // Critical problem

func init() {
	mData.VenueNames = make(map[string][]ds.Venue)
	if _, ok := mapUsers["admin"]; !ok {
		fmt.Println("navigate > Index > Creating admin account")
		bPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
		mapUsers["admin"] = ds.User{"admin", bPassword, "admin", "admin"}
	}
	file, err := os.OpenFile(constants.LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	
	wlog = log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	elog = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func setSession(userName string, response http.ResponseWriter) {
	fmt.Println("setSession")
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

// Signup allows users to sign up
func Signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var myUser ds.User
	var errorString string

	// 1. when you pressed submit
	if req.Method == http.MethodPost {
		// 2. get values from the fields
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		icnumber := req.FormValue("icnumber")
		email := req.FormValue("email")

		// check if username exist/ taken
		if _, ok := mapUsers[username]; ok {
			errorString += "Username already taken"
		}
		fmt.Println(errorString)

		if res, err := utils.VerifyIC(icnumber); res == false {
			errorString += err.Error()
		}
		fmt.Println(errorString)

		if res, err := utils.VerifyEmail(email); res == false {
			errorString += err.Error()
		}
		fmt.Println(errorString)

		if res, err := utils.VerifyPassword(password); res == false {
			errorString += err.Error()
		}
		fmt.Println(errorString)

		// create session
		id, _ := uuid.NewV4()
		myCookie := &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}

		http.SetCookie(res, myCookie)
		mapSessions[myCookie.Value] = username

		bPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			return
		}
		myUser = ds.User{username, bPassword, firstname, lastname}
		mapUsers[username] = myUser
		
		// redirect to main index

		if len(errorString) > 0 {
			fmt.Println("navigate > Signup > Error signing up")
			tpl.ExecuteTemplate(res, "signup.gohtml", errorString)
			return
		}
		fmt.Println("navigate > Signup > No error signing up")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", myUser)
}


// Index k
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("navigate > Index > Start of index")
	if !alreadyLoggedIn(req) {
		// Checks if admin user is already created
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}


	// Renders the template with user struct
	myUser := getUser(res, req)

	// Get all bookings (duplicated code below, TODO: refactor)
	keys := make([]string, 0, len(mData.VenueNames))
	for k := range mData.VenueNames {
		keys = append(keys, k)
	}
	s2 := []ds.Venue{}
	for _, num1 := range keys {
		for _, num2 := range mData.VenueNames[num1] {
			s2 = append(s2, num2)
		}
	}
	mData.VenueAll = s2

	mData.MyUser = myUser

	// Get user-specific bookings
	s1 := []ds.Venue{}
	for _, num := range mData.VenueNames[myUser.Username] {
		s1 = append(s1, num)
	}
	mData.VenueUser = s1

	if req.Method == http.MethodPost {
		location := req.FormValue("date")
		capacity := req.FormValue("capacity")

		id, _ := uuid.NewV4()

		v := ds.Venue{id.String(), location, capacity, "true"}
		s := mData.VenueNames[myUser.Username]
		s = append(s, v)
		fmt.Println("navigate > Index > mData: ", mData)
		fmt.Println("navigate > Index > mData.VenueNames: ", mData.VenueNames)
		fmt.Println("navigate > Index > myUser.Username: ", myUser.Username)
		fmt.Println("navigate > Index > s: ", s)

		mData.VenueNames[myUser.Username] = s

		// Get user-specific bookings
		s1 := []ds.Venue{}
		for _, num := range mData.VenueNames[myUser.Username] {
			s1 = append(s1, num)
		}
		mData.VenueUser = s1

		// Get all bookings
		keys := make([]string, 0, len(mData.VenueNames))
		for k := range mData.VenueNames {
			keys = append(keys, k)
		}
		s2 := []ds.Venue{}
		for _, num1 := range keys {
			for _, num2 := range mData.VenueNames[num1] {
				s2 = append(s2, num2)
			}
		}
		mData.VenueAll = s2

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "index.gohtml", mData)
}


// Login : allows log in
func Login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	
	// process form submission
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		// check if user exist with username
		myUser, ok := mapUsers[username]
		if !ok {
			// http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			tpl.ExecuteTemplate(res, "login.gohtml", "Please sign up for an account first!")
			return
		}
		// Matching of password entered
		err := bcrypt.CompareHashAndPassword(myUser.Password, []byte(password))
		if err != nil {
			wlog.Println("login failed")
			tpl.ExecuteTemplate(res, "login.gohtml", "Username and/or password do not match!")
			return
		}
		// create session
		id, _ := uuid.NewV4()
		myCookie := &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}
		http.SetCookie(res, myCookie)
		mapSessions[myCookie.Value] = username
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

// Logout : logs the user out of the app
func Logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	myCookie, _ := req.Cookie("myCookie")
	// delete the session
	delete(mapSessions, myCookie.Value)
	// remove the cookie
	myCookie = &http.Cookie{
		Name:   "myCookie",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, myCookie)

	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func createCookie(res http.ResponseWriter, req *http.Request) {
	c := &http.Cookie{
		Name: "first_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}
	http.SetCookie(res, c)
}

func getUser(res http.ResponseWriter, req *http.Request) ds.User {
	// get current session cookie
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		id, _ := uuid.NewV4()
		myCookie = &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}
		http.SetCookie(res, myCookie)
	}
	
	// if the user exists already, get user
	var myUser ds.User
	if username, ok := mapSessions[myCookie.Value]; ok {
		myUser = mapUsers[username]
	}
	return myUser
}

// alreadyLoggedIn checks if user is part of active user list
func alreadyLoggedIn(req *http.Request) bool {
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		return false
	}
	username := mapSessions[myCookie.Value]
	_, ok := mapUsers[username]
	return ok
}
