package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "testing")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
