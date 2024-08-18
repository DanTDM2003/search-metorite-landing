package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type Repository interface {
	GetArticles(ctx context.Context, opt GetArticlesOptions) ([]models.Article, paginator.Paginator, error)
	GetOneArticle(ctx context.Context, opt GetOneArticleOptions) (models.Article, error)
	CreateArticle(ctx context.Context, opt CreateArticleOptions) (models.Article, error)
	UpdateArticle(ctx context.Context, opt UpdateArticleOptions, article models.Article) (models.Article, error)
	DeleteArticle(ctx context.Context, id uint) error
}
