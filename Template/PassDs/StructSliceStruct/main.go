package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("one.gohtml"))

}

func main() {

	type sage struct {
		Name  string
		Motto string
	}

	type car struct {
		Model string
		Year  int
		Doors int
	}

	sage1 := sage{
		Name:  "Ronald",
		Motto: "Fuck off World",
	}

	sage2 := sage{
		Name:  "Fonald",
		Motto: "Fuck off  this World",
	}

	car1 := car{
		Model: "Ford",
		Year:  2006,
		Doors: 2,
	}
	car2 := car{
		Model: "Ford",
		Year:  2006,
		Doors: 2,
	}

	car3 := car{
		Model: "Ford",
		Year:  2006,
		Doors: 2,
	}

	type items struct {
		Wisdom    []sage
		Transport []car
	}

	sages := []sage{sage1, sage2}
	cars := []car{car1, car2, car3}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(nf, data)
	if err != nil {
		panic(err)
	}

}
