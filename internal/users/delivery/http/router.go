package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapUserRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.Use(mw.Auth()).Use(mw.UserSession())
	r.GET("", h.GetUsers)
	r.GET("/:id", h.GetOneUser)
	r.POST("", mw.Permission(), h.CreateUser)
	r.PUT("/:id", h.UpdateUser)
	r.DELETE("/:id", mw.Permission(), h.DeleteUser)
	r.PATCH("/:id/promote", mw.Permission(), h.PromoteToAdmin)
	r.PATCH("/:id/demote", mw.Permission(), h.DemoteToUser)
	r.PATCH("/:id/change-password", h.ChangePassword)
}
