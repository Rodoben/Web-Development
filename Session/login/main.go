package main

import (
	"crypto/x/crypto/bcrypt"
	"net/http"
	uuid "satori/go.uuid"
	"text/template"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbUser = map[string]user{}
var dbSession = map[string]string{}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		if _, ok := dbUser[un]; ok {
			http.Error(res, "Username already exist", http.StatusForbidden)
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
		}

		u = user{un, bs, f, l}
		dbUser[un] = u
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(res, "signup.gohtml", u)

}

func login(res http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		u, ok := dbUser[un]
		if !ok {
			http.Error(res, "User does not exist", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "Password entered is wrong", http.StatusForbidden)
			return
		}
		sId := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(res, c)
		dbSession[c.Value] = un
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}
