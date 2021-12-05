// func HandleFunc (pattern string, handler func(ResponseWriter,*Request))
package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}
func main() {
	http.HandleFunc("/cat", c)
	http.HandleFunc("/dog", d)
	http.ListenAndServe(":8080", nil)
}
