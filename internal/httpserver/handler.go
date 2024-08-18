package httpserver

import (
	articleHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/articles/delivery/http"
	articleDB "github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository/database"
	articleRedis "github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository/redis"
	articleUC "github.com/DanTDM2003/search-api-docker-redis/internal/articles/usecase"

	mLHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/delivery/http"
	mLDB "github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository/database"
	mLRedis "github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository/redis"
	mLUC "github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"

	postHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/posts/delivery/http"
	postDB "github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository/database"
	postRedis "github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository/redis"
	postUC "github.com/DanTDM2003/search-api-docker-redis/internal/posts/usecase"

	userHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/users/delivery/http"
	userDB "github.com/DanTDM2003/search-api-docker-redis/internal/users/repository/database"
	userRedis "github.com/DanTDM2003/search-api-docker-redis/internal/users/repository/redis"
	userUC "github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"

	sessionHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/session/delivery/http"
	sessionUC "github.com/DanTDM2003/search-api-docker-redis/internal/session/usecase"

	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
)

func (srv HTTPServer) mapHandlers() error {
	srv.gin.Use(middleware.Recovery())

	// JWT
	jwtManager := pkgJWT.New(srv.secretKey, srv.redis)

	// MeteoriteLanding
	meteoriteLandingRepo := mLDB.New(srv.l, srv.database)
	meteoriteLandingRedisRepo := mLRedis.New(srv.l, srv.redis)
	meteoriteLandingUC := mLUC.New(srv.l, meteoriteLandingRepo, meteoriteLandingRedisRepo)
	meteoriteLandingH := mLHTTP.New(srv.l, meteoriteLandingUC)

	// User
	userRepo := userDB.New(srv.l, srv.database)
	userRedisRepo := userRedis.New(srv.l, srv.redis)
	userUC := userUC.New(srv.l, userRepo, userRedisRepo)
	userH := userHTTP.New(srv.l, userUC)

	// Post
	postRepo := postDB.New(srv.l, srv.database)
	postRedisRepo := postRedis.New(srv.l, srv.redis)
	postUC := postUC.New(srv.l, postRepo, postRedisRepo, userUC)
	postH := postHTTP.New(srv.l, postUC)

	// Article
	articleRepo := articleDB.New(srv.l, srv.database)
	articleRedisRepo := articleRedis.New(srv.l, srv.redis)
	articleUC := articleUC.New(srv.l, articleRepo, articleRedisRepo)
	articleH := articleHTTP.New(srv.l, articleUC)

	// Session
	sessionUC := sessionUC.New(srv.l, userUC, jwtManager)
	sessionH := sessionHTTP.New(srv.l, sessionUC)

	// Middleware
	mw := middleware.New(srv.l, jwtManager, userUC)

	// API Routes
	api := srv.gin.Group("api/v1")
	mLHTTP.MapMeteoriteLandingRoutes(api.Group("meteorite-landings"), meteoriteLandingH, mw)
	userHTTP.MapUserRoutes(api.Group("users"), userH, mw)
	postHTTP.MapPostRoutes(api.Group("posts"), postH, mw)
	articleHTTP.MapArticleRoutes(api.Group("articles"), articleH)
	sessionHTTP.MapSessionRoutes(api.Group("session"), sessionH, mw)

	return nil
}
