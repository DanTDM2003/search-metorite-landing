package database

import (
	"strings"

	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository"
	"gorm.io/gorm"
)

func (repo impleRepository) buildGetPostsCondition(cursor *gorm.DB, opt repository.GetPostsOptions) *gorm.DB {
	var conditions []string
	var params []interface{}

	if opt.AuthorID != 0 {
		conditions = append(conditions, "author_id = ?")
		params = append(params, opt.AuthorID)
	}

	if len(conditions) > 0 {
		cursor = cursor.Where(strings.Join(conditions, " AND "), params...)
	}

	return cursor
}

func (repo impleRepository) buildGetOnePostCondition(cursor *gorm.DB, opt repository.GetOnePostOptions) *gorm.DB {
	var conditions []string
	var params []interface{}

	if opt.ID != 0 {
		conditions = append(conditions, "id = ?")
		params = append(params, opt.ID)
	}

	if opt.AuthorID != 0 {
		conditions = append(conditions, "author_id = ?")
		params = append(params, opt.AuthorID)
	}

	if len(conditions) > 0 {
		cursor = cursor.Where(strings.Join(conditions, " AND "), params...)
	}

	return cursor
}

func (repo impleRepository) buildDeletePostsCondition(cursor *gorm.DB, opt repository.DeletePostsOptions) *gorm.DB {
	var conditions []string
	var params []interface{}

	if len(opt.IDs) > 0 {
		conditions = append(conditions, "id IN (?)")
		params = append(params, opt.IDs)
	}

	if len(opt.AuthorIDs) > 0 {
		conditions = append(conditions, "author_id IN (?)")
		params = append(params, opt.AuthorIDs)
	}

	if len(conditions) > 0 {
		cursor = cursor.Where(strings.Join(conditions, " AND "), params...)
	}

	return cursor
}
