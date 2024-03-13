package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

var tpl template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// takes values only from the body
	// tpl.ExecuteTemplate(w, "index.gohtml", req.PostForm)

	// takes values from url as well
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func init() {
	tpl = *template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
