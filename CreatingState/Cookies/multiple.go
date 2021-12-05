package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.ListenAndServe(":8080", nil)
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-Cookie")
	if err != nil {
		http.Error(res, err.Error(), 404)
	}
	fmt.Fprint(res, "your Cookie", c)

	c1, err := req.Cookie("general")
	if err != nil {
		http.Error(res, err.Error(), 404)
	}
	fmt.Fprint(res, "your Cookie", c1)

	c2, err := req.Cookie("specific")
	if err != nil {
		http.Error(res, err.Error(), 404)
	}
	fmt.Fprint(res, "your Cookie", c2)

}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "my-Cookie",
		Value: "ffbfgbg",
	})
	fmt.Fprintln(res, "COOKIE WRITTEN")
	fmt.Fprintln(res, "In Chrome go to dev tools/app/cookkie")

}

func abundance(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "general",
		Value: "brynynyt",
	})
	http.SetCookie(res, &http.Cookie{
		Name:  "specific",
		Value: "srgtbtbtb",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN")
	fmt.Fprintln(res, "Check THe location")

}
