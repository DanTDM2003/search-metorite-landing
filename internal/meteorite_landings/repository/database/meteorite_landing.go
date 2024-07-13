package database

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	meteoriteLandingsTable = "meteorite_landings"
)

func (repo impleRepository) getTable() *gorm.DB {
	return repo.db.Table(meteoriteLandingsTable)
}

func (repo impleRepository) GetMetoriteLandings(ctx context.Context, opt repository.GetMeteoriteLandingsOption) ([]models.MeteoriteLanding, paginator.Paginator, error) {
	table := repo.getTable()
	var mLs []models.MeteoriteLanding
	var total int64

	// Count total records
	if err := table.Count(&total).Error; err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.database.GetMeteoriteLandings.db.Count: %v", err)
		return nil, paginator.Paginator{}, err
	}

	cursor := table.
		Limit(int(opt.Limit)).
		Offset(int(opt.PaginatorQuery.Offset()))

	if err := cursor.Find(&mLs).Error; err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.database.GetMeteoriteLandings.db.Find: %v", err)
		return nil, paginator.Paginator{}, err
	}

	return mLs, paginator.Paginator{
		Total:       total,
		Count:       int64(len(mLs)),
		PerPage:     opt.Limit,
		CurrentPage: opt.Page,
	}, nil
}

func (uc impleRepository) GetOneMeteoriteLanding(ctx context.Context, opt repository.GetOneMeteoriteLandingOption) (models.MeteoriteLanding, error) {
	table := uc.getTable()
	var mL models.MeteoriteLanding

	if err := table.Where("id = ?", opt.ID).First(&mL).Error; err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.repository.database.GetOneMeteoriteLanding.db.First: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (uc impleRepository) CreateMeteoriteLanding(ctx context.Context, opt repository.CreateMeteoriteLandingOption) (models.MeteoriteLanding, error) {
	table := uc.getTable()

	mL := uc.buildCreateMeteoriteLandingModel(opt)

	if err := table.Create(&mL).Error; err != nil {
		uc.l.Errorf(ctx, "meteorite_landings.repository.database.CreateMeteoriteLanding.db.Create: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (repo impleRepository) UpdateMeteoriteLanding(ctx context.Context, opt repository.UpdateMeteoriteLandingOption) (models.MeteoriteLanding, error) {
	table := repo.getTable()

	mL := repo.buildUpdateMeteoriteLandingModel(opt)

	if err := table.Where("id = ?", opt.ID).Updates(&mL).Error; err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.database.UpdateMeteoriteLanding.db.Updates: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}
