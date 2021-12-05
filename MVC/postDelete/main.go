package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"postDelete/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user", createUser)
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
</body>
</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Ronald",
		Age:    32,
		Gender: "male",
		Id:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(res, "%s\n", uj)
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(req.Body).Decode(&u)
	u.Id = "007"
	uj, _ := json.Marshal(u)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(res, "%s\n", uj)
}
func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
