package database

import (
	"strings"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
)

func (repo impleRepository) buildGetOneMeteoriteLandingCondition(opt repository.GetOneMeteoriteLandingOption) (string, interface{}) {
	var conditions []string
	var params []interface{}

	if opt.ID != 0 {
		conditions = append(conditions, "id = ?")
		params = append(params, opt.ID)
	}
	condition := strings.Join(conditions, " AND ")

	return condition, params
}
