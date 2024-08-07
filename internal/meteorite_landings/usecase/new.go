package usecase

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Usecase interface {
	GetMeteoriteLandings(ctx context.Context, input GetMeteoriteLandingsInput) (GetMeteoriteLandingsOutput, error)
	GetOneMeteoriteLanding(ctx context.Context, input GetOneMeteoriteLandingInput) (models.MeteoriteLanding, error)
	CreateMeteoriteLanding(ctx context.Context, input CreateMeteoriteLandingInput) (models.MeteoriteLanding, error)
	UpdateMeteoriteLanding(ctx context.Context, input UpdateMeteoriteLandingInput) (models.MeteoriteLanding, error)
	DeleteMeteoriteLanding(ctx context.Context, id uint) error
}

type impleUsecase struct {
	l     pkgLog.Logger
	repo  repository.Repository
	redis repository.RedisRepository
}

func New(
	l pkgLog.Logger,
	repo repository.Repository,
	redisRepo repository.RedisRepository,
) Usecase {
	return &impleUsecase{
		l:     l,
		repo:  repo,
		redis: redisRepo,
	}
}
