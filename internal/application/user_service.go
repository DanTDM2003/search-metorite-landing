package application

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type UserUsecase interface {
	GetOneUser(ctx context.Context, input GetOneUserInput) (models.User, error)
	CreateUser(ctx context.Context, input CreateUserInput) (models.User, error)
}

type GetOneUserInput struct {
	ID    uint
	Email string
}

type CreateUserInput struct {
	Username string
	Email    string
	Password string
}
