package postgres

import (
	"database/sql"

	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
)

type impleRepository struct {
	l  pkgLog.Logger
	db *sql.DB
}

var _ repository.Repository = &impleRepository{}

func New(l pkgLog.Logger, db *sql.DB) repository.Repository {
	return &impleRepository{
		l:  l,
		db: db,
	}
}
