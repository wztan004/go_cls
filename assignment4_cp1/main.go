package main

import (
	"assignment4_cp1/navigate"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	ds "assignment4_cp1/dataStruct"
)

// var tpl *template.Template
var venues = []ds.Venue{}
var mapUsers = map[string]ds.User{}
var mData ds.Data

func init() {
	bPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	mapUsers["admin"] = ds.User{"admin", bPassword, "admin", "admin"}
	mData.VenueNames = make(map[string][]ds.Venue)
}

func main() {
	http.HandleFunc("/", navigate.Index)
	http.HandleFunc("/names", navigate.Names)
	http.HandleFunc("/restricted", navigate.Restricted)
	http.HandleFunc("/signup", navigate.Signup)
	http.HandleFunc("/login", navigate.Login)
	http.HandleFunc("/logout", navigate.Logout)
	http.HandleFunc("/remove", navigate.Remove)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServeTLS(":5221", "security/cert.pem", "security/key.pem", nil)
}
