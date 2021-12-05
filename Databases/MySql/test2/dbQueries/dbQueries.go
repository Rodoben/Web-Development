package dbQueries

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var err error

const (
	username = "root"
	password = "Ronald@123"
	hostname = "localhost:3306"
	dbname   = "ron"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func OpenConnection(res http.ResponseWriter) (string, error) {
	fmt.Fprintln(res, "I am At index")
	db, err = sql.Open("mysql", dsn("ron"))
	if err != nil {
		fmt.Fprintln(res, "Error connecting", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Fprintln(res, err)
	}

	fmt.Fprintln(res, "Connected to database", db)
	return "success", err
}

func MakeTable(tableName string, res http.ResponseWriter) (string, error) {
	s := "CREATE TABLE " + tableName + " (name VARCHAR(20));"
	stmt, err := db.Prepare(s)
	check(res, "Preparation failed", err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(res, "executionFailed", err)

	n, err := r.RowsAffected()
	check(res, "Rows failed", err)
	_ = n
	fmt.Println("CREATED TABLE ", tableName)

	return s, err
}

func InsertRec(res http.ResponseWriter) {
	stmt, err := db.Prepare(`INSERT INTO ron.customer VALUES ("James");`)
	check(res, "preparing error", err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(res, "stmt error", err)

	n, err := r.RowsAffected()
	check(res, "preparing error", err)

	fmt.Println(n)
}

func check(res http.ResponseWriter, e string, err error) {
	if err != nil {
		fmt.Fprintln(res, e, err)
	}
}
