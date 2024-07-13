package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type Repository interface {
	GetMetoriteLandings(ctx context.Context, opt GetMeteoriteLandingsOption) ([]models.MeteoriteLanding, paginator.Paginator, error)
}
