package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type hotdog int

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))

}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)

}

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Form,
	}
	tpl.ExecuteTemplate(res, "index.gohtml", data)

}
