package users

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

type GetOneUserInput struct {
	ID    uint
	Email string
}

type CreateUserInput struct {
	Username string
	Email    string
	Password string
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
