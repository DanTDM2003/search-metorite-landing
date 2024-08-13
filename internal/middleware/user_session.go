package middleware

import (
	"strconv"

	"github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (m Middleware) UserSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		payload, ok := pkgJWT.GetPayloadFromContext(ctx)
		if !ok {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		sessUserID, err := strconv.Atoi(payload.StandardClaims.Subject)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		sessUser, err := m.userUC.GetOneUser(ctx, usecase.GetOneUserInput{
			ID: uint(sessUserID),
		})
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		ctx = pkgJWT.SetUserToContext(ctx, sessUser)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
