package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl template.Template

func init() {
	tpl = *template.Must(template.ParseGlob("Templates/*"))
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "Introduction to Web dev with golang")
	if err != nil {
		log.Fatalln("Error!")
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln("Error!")
	}
}

func me(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "me.gohtml", "Amber")
	if err != nil {
		log.Fatalln("Error!")
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}
