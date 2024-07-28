package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type UseCase interface {
}

type impleUseCase struct {
	l     pkgLog.Logger
	repo  repository.Repository
	redis repository.RedisRepository
}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	redis repository.RedisRepository,
) UseCase {
	return &impleUseCase{
		l:     l,
		repo:  repo,
		redis: redis,
	}
}
