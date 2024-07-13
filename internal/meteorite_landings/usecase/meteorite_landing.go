package usecase

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
)

func (uc impleUsecase) GetMeteoriteLandings(ctx context.Context, input GetMeteoriteLandingsInput) (GetMeteoriteLandingsOutput, error) {
	mLs, pag, err := uc.repo.GetMetoriteLandings(ctx, repository.GetMeteoriteLandingsOption{
		PaginatorQuery: input.PaginatorQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.usecase.GetMeteoriteLandings.repo.GetMetoriteLandings: %v", err)
		return GetMeteoriteLandingsOutput{}, err
	}

	return GetMeteoriteLandingsOutput{
		MeteoriteLandings: mLs,
		Paginator:         pag,
	}, err
}
