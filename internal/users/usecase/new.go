package usecase

import (
	"context"

	application "github.com/DanTDM2003/search-api-docker-redis/internal/application/user"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	pkgPassword "github.com/DanTDM2003/search-api-docker-redis/pkg/password"
)

type Usecase interface {
	GetUsers(ctx context.Context, input GetUsersInput) (GetUsersOutput, error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (models.User, error)
	DeleteUser(ctx context.Context, id uint) error
	PromoteToAdmin(ctx context.Context, id uint) (models.User, error)
	DemoteToUser(ctx context.Context, id uint) (models.User, error)
	ChangePassword(ctx context.Context, input ChangePasswordInput) error
}

type impleUsecase struct {
	l               pkgLog.Logger
	repo            repository.Repository
	redisRepo       repository.RedisRepository
	passwordManager pkgPassword.Password
}

var _ Usecase = &impleUsecase{}
var _ application.UserUsecase = &impleUsecase{}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	reidsRepo repository.RedisRepository,
	passwordManager pkgPassword.Password,
) *impleUsecase {
	return &impleUsecase{
		l:               l,
		repo:            repo,
		redisRepo:       reidsRepo,
		passwordManager: passwordManager,
	}
}
