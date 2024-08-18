package database

import (
	"strings"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
	"gorm.io/gorm"
)

func (repo impleRepository) buildGetMeteoriteLandingsCondition(cursor *gorm.DB, opt repository.GetMeteoriteLandingsOption) *gorm.DB {
	var conditions []string
	var params []interface{}

	if !opt.Year.IsZero() {
		conditions = append(conditions, "year = ?")
		params = append(params, opt.Year)
	}

	if opt.Recclass != "" {
		conditions = append(conditions, "recclass = ?")
		params = append(params, opt.Recclass)
	}

	if opt.Mass != 0 {
		conditions = append(conditions, "mass = ?")
		params = append(params, opt.Mass)
	}

	if len(conditions) > 0 {
		cursor = cursor.Where(strings.Join(conditions, " AND "), params...)
	}

	return cursor
}

func (repo impleRepository) buildGetOneMeteoriteLandingCondition(cursor *gorm.DB, opt repository.GetOneMeteoriteLandingOption) *gorm.DB {
	var conditions []string
	var params []interface{}

	if opt.ID != 0 {
		conditions = append(conditions, "id = ?")
		params = append(params, opt.ID)
	}

	if opt.Name != "" {
		conditions = append(conditions, "name = ?")
		params = append(params, opt.Name)
	}

	if len(conditions) > 0 {
		cursor = cursor.Where(strings.Join(conditions, " AND "), params...)
	}

	return cursor
}
