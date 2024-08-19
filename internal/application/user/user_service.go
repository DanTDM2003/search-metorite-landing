package user

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type UserUsecase interface {
	GetOneUser(ctx context.Context, input GetOneUserInput) (models.User, error)
	CreateUser(ctx context.Context, input CreateUserInput) (models.User, error)
}