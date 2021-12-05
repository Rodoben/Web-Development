package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  string
}

type doubleZero struct {
	person
	LiscenceToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := doubleZero{
		person{
			Name: "Ronald",
			Age:  "56",
		}, false,
	}
	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
