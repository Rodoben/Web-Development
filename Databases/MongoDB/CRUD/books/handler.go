package books

import (
	"CRUD/config"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, http.StatusText(500), http.StatusBadGateway)
		return
	}

	bks, err := AllBooks()
	if err != nil {
		http.Error(res, http.StatusText(405), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(res, "books.gohtml", bks)

}
func Show(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(500), http.StatusBadGateway)
		return
	}

	bk, err := OneBook(req)
	if err != nil {
		http.Error(res, http.StatusText(405), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(res, "show.gohtml", bk)

}
func Create(res http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(res, "create.gohtml", nil)
}
func CreateProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, http.StatusText(500), http.StatusBadGateway)
		return
	}

	bk, err := PutBook(req)
	if err != nil {
		http.Error(res, http.StatusText(405), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(res, "created.gohtml", bk)

}
func Update(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	bk, err := OneBook(req)
	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(res, "update.gohtml", bk)
}
func UpdateProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := UpdateBook(req)
	if err != nil {
		http.Error(res, http.StatusText(406), http.StatusBadRequest)
		return
	}

	config.TPL.ExecuteTemplate(res, "updated.gohtml", bk)
}
func DeleteProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteBooks(req)
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(res, req, "/books", http.StatusSeeOther)
}
