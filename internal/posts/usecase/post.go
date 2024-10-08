package usecase

import (
	"context"
	"errors"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts"
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (uc impleUsecase) GetPosts(ctx context.Context, input posts.GetPostsInput) (posts.GetPostsOutput, error) {
	ps, pag, err := uc.repo.GetPosts(ctx, repository.GetPostsOptions{
		GetPostsFilter: repository.GetPostsFilter{
			AuthorID: input.AuthorID,
		},
		PaginatorQuery: input.PaginatorQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "posts.usecase.GetPosts.repo.GetPosts: %v", err)
		return posts.GetPostsOutput{}, err
	}

	return posts.GetPostsOutput{
		Posts:     ps,
		Paginator: pag,
	}, nil
}

func (uc impleUsecase) GetOnePost(ctx context.Context, input posts.GetOnePostInput) (models.Post, error) {
	post, err := uc.redisRepo.GetPost(ctx, input.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			post, err := uc.repo.GetOnePost(ctx, repository.GetOnePostOptions{
				ID:       input.ID,
				AuthorID: input.AuthorID,
			})
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					uc.l.Warnf(ctx, "posts.usecase.GetOnePost.repo.GetOnePost: %v", err)
					return models.Post{}, ErrPostNotFound
				}
				uc.l.Errorf(ctx, "posts.usecase.GetOnePost.repo.GetOnePost: %v", err)
				return models.Post{}, err
			}

			if err := uc.redisRepo.SetPost(ctx, post); err != nil {
				uc.l.Errorf(ctx, "posts.usecase.GetOnePost.redisRepo.SetPost: %v", err)
			}
			return post, nil
		} else {
			uc.l.Errorf(ctx, "posts.usecase.GetOnePost.redisRepo.GetPost: %v", err)
			return models.Post{}, err
		}
	}

	return post, nil
}

func (uc impleUsecase) CreatePost(ctx context.Context, input posts.CreatePostInput) (models.Post, error) {
	userService := uc.locator.GetService(serviceLocator.UserService).(users.Usecase)
	_, err := userService.GetOneUser(ctx, users.GetOneUserInput{
		ID: input.AuthorID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "posts.usecase.CreatePost.userUC.GetOneUser: %v", err)
			return models.Post{}, err
		}
		uc.l.Errorf(ctx, "posts.usecase.CreatePost.userUC.GetOneUser: %v", err)
		return models.Post{}, err
	}

	post, err := uc.repo.CreatePost(ctx, repository.CreatePostOptions{
		AuthorID: input.AuthorID,
		Title:    input.Title,
		Content:  input.Content,
	})
	if err != nil {
		uc.l.Errorf(ctx, "posts.usecase.CreatePost.repo.CreatePost: %v", err)
		return models.Post{}, err
	}

	if err := uc.redisRepo.SetPost(ctx, post); err != nil {
		uc.l.Errorf(ctx, "posts.usecase.CreatePost.redisRepo.SetPost: %v", err)
		return models.Post{}, err
	}

	return post, nil
}

func (uc impleUsecase) UpdatePost(ctx context.Context, input posts.UpdatePostInput) (models.Post, error) {
	post, err := uc.repo.GetOnePost(ctx, repository.GetOnePostOptions{
		ID:       input.ID,
		AuthorID: input.AuthorID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "posts.usecase.UpdatePost.repo.GetOnePost: %v", err)
			return models.Post{}, ErrPostNotFound
		}
		uc.l.Errorf(ctx, "posts.usecase.UpdatePost.repo.GetOnePost: %v", err)
		return models.Post{}, err
	}

	post, err = uc.repo.UpdatePost(ctx, repository.UpdatePostOptions{
		Title:   input.Title,
		Content: input.Content,
	}, post)
	if err != nil {
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
			uc.l.Warnf(ctx, "posts.usecase.DeletePost.repo.GetOnePost: %v", err)
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
