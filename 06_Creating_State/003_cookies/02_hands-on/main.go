package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("count-cookie")

	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "count-cookie",
			Value: "0",
		})
	}

	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	c.Value = strconv.Itoa(count)
	http.SetCookie(w, c)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("count-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Printf("Count: %v\n", c.Value)
}

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
