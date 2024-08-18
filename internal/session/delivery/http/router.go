package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapSessionRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.POST("/signin", h.SignIn)
	r.POST("/refresh", h.Refresh)
	r.POST("/signup", h.SignUp)
	r.DELETE("/signout", mw.Auth(), h.SignOut)
}
