// The preffered method for serving a file is http.fileServer
// Allows to read everything from a directory

package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dogg", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(res, `<img src="/bajrang.jpg">`)
}
