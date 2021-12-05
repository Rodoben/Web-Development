package main

import (
	"net/http"
	uuid "satori/go.uuid"
)

func getUser(res http.ResponseWriter, req *http.Request) user {
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
	return u
}

func AlreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSession[c.Value]
	_, ok := dbUser[un]
	return ok
}
