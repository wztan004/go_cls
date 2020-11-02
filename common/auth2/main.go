// example log in log out server
// https://mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout/

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))






// login handler



// logout handler



// index page

const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`



// internal page

const internalPage = `
<h1>Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`




// When a user visits the default page, this handler renders the static HTML
// (var: indexPage) using Fprintf. This static page includes username and 
// password fields
func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("indexPageHandler")
	fmt.Fprintf(response, indexPage)
}

// While the user is still on the default page, this handler fetches the form
// fields, and passes them to setSession() to set cookie
// After setting the cookie, redirect to the internal page 
func loginHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("loginHandler")
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, response)
		redirectTarget = "/internal"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// Sets the cookie on the browser based on the username field
// It doesn't seem to be strictly "session", but it expire field is unset
func setSession(userName string, response http.ResponseWriter) {
	fmt.Println("setSession")
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

// Checks for the session's username with getUserName().
// Shows the webpage containing the user's name
// If blank, redirect to "/"
func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("internalPageHandler")
	userName := getUserName(request)
	if userName != "" {
		fmt.Fprintf(response, internalPage, userName)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}

// Get the username from the cookie
func getUserName(request *http.Request) (userName string) {
	fmt.Println("getUserName")
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("logoutHandler")
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

// reset by setting value of cookie to blank
func clearSession(response http.ResponseWriter) {
	fmt.Println("clearSession")
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}



// server main method

var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}