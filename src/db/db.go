package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func LoadDB() (*DB, error) {
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=postgres dbname=rinha sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}
