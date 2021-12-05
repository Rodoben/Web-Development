package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-cookie",
		Value: "x123xcv",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN , CHECK BROWSER")
	fmt.Fprintln(res, "in chrome got ot dev tools/application/cookie")
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(res, "YOur Cookie:", c)
}
