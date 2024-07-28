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

func (uc impleRepository) buildUpdateMeteoriteLandingModel(opt repository.UpdateMeteoriteLandingOption, mL models.MeteoriteLanding) models.MeteoriteLanding {
	update := models.MeteoriteLanding{
		ID:        mL.ID,
		CreatedAt: mL.CreatedAt,
	}

	if opt.Name != "" {
		update.Name = opt.Name
	} else {
		update.Name = mL.Name
	}

	if opt.NameType != "" {
		update.NameType = opt.NameType
	} else {
		update.NameType = mL.NameType
	}

	if !opt.Year.IsZero() {
		update.Year = opt.Year
	} else {
		update.Year = mL.Year
	}

	if opt.Recclass != "" {
		update.Recclass = opt.Recclass
	} else {
		update.Recclass = mL.Recclass
	}

	if opt.Mass != 0 {
		update.Mass = opt.Mass
	} else {
		update.Mass = mL.Mass
	}

	if opt.Fall != "" {
		update.Fall = opt.Fall
	} else {
		update.Fall = mL.Fall
	}

	if opt.Reclat != 0 {
		update.Reclat = opt.Reclat
	} else {
		update.Reclat = mL.Reclat
	}

	if opt.Reclong != 0 {
		update.Reclong = opt.Reclong
	} else {
		update.Reclong = mL.Reclong
	}

	if opt.GeoLocation.Latitude != 0 {
		update.GeoLocation.Latitude = opt.GeoLocation.Latitude
	} else {
		update.GeoLocation.Latitude = mL.GeoLocation.Latitude
	}

	if opt.GeoLocation.Longitude != 0 {
		update.GeoLocation.Longitude = opt.GeoLocation.Longitude
	} else {
		update.GeoLocation.Longitude = mL.GeoLocation.Longitude
	}

	if opt.GeoLocation.NeedsRecoding {
		update.GeoLocation.NeedsRecoding = opt.GeoLocation.NeedsRecoding
	} else {
		update.GeoLocation.NeedsRecoding = mL.GeoLocation.NeedsRecoding
	}

	return update
}
