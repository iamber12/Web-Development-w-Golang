package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Status struct {
	Code    string `json:"Code"`
	Descrip string `json:"Descrip"`
}

type Statuses []Status

func main() {
	// http.HandleFunc("/", index)
	// http.ListenAndServe(":8080", nil)
	var data Statuses
	rcvd := `[{"Code":"200","Descrip":"StatusOK"},{"Code":"301","Descrip":"StatusMovedPermanently"}]`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, elem := range data {
		fmt.Println("Code: ", elem.Code)
		fmt.Println("Description: ", elem.Descrip)
	}
}
