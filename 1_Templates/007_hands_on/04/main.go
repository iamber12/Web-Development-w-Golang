package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type column struct {
	Date time.Time
	Open float64
}

var table []column

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "table.csv", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
