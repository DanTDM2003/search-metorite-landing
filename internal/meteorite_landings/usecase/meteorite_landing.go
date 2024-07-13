package usecase

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

func (uc impleUsecase) GetMeteoriteLandings(ctx context.Context) ([]models.MeteoriteLanding, error) {
	mLs, err := uc.repo.GetMetoriteLandings(ctx)
	if err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.usecase.GetMeteoriteLandings.repo.GetMetoriteLandings: %v", err)
		return nil, err
	}

	return mLs, err
}
