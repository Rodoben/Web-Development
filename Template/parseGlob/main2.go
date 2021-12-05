package main

import (
	"html/template"
	"os"
)

func main() {

	tpl, err := template.ParseGlob("*.gohtml")
	if err != nil {
		panic(err)
	}
	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	panic(err)
	// }

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		panic(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		panic(err)
	}
}
