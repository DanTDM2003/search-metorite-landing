package httpserver

import (
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	pkgRedis "github.com/DanTDM2003/search-api-docker-redis/pkg/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HTTPServer struct {
	l         pkgLog.Logger
	gin       *gin.Engine
	port      int
	database  *gorm.DB
	redis     *pkgRedis.RedisClient
	secretKey string
}

type Config struct {
	Port      int
	Database  *gorm.DB
	Redis     *pkgRedis.RedisClient
	SecretKey string
}

func New(l pkgLog.Logger, cfg Config) *HTTPServer {
	return &HTTPServer{
		l:         l,
		gin:       gin.Default(),
		port:      cfg.Port,
		database:  cfg.Database,
		redis:     cfg.Redis,
		secretKey: cfg.SecretKey,
	}
}
