package main

import (
	"html/template"
	"net/http"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type user struct {
	Username string
	Password []byte
	First    string
	Last     string
}

type data struct {
	MyUser user
	VenueNames map[string][]venue
	VenueUser []venue
	VenueAll []venue
}

type venue struct {
	ID string
	Date string
	Capacity string
	IsBooked string
}

var mData data

var tpl *template.Template
var mapUsers = map[string]user{}
var mapSessions = map[string]string{}
var venues = []venue{}


func init() {
	// Create credentials for admin
	tpl = template.Must(template.ParseGlob("templates/*"))
	bPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	mapUsers["admin"] = user{"admin", bPassword, "admin", "admin"}
	mData.VenueNames = make(map[string][]venue)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/names", names)
	http.HandleFunc("/restricted", restricted)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/remove", remove)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5221", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
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
	s2 := []venue{}
	for _, num1 := range keys {
		for _, num2 := range mData.VenueNames[num1] {
			s2 = append(s2, num2)
		}
	}
	mData.VenueAll = s2

	mData.MyUser = myUser

	// Get user-specific bookings
	s1 := []venue{}
	for _, num := range mData.VenueNames[myUser.Username] {
		s1 = append(s1, num)
	}
	mData.VenueUser = s1

	if req.Method == http.MethodPost {
		location := req.FormValue("date")
		capacity := req.FormValue("capacity")

		id, _ := uuid.NewV4()
		
		v := venue{id.String(),location,capacity,"true"}
		s := mData.VenueNames[myUser.Username]
		s = append(s,v)
		mData.VenueNames[myUser.Username] = s

		// Get user-specific bookings
		s1 := []venue{}
		for _, num := range mData.VenueNames[myUser.Username] {
			s1 = append(s1, num)
		}
		mData.VenueUser = s1


		// Get all bookings
		keys := make([]string, 0, len(mData.VenueNames))
		for k := range mData.VenueNames {
			keys = append(keys, k)
		}
		s2 := []venue{}
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

func names(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	u := getUser(res, req)

	data := struct {
		Firstname string 
		Lastname string 
	} { 
		u.First,
		u.Last,
	}

	
	fmt.Println("Q",u.First, u.Last)

	if req.Method == http.MethodPost {
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		
		mStruct := mapUsers[u.Username]
		mStruct.First = firstname
		mStruct.Last = lastname

		myUser := user{u.Username, u.Password, firstname, lastname}
		mapUsers[u.Username] = myUser

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "names.gohtml", data)
}

func restricted(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(res, req)

	if u.Username != "admin" {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	userList := make([]string, 0, len(mapUsers))

	for i,_ := range mapUsers {
		userList = append(userList, i)
	}

	data := struct {
		Users []string 
	} { 
		userList,
	}
	fmt.Println(mapUsers)

	if req.Method == http.MethodPost {
		userid := req.FormValue("userid") //username

		for i,_ := range mapUsers {
			if i == userid {
				delete(mapUsers,i)
			} 
		}

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "restricted.gohtml", data)
}

func remove(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	
	u := getUser(res, req)
	mData.MyUser = u

	if req.Method == http.MethodPost {
		id := req.FormValue("id")
		s := []venue{}
		for _, num := range mData.VenueNames[u.Username] {
			if num.ID != id {
				s = append(s,num)
			}
		}
		mData.VenueNames[u.Username] = s
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "remove.gohtml", mData)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var myUser user

	// 1. when you pressed submit
	if req.Method == http.MethodPost {
		// 2. get values from the fields
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		if username != "" {
			// check if username exist/ taken
			if _, ok := mapUsers[username]; ok {
				http.Error(res, "Username already taken", http.StatusForbidden)
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
			bPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
			if err != nil {
				http.Error(res, "Internal server error", http.StatusInternalServerError)
				return
			}
			myUser = user{username, bPassword, firstname, lastname}
			mapUsers[username] = myUser
		}
		// redirect to main index

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", myUser)
}

func login(res http.ResponseWriter, req *http.Request) {
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
			http.Error(res, "Username and/or password do not match", http.
				StatusForbidden)
			return
		}
		// Matching of password entered
		err := bcrypt.CompareHashAndPassword(myUser.Password, []byte(password))
		if err != nil {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
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

func logout(res http.ResponseWriter, req *http.Request) {
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

func getUser(res http.ResponseWriter, req *http.Request) user {
	// get current session cookie
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		id, _ := uuid.NewV4()
		myCookie = &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}
	}
	http.SetCookie(res, myCookie)
	// if the user exists already, get user
	var myUser user
	if username, ok := mapSessions[myCookie.Value]; ok {
		myUser = mapUsers[username]
	}
	return myUser
}

func alreadyLoggedIn(req *http.Request) bool {
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		return false
	}
	username := mapSessions[myCookie.Value]

	_, ok := mapUsers[username]
	return ok
}
