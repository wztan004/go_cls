package route

import (
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	"fmt"
	"net/http"
	"log"
	"strings"
)

// ChangeName allows the user to change names
func ChangeName(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
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
		result := utils.ReadFile(`confidential\users.csv`)

		EditNames(`confidential\users.csv`, mUserClient.Username, result, firstname, lastname)

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "names.gohtml", data)
}


func EditNames(path string, username string, result [][]string, firstname string, lastname string) {
	var returnRes [][]string
	
	for _, v := range result {
		if (username == v[4]) {
			v[2] = firstname
			v[3] = lastname
			returnRes = append(returnRes, v)
		} else {
			returnRes = append(returnRes, v)
		}
	}

	utils.CreateNewUserCSV(path, returnRes)
}


// Restricted allows the admin to access privilged pages
func Restricted(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if mUserClient.Username != "admin" {
		elog.Fatalln("Unauthorized access, closing server")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var mData data
	mData.MyUser = mUserClient

	fileRes1 := utils.ReadFile(`confidential\venues_202009.csv`)
	fileRes2 := utils.ReadFile(`confidential\venues_202010.csv`)

	venueCSVStruct := func(i []string) datastruct.Venue {
		var mVenue datastruct.Venue
		mVenue.Date = i[0]
		mVenue.Type = i[1]
		mVenue.Capacity = i[2]
		mVenue.BookedBy = i[3]
		mVenue.Username = i[4]
		return mVenue
	}

	updateVenueCSV := func(fileResList ...[][]string) {
		for _, j := range fileResList {
			for _, i := range j {
				if i[4] == "not booked" {
					mData.VenueUnbooked = append(mData.VenueUnbooked, venueCSVStruct(i))
				} else if i[4] == mUserClient.Username {
					mData.VenueUser = append(mData.VenueUser, venueCSVStruct(i))
				}
				mData.VenueAll = append(mData.VenueAll, venueCSVStruct(i))
			}
		}
	}

	updateVenueCSV(fileRes1, fileRes2)

	allUsernames := func() (int, []string) {
		res := utils.ReadFile(`confidential\users.csv`)
		var userList []string
		userNum := 0
		for _, v := range res {
			userList = append(userList, v[4]) //TODO vulnerable to changes
			userNum++
		}
		return userNum, userList
	}
	userNum, userList := allUsernames()

	if req.Method == http.MethodPost {
		// Remove users
		userid := req.FormValue("userid") //username

		for i, _ := range mapUsers {
			if i == userid {
				delete(mapUsers, i)
			}
		}

		http.Redirect(res, req, "/restricted", http.StatusSeeOther)
	}

	sessionList, err := mLinkedList.GetAllID()
	if err != nil {
		fmt.Println(err)
	}

	dataToTemplate := struct {
		MData       data
		UserList    []string
		UserNum     int
		SessionList []string
	}{
		mData,
		userList,
		userNum,
		sessionList,
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("userid")
		mLinkedList.Remove(username)
		http.Redirect(res, req, "/restricted", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "restricted.gohtml", dataToTemplate)
}


func Remove(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Remove()")
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		toIndexPage(res, req)
		return
	}

	if mUserClient.Username != "admin" {
		elog.Fatalln("Unauthorized access, closing server")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var mData data
	mData.MyUser = mUserClient

	fileRes1 := utils.ReadFile(`confidential\venues_202009.csv`)
	fileRes2 := utils.ReadFile(`confidential\venues_202010.csv`)

	venueCSVStruct := func(i []string) datastruct.Venue {
		var mVenue datastruct.Venue
		mVenue.Date = i[0]
		mVenue.Type = i[1]
		mVenue.Capacity = i[2]
		mVenue.BookedBy = i[3]
		mVenue.Username = i[4]
		return mVenue
	}

	updateVenueCSV := func(fileResList ...[][]string) {
		for _, j := range fileResList {
			for _, i := range j {
				if i[4] == "not booked" {
					mData.VenueUnbooked = append(mData.VenueUnbooked, venueCSVStruct(i))
				} else if i[4] == mUserClient.Username {
					mData.VenueUser = append(mData.VenueUser, venueCSVStruct(i))
				}
				mData.VenueAll = append(mData.VenueAll, venueCSVStruct(i))
			}
		}
	}

	updateVenueCSV(fileRes1, fileRes2)

	dataToTemplate := struct {
		MData       data
	}{
		mData,
	}

	if req.Method == http.MethodPost {
		date := strings.TrimSpace(req.FormValue("date"))
		venueType := strings.TrimSpace(req.FormValue("venueType"))
		capacity := strings.TrimSpace(req.FormValue("capacity"))

		// Goes through each CSV to see if the requested venue is in each CSV
		// Breaks off the loop once it's found
		for _, k := range([]string{
			`confidential\venues_202009.csv`,
			`confidential\venues_202010.csv`,
		}) {
			hasBooked, err := EditVenue("remove", k, date, venueType, capacity, mUserClient)
			if err != nil {
				log.Fatalln(err)
			}
			if hasBooked == true {
				http.Redirect(res, req, "/remove", http.StatusSeeOther)
				return
			}
		}

		http.Error(res, "Check your input again. You can only enter available venues", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(res, "remove.gohtml", dataToTemplate)
}


func toIndexPage(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/", http.StatusSeeOther)
	return
}