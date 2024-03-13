package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/dcd", dcd)
	http.HandleFunc("/encd", encd)
	http.HandleFunc("/unmshl", unmshl)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>FOO</title>
		</head>
		<body>
		You are at foo
		</body>
		</html>`
	w.Write([]byte(s))
}

func unmshl(w http.ResponseWriter, req *http.Request) {
	var data person
	rcvd := `{"Fname":"James","Lname":"Bond","Items":["Suit","Gun","Wry sense of humor"]}`
	if err := json.Unmarshal([]byte(rcvd), &data); err != nil {
		log.Println(err)
	}

	fmt.Fprintln(w, data)
}

func dcd(w http.ResponseWriter, req *http.Request) {
	var data person
	rcvd := `{"Fname":"James","Lname":"Bond","Items":["Suit","Gun","Wry sense of humor"]}`
	err := json.NewDecoder(strings.NewReader(rcvd)).Decode(&data)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, data)
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
	// can also use  -
	// fmt.Fprintln(w, string(json))
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
