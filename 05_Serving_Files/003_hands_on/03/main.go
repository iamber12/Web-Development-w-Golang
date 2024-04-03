package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func loadInit(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("starting-files/public"))
	http.HandleFunc("/", loadInit)
	http.Handle("/pics/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
