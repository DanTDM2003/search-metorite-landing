package http

import "github.com/gin-gonic/gin"

func MapMeteoriteLandingRoutes(r *gin.RouterGroup, h Handler) {
	r.Use()
	r.GET("/meteorite-landings", h.GetMeteoriteLandings)
}
