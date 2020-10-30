package main

import (
	nav "assignment4_cp1/Navigation"
	"html/template"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	ds "assignment4_cp1/dataStruct"
)

var tpl *template.Template
var venues = []ds.Venue{}
var mapUsers = map[string]ds.User{}
var mData ds.Data

func init() {
	// Create credentials for admin
	tpl = template.Must(template.ParseGlob("templates/*"))
	bPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	mapUsers["admin"] = ds.User{"admin", bPassword, "admin", "admin"}
	mData.VenueNames = make(map[string][]ds.Venue)
}

func main() {
	http.HandleFunc("/", nav.Index)
	http.HandleFunc("/names", nav.Names)
	http.HandleFunc("/restricted", nav.Restricted)
	http.HandleFunc("/signup", nav.Signup)
	http.HandleFunc("/login", nav.Login)
	http.HandleFunc("/logout", nav.Logout)
	http.HandleFunc("/remove", nav.Remove)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServeTLS(":5221", "Https/cert.pem", "Https/key.pem", nil)
}
