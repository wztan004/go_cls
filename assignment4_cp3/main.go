package main

import (
	"assignment4_cp3/route"
	"net/http"
	ds "assignment4_cp3/datastruct"
	"log"
)

var venues = []ds.Venue{}
var mData ds.Data

func init() {
	mData.VenueNames = make(map[string][]ds.Venue)
}

func main() {
	http.HandleFunc("/", route.Index)
	http.HandleFunc("/names", route.ChangeName)
	http.HandleFunc("/restricted", route.Restricted)
	http.HandleFunc("/signup", route.Signup)
	http.HandleFunc("/login", route.Login)
	http.HandleFunc("/logout", route.Logout)
	http.HandleFunc("/remove", route.Remove)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServeTLS(":5221", "confidential/cert.pem", "confidential/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
