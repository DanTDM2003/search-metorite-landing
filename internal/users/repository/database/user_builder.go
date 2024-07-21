package database

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
)

func (repo impleRepository) buildCreateUserModel(opt repository.CreateUserOptions) models.User {
	return models.User{
		Username: opt.Username,
		Password: opt.Password,
		Email:    opt.Email,
		Role:     opt.Role,
	}
}

func (repo impleRepository) buildUpdateUserModel(opt repository.UpdateUserOptions, user models.User) models.User {
	update := models.User{
		ID:        user.ID,
		Username:  opt.Username,
		Password:  opt.Password,
		Email:     opt.Email,
		CreatedAt: user.CreatedAt,
	}

	if opt.Username != "" {
		update.Username = opt.Username
	} else {
		update.Username = user.Username
	}

	if opt.Password != "" {
		update.Password = opt.Password
	} else {
		update.Password = user.Password
	}

	if opt.Role != "" {
		update.Role = opt.Role
	} else {
		update.Role = user.Role
	}

	if opt.Email != "" {
		update.Email = opt.Email
	} else {
		update.Email = user.Email
	}

	return update
}
