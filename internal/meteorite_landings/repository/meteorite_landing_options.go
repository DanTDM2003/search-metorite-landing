package repository

import (
	"time"

	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
)

type GetMeteoriteLandingsFilter struct {
	Year     time.Time
	Recclass string
	Mass     float64
}

type GetMeteoriteLandingsOption struct {
	GetMeteoriteLandingsFilter
	paginator.PaginatorQuery
}

type GetOneMeteoriteLandingOption struct {
	ID   uint
	Name string
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
