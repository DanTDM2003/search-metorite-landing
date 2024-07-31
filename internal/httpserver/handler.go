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
	"github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	postHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/posts/delivery/http"
	postDB "github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository/database"
	postRedis "github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository/redis"
	postUC "github.com/DanTDM2003/search-api-docker-redis/internal/posts/usecase"
	userHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/users/delivery/http"
	userDB "github.com/DanTDM2003/search-api-docker-redis/internal/users/repository/database"
	userRedis "github.com/DanTDM2003/search-api-docker-redis/internal/users/repository/redis"
	userUC "github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
)

func (srv HTTPServer) mapHandlers() error {
	srv.gin.Use(middleware.Recovery())

	jwtManager := pkgJWT.New(srv.secretKey, srv.database)

	middlewares := middleware.New(srv.l, jwtManager)

	meteoriteLandingRepo := mLDB.New(srv.l, srv.database)
	meteoriteLandingRedisRepo := mLRedis.New(srv.l, srv.redis)
	meteoriteLandingUC := mLUC.New(srv.l, meteoriteLandingRepo, meteoriteLandingRedisRepo)
	meteoriteLandingH := mLHTTP.New(srv.l, meteoriteLandingUC)

	userRepo := userDB.New(srv.l, srv.database)
	userRedisRepo := userRedis.New(srv.l, srv.redis)
	userUC := userUC.New(srv.l, userRepo, userRedisRepo)
	userH := userHTTP.New(srv.l, userUC)

	postRepo := postDB.New(srv.l, srv.database)
	postRedisRepo := postRedis.New(srv.l, srv.redis)
	postUC := postUC.New(srv.l, postRepo, postRedisRepo)
	postH := postHTTP.New(srv.l, postUC)

	articleRepo := articleDB.New(srv.l, srv.database)
	articleRedisRepo := articleRedis.New(srv.l, srv.redis)
	articleUC := articleUC.New(srv.l, articleRepo, articleRedisRepo)
	articleH := articleHTTP.New(srv.l, articleUC)

	api := srv.gin.Group("api/v1")
	mLHTTP.MapMeteoriteLandingRoutes(api.Group("meteorite-landings"), meteoriteLandingH, middlewares)
	userHTTP.MapUserRoutes(api.Group("users"), userH, middlewares)
	postHTTP.MapPostRoutes(api.Group("posts"), postH, middlewares)
	articleHTTP.MapArticleRoutes(api.Group("articles"), articleH)

	return nil
}
