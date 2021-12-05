// http. http.StripPrefix - used for striping.

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)

	http.Handle("/resources", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
}

func dog(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(res, `<img src="/bajrang.jpg">`)

}
