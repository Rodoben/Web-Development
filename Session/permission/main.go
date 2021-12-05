package main

import (
	"crypto/x/crypto/bcrypt"
	"html/template"
	"net/http"
	uuid "satori/go.uuid"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	role     string
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)

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

	if u.role != "007" {
		http.Error(res, "007 to enter the bar", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(res, "bar.gohtml", u)

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
			http.Error(res, "Username Does not Exist", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "Password is wrong", http.StatusForbidden)
			return
		}
		sId := uuid.NewV4()
		c := &http.Cookie{Name: "session", Value: sId.String()}
		http.SetCookie(res, c)
		dbSession[c.Value] = un
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)

}

func signup(res http.ResponseWriter, req *http.Request) {
	var u user
	if AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		if _, ok := dbUser[un]; ok {
			http.Error(res, "Username already taken", http.StatusForbidden)
			return
		}
		sId := uuid.NewV4()
		c := &http.Cookie{Name: "session", Value: sId.String()}
		http.SetCookie(res, c)
		dbSession[c.Value] = un

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal Error Happened", http.StatusInternalServerError)
		}
		u = user{UserName: un, Password: bs, First: f, Last: l, role: r}
		dbUser[un] = u
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(res, "signup.gohtml", u)

}

func logout(res http.ResponseWriter, req *http.Request) {
	if !AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	delete(dbSession, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
