package main

import (
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user

	cookie, err := req.Cookie("session")
	if err != nil {
		return u
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	if s, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	s := dbSessions[cookie.Value]
	_, ok := dbUsers[s.un]
	return ok
}

func cleanSessions() {
	for k, v := range dbSessions {
		if time.Since(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}

	dbSessionsCleaned = time.Now()
}
