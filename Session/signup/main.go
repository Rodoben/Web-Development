package main

import (
	"fmt"
	"net/http"
	uuid "satori/go.uuid"
	"text/template"
)

type user struct {
	UserName string
	Password string
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

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Entered index")
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Entered bar")
	u := getUser(res, req)
	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "bar.gohtml", u)

}

func signup(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Entered signup")
	if AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	fmt.Println("after signup")
	if req.Method == http.MethodPost {
		fmt.Println("inside forms")
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		if _, ok := dbUser[un]; ok {
			http.Error(res, "USerName Already Taken", http.StatusForbidden)
			return
		}
		sId := uuid.NewV4()
		c := &http.Cookie{Name: "session", Value: sId.String()}
		http.SetCookie(res, c)
		dbSession[c.Value] = un
		u := user{un, p, f, l}
		dbUser[un] = u
		fmt.Println(dbSession)
		fmt.Println(dbUser)
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(res, "signup.gohtml", nil)

}
