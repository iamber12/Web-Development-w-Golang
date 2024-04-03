package main

import (
	"fmt"
	"log"
	"net/http"
)

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "test-cookie",
		Value: "testing cookies",
	})
}

func read(w http.ResponseWriter, req *http.Request) {
	// we can set multiple cookies
	c, err := req.Cookie("test-cookie")
	if err == http.ErrNoCookie {
		// alternatively redirect to set cookie
		fmt.Println("Cookie not found")
		return
	}
	fmt.Println("Got cookie: ", c)
}

func delete(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("test-cookie")
	if err != nil {
		// alternatively redirect to set cookie
		log.Fatalln("Error fetching cookie")
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	// can set cookie again by redirecting
}

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
