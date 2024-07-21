package usecase

import (
	"context"
	"errors"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/utils"
	"gorm.io/gorm"
)

func (uc impleUsecase) GetUsers(ctx context.Context, input GetUsersInput) (GetUsersOutput, error) {
	users, pag, err := uc.repo.GetUsers(ctx, repository.GetUsersOptions{
		PaginatorQuery: input.PaginatorQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.GetUsers.repo.GetUsers: %v", err)
		return GetUsersOutput{}, err
	}

	return GetUsersOutput{
		Users:     users,
		Paginator: pag,
	}, nil
}

func (uc impleUsecase) GetOneUser(ctx context.Context, input GetOneUserInput) (models.User, error) {
	user, err := uc.repo.GetOneUser(ctx, repository.GetOneUserOptions{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "users.usecase.GetUser.repo.GetOneUser: %v", ErrUserNotFound)
			return models.User{}, ErrUserNotFound
		}
		uc.l.Errorf(ctx, "users.usecase.GetUser.repo.GetOneUser: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (uc impleUsecase) CreateUser(ctx context.Context, input CreateUserInput) (models.User, error) {
	user, err := uc.repo.CreateUser(ctx, input.toOptions())
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.CreateUser.repo.CreateUser: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (uc impleUsecase) UpdateUser(ctx context.Context, input UpdateUserInput) (models.User, error) {
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
		Password: input.Password,
		Email:    input.Email,
	}, user)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.UpdateUser.repo.UpdateUser: %v", err)
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

	return nil
}

func (uc impleUsecase) SignIn(ctx context.Context, input SignInInput) (SignInOutput, error) {
	user, err := uc.repo.GetOneUser(ctx, repository.GetOneUserOptions{
		Email: input.Email,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "users.usecase.SignIn.repo.GetOneUserByEmail: %v", ErrUserNotFound)
			return SignInOutput{}, ErrUserNotFound
		}
		uc.l.Errorf(ctx, "users.usecase.SignIn.repo.GetOneUserByEmail: %v", err)
		return SignInOutput{}, err
	}

	if ok := utils.CheckPasswordHash(input.Password, user.Password); !ok {
		uc.l.Warnf(ctx, "users.usecase.SignIn.user.ComparePassword: %v", err)
		return SignInOutput{}, ErrWrongPassword
	}

	return SignInOutput{
		User: user,
	}, nil
}

func (uc impleUsecase) SignUp(ctx context.Context, input SignUpInput) (SignUpOutput, error) {
	user, err := uc.repo.CreateUser(ctx, input.toOptions())
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.SignUp.repo.CreateUser: %v", err)
		return SignUpOutput{}, err
	}

	return SignUpOutput{
		User: user,
	}, nil
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

	if user.Role == UserRoleAdmin {
		uc.l.Warnf(ctx, "users.usecase.PromoteToAdmin.user.CheckRole: %v", ErrUserAlreadyAdmin)
		return models.User{}, ErrUserAlreadyAdmin
	}

	user, err = uc.repo.UpdateUser(ctx, repository.UpdateUserOptions{
		Role: UserRoleAdmin,
	}, user)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.PromoteToAdmin.repo.UpdateUser: %v", err)
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

	if user.Role == UserRoleUser {
		uc.l.Warnf(ctx, "users.usecase.DemoteToUser.user.CheckRole: %v", ErrUserAlreadyUser)
		return models.User{}, ErrUserAlreadyUser
	}

	user, err = uc.repo.UpdateUser(ctx, repository.UpdateUserOptions{
		Role: UserRoleUser,
	}, user)
	if err != nil {
		uc.l.Errorf(ctx, "users.usecase.DemoteToUser.repo.UpdateUser: %v", err)
		return models.User{}, err
	}

	return user, nil
}
