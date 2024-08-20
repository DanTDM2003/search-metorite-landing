package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h handler) GetArticles(c *gin.Context) {
	ctx := c.Request.Context()

	_, pagQuery, err := h.processGetArticlesReq(c)
	if err != nil {
		h.l.Warnf(ctx, "articles.http.GetArticles.processGetArticlesReq: %v", err)
		response.Error(c, err)
		return
	}

	output, err := h.uc.GetArticles(ctx, articles.GetArticlesInput{
		PaginatorQuery: pagQuery,
	})
	if err != nil {
		h.l.Errorf(ctx, "articles.http.GetArticles.uc.GetArticles: %v", err)
		response.Error(c, err)
		return
	}

	response.Success(c, h.newGetArticlesResp(output))
}

func (h handler) GetOneArticle(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processGetOneArticleReq(c)
	if err != nil {
		h.l.Warnf(ctx, "articles.http.GetOneArticle.processGetOneArticleReq: %v", err)
		response.Error(c, err)
		return
	}

	article, err := h.uc.GetOneArticle(ctx, articles.GetOneArticleInput{
		Slug: req.Slug,
	})
	if err != nil {
		h.l.Errorf(ctx, "articles.http.GetOneArticle.uc.GetOneArticle: %v", err)
		response.Error(c, err)
		return
	}

	response.Success(c, h.newGetOneArticleResp(article))
}

func (h handler) CreateArticle(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processCreateArticleReq(c)
	if err != nil {
		h.l.Warnf(ctx, "articles.http.CreateArticle.processCreateArticleReq: %v", err)
		response.Error(c, err)
		return
	}

	article, err := h.uc.CreateArticle(ctx, articles.CreateArticleInput{
		Title:   req.Title,
		Content: req.Content,
		Tag:     req.Tag,
	})
	if err != nil {
		h.l.Errorf(ctx, "articles.http.CreateArticle.uc.CreateArticle: %v", err)
		response.Error(c, err)
		return
	}

	response.Success(c, h.newCreateArticleResp(article))
}

func (h handler) UpdateArticle(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processUpdateArticleReq(c)
	if err != nil {
		h.l.Warnf(ctx, "articles.http.UpdateArticle.processUpdateArticleReq: %v", err)
		response.Error(c, err)
		return
	}

	article, err := h.uc.UpdateArticle(ctx, articles.UpdateArticleInput{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		Tag:     req.Tag,
	})
	if err != nil {
		h.l.Errorf(ctx, "articles.http.UpdateArticle.uc.UpdateArticle: %v", err)
		response.Error(c, err)
		return
	}

	response.Success(c, h.newUpdateArticleResp(article))
}

func (h handler) DeleteArticle(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processDeleteArticleReq(c)
	if err != nil {
		h.l.Warnf(ctx, "articles.http.DeleteArticle.processDeleteArticleReq: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.DeleteArticle(ctx, req.ID)
	if err != nil {
		h.l.Errorf(ctx, "articles.http.DeleteArticle.uc.DeleteArticle: %v", err)
		response.Error(c, err)
		return
	}

	response.Success(c, nil)
}
