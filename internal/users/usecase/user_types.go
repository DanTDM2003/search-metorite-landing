package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type GetUsersFilter struct {
	Username string
	Role     string
}

type GetUsersInput struct {
	GetUsersFilter
	paginator.PaginatorQuery
}

type GetUsersOutput struct {
	Users     []models.User
	Paginator paginator.Paginator
}

type UpdateUserInput struct {
	ID       uint
	Username string
	Email    string
}

type ChangePasswordInput struct {
	ID          uint
	OldPassword string
	NewPassword string
}
