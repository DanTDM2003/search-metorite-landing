package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

func (redisRepo impleRedisRepository) GetPost(ctx context.Context, id uint) (models.Post, error) {
	key := fmt.Sprintf("post:%d", id)
	val, err := redisRepo.redis.Get(ctx, key)
	if err != nil {
		redisRepo.l.Errorf(ctx, "posts.repository.redis.GetPost.redis.Get: %v", err)
		return models.Post{}, err
	}

	var post models.Post
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		redisRepo.l.Errorf(ctx, "posts.repository.redis.GetPost.json.Unmarshal: %v", err)
		return models.Post{}, err
	}

	return post, nil
}

func (redisRepo impleRedisRepository) SetPost(ctx context.Context, post models.Post) error {
	val, err := json.Marshal(post)
	if err != nil {
		redisRepo.l.Errorf(ctx, "posts.repository.redis.SetPost.json.Marshal: %v", err)
		return err
	}

	key := fmt.Sprintf("post:%d", post.ID)
	err = redisRepo.redis.Set(ctx, key, val)
	if err != nil {
		redisRepo.l.Errorf(ctx, "posts.repository.redis.SetPost.redis.Set: %v", err)
		return err
	}

	return nil
}

func (redisRepo impleRedisRepository) DeletePost(ctx context.Context, id uint) error {
	key := fmt.Sprintf("post:%d", id)
	err := redisRepo.redis.Del(ctx, key)
	if err != nil {
		redisRepo.l.Errorf(ctx, "posts.repository.redis.DeletePost.redis.Del: %v", err)
		return err
	}

	return nil
}
