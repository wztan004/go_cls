// Package route includes routing details, including authentication and key
// options (e.g. book venues) within the web app.
package route

import (
	"assignment4_cp3/datastruct"
	"assignment4_cp3/utils"
	"assignment4_cp3/constants"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	tpl *template.Template = template.Must(template.ParseGlob("templates/*"))
	mLinkedList datastruct.SessionLinkedList
	wlog *log.Logger // Be concerned
	elog *log.Logger // Error problem
	clog *log.Logger // Critical problem
)

func init() {
	file, err := os.OpenFile(constants.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	wlog = log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	elog = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	clog = log.New(io.MultiWriter(file, os.Stderr), "CRITICAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}



// Index is the home page. If the user is logged in, it will show the user's
// current booked venues and main actions to take in the system.
func Index(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, _ := alreadyLoggedIn(req)
	if !isLoggedIn {
		http.Redirect(res, req, "/login", 302)
		return
	}
	http.Redirect(res, req, "/book", 302)
	return
}




// Signup starts account registration, validates their fields, saves the
// profile and then redirect to the home page.
func Signup(res http.ResponseWriter, req *http.Request) {
	// 1. If logged in, redirect to Index().
	isLoggedIn, _ := alreadyLoggedIn(req)
	if isLoggedIn {
		toIndexPage(res,req)
		return
	}

	// 2. Perform form validation with regular expressions.
	var errorString []string
	if req.Method == http.MethodPost {
		username := strings.TrimSpace(req.FormValue("username"))
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		ic := req.FormValue("icnumber")
		email := req.FormValue("email")
		if username == "" {
			errorString = append(errorString, "Enter a username.")
		}
		if _, err := utils.GetUserCSV(username); err == nil {
			errorString = append(errorString,"Username is already taken. Choose another one.")
		}
		if res, err := utils.VerifyIC(ic); res == false {
			errorString = append(errorString,err.Error())
		}
		if res, err := utils.VerifyEmail(email); res == false {
			errorString = append(errorString,err.Error())
		}
		if res, err := utils.VerifyPassword(password); res == false {
			errorString = append(errorString,err.Error())
		}
		if len(errorString) > 0 {
			tpl.ExecuteTemplate(res, "signup.gohtml", errorString)
			return
		}

		// 3. CSV User: Update
		bPassword := utils.CreateChecksum(password)
		utils.WriteCSV(constants.UserFile, []string{ic, email, firstname, lastname, username, bPassword})
		startSession(res, req, username)
		toIndexPage(res,req)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

// Login allows user to authenticate. If successful, start a session and
// redirect to home page.
func Login(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, _ := alreadyLoggedIn(req)
	if isLoggedIn {
		toIndexPage(res,req)
		return
	}
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		ok := authenticateUser(username, password)
		if ok {
			startSession(res, req, username)
			toIndexPage(res,req)
			return
		} 
		tpl.ExecuteTemplate(res, "login.gohtml", "Not the right username/password.")
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", "")
}

// Logout removes the session in the server and expires the client cookie.
func Logout(res http.ResponseWriter, req *http.Request) {
	isLoggedIn, mUserClient := alreadyLoggedIn(req)
	if !isLoggedIn {
		toIndexPage(res,req)
		return
	}

	mLinkedList.RemoveSession(mUserClient.Username)

	myCookie := &http.Cookie{
		Name:   constants.CookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, myCookie)
	toIndexPage(res,req)
}

