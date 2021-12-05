package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func foo(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
	<form method="POST">
	<label for="fname">First name:</label><br><br>
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)
}
