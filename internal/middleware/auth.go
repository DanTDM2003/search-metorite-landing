package middleware

import (
	"strings"

	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (m Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		if token == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		claims, err := m.jwtManager.VerifyAccessToken(token)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = pkgJWT.SetPayloadToContext(ctx, claims)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
