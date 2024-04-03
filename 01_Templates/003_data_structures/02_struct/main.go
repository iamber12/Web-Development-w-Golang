package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name    string
	Country string
}

type car struct {
	Company string
	Year    string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//sages := []string{"Gandhi", "MLK", "Jesus", "Buddha"}

	gandhi := sage{"Gandhi", "India"}
	buddha := sage{"Buddha", "Sikkim"}

	honda := car{"Honda", "Germany"}
	mg := car{"MG", "UK"}

	sages := []sage{gandhi, buddha}
	cars := []car{honda, mg}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
