package main

import (
	"fmt"
	"net/http"
	uuid "satori/go.uuid"
	"time"
)

func getUser(res http.ResponseWriter, req *http.Request) user {
	c, err := req.Cookie("session")
	if err != nil {
		sId := uuid.NewV4()
		c = &http.Cookie{Name: "session", Value: sId.String()}

	}
	c.MaxAge = sessionLength
	http.SetCookie(res, c)

	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

func AlreadyLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	c.MaxAge = sessionLength
	http.SetCookie(res, c)
	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 1) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	showSessions()
}

func showSessions() {
	fmt.Println("****************")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
