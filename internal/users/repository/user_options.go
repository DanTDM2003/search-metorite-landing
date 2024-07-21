package repository

import "github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"

type GetUsersFilter struct {
	Role string
}

type GetUsersOptions struct {
	GetUsersFilter
	paginator.PaginatorQuery
}

type GetOneUserOptions struct {
	ID       uint
	Username string
	Email    string
}

type CreateUserOptions struct {
	Username string
	Email    string
	Password string
	Role     string
}

type UpdateUserOptions struct {
	Username string
	Email    string
	Password string
	Role     string
}
