package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type RedisRepository interface {
	GetPost(ctx context.Context, id uint) (models.Post, error)
	SetPost(ctx context.Context, post models.Post) error
	DeletePost(ctx context.Context, id uint) error
}
