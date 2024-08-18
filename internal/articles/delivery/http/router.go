package http

import "github.com/gin-gonic/gin"

func MapArticleRoutes(r *gin.RouterGroup, h Handler) {
	r.GET("", h.GetArticles)
	r.GET("/:slug", h.GetOneArticle)
	r.POST("", h.CreateArticle)
	r.PUT("/:id", h.UpdateArticle)
	r.DELETE("/:id", h.DeleteArticle)
}
