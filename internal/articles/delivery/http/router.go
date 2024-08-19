package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapArticleRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.Use(mw.Auth()).Use(mw.UserSession())
	r.GET("", h.GetArticles)
	r.GET("/:slug", h.GetOneArticle)
	r.POST("", mw.Permission(), h.CreateArticle)
	r.PUT("/:id", mw.Permission(), h.UpdateArticle)
	r.DELETE("/:id", mw.Permission(), h.DeleteArticle)
}
