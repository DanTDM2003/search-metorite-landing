package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type GetArticlesFilter struct {
	Author string
}

type GetArticlesInput struct {
	GetArticlesFilter
	paginator.PaginatorQuery
}

type GetArticlesOutput struct {
	Articles []models.Article
	paginator.Paginator
}

type GetOneArticleInput struct {
	ID   uint
	Slug string
}

type CreateArticleInput struct {
	AuthorID uint
	Title    string
	Content  string
	Tag      string
}

type UpdateArticleInput struct {
	ID       uint
	AuthorID uint
	Title    string
	Content  string
	Tag      string
}
