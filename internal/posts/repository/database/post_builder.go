package database

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository"
)

func (repo impleRepository) buildCreatePostModel(opt repository.CreatePostOptions) models.Post {
	return models.Post{
		Title:    opt.Title,
		Content:  opt.Content,
		AuthorID: opt.AuthorID,
	}
}

func (repo impleRepository) buildUpdatePostModel(opt repository.UpdatePostOptions, post models.Post) models.Post {
	update := models.Post{
		ID:        post.ID,
		AuthorID:  post.AuthorID,
		ViewCount: post.ViewCount,
		Rating:    post.Rating,
		CreatedAt: post.CreatedAt,
	}

	if opt.Title != "" {
		update.Title = opt.Title
	} else {
		update.Title = post.Title
	}

	if opt.Content != "" {
		update.Content = opt.Content
	} else {
		update.Content = post.Content
	}

	return update
}
