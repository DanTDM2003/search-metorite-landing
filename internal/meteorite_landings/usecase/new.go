package usecase

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Usecase interface {
	GetMeteoriteLandings(ctx context.Context, input GetMeteoriteLandingsInput) (GetMeteoriteLandingsOutput, error)
}

type impleUsecase struct {
	l    pkgLog.Logger
	repo repository.Repository
}

func New(l pkgLog.Logger, repo repository.Repository) Usecase {
	return &impleUsecase{
		l:    l,
		repo: repo,
	}
}
