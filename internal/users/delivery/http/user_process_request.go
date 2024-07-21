package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/gin-gonic/gin"
)

func (h handler) processGetUsersReq(c *gin.Context) (getUsersReq, paginator.PaginatorQuery, error) {
	ctx := c.Request.Context()

	var req getUsersReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processGetUsersReq.ShouldBindQuery: %v", err)
		return getUsersReq{}, paginator.PaginatorQuery{}, errWrongQuery
	}

	var pagQuery paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pagQuery); err != nil {
		h.l.Warnf(ctx, "http.handler.processGetUsersReq.ShouldBindQuery: %v", err)
		return getUsersReq{}, paginator.PaginatorQuery{}, errWrongQuery
	}

	pagQuery.Adjust()

	return req, pagQuery, nil
}

func (h handler) processGetOneUserReq(c *gin.Context) (getOneUserReq, error) {
	ctx := c.Request.Context()

	var req getOneUserReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processGetOneUserReq.ShouldBindUri: %v", err)
		return getOneUserReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processCreateUserReq(c *gin.Context) (createUserReq, error) {
	ctx := c.Request.Context()

	var req createUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processCreateUserReq.ShouldBindJSON: %v", err)
		return createUserReq{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "http.handler.processCreateUserReq.validate: %v", err)
		return createUserReq{}, err
	}

	return req, nil
}

func (h handler) processUpdateUserReq(c *gin.Context) (updateUserReq, error) {
	ctx := c.Request.Context()

	var req updateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processUpdateUserReq.ShouldBindJSON: %v", err)
		return updateUserReq{}, errWrongBody
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processUpdateUserReq.ShouldBindUri: %v", err)
		return updateUserReq{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "http.handler.processUpdateUser.validate: %v", err)
		return updateUserReq{}, err
	}

	return req, nil
}

func (h handler) processDeleteUserReq(c *gin.Context) (deleteUserReq, error) {
	ctx := c.Request.Context()

	var req deleteUserReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processDeleteUserReq.ShouldBindUri: %v", err)
		return deleteUserReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processSignInReq(c *gin.Context) (signInReq, error) {
	ctx := c.Request.Context()

	var req signInReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processSignInReq.ShouldBindJSON: %v", err)
		return signInReq{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "http.handler.processSignInReq.validate: %v", err)
		return signInReq{}, err
	}

	return req, nil
}

func (h handler) processSignUpReq(c *gin.Context) (signUpReq, error) {
	ctx := c.Request.Context()

	var req signUpReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processSignUpReq.ShouldBindJSON: %v", err)
		return signUpReq{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "http.handler.processSignUpReq.validate: %v", err)
		return signUpReq{}, err
	}

	return req, nil
}

func (h handler) processPromoteToAdminReq(c *gin.Context) (promoteToAdminReq, error) {
	ctx := c.Request.Context()

	var req promoteToAdminReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processPromoteToAdminReq.ShouldBindUri: %v", err)
		return promoteToAdminReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processDemoteToUserReq(c *gin.Context) (demoteToUserReq, error) {
	ctx := c.Request.Context()

	var req demoteToUserReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "http.handler.processDemoteToUserReq.ShouldBindUri: %v", err)
		return demoteToUserReq{}, errWrongQuery
	}

	return req, nil
}
