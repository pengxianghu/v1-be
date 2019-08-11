package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:990085@tcp(119.23.70.24:25003)/v1-be?parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	err = dbConn.Ping()
	if err != nil {
		panic(err.Error())
	}
}