package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/usecase"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetPosts(c *gin.Context)
	GetOnePost(c *gin.Context)
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}

type handler struct {
	l  pkgLog.Logger
	uc usecase.UseCase
}

func New(
	l pkgLog.Logger,
	uc usecase.UseCase,
) Handler {
	return &handler{
		l:  l,
		uc: uc,
	}
}
