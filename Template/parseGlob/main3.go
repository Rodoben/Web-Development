package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		panic(err)
	}

	err1 := tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		panic(err1)
	}

	err2 := tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		panic(err2)
	}

	nf, err := os.Create("index1.html")
	if err != nil {
		panic(err)
	}
	err = tpl.ExecuteTemplate(nf, "one.gohtml", 42)
	if err != nil {
		panic(err)
	}

}
