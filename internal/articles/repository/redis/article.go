package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

func (redis impleRedisRepository) GetArticle(ctx context.Context, slug string) (models.Article, error) {
	key := "article:" + slug
	val, err := redis.redis.Get(ctx, key)
	if err != nil {
		redis.l.Errorf(ctx, "articles.repository.redis.GetArticle.redis.Get: %v", err)
		return models.Article{}, err
	}

	var article models.Article
	err = json.Unmarshal([]byte(val), &article)
	if err != nil {
		redis.l.Errorf(ctx, "articles.repository.redis.GetArticle.json.Unmarshal: %v", err)
		return models.Article{}, err
	}

	return article, nil
}

func (redis impleRedisRepository) SetArticle(ctx context.Context, article models.Article) error {
	val, err := json.Marshal(article)
	if err != nil {
		redis.l.Errorf(ctx, "articles.repository.redis.SetArticle.json.Marshal: %v", err)
		return err
	}

	key := fmt.Sprintf("article:%d", article.ID)
	err = redis.redis.Set(ctx, key, val)
	if err != nil {
		redis.l.Errorf(ctx, "articles.repository.redis.SetArticle.redis.Set: %v", err)
		return err
	}

	return nil
}

func (redis impleRedisRepository) DeleteArticle(ctx context.Context, id uint) error {
	key := fmt.Sprintf("article:%d", id)
	err := redis.redis.Del(ctx, key)
	if err != nil {
		redis.l.Errorf(ctx, "articles.repository.redis.DeleteArticle.redis.Del: %v", err)
		return err
	}

	return nil
}
