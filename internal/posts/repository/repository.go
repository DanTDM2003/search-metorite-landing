package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type Repository interface {
	GetPosts(ctx context.Context, opt GetPostsOptions) ([]models.Post, paginator.Paginator, error)
	GetOnePost(ctx context.Context, opt GetOnePostOptions) (models.Post, error)
	CreatePost(ctx context.Context, opt CreatePostOptions) (models.Post, error)
	UpdatePost(ctx context.Context, opt UpdatePostOptions, post models.Post) (models.Post, error)
	DeletePost(ctx context.Context, id uint) error
	DeletePosts(ctx context.Context, opt DeletePostsOptions) error
}
