package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "dbname=racehub sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
