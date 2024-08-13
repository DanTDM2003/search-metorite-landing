package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Usecase interface {
}

type impleUsecase struct {
	l     pkgLog.Logger
	repo  repository.Repository
	redis repository.RedisRepository
}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	redis repository.RedisRepository,
) Usecase {
	return &impleUsecase{
		l:     l,
		repo:  repo,
		redis: redis,
	}
}
