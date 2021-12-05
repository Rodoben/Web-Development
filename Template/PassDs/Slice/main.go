package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("one.gohtml"))
}

func main() {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

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
		fmt.Println("coudnt crete a file")
	}

}
