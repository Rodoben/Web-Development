package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dpgpic", dogpic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/bajrang.jpg">`)
}

func dogpic(w http.ResponseWriter, res *http.Request) {
	f, err := os.Open("bajrang.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	log.Println(w, res, f.Name(), fi.ModTime())
	http.ServeContent(w, res, f.Name(), fi.ModTime(), f)
}
