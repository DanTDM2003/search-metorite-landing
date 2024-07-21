package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type Repository interface {
	GetOneUser(ctx context.Context, opt GetOneUserOptions) (models.User, error)
	GetUsers(ctx context.Context, opt GetUsersOptions) ([]models.User, paginator.Paginator, error)
	CreateUser(ctx context.Context, opt CreateUserOptions) (models.User, error)
	UpdateUser(ctx context.Context, opt UpdateUserOptions, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, id uint) error
}
