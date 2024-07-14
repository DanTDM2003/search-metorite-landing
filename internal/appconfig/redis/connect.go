package redis

import (
	"github.com/DanTDM2003/search-api-docker-redis/config"

	pkgRedis "github.com/DanTDM2003/search-api-docker-redis/pkg/redis"
)

type Redis struct {
	RedisClient *pkgRedis.RedisClient
}

func Connect(cfg config.RedisConfig) (*Redis, error) {
	opt := pkgRedis.NewRedisOptions().SetOptions(cfg)

	client, err := pkgRedis.Connect(opt)
	if err != nil {
		return nil, err
	}

	cl := pkgRedis.NewRedisClient(client)

	return &Redis{
		RedisClient: cl,
	}, nil
}
