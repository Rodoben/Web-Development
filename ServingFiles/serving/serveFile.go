//http.ServeFile(response, *request, filename)

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dogpic", dogpic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">	`)
}

func dogpic(w http.ResponseWriter, res *http.Request) {
	http.ServeFile(w, res, "bajrang.jpg")
}
