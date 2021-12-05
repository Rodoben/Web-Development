package controllers

import (
	"ex3/model"
	"ex3/session"
	"net/http"
	uuid "satori/go.uuid"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (c Controller) SignUP(res http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var u model.User

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		if _, ok := session.Users[un]; ok {
			http.Error(res, "Username Already Taken", http.StatusForbidden)
			return
		}

		sID := uuid.NewV4()
		ck := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		ck.MaxAge = session.Length
		http.SetCookie(res, ck)
		session.Sessions[ck.Value] = model.Session{
			UserName:     un,
			LastActivity: time.Now(),
		}

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		u = model.User{un, bs, f, l, r}
		session.Users[un] = u
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return

	}
	session.Show()
	c.tpl.ExecuteTemplate(res, "signup.gohtml", u)

}

func (c Controller) Login(res http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var u model.User

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		u, ok := session.Users[un]
		if !ok {
			http.Error(res, "Username and password do not match", http.StatusSeeOther)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "USername or Password mismatched", http.StatusForbidden)
			return
		}
		sID := uuid.NewV4()
		ck := &http.Cookie{Name: "session", Value: sID.String()}
		ck.MaxAge = session.Length
		http.SetCookie(res, ck)
		session.Sessions[ck.Value] = model.Session{un, time.Now()}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	session.Show()
	c.tpl.ExecuteTemplate(res, "login.gohtml", u)
}

func (c Controller) Logout(res http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	ck, _ := req.Cookie("session")
	delete(session.Sessions, ck.Value)
	ck = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, ck)

	if time.Now().Sub(session.LastCleaned) > (time.Second * 30) {
		session.Clean()
	}
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}
