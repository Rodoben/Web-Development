package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type hotdog int

var tpl *template.Template

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method      string
		URL         *url.URL
		Submissions map[string][]string
		Header      http.Header
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
	}
	tpl.ExecuteTemplate(res, "index.gohtml", data)

}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
