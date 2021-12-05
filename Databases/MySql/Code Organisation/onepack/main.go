package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float64
}

var db *sql.DB
var tpl *template.Template

const (
	username = "root"
	password = "Ronald@123"
	dbName   = "library"
	hostName = "localhost:3306"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostName, dbName)
}

func init() {
	var err error

	db, err = sql.Open("mysql", dsn())
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/books", bookIndex)
	http.HandleFunc("/books/show", bookShow)
	http.HandleFunc("/books/create", bookCreateForm)
	http.HandleFunc("/books/create/process", bookCreateProcess)
	http.HandleFunc("/books/update", booksUpdateForm)
	http.HandleFunc("/books/update/process", booksUpdateProcess)
	http.HandleFunc("/books/delete/process", bookDeleteProcess)
	http.ListenAndServe(":8080", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/books", http.StatusSeeOther)
}

func bookIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, http.StatusText(505), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("select * from library.books")
	if err != nil {
		http.Error(res, "Unable to fetch data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bks := []Book{}

	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			http.Error(res, "Unable to fetch data", http.StatusInternalServerError)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(res, http.StatusText(500), 500)
		return
	}

	tpl.ExecuteTemplate(res, "books.gohtml", bks)

}

func bookShow(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "not aloowed", http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	fmt.Println(isbn)
	if isbn == "" {
		http.Error(res, "bad request", http.StatusBadRequest)
		return
	}

	row := db.QueryRow(`select * from library.books where isbn ="` + isbn + `"`)

	bk := Book{}
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(res, req)
		return
	case err != nil:
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(res, "show.gohtml", bk)
}

func booksUpdateForm(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "not aloowed", http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	fmt.Println(isbn)
	if isbn == "" {
		http.Error(res, "bad request", http.StatusBadRequest)
		return
	}

	row := db.QueryRow(`select * from library.books where isbn ="` + isbn + `"`)

	bk := Book{}
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
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

func booksUpdateProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, http.StatusText(405), http.StatusSeeOther)
		return
	}
	bk := Book{}
	bk.Isbn = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("author")
	p := req.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	bk.Price = float64(f64)
	fmt.Println("values", bk.Isbn, bk.Author, bk.Price, bk.Title)
	var s string = `UPDATE library.books SET isbn="` + bk.Isbn + `",title="` + bk.Title + `",author="` + bk.Author + `",price=` + p + ` where isbn= "` + bk.Isbn + `"`
	fmt.Println(s)
	stmt, err := db.Prepare(s)
	if err != nil {
		fmt.Println("query error", err)
	}
	r, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
	}
	n, err := r.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	tpl.ExecuteTemplate(res, "updated.gohtml", bk)
}

func bookCreateForm(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "create.gohtml", nil)
}

func bookCreateProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk := Book{}
	bk.Isbn = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("author")
	p := req.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	f64, err := strconv.ParseFloat(p, 34)
	if err != nil {
		http.Error(res, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float64(f64)

	//insert a value
	fmt.Println(bk)

	stmt, err := db.Prepare(`INSERT INTO library.books (isbn, title, author, price) VALUES(?,?,?,?);`)
	if err != nil {
		fmt.Println("error in preparing", err)
	}
	r, err := stmt.Exec(bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		fmt.Println(err)
	}
	n, err := r.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	// confirm insertion
	tpl.ExecuteTemplate(res, "created.gohtml", bk)
}

func bookDeleteProcess(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, http.StatusText(505), http.StatusMethodNotAllowed)
		return
	}

	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	_, err := db.Exec(`DELETE FROM library.books WHERE isbn="` + isbn + `"`)
	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, "/books", http.StatusSeeOther)

}
