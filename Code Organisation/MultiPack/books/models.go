package books

import (
	"MultiPack/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float64
}

func AllBooks() ([]Book, error) {
	rows, err := config.DB.Query("SELECT * FROM library.books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bks := make([]Book, 0)

	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	return bks, nil
}

func OneBook(req *http.Request) (Book, error) {
	bk := Book{}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("BAD REQUEST")
	}
	var s string = `SELECT * FROM library.books WHERE isbn = "` + isbn + `"`
	row := config.DB.QueryRow(s)
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func InsertBook(req *http.Request) (Book, error) {
	bk := Book{}
	bk.Isbn = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("author")
	p := req.FormValue("price")
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete.")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("Price must be a number")
	}
	bk.Price = float64(f64)

	stmt, err := config.DB.Prepare("INSERT into library.books (isbn,title,author,price) VALUES (?,?,?,?)")
	if err != nil {
		return bk, errors.New("Error in Prepartion")
	}
	r, err := stmt.Exec(bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, errors.New("Execution Error")
	}
	n, err := r.RowsAffected()
	if err != nil {
		fmt.Println(n)
	}
	return bk, nil
}

func UpdateBook(req *http.Request) (Book, error) {
	bk := Book{}
	bk.Isbn = req.FormValue("isbn")
	bk.Title = req.FormValue("title")
	bk.Author = req.FormValue("author")
	p := req.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("Fields cant be empty")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("price can only be number")
	}
	bk.Price = float64(f64)
	var s string = `UPDATE library.books SET isbn="` + bk.Isbn + `",title="` + bk.Title + `",author="` + bk.Author + `",price=` + p + ` where isbn= "` + bk.Isbn + `"`
	fmt.Println(s)
	stmt, err := config.DB.Prepare(s)
	if err != nil {
		return bk, errors.New("Error in preparation")
	}
	r, err := stmt.Exec()
	if err != nil {
		return bk, errors.New("Error in Execut")
	}
	n, err := r.RowsAffected()
	if err != nil {
		return bk, errors.New("Must be a number not string")
	}
	fmt.Println(n)
	return bk, nil

}

func DeleteBook(req *http.Request) error {
	isbn := req.FormValue("isbn")

	if isbn == "" {
		return errors.New("400. Bad Request.")
	}
	var s string = `DELETE FROM library.books WHERE isbn ="` + isbn + `"`
	_, err := config.DB.Exec(s)
	if err != nil {
		return errors.New("Internal Error")
	}
	return nil
}
