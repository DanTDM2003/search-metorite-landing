package middleware

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (m Middleware) Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user role from context
		role := pkgJWT.GetRoleFromContext(c.Request.Context())

		// Check if user role is superadmin
		if role != models.UserSuperAdmin && role != models.UserRoleAdmin {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
