package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Connection *sql.DB

func InitDB(connStr string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	Connection = db
}