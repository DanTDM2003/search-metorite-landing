package database

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

func (repo impleRepository) buildCrearteArticleModel(opt repository.CreateArticleOptions) models.Article {
	return models.Article{
		AuthorID: opt.AuthorID,
		Title:    opt.Title,
		Slug:     opt.Slug,
		Content:  opt.Content,
		// Tag:      opt.Tag,
	}
}

func (repo impleRepository) buildUpdateArticleModel(opt repository.UpdateArticleOptions, article models.Article) models.Article {
	return models.Article{
		ID:       opt.ID,
		AuthorID: article.ID,
		Title:    opt.Title,
		Slug:     opt.Slug,
		Content:  opt.Content,
		// Tag:      opt.Tag,
	}
}
