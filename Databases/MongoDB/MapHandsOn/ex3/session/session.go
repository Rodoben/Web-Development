package session

import (
	"ex3/model"
	"fmt"
	"net/http"
	uuid "satori/go.uuid"
	"time"
)

const Length int = 30

var Users = map[string]model.User{}
var Sessions = map[string]model.Session{}
var LastCleaned time.Time = time.Now()

func GetUser(res http.ResponseWriter, req *http.Request) model.User {
	ck, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		ck = &http.Cookie{Name: "session", Value: sID.String()}
	}
	ck.MaxAge = Length
	http.SetCookie(res, ck)

	var u model.User

	if s, ok := Sessions[ck.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
		u = Users[s.UserName]
	}
	return u

}

func AlreadyLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	ck, err := req.Cookie("session")
	if err != nil {
		return false
	}

	s, ok := Sessions[ck.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
	}

	_, ok = Users[s.UserName]
	ck.MaxAge = Length
	http.SetCookie(res, ck)
	return ok

}

func Clean() {
	fmt.Println("BEFORE CLEAN")
	Show()

	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	Show()
}

func Show() {
	fmt.Println("*******")
	for k, v := range Sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
