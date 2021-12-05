package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foo:", req.Method)
}
func bar(res http.ResponseWriter, req *http.Request) {

	fmt.Println("Your request method at bar:", req.Method, "\n\n")

	http.Redirect(res, req, "/", 301)

}
