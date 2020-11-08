
package route

import (
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	"fmt"
	"net/http"
	"log"
	"strings"
)

// ChangeName renders the HTML page that allows changing first name and last
// name.
func ChangeName(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var mData datastruct.VenueAvailability
	mData.MyUser = mUserClient

	dataToTemplate := struct {
		MData       datastruct.VenueAvailability
	}{
		mData,
	}

	if req.Method == http.MethodPost {
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		result := utils.ReadFile(`confidential\users.csv`)

		editNames(`confidential\users.csv`, mUserClient.Username, result, firstname, lastname)

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "names.gohtml", dataToTemplate)
}

// editNames allows the user to change first name and last name.
func editNames(path string, username string, result [][]string, firstname string, lastname string) {
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


// Restricted allows the admin to access privileged pages
func Restricted(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if mUserClient.Username != "admin" {
		elog.Println("Unauthorized access. Attempted by username:", mUserClient.Username)
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var mData datastruct.VenueAvailability
	mData.MyUser = mUserClient
	
	venueCSVStruct := func(i []string) datastruct.Venue {
		var mVenue datastruct.Venue
		mVenue.Date = i[0]
		mVenue.Type = i[1]
		mVenue.Capacity = i[2]
		mVenue.BookedBy = i[3]
		mVenue.Username = i[4]
		return mVenue
	}

	updateVenueCSV := func() {
		for _, i := range utils.ReadMultipleFilesConcurrently() {
			if i[4] == "unbook" {
				mData.VenueUnbook = append(mData.VenueUnbook, venueCSVStruct(i))
			} else if i[4] == mUserClient.Username {
				mData.VenueUser = append(mData.VenueUser, venueCSVStruct(i))
			}
			mData.VenueAll = append(mData.VenueAll, venueCSVStruct(i))
		}
	}
	updateVenueCSV()

	allUsernames := func() (int, []string) {
		res := utils.ReadFile(`confidential\users.csv`)
		var userList []string
		userNum := 0
		for _, v := range res {
			username := v[4]
			userList = append(userList, username)
			userNum++
		}
		return userNum, userList
	}
	userNum, userList := allUsernames()

	sessionList, err := mLinkedList.GetAllID()
	if err != nil {
		fmt.Println(err)
	}

	dataToTemplate := struct {
		MData       datastruct.VenueAvailability
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
		fmt.Println("QQQ-Start", username)
		mLinkedList.RemoveSession(username)
		fmt.Println("QQQ-End", username)
		http.Redirect(res, req, "/restricted", http.StatusSeeOther)
		
		return
	}

	tpl.ExecuteTemplate(res, "restricted.gohtml", dataToTemplate)
}


func UnbookVenue(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		toIndexPage(res, req)
		return
	}

	var mData datastruct.VenueAvailability
	mData.MyUser = mUserClient

	venueCSVStruct := func(i []string) datastruct.Venue {
		var mVenue datastruct.Venue
		mVenue.Date = i[0]
		mVenue.Type = i[1]
		mVenue.Capacity = i[2]
		mVenue.BookedBy = i[3]
		mVenue.Username = i[4]
		return mVenue
	}

	updateVenueCSV := func() {
		for _, i := range utils.ReadMultipleFilesConcurrently() {
			if i[4] == "unbook" {
				mData.VenueUnbook = append(mData.VenueUnbook, venueCSVStruct(i))
			} else if i[4] == mUserClient.Username {
				mData.VenueUser = append(mData.VenueUser, venueCSVStruct(i))
			}
			mData.VenueAll = append(mData.VenueAll, venueCSVStruct(i))
		}
	}
	updateVenueCSV()

	dataToTemplate := struct {
		MData       datastruct.VenueAvailability
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
			hasBooked, err := EditVenue("unbook", k, date, venueType, capacity, mUserClient)
			if err != nil {
				log.Fatalln(err)
			}
			if hasBooked == true {
				http.Redirect(res, req, "/unbook", http.StatusSeeOther)
				return
			}
		}

		http.Error(res, "Check your input again. You can only enter available venues", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(res, "unbook.gohtml", dataToTemplate)
}


func toIndexPage(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/", http.StatusSeeOther)
	return
}