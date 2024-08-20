package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
)

type getPostsReq struct {
	AuthorID uint `form:"author_id"`
}

type getPostsRespItem struct {
	ID        uint              `json:"id"`
	AuthorID  uint              `json:"author_id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	Rating    float64           `json:"rating"`
	ViewCount uint              `json:"view_count"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

type getPostsResp struct {
	Items []getPostsRespItem          `json:"items"`
	Meta  paginator.PaginatorResponse `json:"meta"`
}

func (h handler) NewGetPostsResp(ouPut posts.GetPostsOutput) getPostsResp {
	items := make([]getPostsRespItem, 0, len(ouPut.Posts))
	for _, post := range ouPut.Posts {
		items = append(items, getPostsRespItem{
			ID:        post.ID,
			AuthorID:  post.AuthorID,
			Title:     post.Title,
			Content:   post.Content,
			Rating:    post.Rating,
			ViewCount: post.ViewCount,
			CreatedAt: response.DateTime(post.CreatedAt),
			UpdatedAt: response.DateTime(post.UpdatedAt),
		})

	}

	return getPostsResp{
		Items: items,
		Meta:  ouPut.Paginator.ToResponse(),
	}
}

type getOnePostReq struct {
	ID       uint `uri:"id"`
	AuthorID uint `form:"author_id"`
}

type getOnePostResp struct {
	ID        uint              `json:"id"`
	AuthorID  uint              `json:"author_id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	Rating    float64           `json:"rating"`
	ViewCount uint              `json:"view_count"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) NewGetOnePostResp(post models.Post) getOnePostResp {
	return getOnePostResp{
		ID:        post.ID,
		AuthorID:  post.AuthorID,
		Title:     post.Title,
		Content:   post.Content,
		Rating:    post.Rating,
		ViewCount: post.ViewCount,
		CreatedAt: response.DateTime(post.CreatedAt),
		UpdatedAt: response.DateTime(post.UpdatedAt),
	}
}

type createPostReq struct {
	AuthorID uint   `json:"author_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type createPostResp struct {
	ID        uint              `json:"id"`
	AuthorID  uint              `json:"author_id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	Rating    float64           `json:"rating"`
	ViewCount uint              `json:"view_count"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) NewCreatePostResp(post models.Post) createPostResp {
	return createPostResp{
		ID:        post.ID,
		AuthorID:  post.AuthorID,
		Title:     post.Title,
		Content:   post.Content,
		Rating:    post.Rating,
		ViewCount: post.ViewCount,
		CreatedAt: response.DateTime(post.CreatedAt),
		UpdatedAt: response.DateTime(post.UpdatedAt),
	}
}

type updatePostReq struct {
	ID       uint   `uri:"id"`
	AuthorID uint   `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type updatePostResp struct {
	ID        uint              `json:"id"`
	AuthorID  uint              `json:"author_id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	Rating    float64           `json:"rating"`
	ViewCount uint              `json:"view_count"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) NewUpdatePostResp(post models.Post) updatePostResp {
	return updatePostResp{
		ID:        post.ID,
		AuthorID:  post.AuthorID,
		Title:     post.Title,
		Content:   post.Content,
		Rating:    post.Rating,
		ViewCount: post.ViewCount,
		CreatedAt: response.DateTime(post.CreatedAt),
		UpdatedAt: response.DateTime(post.UpdatedAt),
	}
}

type deletePostReq struct {
	ID uint `uri:"id"`
}
