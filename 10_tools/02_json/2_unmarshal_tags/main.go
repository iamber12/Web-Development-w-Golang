package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type City struct {
	// names can be anything in this DS
	PostalCode string  `json:"Postal"`
	Lat        float64 `json:"Latitude"`
	Long       float64 `json:"Longitude"`
	Address    string  `json:"Address"`
	City       string  `json:"City"`
	State      string  `json:"State"`
	Zip        string  `json:"Zip"`
	Country    string  `json:"Country"`
}

type cities []City

func main() {
	var data cities
	rcvd := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(data)
}
