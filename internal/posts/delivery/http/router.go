package http

import "github.com/gin-gonic/gin"

func MapPostRoutes(r *gin.RouterGroup, h Handler) {
	r.Group("/posts")
	r.GET("", h.GetPosts)
	r.GET("/:id", h.GetOnePost)
	r.POST("", h.CreatePost)
	r.PATCH("/:id", h.UpdatePost)
	r.DELETE("/:id", h.DeletePost)
}
