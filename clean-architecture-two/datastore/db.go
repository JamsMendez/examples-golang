package datastore

import (
	"database/sql"
	"log"
)

func NewDB() *sql.DB {
	// Connection DB
	db, err := sql.Open("mysql", "string connection")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
