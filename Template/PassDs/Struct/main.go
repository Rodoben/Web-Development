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

	type dataEmployee struct {
		Name string
		Age  int
		Sal  float64
	}

	ramesh := dataEmployee{
		Name: "ramesh",
		Age:  24,
		Sal:  10000.2032,
	}

	err := tpl.Execute(os.Stdout, ramesh)
	if err != nil {
		panic(err)
	}
	nf, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(nf, ramesh)
	if err != nil {
		panic(err)
	}

}
