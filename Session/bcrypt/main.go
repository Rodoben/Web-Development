package main

import (
	"crypto/x/crypto/bcrypt"
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
	u := getUser(req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("entered bar")
	u := getUser(req)
	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	var u user
	fmt.Println("entered signup")
	if AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		fmt.Println("entered form")
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		if _, ok := dbUser[un]; ok {
			http.Error(res, "USerNAme Already Taken", http.StatusForbidden)
			return
		}
		sId := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}

		http.SetCookie(res, c)
		dbSession[c.Value] = un

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println(bs)
		u = user{un, string(bs), f, l}
		dbUser[un] = u

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", u)

}
