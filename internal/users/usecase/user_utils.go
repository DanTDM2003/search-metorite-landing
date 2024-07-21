package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/utils"
)

func (input CreateUserInput) toOptions() repository.CreateUserOptions {
	opt := repository.CreateUserOptions{
		Username: input.Username,
		Email:    input.Email,
		Role:     UserRoleUser,
	}

	opt.Password, _ = utils.HashPassword(input.Password)

	return opt
}

func (input SignUpInput) toOptions() repository.CreateUserOptions {
	opt := repository.CreateUserOptions{
		Username: input.Username,
		Email:    input.Email,
		Role:     UserRoleUser,
	}

	opt.Password, _ = utils.HashPassword(input.Password)

	return opt
}
