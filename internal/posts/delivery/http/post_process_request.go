package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/gin-gonic/gin"
)

func (h handler) processGetPostsReq(c *gin.Context) (getPostsReq, paginator.PaginatorQuery, error) {
	ctx := c.Request.Context()

	var req getPostsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Warnf(ctx, "posts.http.processGetPostsReq: %v", err)
		return getPostsReq{}, paginator.PaginatorQuery{}, errWrongQuery
	}

	var pagQuery paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pagQuery); err != nil {
		h.l.Warnf(ctx, "posts.http.processGetPostsReq: %v", err)
		return getPostsReq{}, paginator.PaginatorQuery{}, errWrongQuery
	}

	pagQuery.Adjust()

	return req, pagQuery, nil
}

func (h handler) processGetOnePostReq(c *gin.Context) (getOnePostReq, error) {
	ctx := c.Request.Context()

	var req getOnePostReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Warnf(ctx, "posts.http.processGetOnePostReq: %v", err)
		return getOnePostReq{}, errWrongQuery
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "posts.http.processGetOnePostReq: %v", err)
		return getOnePostReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processCreatePostReq(c *gin.Context) (createPostReq, error) {
	ctx := c.Request.Context()

	var req createPostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "posts.http.processCreatePostReq: %v", err)
		return createPostReq{}, errWrongBody
	}

	return req, nil
}

func (h handler) processUpdatePostReq(c *gin.Context) (updatePostReq, error) {
	ctx := c.Request.Context()

	var req updatePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "posts.http.processUpdatePostReq: %v", err)
		return updatePostReq{}, errWrongBody
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "posts.http.processUpdatePostReq: %v", err)
		return updatePostReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processDeletePostReq(c *gin.Context) (deletePostReq, error) {
	ctx := c.Request.Context()

	var req deletePostReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "posts.http.processDeletePostReq: %v", err)
		return deletePostReq{}, errWrongQuery
	}

	return req, nil
}
