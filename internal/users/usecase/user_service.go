package usecase

import (
	"context"
	"errors"

	userSrv "github.com/DanTDM2003/search-api-docker-redis/internal/application/user"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (uc impleUsecase) GetOneUser(ctx context.Context, input userSrv.GetOneUserInput) (models.User, error) {
	user, err := uc.redisRepo.GetUser(ctx, input.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			user, err := uc.repo.GetOneUser(ctx, repository.GetOneUserOptions{
				ID:    input.ID,
				Email: input.Email,
			})
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					uc.l.Warnf(ctx, "users.usecase.GetUser.repo.GetOneUser: %v", ErrUserNotFound)
					return models.User{}, ErrUserNotFound
				}
				uc.l.Errorf(ctx, "users.usecase.GetUser.repo.GetOneUser: %v", err)
				return models.User{}, err
			}

			if err := uc.redisRepo.SetUser(ctx, user); err != nil {
				uc.l.Errorf(ctx, "users.usecase.GetUser.redis.SetUser: %v", err)
				return models.User{}, err
			}
			return user, nil
		}
		uc.l.Errorf(ctx, "users.usecase.GetUser.redis.GetUser: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (uc impleUsecase) CreateUser(ctx context.Context, input userSrv.CreateUserInput) (models.User, error) {
	_, err := uc.repo.GetOneUser(ctx, repository.GetOneUserOptions{
		Email: input.Email,
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Errorf(ctx, "users.usecase.CreateUser.repo.GetOneUserByEmail: %v", err)
			return models.User{}, err
		}
	} else {
		uc.l.Warnf(ctx, "users.usecase.CreateUser.repo.GetOneUserByEmail: %v", ErrUserEmailExists)
		return models.User{}, ErrUserEmailExists
	}

	hashedPassword, err := uc.passwordManager.HashPassword(input.Password)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.CreateUser.passwordManager.HashPassword: %v", err)
		return models.User{}, err
	}

	user, err := uc.repo.CreateUser(ctx, repository.CreateUserOptions{
		Username: input.Username,
		Email:    input.Email,
		Role:     models.UserRoleUser,
		Password: hashedPassword,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.CreateUser.repo.CreateUser: %v", err)
		return models.User{}, err
	}

	if err := uc.redisRepo.SetUser(ctx, user); err != nil {
		uc.l.Errorf(ctx, "users.usecase.CreateUser.redis.SetUser: %v", err)
		return models.User{}, err
	}

	return user, nil
}
