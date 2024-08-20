package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/session"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/utils"
)

type signInReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (req signInReq) validate() error {
	if err := utils.ValidatePassword(req.Password); err != nil {
		return errInvalidPassword
	}

	if err := utils.ValidateEmail(req.Email); err != nil {
		return errInvalidEmail
	}

	return nil
}

func (req signInReq) toInput() session.SignInInput {
	return session.SignInInput{
		Email:    req.Email,
		Password: req.Password,
	}
}

type user struct {
	ID        uint              `json:"id"`
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Role      string            `json:"role"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

type signInResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         user   `json:"user"`
}

func (h handler) newSignInResp(o session.SignInOutput) signInResp {
	return signInResp{
		AccessToken:  o.AccessToken,
		RefreshToken: o.RefreshToken,
		User: user{
			ID:        o.User.ID,
			Username:  o.User.Username,
			Email:     o.User.Email,
			Role:      o.User.Role,
			CreatedAt: response.DateTime(o.User.CreatedAt),
			UpdatedAt: response.DateTime(o.User.UpdatedAt),
		},
	}
}

type signUpReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (req signUpReq) validate() error {
	if err := utils.ValidatePassword(req.Password); err != nil {
		return errInvalidPassword
	}

	if err := utils.ValidateEmail(req.Email); err != nil {
		return errInvalidEmail
	}

	return nil
}

func (req signUpReq) toInput() session.SignUpInput {
	return session.SignUpInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

type signUpResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         user   `json:"user"`
}

func (h handler) newSignUpResp(o session.SignUpOutput) signUpResp {
	return signUpResp{
		AccessToken:  o.AccessToken,
		RefreshToken: o.RefreshToken,
		User: user{
			ID:        o.User.ID,
			Username:  o.User.Username,
			Email:     o.User.Email,
			Role:      o.User.Role,
			CreatedAt: response.DateTime(o.User.CreatedAt),
			UpdatedAt: response.DateTime(o.User.UpdatedAt),
		},
	}
}

type refreshReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type refreshResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h handler) newRefreshResp(o session.RefreshOutput) refreshResp {
	return refreshResp{
		AccessToken:  o.AccessToken,
		RefreshToken: o.RefreshToken,
	}
}
