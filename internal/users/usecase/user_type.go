package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type GetUsersFilter struct {
	Role string
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
	ID uint
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
	Password string
}

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	Token string
	User  models.User
}

type SignUpInput struct {
	Username string
	Email    string
	Password string
}

type SignUpOutput struct {
	Token string
	User  models.User
}
