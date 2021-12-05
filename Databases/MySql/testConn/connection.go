package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "Ronald@123"
	hostname = "localhost:3306"
	dbname   = "ron"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

var db *sql.DB
var err error

func main() {
	openConnection()
	start()
	insertCustomer()
	update()
	//amigos()

	//read()
	//create()
}

func start() {
	stmt, err := db.Prepare(`show databases`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)
	fmt.Println("CREATED TABLE customer", r)
	n, err := r.RowsAffected()

	check(err)

	fmt.Println("CREATED TABLE customer", n)
}

func openConnection() {
	db, err = sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Println("connection failed")
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("ping failed")
	}
	fmt.Println("established", db)
}

func read() {
	rows, err := db.Query(`SELECT * FROM ron.test1;`)
	if err != nil {
		fmt.Println("scan failed")
	}
	defer rows.Close()

	var name string
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, name)
	}
}
func create() {

	stmt, err := db.Prepare(`CREATE TABLE ron.customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Println("CREATED TABLE customer", n)
}

func insertCustomer() {

	stmt, err := db.Prepare(`INSERT INTO ron.customer VALUES ("James");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Println(n)
}

func update() {
	stmt, err := db.Prepare(`UPDATE ron.test1 SET name="Jimmy" WHERE index = 1;`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Println(n)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
