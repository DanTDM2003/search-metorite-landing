package database

import (
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"gorm.io/gorm"
)

type impleRepository struct {
	log pkgLog.Logger
	db  *gorm.DB
}

func New(l pkgLog.Logger, db *gorm.DB) *impleRepository {
	return &impleRepository{
		log: l,
		db:  db,
	}
}
