package http

import "github.com/gin-gonic/gin"

func MapMeteoriteLandingRoutes(r *gin.RouterGroup, h Handler) {
	r.Use()
	r.GET("/meteorite-landings", h.GetMeteoriteLandings)
	r.GET("/meteorite-landings/:id", h.GetOneMeteoriteLanding)
	r.POST("/meteorite-landings", h.CreateMeteoriteLanding)
	r.PATCH("/meteorite-landings/:id", h.UpdateMeteoriteLanding)
}
