package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/utils"
)

type getUsersReq struct {
	Role string `form:"role"`
}

type getUsersRespItem struct {
	ID        uint              `json:"id"`
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Role      string            `json:"role"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

type getUsersResp struct {
	Data []getUsersRespItem          `json:"data"`
	Meta paginator.PaginatorResponse `json:"meta"`
}

func (h handler) newGetUsersResp(o usecase.GetUsersOutput) getUsersResp {
	var items []getUsersRespItem

	for _, item := range o.Users {
		items = append(items, getUsersRespItem{
			ID:        item.ID,
			Username:  item.Username,
			Email:     item.Email,
			Role:      item.Role,
			CreatedAt: response.DateTime(item.CreatedAt),
			UpdatedAt: response.DateTime(item.UpdatedAt),
		})
	}

	return getUsersResp{
		Data: items,
		Meta: o.Paginator.ToResponse(),
	}
}

type getOneUserReq struct {
	ID uint `uri:"id"`
}

type getOneUserResp struct {
	ID        uint              `json:"id"`
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Role      string            `json:"role"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) newGetOneUserResp(user models.User) getOneUserResp {
	return getOneUserResp{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: response.DateTime(user.CreatedAt),
		UpdatedAt: response.DateTime(user.UpdatedAt),
	}
}

type createUserReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (req createUserReq) validate() error {
	if err := utils.ValidatePassword(req.Password); err != nil {
		return errInvalidPassword
	}

	if err := utils.ValidateEmail(req.Email); err != nil {
		return errInvalidEmail
	}

	return nil
}

func (req createUserReq) toInput() usecase.CreateUserInput {
	return usecase.CreateUserInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

type createUserResp struct {
	ID        uint              `json:"id"`
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Role      string            `json:"role"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) newCreateUserResp(user models.User) createUserResp {
	return createUserResp{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: response.DateTime(user.CreatedAt),
		UpdatedAt: response.DateTime(user.UpdatedAt),
	}
}

type updateUserReq struct {
	ID       uint   `uri:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req updateUserReq) validate() error {
	if req.Password != "" {
		if err := utils.ValidatePassword(req.Password); err != nil {
			return errInvalidPassword
		}
	}

	if req.Email != "" {
		if err := utils.ValidateEmail(req.Email); err != nil {
			return errInvalidEmail
		}
	}

	return nil
}

func (req updateUserReq) toInput() usecase.UpdateUserInput {
	return usecase.UpdateUserInput{
		ID:       req.ID,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

type updateUserResp struct {
	ID        uint              `json:"id"`
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Role      string            `json:"role"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) newUpdateUserResp(user models.User) updateUserResp {
	return updateUserResp{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: response.DateTime(user.CreatedAt),
		UpdatedAt: response.DateTime(user.UpdatedAt),
	}
}

type deleteUserReq struct {
	ID uint `uri:"id"`
}

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

func (req signInReq) toInput() usecase.SignInInput {
	return usecase.SignInInput{
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
	Token string `json:"token"`
	User  user   `json:"user"`
}

func (h handler) newSignInResp(o usecase.SignInOutput) signInResp {
	return signInResp{
		Token: o.Token,
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

func (req signUpReq) toInput() usecase.SignUpInput {
	return usecase.SignUpInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

type signUpResp struct {
	Token string `json:"token"`
	User  user   `json:"user"`
}

func (h handler) newSignUpResp(o usecase.SignUpOutput) signUpResp {
	return signUpResp{
		Token: o.Token,
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

type promoteToAdminReq struct {
	ID uint `uri:"id"`
}

type promoteToAdminResp struct {
	ID        uint              `json:"id"`
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Role      string            `json:"role"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) newPromoteToAdminResp(user models.User) promoteToAdminResp {
	return promoteToAdminResp{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: response.DateTime(user.CreatedAt),
		UpdatedAt: response.DateTime(user.UpdatedAt),
	}
}

type demoteToUserReq struct {
	ID uint `uri:"id"`
}

type demoteToUserResp struct {
	ID        uint              `json:"id"`
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Role      string            `json:"role"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) newDemoteToUserResp(user models.User) demoteToUserResp {
	return demoteToUserResp{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: response.DateTime(user.CreatedAt),
		UpdatedAt: response.DateTime(user.UpdatedAt),
	}
}
