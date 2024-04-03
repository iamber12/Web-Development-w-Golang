package main

import (
	"io"
	"net/http"
	"os"
)

func getToby(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	defer f.Close()
	io.Copy(w, f)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// io.WriteString(w, `
	// <!--not serving from our server-->
	// <img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	// `)

	io.WriteString(w, `
	<!--not serving from our server-->
	<img src="toby.jpg">
	`)
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", getToby)
	http.ListenAndServe(":8080", nil)
}
