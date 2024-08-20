package posts

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type GetPostsFilter struct {
	AuthorID uint
}

type GetPostsInput struct {
	GetPostsFilter
	paginator.PaginatorQuery
}

type GetPostsOutput struct {
	Posts []models.Post
	paginator.Paginator
}

type GetOnePostInput struct {
	ID       uint
	AuthorID uint
}

type CreatePostInput struct {
	AuthorID uint
	Title    string
	Content  string
}

type UpdatePostInput struct {
	ID       uint
	AuthorID uint
	Title    string
	Content  string
}

type DeletePostsInput struct {
	IDs       []uint
	AuthorIDs []uint
}
