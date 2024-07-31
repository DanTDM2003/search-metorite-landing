package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapMeteoriteLandingRoutes(r *gin.RouterGroup, h Handler, m middleware.Middleware) {
	r.Use(m.Auth())
	r.GET("", h.GetMeteoriteLandings)
	r.GET("/:id", h.GetOneMeteoriteLanding)
	r.POST("", h.CreateMeteoriteLanding)
	r.PATCH("/:id", h.UpdateMeteoriteLanding)
	r.DELETE("/:id", h.DeleteMeteoriteLanding)
}
