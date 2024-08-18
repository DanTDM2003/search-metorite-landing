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

	cursor = repo.buildGetPostsCondition(cursor, opt)

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

func (repo impleRepository) GetOnePost(ctx context.Context, opt repository.GetOnePostOptions) (models.Post, error) {
	table := repo.getTable()

	cursor := repo.buildGetOnePostCondition(table, opt)

	var post models.Post
	if err := cursor.First(&post).Error; err != nil {
		repo.l.Errorf(ctx, "posts.repository.database.GetOnePost.db.First: %v", err)
		return models.Post{}, err
	}

	return post, nil
}

func (repo impleRepository) CreatePost(ctx context.Context, opt repository.CreatePostOptions) (models.Post, error) {
	table := repo.getTable()

	create := repo.buildCreatePostModel(opt)
	if err := table.Create(&create).Error; err != nil {
		repo.l.Errorf(ctx, "posts.repository.database.CreatePost.db.Create: %v", err)
		return models.Post{}, err
	}

	return create, nil
}

func (repo impleRepository) UpdatePost(ctx context.Context, opt repository.UpdatePostOptions, post models.Post) (models.Post, error) {
	table := repo.getTable()

	update := repo.buildUpdatePostModel(opt, post)

	if err := table.Where("id = ?", post.ID).Updates(&update).Error; err != nil {
		repo.l.Errorf(ctx, "posts.repository.database.UpdatePost.db.Updates: %v", err)
		return models.Post{}, err
	}

	return update, nil
}

func (repo impleRepository) DeletePost(ctx context.Context, id uint) error {
	table := repo.getTable()

	if err := table.Where("id = ?", id).Delete(&models.Post{}).Error; err != nil {
		repo.l.Errorf(ctx, "posts.repository.database.DeletePost.db.Delete: %v", err)
		return err
	}

	return nil
}

func (repo impleRepository) DeletePosts(ctx context.Context, opt repository.DeletePostsOptions) error {
	table := repo.getTable()

	cursor := repo.buildDeletePostsCondition(table, opt)

	if err := cursor.Delete(&models.Post{}).Error; err != nil {
		repo.l.Errorf(ctx, "posts.repository.database.DeletePosts.db.Delete: %v", err)
		return err
	}

	return nil
}
