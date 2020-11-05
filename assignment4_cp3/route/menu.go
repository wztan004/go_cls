package route

import (
	"fmt"
	"net/http"
	"assignment4_cp3/datastruct"
)

// ChangeName allows the user to change names
func ChangeName(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	b, mUserClient := alreadyLoggedIn(req)
	if !b {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}


	data := struct {
		Firstname string
		Lastname  string
	}{
		mUserClient.Firstname,
		mUserClient.Lastname,
	}


	if req.Method == http.MethodPost {
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")

		mStruct := mapUsers[mUserClient.Username]
		mStruct.Firstname = firstname
		mStruct.Lastname = lastname

		myUser := datastruct.UserClient{mUserClient.Username, mUserClient.Firstname, mUserClient.Lastname}
		mapUsers[myUser.Username] = myUser

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "names.gohtml", data)
}

// Restricted allows the admin to access privilged pages
func Restricted(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	b, mUserClient := alreadyLoggedIn(req)
	if !b {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if mUserClient.Username != "admin" {
		elog.Fatalln("Unauthorized access, closing server")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	userList := make([]string, 0, len(mapUsers))

	for i, _ := range mapUsers {
		userList = append(userList, i)
	}

	data := struct {
		Users []string
	}{
		userList,
	}
	fmt.Println(mapUsers)

	if req.Method == http.MethodPost {
		userid := req.FormValue("userid") //username

		for i, _ := range mapUsers {
			if i == userid {
				delete(mapUsers, i)
			}
		}

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "restricted.gohtml", data)
}

// Remove allows the admin to remove a user
func Remove(res http.ResponseWriter, req *http.Request) {
	b, mUserClient := alreadyLoggedIn(req)
	if !b {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		id := req.FormValue("id")
		s := []datastruct.Venue{}
		for _, num := range mData.VenueNames[mUserClient.Username] {
			if num.ID != id {
				s = append(s, num)
			}
		}
		mData.VenueNames[mUserClient.Username] = s
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "remove.gohtml", mData)
}