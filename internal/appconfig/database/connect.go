package database

import (
	"database/sql"
	"fmt"

	"github.com/DanTDM2003/search-api-docker-redis/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnection struct {
	DB *gorm.DB
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

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgresConnection{
		DB: gormDB,
	}, nil
}

func Close(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		return
	}
	conn.Close()
}
