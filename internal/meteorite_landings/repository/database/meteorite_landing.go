package database

import (
	"context"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"gorm.io/gorm"
)

const (
	meteoriteLandingsTable = "meteorite_landings"
)

func (repo impleRepository) getTable() *gorm.DB {
	return repo.db.Table(meteoriteLandingsTable)
}

func (repo impleRepository) GetMetoriteLandings(ctx context.Context, opt repository.GetMeteoriteLandingsOptions) ([]models.MeteoriteLanding, paginator.Paginator, error) {
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

	cursor = repo.buildGetMeteoriteLandingsCondition(cursor, opt)

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

func (repo impleRepository) GetOneMeteoriteLanding(ctx context.Context, opt repository.GetOneMeteoriteLandingOptions) (models.MeteoriteLanding, error) {
	table := repo.getTable()

	cursor := repo.buildGetOneMeteoriteLandingCondition(table, opt)

	var mL models.MeteoriteLanding
	if err := cursor.First(&mL).Error; err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.database.GetOneMeteoriteLanding.db.First: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (repo impleRepository) CreateMeteoriteLanding(ctx context.Context, opt repository.CreateMeteoriteLandingOptions) (models.MeteoriteLanding, error) {
	table := repo.getTable()

	mL := repo.buildCreateMeteoriteLandingModel(opt)

	if err := table.Create(&mL).Error; err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.database.CreateMeteoriteLanding.db.Create: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return mL, nil
}

func (repo impleRepository) UpdateMeteoriteLanding(ctx context.Context, opt repository.UpdateMeteoriteLandingOptions, mL models.MeteoriteLanding) (models.MeteoriteLanding, error) {
	table := repo.getTable()

	update := repo.buildUpdateMeteoriteLandingModel(opt, mL)

	if err := table.Where("id = ?", update.ID).Updates(&update).Error; err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.database.UpdateMeteoriteLanding.db.Updates: %v", err)
		return models.MeteoriteLanding{}, err
	}

	return update, nil
}

func (repo impleRepository) DeleteMeteoriteLanding(ctx context.Context, id uint) error {
	table := repo.getTable()

	if err := table.Where("id = ?", id).Delete(&models.MeteoriteLanding{}).Error; err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.database.DeleteMeteoriteLanding.db.Delete: %v", err)
		return err
	}

	return nil
}
