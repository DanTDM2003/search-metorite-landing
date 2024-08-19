package http

import (
	userSrv "github.com/DanTDM2003/search-api-docker-redis/internal/application/user"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	req, pagQuery, err := h.processGetUsersReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.GetUsers.processGetUsersReq: %v", err)
		response.Error(c, err)
		return
	}

	o, err := h.uc.GetUsers(ctx, usecase.GetUsersInput{
		GetUsersFilter: usecase.GetUsersFilter{
			Role: req.Role,
		},
		PaginatorQuery: pagQuery,
	})
	if err != nil {
		h.l.Errorf(ctx, "users.http.GetUsers.uc.GetUsers: %v", err)
		response.Error(c, err)
		return
	}

	response.Success(c, h.newGetUsersResp(o))
}

func (h handler) GetOneUser(c *gin.Context) {
	ctx := c.Request.Context()
	userService := h.locator.GetService(serviceLocator.UserService).(userSrv.UserUsecase)

	req, err := h.processGetOneUserReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.GetOneUser.processGetOneUserReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := userService.GetOneUser(ctx, userSrv.GetOneUserInput{
		ID: req.ID,
	})
	if err != nil {
		h.l.Errorf(ctx, "users.http.GetOneUser.uc.GetOneUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newGetOneUserResp(u))
}

func (h handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	userService := h.locator.GetService("userUsecase").(userSrv.UserUsecase)

	req, err := h.processCreateUserReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.CreateUser.processCreateUserReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := userService.CreateUser(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "users.http.CreateUser.uc.CreateUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newCreateUserResp(u))
}

func (h handler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processUpdateUserReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.UpdateUser.processUpdateUserReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.UpdateUser(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "users.http.UpdateUser.uc.UpdateUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newUpdateUserResp(u))
}

func (h handler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processDeleteUserReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.DeleteUser.processDeleteUserReq: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.DeleteUser(ctx, req.ID)
	if err != nil {
		h.l.Errorf(ctx, "users.http.DeleteUser.uc.DeleteUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, nil)
}

func (h handler) PromoteToAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processPromoteToAdminReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.PromoteToAdmin.processPromoteToAdminReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.PromoteToAdmin(ctx, req.ID)
	if err != nil {
		h.l.Errorf(ctx, "users.http.PromoteToAdmin.uc.PromoteToAdmin: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newPromoteToAdminResp(u))
}

func (h handler) DemoteToUser(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processDemoteToUserReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.DemoteToUser.processDemoteToUserReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.DemoteToUser(ctx, req.ID)
	if err != nil {
		h.l.Errorf(ctx, "users.http.DemoteToUser.uc.DemoteToUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newDemoteToUserResp(u))
}

func (h handler) ChangePassword(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processChangePasswordReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.ChangePassword.processChangePasswordReq: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.ChangePassword(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "users.http.ChangePassword.uc.ChangePassword: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, nil)
}
