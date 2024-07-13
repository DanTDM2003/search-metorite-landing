package usecase

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"gorm.io/gorm"
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

func (uc impleUsecase) GetOneMeteoriteLanding(ctx context.Context, input GetOneMeteoriteLandingInput) (models.MeteoriteLanding, error) {
	mL, err := uc.repo.GetOneMeteoriteLanding(ctx, repository.GetOneMeteoriteLandingOption{
		ID: input.ID,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.l.Warnf(ctx, "meteorite_landings.usecase.GetOneMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", ErrMeteoriteLandingsNotFound)
			return models.MeteoriteLanding{}, ErrMeteoriteLandingsNotFound
		}
		uc.l.Errorf(ctx, "meteorite_landings.usecase.GetOneMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (uc impleUsecase) CreateMeteoriteLanding(ctx context.Context, input CreateMeteoriteLandingInput) (models.MeteoriteLanding, error) {
	mL, err := uc.repo.CreateMeteoriteLanding(ctx, repository.CreateMeteoriteLandingOption{
		Name:     input.Name,
		NameType: input.NameType,
		Year:     input.Year,
		Mass:     input.Mass,
		Recclass: input.Recclass,
		Fall:     input.Fall,
		Reclat:   input.Reclat,
		Reclong:  input.Reclong,
		GeoLocation: repository.GeoLocation{
			Latitude:      input.GeoLocation.Latitude,
			Longitude:     input.GeoLocation.Longitude,
			NeedsRecoding: input.GeoLocation.NeedsRecoding,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.usecase.CreateMeteoriteLanding.repo.CreateMeteoriteLanding: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (uc impleUsecase) UpdateMeteoriteLanding(ctx context.Context, input UpdateMeteoriteLandingInput) (models.MeteoriteLanding, error) {
	mL, err := uc.repo.UpdateMeteoriteLanding(ctx, repository.UpdateMeteoriteLandingOption{
		ID:       input.ID,
		Name:     input.Name,
		NameType: input.NameType,
		Year:     input.Year,
		Mass:     input.Mass,
		Recclass: input.Recclass,
		Fall:     input.Fall,
		Reclat:   input.Reclat,
		Reclong:  input.Reclong,
		GeoLocation: repository.GeoLocation{
			Latitude:      input.GeoLocation.Latitude,
			Longitude:     input.GeoLocation.Longitude,
			NeedsRecoding: input.GeoLocation.NeedsRecoding,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.usecase.UpdateMeteoriteLanding.repo.UpdateMeteoriteLanding: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}
