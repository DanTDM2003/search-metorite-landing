package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

type Repository interface {
	GetMetoriteLandings(ctx context.Context) ([]models.MeteoriteLanding, error)
}
