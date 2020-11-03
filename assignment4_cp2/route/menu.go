package route

import (
	"fmt"
	"net/http"
	ds "assignment4_cp2/dataStruct"
)

// ChangeName allows the user to change names
func ChangeName(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	u := getUser(res, req)

	data := struct {
		Firstname string
		Lastname  string
	}{
		u.First,
		u.Last,
	}

	fmt.Println("Q", u.First, u.Last)

	if req.Method == http.MethodPost {
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")

		mStruct := mapUsers[u.Username]
		mStruct.First = firstname
		mStruct.Last = lastname

		myUser := ds.User{u.Username, u.Password, firstname, lastname}
		mapUsers[u.Username] = myUser

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "names.gohtml", data)
}

// Restricted allows the admin to access privilged pages
func Restricted(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(res, req)

	if u.Username != "admin" {
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
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	u := getUser(res, req)
	mData.MyUser = u

	if req.Method == http.MethodPost {
		id := req.FormValue("id")
		s := []ds.Venue{}
		for _, num := range mData.VenueNames[u.Username] {
			if num.ID != id {
				s = append(s, num)
			}
		}
		mData.VenueNames[u.Username] = s
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "remove.gohtml", mData)
}