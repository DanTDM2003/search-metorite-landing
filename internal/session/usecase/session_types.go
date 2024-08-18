package usecase

import "github.com/DanTDM2003/search-api-docker-redis/internal/models"

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	AccessToken  string
	RefreshToken string
	User         models.User
}

type SignUpInput struct {
	Username string
	Email    string
	Password string
}

type SignUpOutput struct {
	AccessToken  string
	RefreshToken string
	User         models.User
}

type RefreshInput struct {
	RefreshToken string
}

type RefreshOutput struct {
	RefreshToken string
	AccessToken  string
}
