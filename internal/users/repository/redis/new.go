package redis

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	pkgRedis "github.com/DanTDM2003/search-api-docker-redis/pkg/redis"
)

type impleRedisRepository struct {
	l     pkgLog.Logger
	redis *pkgRedis.RedisClient
}

func New(l pkgLog.Logger, redis *pkgRedis.RedisClient) repository.RedisRepository {
	return impleRedisRepository{
		l:     l,
		redis: redis,
	}
}
