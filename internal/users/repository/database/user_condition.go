package database

import (
	"strings"

	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
)

func (repo impleRepository) buildGetOneUserCondition(opt repository.GetOneUserOptions) (string, interface{}) {
	var conditions []string
	var params []interface{}

	if opt.ID != 0 {
		conditions = append(conditions, "id = ?")
		params = append(params, opt.ID)
	}

	if opt.Email != "" {
		conditions = append(conditions, "email = ?")
		params = append(params, opt.Email)
	}

	if opt.Username != "" {
		conditions = append(conditions, "username = ?")
		params = append(params, opt.Username)
	}

	condition := strings.Join(conditions, " AND ")

	return condition, params
}
