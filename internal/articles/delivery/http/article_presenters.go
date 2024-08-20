package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
)

type getArticlesReq struct {
	Author string `form:"author"`
}

type getArticlesRespItem struct {
	ID        uint              `json:"id"`
	AuthorID  uint              `json:"author_id"`
	Title     string            `json:"title"`
	Slug      string            `json:"slug"`
	Content   string            `json:"content"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

type getArticlesResp struct {
	Items []getArticlesRespItem       `json:"items"`
	Meta  paginator.PaginatorResponse `json:"meta"`
}

func (h handler) newGetArticlesResp(o articles.GetArticlesOutput) getArticlesResp {
	items := make([]getArticlesRespItem, 0, len(o.Articles))
	for _, article := range o.Articles {
		items = append(items, getArticlesRespItem{
			ID:        article.ID,
			AuthorID:  article.AuthorID,
			Title:     article.Title,
			Slug:      article.Slug,
			Content:   article.Content,
			CreatedAt: response.DateTime(article.CreatedAt),
			UpdatedAt: response.DateTime(article.UpdatedAt),
		})
	}

	return getArticlesResp{
		Items: items,
		Meta:  o.Paginator.ToResponse(),
	}
}

type getOneArticleReq struct {
	Slug string `uri:"slug"`
}

type getOneArticleResp struct {
	ID        uint              `json:"id"`
	AuthorID  uint              `json:"author_id"`
	Title     string            `json:"title"`
	Slug      string            `json:"slug"`
	Content   string            `json:"content"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) newGetOneArticleResp(article models.Article) getOneArticleResp {
	return getOneArticleResp{
		ID:        article.ID,
		AuthorID:  article.AuthorID,
		Title:     article.Title,
		Slug:      article.Slug,
		Content:   article.Content,
		CreatedAt: response.DateTime(article.CreatedAt),
		UpdatedAt: response.DateTime(article.UpdatedAt),
	}
}

type createArticleReq struct {
	AuthorID uint   `json:"author_id"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Content  string `json:"content"`
	Tag      string `json:"tag"`
}

type createArticleResp struct {
	ID        uint              `json:"id"`
	AuthorID  uint              `json:"author_id"`
	Title     string            `json:"title"`
	Slug      string            `json:"slug"`
	Content   string            `json:"content"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) newCreateArticleResp(article models.Article) createArticleResp {
	return createArticleResp{
		ID:        article.ID,
		AuthorID:  article.AuthorID,
		Title:     article.Title,
		Slug:      article.Slug,
		Content:   article.Content,
		CreatedAt: response.DateTime(article.CreatedAt),
		UpdatedAt: response.DateTime(article.UpdatedAt),
	}
}

type updateArticleReq struct {
	ID       uint   `uri:"id"`
	AuthorID uint   `json:"author_id"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Content  string `json:"content"`
	Tag      string `json:"tag"`
}

type updateArticleResp struct {
	ID        uint              `json:"id"`
	AuthorID  uint              `json:"author_id"`
	Title     string            `json:"title"`
	Slug      string            `json:"slug"`
	Content   string            `json:"content"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

func (h handler) newUpdateArticleResp(article models.Article) updateArticleResp {
	return updateArticleResp{
		ID:        article.ID,
		AuthorID:  article.AuthorID,
		Title:     article.Title,
		Slug:      article.Slug,
		Content:   article.Content,
		CreatedAt: response.DateTime(article.CreatedAt),
		UpdatedAt: response.DateTime(article.UpdatedAt),
	}
}

type deleteArticleReq struct {
	ID uint `uri:"id"`
}
