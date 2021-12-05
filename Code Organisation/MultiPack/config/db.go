package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	UserName = "root"
	Password = "Ronald@123"
	hostName = "localhost:3306"
	dbName   = "library"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", UserName, Password, hostName, dbName)
}

func init() {

	var err error

	DB, err = sql.Open("mysql", dsn())
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

}
