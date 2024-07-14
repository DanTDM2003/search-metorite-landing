package httpserver

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/delivery/http"
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository/database"
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository/redis"
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"
)

func (srv HTTPServer) mapHandlers() error {
	meteoriteLandingRepo := database.New(srv.l, srv.database)
	meteoriteLandingRedisRepo := redis.New(srv.l, srv.redis)
	meteoriteLandingUC := usecase.New(srv.l, meteoriteLandingRepo, meteoriteLandingRedisRepo)
	meteoriteLandingH := http.New(srv.l, meteoriteLandingUC)

	api := srv.gin.Group("api/v1")
	http.MapMeteoriteLandingRoutes(api, meteoriteLandingH)

	return nil
}
