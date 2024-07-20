package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type RedisRepository interface {
	SetMeteoriteLanding(ctx context.Context, mL models.MeteoriteLanding) error
	GetMeteoriteLanding(ctx context.Context, id uint) (models.MeteoriteLanding, error)
	DeleteMeteoriteLanding(ctx context.Context, id uint) error
}
