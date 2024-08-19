package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapPostRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.Use(mw.Auth())
	r.GET("", h.GetPosts)
	r.GET("/:id", h.GetOnePost)
	r.POST("", h.CreatePost)
	r.PUT("/:id", h.UpdatePost)
	r.DELETE("/:id", h.DeletePost)
}
