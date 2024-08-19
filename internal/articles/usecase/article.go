package usecase

import (
	"context"
	"errors"

	userSrv "github.com/DanTDM2003/search-api-docker-redis/internal/application/user"
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	userUC "github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (uc impleUsecase) GetArticles(ctx context.Context, input GetArticlesInput) (GetArticlesOutput, error) {
	articles, pag, err := uc.repo.GetArticles(ctx, repository.GetArticlesOptions{
		GetArticlesFilter: repository.GetArticlesFilter{
			Author: input.Author,
		},
		PaginatorQuery: input.PaginatorQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "articles.usecase.GetArticles.repo.GetArticles: %v", err)
		return GetArticlesOutput{}, err
	}

	return GetArticlesOutput{
		Articles:  articles,
		Paginator: pag,
	}, nil
}

func (uc impleUsecase) GetOneArticle(ctx context.Context, input GetOneArticleInput) (models.Article, error) {
	article, err := uc.redis.GetArticle(ctx, input.Slug)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			article, err := uc.repo.GetOneArticle(ctx, repository.GetOneArticleOptions{
				Slug: input.Slug,
			})
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					uc.l.Warnf(ctx, "articles.usecase.GetOneArticle.repo.GetOneArticle: %v", err)
					return models.Article{}, ErrArticleNotFound
				}
				uc.l.Errorf(ctx, "articles.usecase.GetOneArticle.repo.GetOneArticle: %v", err)
				return models.Article{}, err
			}

			if err := uc.redis.SetArticle(ctx, article); err != nil {
				uc.l.Errorf(ctx, "articles.usecase.GetOneArticle.redis.SetArticle: %v", err)
				return models.Article{}, err
			}
			return article, nil
		} else {
			uc.l.Errorf(ctx, "articles.usecase.GetOneArticle.redis.GetArticle: %v", err)
			return models.Article{}, err
		}
	}

	return article, nil
}

func (uc impleUsecase) CreateArticle(ctx context.Context, input CreateArticleInput) (models.Article, error) {
	slug := utils.Slugify(input.Title)
	_, err := uc.repo.GetOneArticle(ctx, repository.GetOneArticleOptions{
		Slug: slug,
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Errorf(ctx, "articles.usecase.CreateArticle.repo.GetOneArticle: %v", err)
			return models.Article{}, err
		}
	} else {
		uc.l.Errorf(ctx, "articles.usecase.CreateArticle.repo.GetOneArticle: %v", ErrArticleTitleAlreadyUsed)
		return models.Article{}, ErrArticleTitleAlreadyUsed
	}

	userService := uc.locator.GetService(serviceLocator.UserService).(userSrv.UserUsecase)
	_, err = userService.GetOneUser(ctx, userSrv.GetOneUserInput{
		ID: input.AuthorID,
	})
	if err != nil {
		if errors.Is(err, userUC.ErrUserNotFound) {
			uc.l.Warnf(ctx, "articles.usecase.CreateArticle.userService.GetOneUser: %v", err)
			return models.Article{}, err
		}
		uc.l.Errorf(ctx, "articles.usecase.CreateArticle.userService.GetOneUser: %v", err)
		return models.Article{}, err
	}

	article, err := uc.repo.CreateArticle(ctx, repository.CreateArticleOptions{
		AuthorID: input.AuthorID,
		Title:    input.Title,
		Slug:     slug,
		Content:  input.Content,
		Tag:      input.Tag,
	})
	if err != nil {
		uc.l.Errorf(ctx, "articles.usecase.CreateArticle.repo.CreateArticle: %v", err)
		return models.Article{}, err
	}

	if err := uc.redis.SetArticle(ctx, article); err != nil {
		uc.l.Errorf(ctx, "articles.usecase.CreateArticle.redis.SetArticle: %v", err)
		return models.Article{}, err
	}

	return article, nil
}

func (uc impleUsecase) UpdateArticle(ctx context.Context, input UpdateArticleInput) (models.Article, error) {
	article, err := uc.repo.GetOneArticle(ctx, repository.GetOneArticleOptions{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "articles.usecase.UpdateArticle.repo.GetOneArticle: %v", err)
			return models.Article{}, ErrArticleNotFound
		}
		uc.l.Errorf(ctx, "articles.usecase.UpdateArticle.repo.GetOneArticle: %v", err)
		return models.Article{}, err
	}

	if input.Title != "" {
		slug := utils.Slugify(input.Title)
		_, err = uc.repo.GetOneArticle(ctx, repository.GetOneArticleOptions{
			Slug: slug,
		})
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				uc.l.Errorf(ctx, "articles.usecase.UpdateArticle.repo.GetOneArticle: %v", err)
				return models.Article{}, err
			}
		} else {
			uc.l.Errorf(ctx, "articles.usecase.UpdateArticle.repo.GetOneArticle: %v", ErrArticleTitleAlreadyUsed)
			return models.Article{}, ErrArticleTitleAlreadyUsed
		}
	}

	article, err = uc.repo.UpdateArticle(ctx, repository.UpdateArticleOptions{
		ID:      input.ID,
		Title:   input.Title,
		Content: input.Content,
	}, article)
	if err != nil {
		uc.l.Errorf(ctx, "articles.usecase.UpdateArticle.repo.UpdateArticle: %v", err)
		return models.Article{}, err
	}

	if err := uc.redis.SetArticle(ctx, article); err != nil {
		uc.l.Errorf(ctx, "articles.usecase.UpdateArticle.redis.SetArticle: %v", err)
		return models.Article{}, err
	}

	return article, nil
}

func (uc impleUsecase) DeleteArticle(ctx context.Context, id uint) error {
	_, err := uc.repo.GetOneArticle(ctx, repository.GetOneArticleOptions{
		ID: id,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "articles.usecase.DeleteArticle.repo.GetOneArticle: %v", err)
			return ErrArticleNotFound
		}
		uc.l.Errorf(ctx, "articles.usecase.DeleteArticle.repo.GetOneArticle: %v", err)
		return err
	}

	if err := uc.repo.DeleteArticle(ctx, id); err != nil {
		uc.l.Errorf(ctx, "articles.usecase.DeleteArticle.repo.DeleteArticle: %v", err)
		return err
	}

	if err := uc.redis.DeleteArticle(ctx, id); err != nil {
		uc.l.Errorf(ctx, "articles.usecase.DeleteArticle.redis.DeleteArticle: %v", err)
		return err
	}

	return nil
}
