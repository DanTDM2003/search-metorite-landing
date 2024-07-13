package database

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/repository"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

func (uc impleRepository) buildCreateMeteoriteLandingModel(opt repository.CreateMeteoriteLandingOption) models.MeteoriteLanding {
	return models.MeteoriteLanding{
		Name:     opt.Name,
		NameType: opt.NameType,
		Year:     opt.Year,
		Recclass: opt.Recclass,
		Mass:     opt.Mass,
		Fall:     opt.Fall,
		Reclat:   opt.Reclat,
		Reclong:  opt.Reclong,
		GeoLocation: models.GeoLocation{
			Latitude:      opt.GeoLocation.Latitude,
			Longitude:     opt.GeoLocation.Longitude,
			NeedsRecoding: opt.GeoLocation.NeedsRecoding,
		},
	}
}

func (uc impleRepository) buildUpdateMeteoriteLandingModel(opt repository.UpdateMeteoriteLandingOption) models.MeteoriteLanding {
	return models.MeteoriteLanding{
		Name:     opt.Name,
		NameType: opt.NameType,
		Year:     opt.Year,
		Recclass: opt.Recclass,
		Mass:     opt.Mass,
		Fall:     opt.Fall,
		Reclat:   opt.Reclat,
		Reclong:  opt.Reclong,
		GeoLocation: models.GeoLocation{
			Latitude:      opt.GeoLocation.Latitude,
			Longitude:     opt.GeoLocation.Longitude,
			NeedsRecoding: opt.GeoLocation.NeedsRecoding,
		},
	}
}
