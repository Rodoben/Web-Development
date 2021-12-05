package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*"))

}

func main() {
	xs := []string{"zero", "one", "two", "three"}
	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Benjamin",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xs)
	if err != nil {
		log.Panic(err)
	}
	err1 := tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", data)
	if err1 != nil {
		log.Panic(err1)
	}
}
