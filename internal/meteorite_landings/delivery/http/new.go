package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetMeteoriteLandings(c *gin.Context)
	GetOneMeteoriteLanding(c *gin.Context)
	CreateMeteoriteLanding(c *gin.Context)
	UpdateMeteoriteLanding(c *gin.Context)
}

type handler struct {
	l  pkgLog.Logger
	uc usecase.Usecase
}

func New(l pkgLog.Logger, uc usecase.Usecase) Handler {
	return &handler{
		l:  l,
		uc: uc,
	}
}
