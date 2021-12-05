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

	ronald := dataEmployee{
		Name: "Ronald",
		Age:  24,
		Sal:  100021.321,
	}

	surbhi := dataEmployee{
		Name: "Surbhi",
		Age:  24,
		Sal:  100255.336,
	}

	Sam := dataEmployee{
		Name: "sam",
		Age:  54,
		Sal:  658963.255,
	}

	dE := []dataEmployee{ronald, surbhi, Sam}

	err := tpl.Execute(os.Stdout, dE)
	if err != nil {
		panic(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, dE)

}
