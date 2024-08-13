package usecase

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Usecase interface {
	GetUsers(ctx context.Context, input GetUsersInput) (GetUsersOutput, error)
	GetOneUser(ctx context.Context, input GetOneUserInput) (models.User, error)
	CreateUser(ctx context.Context, input CreateUserInput) (models.User, error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (models.User, error)
	DeleteUser(ctx context.Context, id uint) error
	SignIn(ctx context.Context, input SignInInput) (SignInOutput, error)
	SignUp(ctx context.Context, input SignUpInput) (SignUpOutput, error)
	PromoteToAdmin(ctx context.Context, id uint) (models.User, error)
	DemoteToUser(ctx context.Context, id uint) (models.User, error)
}

type impleUsecase struct {
	l          pkgLog.Logger
	repo       repository.Repository
	redisRepo  repository.RedisRepository
	jwtManager pkgJWT.Manager
}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	reidsRepo repository.RedisRepository,
	jwtManager pkgJWT.Manager,
) Usecase {
	return &impleUsecase{
		l:          l,
		repo:       repo,
		redisRepo:  reidsRepo,
		jwtManager: jwtManager,
	}
}
