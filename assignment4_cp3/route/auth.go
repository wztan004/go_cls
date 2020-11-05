// Package route the app to navigate around
package route

import (
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	"assignment4_cp3/constants"
	"fmt"
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	
	"log"
	"os"
	
	"io"
)

var tpl *template.Template = template.Must(template.ParseGlob("templates/*"))
var mapUsers = map[string]datastruct.User{}
var mapSessions = map[string]string{}
var mData datastruct.Data
var mLinkedList datastruct.LinkedList
var wlog *log.Logger // Be concerned
var elog *log.Logger // Critical problem

func init() {
	utils.InitializeUsers()
	mData.VenueNames = make(map[string][]datastruct.Venue)
	// if _, ok := mapUsers["admin"]; !ok {
	// 	fmt.Println("navigate > Index > Creating admin account")
	// 	bPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	// 	mapUsers["admin"] = ds.User{"admin", bPassword, "admin", "admin"}
	// }
	file, err := os.OpenFile(constants.LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	
	wlog = log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	elog = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// func setSession(userName string, response http.ResponseWriter) {
// 	fmt.Println("setSession")
// 	value := map[string]string{
// 		"name": userName,
// 	}
// 	if encoded, err := cookieHandler.Encode("session", value); err == nil {
// 		cookie := &http.Cookie{
// 			Name:  "session",
// 			Value: encoded,
// 			Path:  "/",
// 		}
// 		http.SetCookie(response, cookie)
// 	}
// }

// Signup allows users to sign up
func Signup(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Signup")
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var myUser datastruct.User
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
		myUser = datastruct.User{username, bPassword, firstname, lastname}
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
	fmt.Println("Index()")
	if !alreadyLoggedIn(req) {
		fmt.Println("Index() 1")
		http.Redirect(res, req, "/login", 302)
		return
	}

	fmt.Println("Index() 2")


	// Renders the template with user struct
	myUser := getUser(res, req)

	// Get all bookings (duplicated code below, TODO: refactor)
	keys := make([]string, 0, len(mData.VenueNames))
	for k := range mData.VenueNames {
		keys = append(keys, k)
	}
	s2 := []datastruct.Venue{}
	for _, num1 := range keys {
		for _, num2 := range mData.VenueNames[num1] {
			s2 = append(s2, num2)
		}
	}
	mData.VenueAll = s2

	mData.MyUser = myUser

	// Get user-specific bookings
	s1 := []datastruct.Venue{}
	for _, num := range mData.VenueNames[myUser.Username] {
		s1 = append(s1, num)
	}
	mData.VenueUser = s1

	if req.Method == http.MethodPost {
		location := req.FormValue("date")
		capacity := req.FormValue("capacity")

		id, _ := uuid.NewV4()

		v := datastruct.Venue{id.String(), location, capacity, "true"}
		s := mData.VenueNames[myUser.Username]
		s = append(s, v)
		fmt.Println("navigate > Index > mData: ", mData)
		fmt.Println("navigate > Index > mData.VenueNames: ", mData.VenueNames)
		fmt.Println("navigate > Index > myUser.Username: ", myUser.Username)
		fmt.Println("navigate > Index > s: ", s)

		mData.VenueNames[myUser.Username] = s

		// Get user-specific bookings
		s1 := []datastruct.Venue{}
		for _, num := range mData.VenueNames[myUser.Username] {
			s1 = append(s1, num)
		}
		mData.VenueUser = s1

		// Get all bookings
		keys := make([]string, 0, len(mData.VenueNames))
		for k := range mData.VenueNames {
			keys = append(keys, k)
		}
		s2 := []datastruct.Venue{}
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
	fmt.Println("Login")
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", 302)
	}

	// process form submission
	if req.Method == http.MethodPost {

		ok := authenticateUser(res, req)

		if ok {
			fmt.Println("Login success")
			http.Redirect(res, req, "/", 302)
		} else {
			fmt.Println("Login failed")
			http.Redirect(res, req, "/login", 302)
		}
		
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

// Logout : logs the user out of the app
func Logout(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Logout()")
	if !alreadyLoggedIn(req) {
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



func getUser(res http.ResponseWriter, req *http.Request) datastruct.User {
	fmt.Println("getUser")
	// get current session cookie
	myCookie, err := req.Cookie("_cookie")
	if err != nil {
		fmt.Println("getUser, creating cookie...")
		id, _ := uuid.NewV4()
		myCookie = &http.Cookie{
			Name:  "_cookie",
			Value: id.String(),
		}
		http.SetCookie(res, myCookie)
	}
	
	// if the user exists already, get user
	var myUser datastruct.User
	if username, ok := mapSessions[myCookie.Value]; ok {
		myUser = mapUsers[username]
	}
	return myUser
}

// alreadyLoggedIn checks if user is part of active user list
func alreadyLoggedIn(req *http.Request) bool {
	fmt.Println("alreadyLoggedIn()")
	myCookie, err := req.Cookie("_cookie")
	if err != nil {
		return false
	}
	return mLinkedList.CheckSessionID(myCookie.Value)
}



// authenticateUser (Still in testing: sub-function issue)
// encryption is non-deterministic
func authenticateUser(res http.ResponseWriter, req *http.Request) bool {
	fmt.Println("authenticateUser")
	username := req.FormValue("username")
	password := req.FormValue("password")
	user, err := utils.GetUserCSV(username)
	if err != nil {
		// user does not exist
		return false
	}

	bPassword := utils.Encrypt(password)

	if (string(user.Password) == bPassword) {
		// create session
		mSession := utils.CreateSession(user.Username)

		// update session linked list
		mLinkedList.EnqueueSession(mSession)
		x, err := mLinkedList.GetAllID()
		fmt.Println("1", x, err)

		// create cookie
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    mSession.SessionUUID,
			HttpOnly: true,
		}
		fmt.Println("cookie is set", cookie)

		http.SetCookie(res, &cookie)
		return true
	}
	return false
}