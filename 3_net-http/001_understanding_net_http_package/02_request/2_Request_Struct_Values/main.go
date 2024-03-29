package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

var tpl template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Parseform - required to read values from the PostForm
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		Submissions   url.Values
		URL           *url.URL
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.Form,
		req.URL,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = *template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
