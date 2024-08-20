package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetArticles(c *gin.Context)
	GetOneArticle(c *gin.Context)
	CreateArticle(c *gin.Context)
	UpdateArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
}

type handler struct {
	l  pkgLog.Logger
	uc articles.Usecase
}

func New(
	l pkgLog.Logger,
	uc articles.Usecase,
) Handler {
	return &handler{
		l:  l,
		uc: uc,
	}
}
