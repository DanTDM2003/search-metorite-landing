package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapUserRoutes(r *gin.RouterGroup, h Handler, m middleware.Middleware) {
	r.Use(m.Auth())
	r.GET("", h.GetUsers)
	r.GET("/:id", h.GetOneUser)
	r.POST("", h.CreateUser)
	r.PATCH("/:id", h.UpdateUser)
	r.DELETE("/:id", h.DeleteUser)
	r.POST("/signin", h.SignIn)
	r.POST("/signup", h.SignUp)
	r.PATCH("/:id/promote", h.PromoteToAdmin)
	r.PATCH("/:id/demote", h.DemoteToUser)
}
