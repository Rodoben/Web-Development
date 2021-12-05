package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	err1 := tpl.ExecuteTemplate(os.Stdout, "file1.gohtml", 45)
	if err1 != nil {
		log.Fatalln(err1)
	}

	err := tpl.ExecuteTemplate(os.Stdout, "file2.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

}
