package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type RedisRepository interface {
	GetUser(ctx context.Context, id uint) (models.User, error)
	SetUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id uint) error
}
