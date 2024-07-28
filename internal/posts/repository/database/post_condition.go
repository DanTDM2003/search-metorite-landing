package database

import (
	"strings"

	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/repository"
)

func (repo impleRepository) buildGetOnePostCondition(opt repository.GetOnePostOptions) (string, interface{}) {
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

	condition := strings.Join(conditions, " AND ")

	return condition, params
}

func (repo impleRepository) buildDeletePostsCondition(opt repository.DeletePostsOptions) (string, interface{}) {
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

	condition := strings.Join(conditions, " AND ")

	return condition, params
}
