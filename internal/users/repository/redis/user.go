package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

func (redisRepo impleRedisRepository) GetUser(ctx context.Context, id uint) (models.User, error) {
	key := fmt.Sprintf("user:%d", id)
	val, err := redisRepo.redis.Get(ctx, key)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.GetUser.redis.Get: %v", err)
		return models.User{}, err
	}

	var user models.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.GetUser.json.Unmarshal: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (redisRepo impleRedisRepository) SetUser(ctx context.Context, user models.User) error {
	val, err := json.Marshal(user)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.SetUser.json.Marshal: %v", err)
		return err
	}

	key := fmt.Sprintf("user:%d", user.ID)
	err = redisRepo.redis.Set(ctx, key, val)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.SetUser.redis.Set: %v", err)
		return err
	}

	return nil
}

func (redisRepo impleRedisRepository) DeleteUser(ctx context.Context, id uint) error {
	key := fmt.Sprintf("user:%d", id)
	err := redisRepo.redis.Del(ctx, key)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.DeleteUser.redis.Del: %v", err)
		return err
	}

	return nil
}
