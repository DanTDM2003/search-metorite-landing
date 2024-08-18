package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapMeteoriteLandingRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.Use(mw.Auth())
	r.GET("", h.GetMeteoriteLandings)
	r.GET("/:id", h.GetOneMeteoriteLanding)
	r.POST("", h.CreateMeteoriteLanding)
	r.PUT("/:id", h.UpdateMeteoriteLanding)
	r.DELETE("/:id", h.DeleteMeteoriteLanding)
}
