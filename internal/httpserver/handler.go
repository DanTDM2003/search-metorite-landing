package httpserver

import (
	mLHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/delivery/http"
	mLDB "github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository/database"
	mLRedis "github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository/redis"
	mLUC "github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"
	middlware "github.com/DanTDM2003/search-api-docker-redis/internal/middleware"
	postHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/posts/delivery/http"
	postDB "github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository/database"
	postRedis "github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository/redis"
	postUC "github.com/DanTDM2003/search-api-docker-redis/internal/posts/usecase"
	userHTTP "github.com/DanTDM2003/search-api-docker-redis/internal/users/delivery/http"
	userDB "github.com/DanTDM2003/search-api-docker-redis/internal/users/repository/database"
	userRedis "github.com/DanTDM2003/search-api-docker-redis/internal/users/repository/redis"
	userUC "github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
)

func (srv HTTPServer) mapHandlers() error {
	srv.gin.Use(middlware.Recovery())

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

	api := srv.gin.Group("api/v1")
	mLHTTP.MapMeteoriteLandingRoutes(api.Group("meteorite-landings"), meteoriteLandingH)
	userHTTP.MapUserRoutes(api.Group("users"), userH)
	postHTTP.MapPostRoutes(api.Group("posts"), postH)

	return nil
}
