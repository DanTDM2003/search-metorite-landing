package repository

import (
	"time"

	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type GetMeteoriteLandingsOption struct {
	paginator.PaginatorQuery
}

type GetOneMeteoriteLandingOption struct {
	ID uint
}

type GeoLocation struct {
	Latitude      float64
	Longitude     float64
	NeedsRecoding bool
}

type CreateMeteoriteLandingOption struct {
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

type UpdateMeteoriteLandingOption struct {
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
