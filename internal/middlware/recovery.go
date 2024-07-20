package middlware

import (
	"log"

	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[Panic Recovered] %v\n", r)
				log.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
				response.PanicError(c, r)
			}
		}()
		c.Next()
	}
}
