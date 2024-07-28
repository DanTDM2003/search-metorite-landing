package usecase

import (
	"context"
	"errors"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (uc impleUsecase) GetPosts(ctx context.Context, input GetPostsInput) (GetPostsOutput, error) {
	posts, pag, err := uc.repo.GetPosts(ctx, repository.GetPostsOptions{
		GetPostsFilter: repository.GetPostsFilter{
			AuthorID: input.AuthorID,
		},
		PaginatorQuery: input.PaginatorQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "posts.usecase.GetPosts.repo.GetPosts: %v", err)
		return GetPostsOutput{}, err
	}

	return GetPostsOutput{
		Posts:     posts,
		Paginator: pag,
	}, nil
}

func (uc impleUsecase) GetOnePost(ctx context.Context, input GetOnePostInput) (models.Post, error) {
	post, err := uc.redisRepo.GetPost(ctx, input.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			post, err = uc.repo.GetOnePost(ctx, repository.GetOnePostOptions{
				ID:       input.ID,
				AuthorID: input.AuthorID,
			})
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					uc.l.Warnf(ctx, "posts.usecase.GetOnePost.repo.GetOnePost: %v", ErrPostNotFound)
					return models.Post{}, ErrPostNotFound
				}
				uc.l.Errorf(ctx, "posts.usecase.GetOnePost.repo.GetOnePost: %v", err)
				return models.Post{}, err
			}

			if err := uc.redisRepo.SetPost(ctx, post); err != nil {
				uc.l.Errorf(ctx, "posts.usecase.GetOnePost.redisRepo.SetPost: %v", err)
			}
		} else {
			uc.l.Errorf(ctx, "posts.usecase.GetOnePost.redisRepo.GetPost: %v", err)
			return models.Post{}, err
		}
	}

	return post, nil
}

func (uc impleUsecase) CreatePost(ctx context.Context, input CreatePostInput) (models.Post, error) {
	post, err := uc.repo.CreatePost(ctx, repository.CreatePostOptions{
		AuthorID: input.AuthorID,
		Title:    input.Title,
		Content:  input.Content,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			uc.l.Warnf(ctx, "posts.usecase.CreatePost.repo.CreatePost: %v", ErrAuthorNotFound)
			return models.Post{}, ErrAuthorNotFound
		}
		uc.l.Errorf(ctx, "posts.usecase.CreatePost.repo.CreatePost: %v", err)
		return models.Post{}, err
	}

	if err := uc.redisRepo.SetPost(ctx, post); err != nil {
		uc.l.Errorf(ctx, "posts.usecase.CreatePost.redisRepo.SetPost: %v", err)
		return models.Post{}, err
	}

	return post, nil
}

func (uc impleUsecase) UpdatePost(ctx context.Context, input UpdatePostInput) (models.Post, error) {
	post, err := uc.repo.GetOnePost(ctx, repository.GetOnePostOptions{
		ID:       input.ID,
		AuthorID: input.AuthorID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "posts.usecase.UpdatePost.repo.GetOnePost: %v", ErrPostNotFound)
			return models.Post{}, ErrPostNotFound
		}
		uc.l.Errorf(ctx, "posts.usecase.UpdatePost.repo.GetOnePost: %v", err)
		return models.Post{}, err
	}

	post, err = uc.repo.UpdatePost(ctx, repository.UpdatePostOptions{
		AuthorID: input.AuthorID,
		Title:    input.Title,
		Content:  input.Content,
	}, post)
	if err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			uc.l.Warnf(ctx, "posts.usecase.UpdatePost.repo.UpdatePost: %v", ErrAuthorNotFound)
			return models.Post{}, ErrAuthorNotFound
		}
		uc.l.Errorf(ctx, "posts.usecase.UpdatePost.repo.UpdatePost: %v", err)
		return models.Post{}, err
	}

	if err := uc.redisRepo.SetPost(ctx, post); err != nil {
		uc.l.Errorf(ctx, "posts.usecase.UpdatePost.redisRepo.SetPost: %v", err)
	}

	return post, nil
}

func (uc impleUsecase) DeletePost(ctx context.Context, id uint) error {
	post, err := uc.repo.GetOnePost(ctx, repository.GetOnePostOptions{
		ID: id,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "posts.usecase.DeletePost.repo.GetOnePost: %v", ErrPostNotFound)
			return ErrPostNotFound
		}
		uc.l.Errorf(ctx, "posts.usecase.DeletePost.repo.GetOnePost: %v", err)
		return err
	}

	if err := uc.repo.DeletePost(ctx, post.ID); err != nil {
		uc.l.Errorf(ctx, "posts.usecase.DeletePost.repo.DeletePost: %v", err)
		return err
	}

	if err := uc.redisRepo.DeletePost(ctx, post.ID); err != nil {
		uc.l.Errorf(ctx, "posts.usecase.DeletePost.redisRepo.DeletePost: %v", err)
	}

	return nil
}
