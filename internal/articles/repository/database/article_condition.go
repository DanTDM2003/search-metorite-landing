package database

import (
	"strings"

	"github.com/DanTDM2003/search-api-docker-redis/internal/articles/repository"
)

func (repo impleRepository) buildGetAritclesCondition(opt repository.GetArticlesOptions) (string, interface{}) {
	var conditions []string
	var params []interface{}

	if opt.Author != "" {
		conditions = append(conditions, "author = ?")
		params = append(params, opt.Author)
	}

	condition := strings.Join(conditions, " AND ")

	return condition, params
}

func (repo impleRepository) buildGetOneArticleCondition(opt repository.GetOneArticleOptions) (string, interface{}) {
	var conditions []string
	var params []interface{}

	if opt.ID != 0 {
		conditions = append(conditions, "id = ?")
		params = append(params, opt.ID)
	}

	if opt.Slug != "" {
		conditions = append(conditions, "slug = ?")
		params = append(params, opt.Slug)
	}

	if opt.Author != "" {
		conditions = append(conditions, "author = ?")
		params = append(params, opt.Author)
	}

	condition := strings.Join(conditions, " AND ")

	return condition, params
}
