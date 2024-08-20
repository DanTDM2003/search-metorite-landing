package users

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type Usecase interface {
	GetUsers(ctx context.Context, input GetUsersInput) (GetUsersOutput, error)
	GetOneUser(ctx context.Context, input GetOneUserInput) (models.User, error)
	CreateUser(ctx context.Context, input CreateUserInput) (models.User, error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (models.User, error)
	DeleteUser(ctx context.Context, id uint) error
	PromoteToAdmin(ctx context.Context, id uint) (models.User, error)
	DemoteToUser(ctx context.Context, id uint) (models.User, error)
	ChangePassword(ctx context.Context, input ChangePasswordInput) error
}
