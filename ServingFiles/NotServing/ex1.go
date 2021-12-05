package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset= utf-8")
	c := "Rodo"
	s := "<!--not serving from our server-->" +
		"<h1>" + c + "</h1>" +
		"<img src=" +
		"https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg" + " >"
	fmt.Println(s)
	io.WriteString(w, s)

}
