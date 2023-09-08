package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	connStr := "host=db port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}