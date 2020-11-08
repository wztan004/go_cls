
package route

import (
	"assignment4_cp3/constants"
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	"net/http"
	"errors"
	"log"
	"strings"
)


// ChangeName renders the HTML page that allows changing first name and last
// name.
func ChangeName(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		toIndexPage(res,req)
		return
	}

	var mData datastruct.VenueAvailability
	var errMsg string
	mData.MyUser = mUserClient

	dataToTemplate := struct {
		MData       datastruct.VenueAvailability
		ErrorMessage string
	}{
		mData,
		errMsg,
	}

	if req.Method == http.MethodPost {
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		if len(firstname) == 0 || len(lastname) == 0 {
			dataToTemplate.ErrorMessage = "You must enter something as your first name and last name."
			tpl.ExecuteTemplate(res, "names.gohtml", dataToTemplate)
			return
		}
		editNames(mUserClient.Username, firstname, lastname)
		toIndexPage(res,req)
		return
	}

	tpl.ExecuteTemplate(res, "names.gohtml", dataToTemplate)
}

// Restricted renders a HTML page that allows the admin to access privileged
// pages
func Restricted(res http.ResponseWriter, req *http.Request) {
	// Checks if user is logged in and renders data
	// If not, redirect to home page
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		toIndexPage(res,req)
		return
	}

	if mUserClient.Username != "admin" {
		elog.Println("Unauthorized access. Attempted by username:", mUserClient.Username)
		toIndexPage(res,req)
		return
	}

	var mData datastruct.VenueAvailability
	mData.MyUser = mUserClient

	updateTemplateData(&mData, &mUserClient)

	allUsernames := func() (int, []string) {
		res := utils.ReadFile(constants.UserFile)
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
		wlog.Println(err)
	}

	var errMsg string
	dataToTemplate := struct {
		MData       datastruct.VenueAvailability
		UserList    []string
		UserNum     int
		SessionList []string
		ErrMsg		string
	}{
		mData,
		userList,
		userNum,
		sessionList,
		errMsg,
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("userid")
		mLinkedList.RemoveSession(username)
		_, err := utils.GetUserCSV(username)
		if err != nil {
			dataToTemplate.ErrMsg = "User does not exist!" //TODO current implementation does not search for sessions
			tpl.ExecuteTemplate(res, "restricted.gohtml", dataToTemplate)
			return
		}
		tpl.ExecuteTemplate(res, "restricted.gohtml", dataToTemplate)
		return
	}

	tpl.ExecuteTemplate(res, "restricted.gohtml", dataToTemplate)
}

// BookVenue show the user's current booked venues and main actions to take 
// in the system.
func BookVenue(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		http.Redirect(res, req, "/login", 302)
		return
	}
	var mData datastruct.VenueAvailability
	var errMsg string
	dataToTemplate := struct {
		MData       datastruct.VenueAvailability
		ErrMsg		string
	}{
		mData,
		errMsg,
	}
	mData.MyUser = mUserClient

	updateTemplateData(&mData, &mUserClient)
	dataToTemplate.MData = mData

	if req.Method == http.MethodPost {
		date := strings.TrimSpace(req.FormValue("date"))
		venueType := strings.TrimSpace(req.FormValue("venueType"))
		capacity := strings.TrimSpace(req.FormValue("capacity"))

		// Goes through each CSV to see if the requested venue is in each CSV
		// Once found, update the venue info and break off the loop.
		for _, k := range []string{
			constants.LatestMthLess1,
			constants.LatestMth,
		} {
			hasBooked, err := EditVenue("book", k, date, venueType, capacity, mUserClient)
			if err != nil {
				log.Fatalln(err)
			}
			if hasBooked == true {
				toIndexPage(res,req)
				return
			}
		}
		dataToTemplate.ErrMsg = "Check your entry again. You can only enter available venues"
		tpl.ExecuteTemplate(res, "book.gohtml", dataToTemplate)
		return
	}
	tpl.ExecuteTemplate(res, "book.gohtml", dataToTemplate)
}


// UnbookVenue renders a HTML page that allows the a user to unbook own
// bookings.
func UnbookVenue(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		toIndexPage(res, req)
		return
	}
	var mData datastruct.VenueAvailability
	mData.MyUser = mUserClient
	updateTemplateData(&mData, &mUserClient)

	var errMsg string
	dataToTemplate := struct {
		MData       datastruct.VenueAvailability
		ErrMsg		string
	}{
		mData,
		errMsg,
	}
	if req.Method == http.MethodPost {
		date := strings.TrimSpace(req.FormValue("date"))
		venueType := strings.TrimSpace(req.FormValue("venueType"))
		capacity := strings.TrimSpace(req.FormValue("capacity"))
		// Goes through each CSV to see if the requested venue is in each CSV
		// Breaks off the loop once it's found
		for _, k := range([]string{
			constants.LatestMthLess1,
			constants.LatestMth,
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
		dataToTemplate.ErrMsg = "Check your input again. You can only enter booked venues"
		tpl.ExecuteTemplate(res, "unbook.gohtml", dataToTemplate)
		return
	}

	tpl.ExecuteTemplate(res, "unbook.gohtml", dataToTemplate)
}

// EditVenue allows the user to either add or unbook a venue. Adding a venue can
// only be applied to unbooked venues, removing a venue can only be applied to
// booked venues. It edits the last two fields of the CSV file: email and
// username. Adding a venue would set the two fields to user-specific info,
// while removing a venue would set them to "unbook".
//
// Action should be either "book" or "unbook".
// Path is the relative path to the CSV venue file.
// Date follows the format of "YYYYMMDD".
// Capacity is an integer but formatted in string.
// mUserClient is the user's info.
func EditVenue(action string, path string, date string, venueType string, capacity string, mUserClient datastruct.UserClient) (bool, error) {
	if action != "book" && action != "unbook" {
		return false, errors.New("Choose either 'book' or 'unbook'")
	}

	// 1. Read a CSV file
	res := utils.ReadFile(path)

	// 2. If record matches, update the array
	var returnRes [][]string

	toReturn := false
	for _, v := range res {
		if date == v[0] && strings.ToLower(venueType) == strings.ToLower(v[1]) && capacity == v[2] {
			if action == "book" {
				v[3] = mUserClient.Email
				v[4] = mUserClient.Username
				toReturn = true
			} else if action == "unbook" && (mUserClient.Username == v[4] || mUserClient.Username == "admin") {
				v[3] = "unbook"
				v[4] = "unbook"
				toReturn = true
			}
			returnRes = append(returnRes, v)
		} else {
			returnRes = append(returnRes, v)
		}
	}

	// 3. Save the array into the same CSV
	if toReturn {
		err := createNewBookingCSV(path, returnRes)
		if err != nil {
			log.Fatalln("Make sure the file format is correct")
		}
	}

	// 4. Return function
	return toReturn, nil
}