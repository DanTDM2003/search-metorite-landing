package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/users"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	pkgPassword "github.com/DanTDM2003/search-api-docker-redis/pkg/password"
)

type impleUsecase struct {
	l               pkgLog.Logger
	repo            repository.Repository
	redisRepo       repository.RedisRepository
	passwordManager pkgPassword.Password
}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	reidsRepo repository.RedisRepository,
	passwordManager pkgPassword.Password,
) users.Usecase {
	return &impleUsecase{
		l:               l,
		repo:            repo,
		redisRepo:       reidsRepo,
		passwordManager: passwordManager,
	}
}
