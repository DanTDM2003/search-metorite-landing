package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	req, pagQuery, err := h.processGetUsersReq(c)
	if err != nil {
		h.l.Warnf(ctx, "http.handler.GetUsers.processGetUsersReq: %v", err)
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
		h.l.Errorf(ctx, "http.handler.GetUsers.uc.GetUsers: %v", err)
		response.Error(c, err)
		return
	}

	response.Success(c, h.newGetUsersResp(o))
}

func (h handler) GetOneUser(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processGetOneUserReq(c)
	if err != nil {
		h.l.Warnf(ctx, "http.handler.GetOneUser.processGetOneUserReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.GetOneUser(ctx, usecase.GetOneUserInput{
		ID: req.ID,
	})
	if err != nil {
		h.l.Errorf(ctx, "http.handler.GetOneUser.uc.GetOneUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newGetOneUserResp(u))
}

func (h handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processCreateUserReq(c)
	if err != nil {
		h.l.Warnf(ctx, "http.handler.CreateUser.processCreateUserReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.CreateUser(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.handler.CreateUser.uc.CreateUser: %v", err)
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
		h.l.Warnf(ctx, "http.handler.UpdateUser.processUpdateUserReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.UpdateUser(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.handler.UpdateUser.uc.UpdateUser: %v", err)
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
		h.l.Warnf(ctx, "http.handler.DeleteUser.processDeleteUserReq: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.DeleteUser(ctx, req.ID)
	if err != nil {
		h.l.Errorf(ctx, "http.handler.DeleteUser.uc.DeleteUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, nil)
}

func (h handler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processSignInReq(c)
	if err != nil {
		h.l.Warnf(ctx, "http.handler.SignIn.processSignInReq: %v", err)
		response.Error(c, err)
		return
	}

	o, err := h.uc.SignIn(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.handler.SignIn.uc.SignIn: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newSignInResp(o))
}

func (h handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processSignUpReq(c)
	if err != nil {
		h.l.Warnf(ctx, "http.handler.SignUp.processSignUpReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.SignUp(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.handler.SignUp.uc.SignUp: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newSignUpResp(u))
}

func (h handler) PromoteToAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processPromoteToAdminReq(c)
	if err != nil {
		h.l.Warnf(ctx, "http.handler.PromoteToAdmin.processPromoteToAdminReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.PromoteToAdmin(ctx, req.ID)
	if err != nil {
		h.l.Errorf(ctx, "http.handler.PromoteToAdmin.uc.PromoteToAdmin: %v", err)
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
		h.l.Warnf(ctx, "http.handler.DemoteToUser.processDemoteToUserReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.DemoteToUser(ctx, req.ID)
	if err != nil {
		h.l.Errorf(ctx, "http.handler.DemoteToUser.uc.DemoteToUser: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newDemoteToUserResp(u))
}
