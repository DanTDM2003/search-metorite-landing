package redis

import (
	"github.com/DanTDM2003/search-api-docker-redis/config"
	"github.com/redis/go-redis/v9"
)

type RedisOptions struct {
	opt *redis.Options
}

func NewRedisOptions() RedisOptions {
	return RedisOptions{
		opt: &redis.Options{},
	}
}

func (redisOpt RedisOptions) SetOptions(opt config.RedisConfig) *redis.Options {
	redisOpt.opt.Addr = opt.Addr
	redisOpt.opt.Password = opt.Password
	redisOpt.opt.DB = opt.DB
	return redisOpt.opt
}
