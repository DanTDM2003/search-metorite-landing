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
	var meteoriteLandings []models.MeteoriteLanding
	var total int64

	// Count total records
	if err := table.Count(&total).Error; err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.database.GetMeteoriteLandings.db.Count: %v", err)
		return nil, paginator.Paginator{}, err
	}

	cursor := table.
		Limit(int(opt.Limit)).
		Offset(int(opt.PaginatorQuery.Offset()))

	if err := cursor.Find(&meteoriteLandings).Error; err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.database.GetMeteoriteLandings.db.Find: %v", err)
		return nil, paginator.Paginator{}, err
	}

	return meteoriteLandings, paginator.Paginator{
		Total:       total,
		Count:       int64(len(meteoriteLandings)),
		PerPage:     opt.Limit,
		CurrentPage: opt.Page,
	}, nil
}
