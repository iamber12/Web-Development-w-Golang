package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func loadInit(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("Error")
	}
}

func main() {
	fs := http.FileServer(http.Dir("public/pic"))
	http.HandleFunc("/", loadInit)
	http.Handle("/public/pic/", http.StripPrefix("/public/pic", fs))
	http.ListenAndServe(":8080", nil)
}
