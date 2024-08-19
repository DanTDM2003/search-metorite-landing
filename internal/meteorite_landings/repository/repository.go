package repository

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type Repository interface {
	GetMetoriteLandings(ctx context.Context, opt GetMeteoriteLandingsOptions) ([]models.MeteoriteLanding, paginator.Paginator, error)
	GetOneMeteoriteLanding(ctx context.Context, opt GetOneMeteoriteLandingOptions) (models.MeteoriteLanding, error)
	CreateMeteoriteLanding(ctx context.Context, opt CreateMeteoriteLandingOptions) (models.MeteoriteLanding, error)
	UpdateMeteoriteLanding(ctx context.Context, opt UpdateMeteoriteLandingOptions, mL models.MeteoriteLanding) (models.MeteoriteLanding, error)
	DeleteMeteoriteLanding(ctx context.Context, id uint) error
}
