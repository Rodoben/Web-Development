package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", formdata)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func formdata(res http.ResponseWriter, req *http.Request) {
	p := req.FormValue("first")
	l := req.FormValue("last")
	q := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(res, "index.gohtml", person{FirstName: p, LastName: l, Subscribed: q})
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)

	}
}
