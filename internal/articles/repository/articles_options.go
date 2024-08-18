package repository

import "github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"

type GetArticlesFilter struct {
	Author string
}

type GetArticlesOptions struct {
	GetArticlesFilter
	paginator.PaginatorQuery
}

type GetOneArticleOptions struct {
	ID     uint
	Slug   string
	Author string
}

type CreateArticleOptions struct {
	AuthorID uint
	Title    string
	Slug     string
	Content  string
	Tag      string
}

type UpdateArticleOptions struct {
	ID      uint
	Title   string
	Slug    string
	Content string
	Tag     string
}
