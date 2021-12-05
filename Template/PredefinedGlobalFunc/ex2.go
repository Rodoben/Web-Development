package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*"))
}

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {
	u1 := user{Name: "Ronald", Motto: "Work Hard", Admin: false}
	u2 := user{Name: "", Motto: "Work Hard", Admin: false}
	u3 := user{Name: "Neha", Motto: "Work Hard", Admin: true}
	u4 := user{Name: "Peha", Motto: "Work Hard", Admin: true}

	users := []user{u1, u2, u3, u4}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl3.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}

	g1 := struct {
		Score1 int
		Score2 int
	}{
		5,
		7,
	}
	err1 := tpl.ExecuteTemplate(os.Stdout, "tpl4.gohtml", g1)
	if err1 != nil {
		log.Fatalln(err1)
	}

}
