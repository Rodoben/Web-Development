package main

import (
	"fmt"
	"net/http"
	uuid "satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session-id")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{Name: "session-id", Value: id.String()}
	}
	http.SetCookie(res, c)
	fmt.Println(c)
}
