package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		// encrypt password
		password, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.MinCost)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		// check if username is taken
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username is taken", http.StatusForbidden)
		}

		// set session
		id := uuid.NewString()
		cookie := &http.Cookie{
			Name:  "session",
			Value: id,
		}
		http.SetCookie(w, cookie)

		u := user{un, password, fn, ln}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
