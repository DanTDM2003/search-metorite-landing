package articles

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type Usecase interface {
	GetArticles(ctx context.Context, input GetArticlesInput) (GetArticlesOutput, error)
	GetOneArticle(ctx context.Context, input GetOneArticleInput) (models.Article, error)
	CreateArticle(ctx context.Context, input CreateArticleInput) (models.Article, error)
	UpdateArticle(ctx context.Context, input UpdateArticleInput) (models.Article, error)
	DeleteArticle(ctx context.Context, id uint) error
}
