package main

import (
	"log"

	"github.com/DanTDM2003/search-api-docker-redis/config"
	"github.com/DanTDM2003/search-api-docker-redis/internal/appconfig/database"
	"github.com/DanTDM2003/search-api-docker-redis/internal/appconfig/redis"
	"github.com/DanTDM2003/search-api-docker-redis/internal/httpserver"
	pkgDatabase "github.com/DanTDM2003/search-api-docker-redis/pkg/database"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

func main() {
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

	err = pkgDatabase.InitDatabase(conn)
	if err != nil {
		log.Fatalf("Could not initialize the database: %v", err)
		panic(err)
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
		log.Fatalf("Could not run the server: %v", err)
		panic(err)
	}
}
