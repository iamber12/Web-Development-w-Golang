package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Item struct {
	Name  string
	Price float64
}

type Meal struct {
	Items []Item
	Type  string
}

type Restaurant struct {
	Meals []Meal
	Name  string
}

type restaurants []Restaurant

func main() {
	r := restaurants{
		Restaurant{
			Name: "Hilton",
			Meals: []Meal{
				{
					Type: "Breakfast",
					Items: []Item{
						{
							Name:  "Oatmeal",
							Price: 4.95,
						},
						{
							Name:  "Cheerios",
							Price: 3.95,
						},
						{
							Name:  "Juice Orange",
							Price: 2.95,
						},
					},
				},
				{
					Type: "Lunch",
					Items: []Item{
						{
							Name:  "Hamburger",
							Price: 6.95,
						},
						{
							Name:  "Cheese Melted Sandwich",
							Price: 3.95,
						},
						{
							Name:  "French Fries",
							Price: 2.95,
						},
					},
				},
				{
					Type: "Dinner",
					Items: []Item{
						{
							Name:  "Pasta Bolognese",
							Price: 7.95,
						},
						{
							Name:  "Steak",
							Price: 13.95,
						},
						{
							Name:  "Bistro Potatoe",
							Price: 6.95,
						},
					},
				},
			},
		},
		Restaurant{
			Name: "Country Inn",
			Meals: []Meal{
				{
					Type: "Breakfast",
					Items: []Item{
						{
							Name:  "Oatmeal",
							Price: 4.95,
						},
						{
							Name:  "Cheerios",
							Price: 3.95,
						},
						{
							Name:  "Juice Orange",
							Price: 2.95,
						},
					},
				},
				{
					Type: "Lunch",
					Items: []Item{
						{
							Name:  "Hamburger",
							Price: 6.95,
						},
						{
							Name:  "Cheese Melted Sandwich",
							Price: 3.95,
						},
						{
							Name:  "French Fries",
							Price: 2.95,
						},
					},
				},
				{
					Type: "Dinner",
					Items: []Item{
						{
							Name:  "Pasta Bolognese",
							Price: 7.95,
						},
						{
							Name:  "Steak",
							Price: 13.95,
						},
						{
							Name:  "Bistro Potatoe",
							Price: 6.95,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
