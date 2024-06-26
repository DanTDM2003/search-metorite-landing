package main

import "database/sql"

type PostgresConnection struct {
	db *sql.DB
}

func NewPostgresConnection() (*PostgresConnection, error) {
	connStr := "user=postgres dbname=postgres password=gosearchapi sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresConnection{
		db: db,
	}, nil
}
