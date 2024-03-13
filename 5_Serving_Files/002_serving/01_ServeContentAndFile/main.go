package main

import (
	"io"
	"net/http"
)

func getToby(w http.ResponseWriter, req *http.Request) {
	// f, err := os.Open("toby.jpg")
	// if err != nil {
	// 	http.Error(w, "File not found", 404)
	// 	return
	// }

	// defer f.Close()

	// fi, err := f.Stat()
	// if err != nil {
	// 	http.Error(w, "File not found", 404)
	// 	return
	// }
	// http.ServeContent(w, req, f.Name(), fi.ModTime(), f)

	// OR
	http.ServeFile(w, req, "toby.jpg")
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">`)
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", getToby)
	http.ListenAndServe(":8080", nil)
}
