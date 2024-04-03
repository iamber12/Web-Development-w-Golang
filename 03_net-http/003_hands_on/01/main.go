package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "For dog")
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Index page")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I am Amber")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}
