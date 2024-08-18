package http

import (
	pkgErrors "github.com/DanTDM2003/search-api-docker-redis/pkg/errors"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/gin-gonic/gin"
)

func (h handler) processGetUsersReq(c *gin.Context) (getUsersReq, paginator.PaginatorQuery, error) {
	ctx := c.Request.Context()

	_, ok := pkgJWT.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "users.http.processGetUsersReq.ShouldBindQuery: %v", pkgErrors.NewUnauthorizedHTTPError())
		return getUsersReq{}, paginator.PaginatorQuery{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req getUsersReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processGetUsersReq.ShouldBindQuery: %v", err)
		return getUsersReq{}, paginator.PaginatorQuery{}, errWrongQuery
	}

	var pagQuery paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pagQuery); err != nil {
		h.l.Warnf(ctx, "users.http.processGetUsersReq.ShouldBindQuery: %v", err)
		return getUsersReq{}, paginator.PaginatorQuery{}, errWrongQuery
	}

	pagQuery.Adjust()

	return req, pagQuery, nil
}

func (h handler) processGetOneUserReq(c *gin.Context) (getOneUserReq, error) {
	ctx := c.Request.Context()

	var req getOneUserReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processGetOneUserReq.ShouldBindUri: %v", err)
		return getOneUserReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processCreateUserReq(c *gin.Context) (createUserReq, error) {
	ctx := c.Request.Context()

	var req createUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processCreateUserReq.ShouldBindJSON: %v", err)
		return createUserReq{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "users.http.processCreateUserReq.validate: %v", err)
		return createUserReq{}, err
	}

	return req, nil
}

func (h handler) processUpdateUserReq(c *gin.Context) (updateUserReq, error) {
	ctx := c.Request.Context()

	var req updateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processUpdateUserReq.ShouldBindJSON: %v", err)
		return updateUserReq{}, errWrongBody
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processUpdateUserReq.ShouldBindUri: %v", err)
		return updateUserReq{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "users.http.processUpdateUser.validate: %v", err)
		return updateUserReq{}, err
	}

	return req, nil
}

func (h handler) processDeleteUserReq(c *gin.Context) (deleteUserReq, error) {
	ctx := c.Request.Context()

	var req deleteUserReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processDeleteUserReq.ShouldBindUri: %v", err)
		return deleteUserReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processPromoteToAdminReq(c *gin.Context) (promoteToAdminReq, error) {
	ctx := c.Request.Context()

	var req promoteToAdminReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processPromoteToAdminReq.ShouldBindUri: %v", err)
		return promoteToAdminReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processDemoteToUserReq(c *gin.Context) (demoteToUserReq, error) {
	ctx := c.Request.Context()

	var req demoteToUserReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processDemoteToUserReq.ShouldBindUri: %v", err)
		return demoteToUserReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processChangePasswordReq(c *gin.Context) (changePasswordReq, error) {
	ctx := c.Request.Context()

	var req changePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processChangePasswordReq.ShouldBindJSON: %v", err)
		return changePasswordReq{}, errWrongBody
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processChangePasswordReq.ShouldBindUri: %v", err)
		return changePasswordReq{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "users.http.processChangePasswordReq.validate: %v", err)
		return changePasswordReq{}, err
	}

	return req, nil
}
