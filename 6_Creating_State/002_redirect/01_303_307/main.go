package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at bar: ", req.Method, "\n\n")

	/****** 303 *******/
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusSeeOther)
	// above can also be written as -
	//http.Redirect(w, req, "/", http.StatusSeeOther)
	/****** 303 *******/

	/****** 307 *******/
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
	/****** 307 *******/
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
