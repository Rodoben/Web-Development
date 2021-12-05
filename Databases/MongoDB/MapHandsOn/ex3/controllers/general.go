package controllers

import (
	"ex3/session"
	"html/template"
	"net/http"
)

type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(res http.ResponseWriter, req *http.Request) {
	u := session.GetUser(res, req)
	session.Show()
	c.tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func (c Controller) Bar(res http.ResponseWriter, req *http.Request) {
	u := session.GetUser(res, req)
	if !session.AlreadyLoggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	if u.Role != "007" {
		http.Error(res, "You Must be 007 to enter the Bar", http.StatusForbidden)
	}

	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(res, "bar.gohtml", u)
}
