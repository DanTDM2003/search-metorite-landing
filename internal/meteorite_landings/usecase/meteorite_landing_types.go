package usecase

import (
	"time"

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

type GetOneMeteoriteLandingInput struct {
	ID uint
}

type GeoLocation struct {
	Latitude      float64
	Longitude     float64
	NeedsRecoding bool
}

type CreateMeteoriteLandingInput struct {
	Name        string
	NameType    string
	Year        time.Time
	Mass        float64
	Recclass    string
	Fall        string
	Reclat      float64
	Reclong     float64
	GeoLocation GeoLocation
}

type UpdateMeteoriteLandingInput struct {
	ID          uint
	Name        string
	NameType    string
	Year        time.Time
	Mass        float64
	Recclass    string
	Fall        string
	Reclat      float64
	Reclong     float64
	GeoLocation GeoLocation
}
