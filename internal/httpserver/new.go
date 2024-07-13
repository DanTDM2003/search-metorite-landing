package httpserver

import (
	"database/sql"

	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	l        pkgLog.Logger
	gin      *gin.Engine
	port     int
	database *sql.DB
}

type Config struct {
	Port     int
	Database *sql.DB
}

func New(l pkgLog.Logger, cfg Config) *HTTPServer {
	return &HTTPServer{
		l:        l,
		gin:      gin.Default(),
		port:     cfg.Port,
		database: cfg.Database,
	}
}
