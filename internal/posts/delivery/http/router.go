package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapPostRoutes(r *gin.RouterGroup, h Handler, m middleware.Middleware) {
	r.Use(m.Auth())
	r.GET("", h.GetPosts)
	r.GET("/:id", h.GetOnePost)
	r.POST("", h.CreatePost)
	r.PUT("/:id", h.UpdatePost)
	r.DELETE("/:id", h.DeletePost)
}
