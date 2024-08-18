package database

import (
	"strings"

	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	"gorm.io/gorm"
)

func (repo impleRepository) buildGetUsersCondition(cursor *gorm.DB, opt repository.GetUsersOptions) *gorm.DB {
	var conditions []string
	var params []interface{}

	if opt.Username != "" {
		conditions = append(conditions, "username = ?")
		params = append(params, opt.Username)
	}

	if opt.Role != "" {
		conditions = append(conditions, "role = ?")
		params = append(params, opt.Role)
	}

	if len(conditions) > 0 {
		cursor = cursor.Where(strings.Join(conditions, " AND "), params...)
	}

	return cursor
}

func (repo impleRepository) buildGetOneUserCondition(cursor *gorm.DB, opt repository.GetOneUserOptions) *gorm.DB {
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

	if len(conditions) > 0 {
		cursor = cursor.Where(strings.Join(conditions, " AND "), params...)
	}

	return cursor
}
