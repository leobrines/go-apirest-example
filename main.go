package main

import(
	//"github.com/go-sql-driver/mysql"
	"database/sql"
)

var db *sql.DB

func main () {
	webserver := new(Webserver)
	webserver.Configure()
	webserver.Start()
}

