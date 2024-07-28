package database

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"gorm.io/gorm"
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
