package main

import (
	"os"
	"text/template"
)

func main() {
	// just to parse one file and execute it to the terminal
	tpl, err := template.ParseFiles("one.gohtml")
	if err != nil {
		panic(err)
	}
	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	panic(err)
	// }

	nf, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer nf.Close()
	err = tpl.Execute(nf, nil)
	if err != nil {
		panic(err)
	}
}
