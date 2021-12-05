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
	sages := map[string]string{

		"abc":    "abc",
		"nnn":    "cccc",
		"cckk":   "kkcc",
		"kkllcc": "ccckdk",
	}

	//err := tpl.Execute(os.Stdout, sages)
	//if err != nil {
	///	panic(err)
	//}

	nf, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(nf, sages)
	if err != nil {
		panic(err)
	}

}
