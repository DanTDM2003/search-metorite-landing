package usecase

import (
	"context"
	"errors"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (uc impleUsecase) GetMeteoriteLandings(ctx context.Context, input GetMeteoriteLandingsInput) (GetMeteoriteLandingsOutput, error) {
	mLs, pag, err := uc.repo.GetMetoriteLandings(ctx, repository.GetMeteoriteLandingsOptions{
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
	mL, err := uc.redis.GetMeteoriteLanding(ctx, input.ID)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			mL, err := uc.repo.GetOneMeteoriteLanding(ctx, repository.GetOneMeteoriteLandingOptions{
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

			if err := uc.redis.SetMeteoriteLanding(ctx, mL); err != nil {
				uc.l.Errorf(ctx, "meteorite_landings.usecase.GetOneMeteoriteLanding.redis.SetMeteoriteLanding: %v", err)
				return models.MeteoriteLanding{}, err
			}
			return mL, nil
		}
		uc.l.Errorf(ctx, "meteorite_landings.usecase.GetOneMeteoriteLanding.redis.GetMeteoriteLanding: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (uc impleUsecase) CreateMeteoriteLanding(ctx context.Context, input CreateMeteoriteLandingInput) (models.MeteoriteLanding, error) {
	_, err := uc.repo.GetOneMeteoriteLanding(ctx, repository.GetOneMeteoriteLandingOptions{
		Name: input.Name,
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Errorf(ctx, "meteorite_landings.usecase.CreateMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", err)
			return models.MeteoriteLanding{}, err
		}
	} else {
		uc.l.Warnf(ctx, "meteorite_landings.usecase.CreateMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", ErrMeteoriteLandingAlreadyExists)
		return models.MeteoriteLanding{}, ErrMeteoriteLandingAlreadyExists
	}

	mL, err := uc.repo.CreateMeteoriteLanding(ctx, repository.CreateMeteoriteLandingOptions{
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

	if err := uc.redis.SetMeteoriteLanding(ctx, mL); err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.usecase.CreateMeteoriteLanding.redis.SetMeteoriteLanding: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (uc impleUsecase) UpdateMeteoriteLanding(ctx context.Context, input UpdateMeteoriteLandingInput) (models.MeteoriteLanding, error) {
	mL, err := uc.repo.GetOneMeteoriteLanding(ctx, repository.GetOneMeteoriteLandingOptions{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "meteorite_landings.usecase.UpdateMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", ErrMeteoriteLandingsNotFound)
			return models.MeteoriteLanding{}, ErrMeteoriteLandingsNotFound
		}
		uc.l.Errorf(ctx, "meteorite_landings.usecase.UpdateMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", err)
		return models.MeteoriteLanding{}, err
	}

	if input.Name != "" {
		_, err = uc.repo.GetOneMeteoriteLanding(ctx, repository.GetOneMeteoriteLandingOptions{
			Name: input.Name,
		})
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				uc.l.Errorf(ctx, "meteorite_landings.usecase.UpdateMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", err)
				return models.MeteoriteLanding{}, err
			}
		} else {
			uc.l.Warnf(ctx, "meteorite_landings.usecase.UpdateMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", ErrMeteoriteLandingAlreadyExists)
			return models.MeteoriteLanding{}, ErrMeteoriteLandingAlreadyExists
		}
	}

	mL, err = uc.repo.UpdateMeteoriteLanding(ctx, repository.UpdateMeteoriteLandingOptions{
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
	}, mL)
	if err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.usecase.UpdateMeteoriteLanding.repo.UpdateMeteoriteLanding: %v", err)
		return models.MeteoriteLanding{}, err
	}

	if err := uc.redis.SetMeteoriteLanding(ctx, mL); err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.usecase.UpdateMeteoriteLanding.redis.SetMeteoriteLanding: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (uc impleUsecase) DeleteMeteoriteLanding(ctx context.Context, id uint) error {
	_, err := uc.repo.GetOneMeteoriteLanding(ctx, repository.GetOneMeteoriteLandingOptions{
		ID: id,
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			uc.l.Warnf(ctx, "meteorite_landings.usecase.DeleteMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", ErrMeteoriteLandingsNotFound)
			return ErrMeteoriteLandingsNotFound
		}
		uc.l.Errorf(ctx, "meteorite_landings.usecase.DeleteMeteoriteLanding.repo.GetOneMeteoriteLanding: %v", err)
		return err
	}

	if err := uc.repo.DeleteMeteoriteLanding(ctx, id); err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.usecase.DeleteMeteoriteLanding.repo.DeleteMeteoriteLanding: %v", err)
		return err
	}

	if err := uc.redis.DeleteMeteoriteLanding(ctx, id); err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.usecase.DeleteMeteoriteLanding.redis.DeleteMeteoriteLanding: %v", err)
		return err
	}

	return nil
}
