package postgres

import (
	"database/sql"
	"fmt"

	"github.com/DanTDM2003/search-api-docker-redis/config"
	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	DB *sql.DB
}

func Connect(cfg config.PostgresConfig) (*PostgresConnection, error) {
	connStr := fmt.Sprintf(
		`
			host=%s
			port=%s
			user=%s
			password=%s
			dbname=%s
			sslmode=%s
		`,
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresConnection{
		DB: db,
	}, nil
}

func Close(db *sql.DB) {
	db.Close()
}
