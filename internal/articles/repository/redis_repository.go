package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type RedisRepository interface {
	GetArticle(ctx context.Context, slug string) (models.Article, error)
	SetArticle(ctx context.Context, article models.Article) error
	DeleteArticle(ctx context.Context, id uint) error
}
