package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/session/usecase"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
	SignOut(c *gin.Context)
	Refresh(c *gin.Context)
}

type handler struct {
	l  pkgLog.Logger
	uc usecase.Usecase
}

func New(
	l pkgLog.Logger,
	uc usecase.Usecase,
) Handler {
	return &handler{
		l:  l,
		uc: uc,
	}
}
