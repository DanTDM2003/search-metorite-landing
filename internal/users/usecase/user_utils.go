package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/application"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/utils"
)

func buildCreateUserOptions(input application.CreateUserInput) repository.CreateUserOptions {
	opt := repository.CreateUserOptions{
		Username: input.Username,
		Email:    input.Email,
		Role:     models.UserRoleUser,
	}

	opt.Password, _ = utils.HashPassword(input.Password)

	return opt
}
