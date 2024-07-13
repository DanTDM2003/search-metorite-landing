package main

import (
	"log"

	"github.com/DanTDM2003/search-api-docker-redis/config"
	"github.com/DanTDM2003/search-api-docker-redis/internal/appconfig/postgres"
	"github.com/DanTDM2003/search-api-docker-redis/internal/httpserver"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Could not load the configuration: %v", err)
		panic(err)
	}

	conn, err := postgres.Connect(cfg.Postgres)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		panic(err)
	}
	defer postgres.Close(conn.DB)

	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	srv := httpserver.New(l, httpserver.Config{
		Port:     cfg.HTTPServer.Port,
		Database: conn.DB,
	})
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
