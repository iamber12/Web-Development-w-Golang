package main

import (
	"io"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	// retrieve value from url
	// form gets precedence over url
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	// change method to get to get the form value in the body
	io.WriteString(w, `
	<form method="get">
		<input type="text" name="q">
		<input type="submit">
	</form>
	`+v)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
