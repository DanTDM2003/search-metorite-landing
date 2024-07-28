package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/usecase"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h handler) GetPosts(c *gin.Context) {
	ctx := c.Request.Context()

	req, pagQuery, err := h.processGetPostsReq(c)
	if err != nil {
		h.l.Warnf(ctx, "posts.http.GetPosts.processGetPostsReq: %v", err)
		response.Error(c, err)
		return
	}

	o, err := h.uc.GetPosts(ctx, usecase.GetPostsInput{
		GetPostsFilter: usecase.GetPostsFilter{
			AuthorID: req.AuthorID,
		},
		PaginatorQuery: pagQuery,
	})
	if err != nil {
		h.l.Errorf(ctx, "posts.http.GetPosts.uc.GetPosts: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.NewGetPostsResp(o))
}

func (h handler) GetOnePost(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processGetOnePostReq(c)
	if err != nil {
		h.l.Warnf(ctx, "posts.http.GetOnePost.processGetOnePostReq: %v", err)
		response.Error(c, err)
		return
	}

	o, err := h.uc.GetOnePost(ctx, usecase.GetOnePostInput{
		ID:       req.ID,
		AuthorID: req.AuthorID,
	})
	if err != nil {
		h.l.Errorf(ctx, "posts.http.GetOnePost.uc.GetOnePost: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.NewGetOnePostResp(o))
}

func (h handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processCreatePostReq(c)
	if err != nil {
		h.l.Warnf(ctx, "posts.http.CreatePost.processCreatePostReq: %v", err)
		response.Error(c, err)
		return
	}

	o, err := h.uc.CreatePost(ctx, usecase.CreatePostInput{
		AuthorID: req.AuthorID,
		Title:    req.Title,
		Content:  req.Content,
	})
	if err != nil {
		h.l.Errorf(ctx, "posts.http.CreatePost.uc.CreatePost: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.NewCreatePostResp(o))
}

func (h handler) UpdatePost(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processUpdatePostReq(c)
	if err != nil {
		h.l.Warnf(ctx, "posts.http.UpdatePost.processUpdatePostReq: %v", err)
		response.Error(c, err)
		return
	}

	o, err := h.uc.UpdatePost(ctx, usecase.UpdatePostInput{
		ID:       req.ID,
		AuthorID: req.AuthorID,
		Title:    req.Title,
		Content:  req.Content,
	})
	if err != nil {
		h.l.Errorf(ctx, "posts.http.UpdatePost.uc.UpdatePost: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.NewUpdatePostResp(o))
}

func (h handler) DeletePost(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processDeletePostReq(c)
	if err != nil {
		h.l.Warnf(ctx, "posts.http.DeletePost.processDeletePostReq: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.DeletePost(ctx, req.ID)
	if err != nil {
		h.l.Errorf(ctx, "posts.http.DeletePost.uc.DeletePost: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, nil)
}
