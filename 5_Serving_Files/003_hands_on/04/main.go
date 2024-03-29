package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func loadInit(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("public/pics"))
	http.HandleFunc("/", loadInit)
	http.Handle("/pics/", http.StripPrefix("/pics", fs))
	http.ListenAndServe(":8080", nil)
}
