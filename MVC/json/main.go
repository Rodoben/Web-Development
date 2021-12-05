package main

import (
	"encoding/json"
	"fmt"
	"json/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	http.ListenAndServe(":8080", r)
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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
	res.Header().Set("Content-Type", "text/html; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(s))
}

func getUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "JAmes Bond",
		Gender: "male",
		Age:    32,
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
