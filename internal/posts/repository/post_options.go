package repository

import "github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"

type GetPostsFilter struct {
	AuthorID uint
}

type GetPostsOptions struct {
	GetPostsFilter
	paginator.PaginatorQuery
}
