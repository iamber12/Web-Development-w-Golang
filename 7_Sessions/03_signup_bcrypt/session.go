package main

import (
	"net/http"
)

func getUser(req *http.Request) user {
	var u user

	cookie, err := req.Cookie("session")
	if err != nil {
		return u
	}

	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok
}
