package usecase

import (
	"context"
	"errors"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (uc impleUsecase) GetUsers(ctx context.Context, input users.GetUsersInput) (users.GetUsersOutput, error) {
	u, pag, err := uc.repo.GetUsers(ctx, repository.GetUsersOptions{
		GetUsersFilter: repository.GetUsersFilter{
			Username: input.Username,
			Role:     input.Role,
		},
		PaginatorQuery: input.PaginatorQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.GetUsers.repo.GetUsers: %v", err)
		return users.GetUsersOutput{}, err
	}

	return users.GetUsersOutput{
		Users:     u,
		Paginator: pag,
	}, nil
}

func (uc impleUsecase) UpdateUser(ctx context.Context, input users.UpdateUserInput) (models.User, error) {
	user, err := uc.repo.GetOneUser(ctx, repository.GetOneUserOptions{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "users.usecase.UpdateUser.repo.GetOneUser: %v", ErrUserNotFound)
			return models.User{}, ErrUserNotFound
		}
		uc.l.Errorf(ctx, "users.usecase.UpdateUser.repo.GetOneUser: %v", err)
		return models.User{}, err
	}

	user, err = uc.repo.UpdateUser(ctx, repository.UpdateUserOptions{
		Username: input.Username,
		Email:    input.Email,
	}, user)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.UpdateUser.repo.UpdateUser: %v", err)
		return models.User{}, err
	}

	if err := uc.redisRepo.SetUser(ctx, user); err != nil {
		uc.l.Errorf(ctx, "users.usecase.UpdateUser.redis.SetUser: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (uc impleUsecase) GetOneUser(ctx context.Context, input users.GetOneUserInput) (models.User, error) {
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

func (uc impleUsecase) CreateUser(ctx context.Context, input users.CreateUserInput) (models.User, error) {
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

func (uc impleUsecase) DeleteUser(ctx context.Context, id uint) error {
	_, err := uc.repo.GetOneUser(ctx, repository.GetOneUserOptions{
		ID: id,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.l.Warnf(ctx, "users.usecase.DeleteUser.repo.GetOneUser: %v", ErrUserNotFound)
			return ErrUserNotFound
		}
		uc.l.Errorf(ctx, "users.usecase.DeleteUser.repo.GetOneUser: %v", err)
		return err
	}

	err = uc.repo.DeleteUser(ctx, id)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.DeleteUser.repo.DeleteUser: %v", err)
		return err
	}

	if err := uc.redisRepo.DeleteUser(ctx, id); err != nil {
		uc.l.Errorf(ctx, "users.usecase.DeleteUser.redis.DeleteUser: %v", err)
		return err
	}

	return nil
}

func (uc impleUsecase) PromoteToAdmin(ctx context.Context, id uint) (models.User, error) {
	user, err := uc.repo.GetOneUser(ctx, repository.GetOneUserOptions{
		ID: id,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "users.usecase.PromoteToAdmin.repo.GetOneUser: %v", ErrUserNotFound)
			return models.User{}, ErrUserNotFound
		}
		uc.l.Errorf(ctx, "users.usecase.PromoteToAdmin.repo.GetOneUser: %v", err)
		return models.User{}, err
	}

	if user.Role == models.UserRoleAdmin {
		uc.l.Warnf(ctx, "users.usecase.PromoteToAdmin.user.CheckRole: %v", ErrUserAlreadyAdmin)
		return models.User{}, ErrUserAlreadyAdmin
	}

	user, err = uc.repo.UpdateUser(ctx, repository.UpdateUserOptions{
		Role: models.UserRoleAdmin,
	}, user)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.PromoteToAdmin.repo.UpdateUser: %v", err)
		return models.User{}, err
	}

	if err := uc.redisRepo.SetUser(ctx, user); err != nil {
		uc.l.Errorf(ctx, "users.usecase.PromoteToAdmin.redis.SetUser: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (uc impleUsecase) DemoteToUser(ctx context.Context, id uint) (models.User, error) {
	user, err := uc.repo.GetOneUser(ctx, repository.GetOneUserOptions{
		ID: id,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "users.usecase.DemoteToUser.repo.GetOneUser: %v", ErrUserNotFound)
			return models.User{}, ErrUserNotFound
		}
		uc.l.Errorf(ctx, "users.usecase.DemoteToUser.repo.GetOneUser: %v", err)
		return models.User{}, err
	}

	if user.Role == models.UserRoleUser {
		uc.l.Warnf(ctx, "users.usecase.DemoteToUser.user.CheckRole: %v", ErrUserAlreadyUser)
		return models.User{}, ErrUserAlreadyUser
	}

	user, err = uc.repo.UpdateUser(ctx, repository.UpdateUserOptions{
		Role: models.UserRoleUser,
	}, user)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.DemoteToUser.repo.UpdateUser: %v", err)
		return models.User{}, err
	}

	if err := uc.redisRepo.SetUser(ctx, user); err != nil {
		uc.l.Errorf(ctx, "users.usecase.DemoteToUser.redis.SetUser: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (uc impleUsecase) ChangePassword(ctx context.Context, input users.ChangePasswordInput) error {
	user, err := uc.repo.GetOneUser(ctx, repository.GetOneUserOptions{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "users.usecase.ChangePassword.repo.GetOneUser: %v", ErrUserNotFound)
			return ErrUserNotFound
		}
		uc.l.Errorf(ctx, "users.usecase.ChangePassword.repo.GetOneUser: %v", err)
		return err
	}

	if ok := uc.passwordManager.CheckPasswordHash(input.OldPassword, user.Password); !ok {
		uc.l.Warnf(ctx, "users.usecase.ChangePassword.user.ComparePassword: %v", err)
		return ErrWrongPassword
	}

	hashedPassword, err := uc.passwordManager.HashPassword(input.NewPassword)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.ChangePassword.user.HashPassword: %v", err)
		return err
	}

	user, err = uc.repo.UpdateUser(ctx, repository.UpdateUserOptions{
		Password: hashedPassword,
	}, user)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.ChangePassword.repo.UpdateUser: %v", err)
		return err
	}

	if err := uc.redisRepo.SetUser(ctx, user); err != nil {
		uc.l.Errorf(ctx, "users.usecase.ChangePassword.redis.SetUser: %v", err)
		return err
	}

	return nil
}
