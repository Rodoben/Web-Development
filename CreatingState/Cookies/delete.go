package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/read", read)
	http.HandleFunc("/set", set)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("index")
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{Name: "ronald", Value: "btggnb"})
	fmt.Println("COKKIE SET")
}
func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("ronald")
	if err != nil {
		http.Redirect(res, req, "/", 303)
	}
	fmt.Println("your Cokkie", c)
}
func expire(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("ronald")
	if err != nil {
		http.Redirect(res, req, "/", 303)
		return
	}

	c.MaxAge = -1
	http.SetCookie(res, c)

}
