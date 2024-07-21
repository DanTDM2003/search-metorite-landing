package http

import "github.com/gin-gonic/gin"

func MapMeteoriteLandingRoutes(r *gin.RouterGroup, h Handler) {
	r.Use()
	r.GET("", h.GetMeteoriteLandings)
	r.GET("/:id", h.GetOneMeteoriteLanding)
	r.POST("", h.CreateMeteoriteLanding)
	r.PATCH("/:id", h.UpdateMeteoriteLanding)
	r.DELETE("/:id", h.DeleteMeteoriteLanding)
}
