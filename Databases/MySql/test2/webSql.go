package main

import (
	"databaseQuery/test2/dbQueries"
	"fmt"

	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/close", close)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/drop", drop)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	o, err := dbQueries.OpenConnection(res)
	if err != nil {
		fmt.Fprintln(res, "error opening conn", err)
	}
	fmt.Fprintln(res, o)
}
func close(res http.ResponseWriter, req *http.Request) {

}
func create(res http.ResponseWriter, req *http.Request) {
	tn := "mars"
	m, err := dbQueries.MakeTable(tn, res)
	if err != nil {
		fmt.Println(res, "error", err)
	}
	fmt.Fprintln(res, m)
}
func insert(res http.ResponseWriter, req *http.Request) {
	dbQueries.InsertRec(res)
}
func update(res http.ResponseWriter, req *http.Request) {

}
func drop(res http.ResponseWriter, req *http.Request) {

}
func delete(res http.ResponseWriter, req *http.Request) {

}
