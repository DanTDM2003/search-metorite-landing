package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/gin-gonic/gin"
)

func (h handler) processGetArticlesReq(c *gin.Context) (getArticlesReq, paginator.PaginatorQuery, error) {
	ctx := c.Request.Context()

	var req getArticlesReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Warnf(ctx, "articles.http.processGetArticlesReq.ShouldBindQuery: %v", err)
		return getArticlesReq{}, paginator.PaginatorQuery{}, err
	}

	var pagQuery paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pagQuery); err != nil {
		h.l.Warnf(ctx, "articles.http.processGetArticlesReq.ShouldBindQuery: %v", err)
		return getArticlesReq{}, paginator.PaginatorQuery{}, err
	}

	pagQuery.Adjust()

	return req, pagQuery, nil
}

func (h handler) processGetOneArticleReq(c *gin.Context) (getOneArticleReq, error) {
	ctx := c.Request.Context()

	var req getOneArticleReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "articles.http.processGetOneArticleReq.ShouldBindUri: %v", err)
		return getOneArticleReq{}, err
	}

	return req, nil
}

func (h handler) processCreateArticleReq(c *gin.Context) (createArticleReq, error) {
	ctx := c.Request.Context()

	var req createArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "articles.http.processCreateArticleReq.ShouldBindJSON: %v", err)
		return createArticleReq{}, err
	}

	return req, nil
}

func (h handler) processUpdateArticleReq(c *gin.Context) (updateArticleReq, error) {
	ctx := c.Request.Context()

	var req updateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "articles.http.processUpdateArticleReq.ShouldBindJSON: %v", err)
		return updateArticleReq{}, err
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "articles.http.processUpdateArticleReq.ShouldBindUri: %v", err)
		return updateArticleReq{}, err
	}

	return req, nil
}

func (h handler) processDeleteArticleReq(c *gin.Context) (deleteArticleReq, error) {
	ctx := c.Request.Context()

	var req deleteArticleReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "articles.http.processDeleteArticleReq.ShouldBindUri: %v", err)
		return deleteArticleReq{}, err
	}

	return req, nil
}
