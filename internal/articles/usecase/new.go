package usecase

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Usecase interface {
	GetArticles(ctx context.Context, input GetArticlesInput) (GetArticlesOutput, error)
	GetOneArticle(ctx context.Context, input GetOneArticleInput) (models.Article, error)
	CreateArticle(ctx context.Context, input CreateArticleInput) (models.Article, error)
	UpdateArticle(ctx context.Context, input UpdateArticleInput) (models.Article, error)
	DeleteArticle(ctx context.Context, id uint) error
}

type impleUsecase struct {
	l       pkgLog.Logger
	repo    repository.Repository
	redis   repository.RedisRepository
	locator *serviceLocator.ServiceLocator
}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	redis repository.RedisRepository,
) Usecase {
	return &impleUsecase{
		l:       l,
		repo:    repo,
		redis:   redis,
		locator: serviceLocator.GetServiceLocator(),
	}
}
