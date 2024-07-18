package main

import (
	"log"

	"github.com/DanTDM2003/search-api-docker-redis/config"
	"github.com/DanTDM2003/search-api-docker-redis/internal/appconfig/database"
	"github.com/DanTDM2003/search-api-docker-redis/internal/appconfig/redis"
	"github.com/DanTDM2003/search-api-docker-redis/internal/httpserver"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/utils"
)

func main() {
	//
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Could not load the configuration: %v", err)
		panic(err)
	}

	conn, err := database.Connect(cfg.Postgres)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		panic(err)
	}
	defer database.Close(conn.DB)

	exists := conn.DB.Migrator().HasTable(&models.MeteoriteLanding{})
	if !exists {
		utils.InitDatabase(conn)
	}

	redis, err := redis.Connect(cfg.Redis)
	if err != nil {
		log.Fatalf("Could not connect to the redis: %v", err)
		panic(err)
	}
	defer redis.RedisClient.Disconnect()

	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	srv := httpserver.New(l, httpserver.Config{
		Port:     cfg.HTTPServer.Port,
		Database: conn.DB,
		Redis:    redis.RedisClient,
	})
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
