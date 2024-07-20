package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func Connect(opt *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(opt)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

type RedisClient struct {
	redis *redis.Client
}

func NewRedisClient(redis *redis.Client) *RedisClient {
	return &RedisClient{
		redis: redis,
	}
}

func (r RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	return r.redis.Set(ctx, key, value, time.Second*time.Duration(ONE_DAY)).Err()
}

func (r RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.redis.Get(ctx, key).Result()
}

func (r RedisClient) Del(ctx context.Context, key string) error {
	return r.redis.Del(ctx, key).Err()
}

func (r RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.redis.Expire(ctx, key, expiration).Err()
}

func (r RedisClient) Exists(ctx context.Context, key string) (int64, error) {
	return r.redis.Exists(ctx, key).Result()
}

func (r *RedisClient) Disconnect() error {
	return r.redis.Close()
}
