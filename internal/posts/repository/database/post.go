package database

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"gorm.io/gorm"
)

const (
	postsTable = "posts"
)

func (repo impleRepository) getTable() *gorm.DB {
	return repo.db.Table(postsTable)
}

func (repo impleRepository) GetPosts(ctx context.Context, opt repository.GetPostsOptions) ([]models.Post, paginator.Paginator, error) {
	table := repo.getTable()

	var total int64
	if err := table.Count(&total).Error; err != nil {
		repo.l.Errorf(ctx, "posts.repository.database.GetPosts.db.Count: %v", err)
		return nil, paginator.Paginator{}, err
	}

	cursor := table.
		Offset(int(opt.PaginatorQuery.Offset())).
		Limit(int(opt.Limit))

	var posts []models.Post
	if err := cursor.Find(&posts).Error; err != nil {
		repo.l.Errorf(ctx, "posts.repository.database.GetPosts.db.Find: %v", err)
		return nil, paginator.Paginator{}, err
	}

	return posts, paginator.Paginator{
		Total:       total,
		Count:       int64(len(posts)),
		PerPage:     opt.Limit,
		CurrentPage: opt.Page,
	}, nil
}
