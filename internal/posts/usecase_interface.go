package posts

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type Usecase interface {
	GetPosts(ctx context.Context, input GetPostsInput) (GetPostsOutput, error)
	GetOnePost(ctx context.Context, input GetOnePostInput) (models.Post, error)
	CreatePost(ctx context.Context, input CreatePostInput) (models.Post, error)
	UpdatePost(ctx context.Context, input UpdatePostInput) (models.Post, error)
	DeletePost(ctx context.Context, id uint) error
}
