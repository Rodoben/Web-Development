package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

var tpl *template.Template

func (p Person) SomeProcessing() int {
	return 7
}

func (p Person) AgeDbl() int {
	return p.Age * 7
}

func (p Person) TakesArgs(x int) int {
	return x * 2
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	p1 := Person{
		Name: "Ronald",
		Age:  36,
	}
	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
