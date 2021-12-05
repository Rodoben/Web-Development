package main

import (
	"database/sql"
	"net/http"
	"text/template"
	"twopack/models"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/create", booksCreateForm)
	http.HandleFunc("/books/create/process", booksCreateProcess)
	http.HandleFunc("/books/update", booksUpdateForm)
	http.HandleFunc("/books/update/process", booksUpdateProcess)
	http.HandleFunc("/books/delete/process", booksDeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/books", http.StatusSeeOther)
}

func booksIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, http.StatusText(505), http.StatusMethodNotAllowed)
		return
	}
	bks, err := models.AllBooks()
	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(res, "books.gohtml", bks)
}

func booksShow(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, http.StatusText(500), http.StatusMethodNotAllowed)
		return
	}

	bks, err := models.OneBook(req)
	if err != nil {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	tpl.ExecuteTemplate(res, "show.gohtml", bks)
}

func booksCreateForm(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "create.gohtml", nil)
}

func booksCreateProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	bk, err := models.PutBook(req)
	if err != nil {
		http.Error(res, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	tpl.ExecuteTemplate(res, "created.gohtml", bk)
}

func booksUpdateForm(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := models.OneBook(req)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(res, req)
		return
	case err != nil:
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(res, "update.gohtml", bk)
}

func booksUpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := models.UpdateBook(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}

	tpl.ExecuteTemplate(w, "updated.gohtml", bk)
}

func booksDeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := models.DeleteBook(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
