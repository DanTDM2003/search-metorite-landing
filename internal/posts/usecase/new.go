package usecase

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Usecase interface {
	GetPosts(ctx context.Context, input GetPostsInput) (GetPostsOutput, error)
	GetOnePost(ctx context.Context, input GetOnePostInput) (models.Post, error)
	CreatePost(ctx context.Context, input CreatePostInput) (models.Post, error)
	UpdatePost(ctx context.Context, input UpdatePostInput) (models.Post, error)
	DeletePost(ctx context.Context, id uint) error
}

type impleUsecase struct {
	l         pkgLog.Logger
	repo      repository.Repository
	redisRepo repository.RedisRepository
}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	redisRepo repository.RedisRepository,
) Usecase {
	return &impleUsecase{
		l:         l,
		repo:      repo,
		redisRepo: redisRepo,
	}
}
