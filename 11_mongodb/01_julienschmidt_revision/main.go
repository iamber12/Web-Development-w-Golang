package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome")
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	http.ListenAndServe("localhost:8080", r)
}
