package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		c = &http.Cookie{Name: "my-cookie", Value: "0"}
	}

	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	c.Value = strconv.Itoa(count)
	http.SetCookie(res, c)
	io.WriteString(res, c.Value)
}
