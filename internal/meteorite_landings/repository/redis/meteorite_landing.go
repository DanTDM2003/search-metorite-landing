package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

func (redisRepo impleRedisRepository) SetMeteoriteLanding(ctx context.Context, mL models.MeteoriteLanding) error {
	val, err := json.Marshal(mL)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.SetMeteoriteLanding.json.Marshal: %v", err)
		return err
	}

	key := fmt.Sprintf("meteorite_landing:%d", mL.ID)
	err = redisRepo.redis.Set(ctx, key, val)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.GetMeteoriteLanding.redis.Set: %v", err)
		return err
	}

	return nil
}

func (redisRepo impleRedisRepository) GetMeteoriteLanding(ctx context.Context, id uint) (models.MeteoriteLanding, error) {
	key := fmt.Sprintf("meteorite_landing:%d", id)
	val, err := redisRepo.redis.Get(ctx, key)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.GetMeteoriteLanding.redis.Get: %v", err)
		return models.MeteoriteLanding{}, err
	}

	var mL models.MeteoriteLanding
	err = json.Unmarshal([]byte(val), &mL)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.GetMeteoriteLanding.json.Unmarshal: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (redisRepo impleRedisRepository) DeleteMeteoriteLanding(ctx context.Context, id uint) error {
	key := fmt.Sprintf("meteorite_landing:%d", id)
	err := redisRepo.redis.Del(ctx, key)
	if err != nil {
		redisRepo.l.Errorf(ctx, "redis.repository.DeleteMeteoriteLanding.redis.Del: %v", err)
		return err
	}

	return nil
}
