package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type GetMeteoriteLandingsInput struct {
	paginator.PaginatorQuery
}

type GetMeteoriteLandingsOutput struct {
	MeteoriteLandings []models.MeteoriteLanding
	Paginator         paginator.Paginator
}
