package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset= utf-8")
	s := "<!--image doesn't serve-->" +
		"<img src=" + "\bajrang.jpg" + ">"
	io.WriteString(w, s)

}
