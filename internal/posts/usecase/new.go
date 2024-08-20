package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts"
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type impleUsecase struct {
	l         pkgLog.Logger
	repo      repository.Repository
	redisRepo repository.RedisRepository
	locator   *serviceLocator.ServiceLocator
}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	redisRepo repository.RedisRepository,
) posts.Usecase {
	return &impleUsecase{
		l:         l,
		repo:      repo,
		redisRepo: redisRepo,
		locator:   serviceLocator.GetServiceLocator(),
	}
}
