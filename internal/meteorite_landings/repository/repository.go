package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type Repository interface {
	GetMetoriteLandings(ctx context.Context, opt GetMeteoriteLandingsOption) ([]models.MeteoriteLanding, paginator.Paginator, error)
	GetOneMeteoriteLanding(ctx context.Context, opt GetOneMeteoriteLandingOption) (models.MeteoriteLanding, error)
	CreateMeteoriteLanding(ctx context.Context, opt CreateMeteoriteLandingOption) (models.MeteoriteLanding, error)
	UpdateMeteoriteLanding(ctx context.Context, opt UpdateMeteoriteLandingOption) (models.MeteoriteLanding, error)
}
