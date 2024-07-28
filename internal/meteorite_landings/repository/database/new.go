package database

import (
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"gorm.io/gorm"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
)

type impleRepository struct {
	l  pkgLog.Logger
	db *gorm.DB
}

func New(
	l pkgLog.Logger,
	db *gorm.DB,
) repository.Repository {
	return &impleRepository{
		l:  l,
		db: db,
	}
}
