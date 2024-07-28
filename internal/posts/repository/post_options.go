package repository

import "github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"

type GetPostsFilter struct {
	AuthorID uint
}

type GetPostsOptions struct {
	GetPostsFilter
	paginator.PaginatorQuery
}

type GetOnePostOptions struct {
	ID       uint
	AuthorID uint
}

type CreatePostOptions struct {
	AuthorID uint
	Title    string
	Content  string
}

type UpdatePostOptions struct {
	AuthorID uint
	Title    string
	Content  string
}

type DeletePostsOptions struct {
	IDs       []uint
	AuthorIDs []uint
}
