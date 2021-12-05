// stroting on the server

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	var s string

	fmt.Println(req.Method, http.MethodPost)
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("file1")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fmt.Println("/file:", f, "/header:", h, "/err:", err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
		fmt.Println(s)
		fmt.Println("creating file")
		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="file1">
	<input type="submit">
	</form>
	<br>`+s)

}
