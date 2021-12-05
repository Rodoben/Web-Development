package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

const (
	username = "root"
	password = "Ronald@123"
	hostName = "localhost:3306"
	dbName   = "library"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostName, dbName)
}

func init() {
	db, err = sql.Open("mysql", dsn())
	if err != nil {
		fmt.Println(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You're Connected to database")
}
