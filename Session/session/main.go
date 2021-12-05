package main

import (
	"fmt"
	"net/http"
	uuid "satori/go.uuid"
	"text/template"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUser = map[string]user{}
var dbSession = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		sId := uuid.NewV4()
		c = &http.Cookie{Name: "session", Value: sId.String()}
	}
	http.SetCookie(res, c)

	var u user
	if un, ok := dbSession[c.Value]; ok {
		u = dbUser[un]
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		u = user{UserName: un, First: fn, Last: ln}
		dbSession[c.Value] = un
		dbUser[un] = u
	}
	fmt.Println("dbSession", dbSession)
	fmt.Println("dbUser", dbUser)

	tpl.ExecuteTemplate(res, "index.gohtml", u)
}
func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("entered bar")
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSession[c.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUser[un]
	tpl.ExecuteTemplate(res, "bar.gohtml", u)

}
