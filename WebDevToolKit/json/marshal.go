package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>FOO</title>
	</head>
	<body>
	You are at foo
	</body>
	</html>`
	res.Write([]byte(s))
}

func marshal(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Benjamin",
		Items: []string{"suit", "gun", "ronald"},
	}
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	res.Write(j)
}

func encode(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Benjamin",
		Items: []string{"suit", "gun", "ronald"},
	}
	err := json.NewEncoder(res).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
