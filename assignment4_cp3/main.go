package main

import (
	"assignment4_cp3/route"
	"assignment4_cp3/constants"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", route.Index)
	http.HandleFunc("/names", route.ChangeName)
	http.HandleFunc("/restricted", route.Restricted)
	http.HandleFunc("/signup", route.Signup)
	http.HandleFunc("/login", route.Login)
	http.HandleFunc("/logout", route.Logout)
	http.HandleFunc("/remove", route.Remove)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	
	err := http.ListenAndServeTLS(constants.SERVER_ADDRESS, 
		constants.SSH_CERTIFICATE, constants.SSH_KEY, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
