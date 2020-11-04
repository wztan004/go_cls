package main

import (
	"assignment4_cp2/datastruct"
	// "encoding/csv"
	// "encoding/hex"
	// "errors"
	"fmt"
	"assignment4_cp2/utils"
	// "golang.org/x/crypto/bcrypt"
	// "log"
	// "os"
	"time"
)

type User struct {
	IC			string
	Email		string
	Firstname	string
	Lastname	string
	Username	string
}

func main() {
	ll := datastruct.NewLinkedList()
	
	i := utils.CreateUUID()
	s := datastruct.Session{i, "node1", time.Now()}
	ll.EnqueueSession(s)

	i = utils.CreateUUID()
	s = datastruct.Session{i, "node2", time.Now()}
	ll.EnqueueSession(s)

	i  = utils.CreateUUID()
	s = datastruct.Session{i, "node3", time.Now()}
	ll.EnqueueSession(s)

	// !!! loops over username instead of sessionID
	ll.GetAllID()

	fmt.Println(ll.Head.Session.Username)
	fmt.Println(ll.Tail.Session.Username)
}




// Progress

// AuthenticateUser (ON HOLD, TO RESOLVE A SUB FUNCTION)
// Create New Session (DONE)
// Putting New Session To Linked List (ON HOLD, TO RESOLVE A SUB FUNCTION)
// COMPARE TIME (DONE)


// // AuthenticateUser
// func AuthenticateUser(res http.ResponseWriter, req *http.Request) {
// 	username := req.FormValue("username")
// 	password := req.FormValue("password")
// 	user, err := utils.GetUserCSV(username)
// 	if err != nil {
// 		log.Fatalln("Cannot parse HTML user")
// 	}

// 	bPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
// 	if err != nil {
// 		log.Fatalln("Encryption error")
// 	}
// 	if (string(user.Password) == string(bPassword)) {
// 		// create session
// 			// create UUID
// 		CreateSession(user)
// 		// create cookie
// 		// set cookie
// 	} else {
// 		// redirect
// 	}
// }

