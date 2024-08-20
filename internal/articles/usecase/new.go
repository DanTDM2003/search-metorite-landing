package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles"
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type impleUsecase struct {
	l       pkgLog.Logger
	repo    repository.Repository
	redis   repository.RedisRepository
	locator *serviceLocator.ServiceLocator
}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	redis repository.RedisRepository,
) articles.Usecase {
	return &impleUsecase{
		l:       l,
		repo:    repo,
		redis:   redis,
		locator: serviceLocator.GetServiceLocator(),
	}
}
