package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapUserRoutes(r *gin.RouterGroup, h Handler, m middleware.Middleware) {
	r.POST("/signin", h.SignIn)
	r.Use(m.Auth()).Use(m.UserSession())
	r.GET("", h.GetUsers)
	r.GET("/:id", h.GetOneUser)
	r.POST("", h.CreateUser)
	r.PUT("/:id", h.UpdateUser)
	r.DELETE("/:id", h.DeleteUser)
	r.PATCH("/:id/promote", h.PromoteToAdmin)
	r.PATCH("/:id/demote", h.DemoteToUser)
	r.POST("/signup", h.SignUp)
}
