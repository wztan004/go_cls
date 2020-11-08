package main

import (
	"assignment4_cp3/route"
	"assignment4_cp3/constants"
	"net/http"
	"log"
)

func init() {
	initializeUsers()
}

func main() {
	http.HandleFunc("/", route.Index)
	http.HandleFunc("/names", route.ChangeName)
	http.HandleFunc("/restricted", route.Restricted)
	http.HandleFunc("/signup", route.Signup)
	http.HandleFunc("/login", route.Login)
	http.HandleFunc("/logout", route.Logout)
	http.HandleFunc("/book", route.BookVenue)
	http.HandleFunc("/unbook", route.UnbookVenue)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	
	err := http.ListenAndServeTLS(constants.Address, 
		constants.Cert, constants.Key, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
