package database

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"gorm.io/gorm"
)

const (
	articlesTable = "articles"
)

func (repo impleRepository) getTable() *gorm.DB {
	return repo.db.Table(articlesTable)
}

func (repo impleRepository) GetArticles(ctx context.Context, opt repository.GetArticlesOptions) ([]models.Article, paginator.Paginator, error) {
	table := repo.getTable()

	var total int64
	if err := table.Count(&total).Error; err != nil {
		repo.l.Errorf(ctx, "articles.repository.database.GetArticles.db.Count: %v", err)
		return nil, paginator.Paginator{}, err
	}

	cursor := table.
		Offset(int(opt.PaginatorQuery.Offset())).
		Limit(int(opt.Limit))

	cond, params := repo.buildGetAritclesCondition(opt)

	var articles []models.Article
	if err := cursor.Where(cond, params).Find(&articles).Error; err != nil {
		repo.l.Errorf(ctx, "articles.repository.database.GetArticles.db.Find: %v", err)
		return nil, paginator.Paginator{}, err
	}

	return articles, paginator.Paginator{
		Total:       total,
		Count:       int64(len(articles)),
		PerPage:     opt.Limit,
		CurrentPage: opt.Page,
	}, nil
}

func (repo impleRepository) GetOneArticle(ctx context.Context, opt repository.GetOneArticleOptions) (models.Article, error) {
	table := repo.getTable()

	cond, params := repo.buildGetOneArticleCondition(opt)

	var article models.Article
	if err := table.Where(cond, params).First(&article).Error; err != nil {
		repo.l.Errorf(ctx, "articles.repository.database.GetOneArticle.db.First: %v", err)
		return models.Article{}, err
	}

	return article, nil
}

func (repo impleRepository) CreateArticle(ctx context.Context, opt repository.CreateArticleOptions) (models.Article, error) {
	table := repo.getTable()

	create := repo.buildCrearteArticleModel(opt)
	if err := table.Create(&create).Error; err != nil {
		repo.l.Errorf(ctx, "articles.repository.database.CreateArticle.db.Create: %v", err)
		return models.Article{}, err
	}

	return create, nil
}

func (repo impleRepository) UpdateArticle(ctx context.Context, opt repository.UpdateArticleOptions, article models.Article) (models.Article, error) {
	table := repo.getTable()

	update := repo.buildUpdateArticleModel(opt, article)
	if err := table.Where("id = ?", opt.ID).Updates(&update).Error; err != nil {
		repo.l.Errorf(ctx, "articles.repository.database.UpdateArticle.db.Updates: %v", err)
		return models.Article{}, err
	}

	return update, nil
}

func (repo impleRepository) DeleteArticle(ctx context.Context, id uint) error {
	table := repo.getTable()

	if err := table.Where("id = ?", id).Delete(&models.Article{}).Error; err != nil {
		repo.l.Errorf(ctx, "articles.repository.database.DeleteArticle.db.Delete: %v", err)
		return err
	}

	return nil
}
