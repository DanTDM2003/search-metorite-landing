package http

import (
	"time"

	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/utils"
)

type geoLocation struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	NeedsRecoding bool    `json:"needs_recoding"`
}

type GetMeteoriteLandingsRespItem struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	NameType    string            `json:"name_type"`
	Year        response.Date     `json:"year"`
	Mass        float64           `json:"mass"`
	Recclass    string            `json:"recclass"`
	Fall        string            `json:"fall"`
	Reclat      float64           `json:"reclat"`
	Reclong     float64           `json:"reclong"`
	GeoLocation geoLocation       `json:"geo_location"`
	CreatedAt   response.DateTime `json:"created_at"`
	UpdatedAt   response.DateTime `json:"updated_at"`
}

type GetMeteoriteLandingsResp struct {
	Items []GetMeteoriteLandingsRespItem `json:"items"`
	Meta  paginator.PaginatorResponse    `json:"meta"`
}

func (h handler) newGetMeteoriteLandingsResp(o usecase.GetMeteoriteLandingsOutput) GetMeteoriteLandingsResp {
	var items []GetMeteoriteLandingsRespItem

	for _, item := range o.MeteoriteLandings {
		items = append(items, GetMeteoriteLandingsRespItem{
			ID:       item.ID,
			Name:     item.Name,
			NameType: item.NameType,
			Year:     response.Date(item.Year),
			Mass:     item.Mass,
			Recclass: item.Recclass,
			Fall:     item.Fall,
			Reclat:   item.Reclat,
			Reclong:  item.Reclong,
			GeoLocation: geoLocation{
				Latitude:      item.GeoLocation.Latitude,
				Longitude:     item.GeoLocation.Longitude,
				NeedsRecoding: item.GeoLocation.NeedsRecoding,
			},
			CreatedAt: response.DateTime(item.CreatedAt),
			UpdatedAt: response.DateTime(item.UpdatedAt),
		})
	}

	return GetMeteoriteLandingsResp{
		Items: items,
		Meta:  o.Paginator.ToResponse(),
	}
}

type getOneMeteoriteLandingReq struct {
	ID uint `uri:"id"`
}

type getOneMeteoriteLandingResp struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	NameType    string            `json:"name_type"`
	Year        response.Date     `json:"year"`
	Mass        float64           `json:"mass"`
	Recclass    string            `json:"recclass"`
	Fall        string            `json:"fall"`
	Reclat      float64           `json:"reclat"`
	Reclong     float64           `json:"reclong"`
	GeoLocation geoLocation       `json:"geo_location"`
	CreatedAt   response.DateTime `json:"created_at"`
	UpdatedAt   response.DateTime `json:"updated_at"`
}

func (h handler) newGetOneMeteoriteLandingResp(mL models.MeteoriteLanding) getOneMeteoriteLandingResp {
	return getOneMeteoriteLandingResp{
		ID:       mL.ID,
		Name:     mL.Name,
		NameType: mL.NameType,
		Year:     response.Date(mL.Year),
		Mass:     mL.Mass,
		Recclass: mL.Recclass,
		Fall:     mL.Fall,
		Reclat:   mL.Reclat,
		Reclong:  mL.Reclong,
		GeoLocation: geoLocation{
			Latitude:      mL.GeoLocation.Latitude,
			Longitude:     mL.GeoLocation.Longitude,
			NeedsRecoding: mL.GeoLocation.NeedsRecoding,
		},
		CreatedAt: response.DateTime(mL.CreatedAt),
		UpdatedAt: response.DateTime(mL.UpdatedAt),
	}
}

type createMeteoriteLandingReq struct {
	Name        string      `json:"name" binding:"required"`
	NameType    string      `json:"name_type" binding:"required"`
	Year        string      `json:"year" binding:"required"`
	Mass        float64     `json:"mass" binding:"required"`
	Recclass    string      `json:"recclass" binding:"required"`
	Fall        string      `json:"fall" binding:"required"`
	Reclat      float64     `json:"reclat" binding:"required"`
	Reclong     float64     `json:"reclong" binding:"required"`
	GeoLocation geoLocation `json:"geo_location" binding:"required"`
}

func (req createMeteoriteLandingReq) validate() error {
	_, err := time.Parse(utils.StandardDateTime, req.Year)
	if err != nil {
		return errWrongBody
	}

	return nil
}

func (req createMeteoriteLandingReq) toInput() usecase.CreateMeteoriteLandingInput {
	input := usecase.CreateMeteoriteLandingInput{
		Name:     req.Name,
		NameType: req.NameType,
		Mass:     req.Mass,
		Recclass: req.Recclass,
		Fall:     req.Fall,
		Reclat:   req.Reclat,
		Reclong:  req.Reclong,
		GeoLocation: usecase.GeoLocation{
			Latitude:      req.GeoLocation.Latitude,
			Longitude:     req.GeoLocation.Longitude,
			NeedsRecoding: req.GeoLocation.NeedsRecoding,
		},
	}

	input.Year, _ = time.Parse(utils.StandardDateTime, req.Year)

	return input
}

type createMeteoriteLandingResp struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	NameType    string            `json:"name_type"`
	Year        response.Date     `json:"year"`
	Mass        float64           `json:"mass"`
	Recclass    string            `json:"recclass"`
	Fall        string            `json:"fall"`
	Reclat      float64           `json:"reclat"`
	Reclong     float64           `json:"reclong"`
	GeoLocation geoLocation       `json:"geo_location"`
	CreatedAt   response.DateTime `json:"created_at"`
	UpdatedAt   response.DateTime `json:"updated_at"`
}

func (h handler) newCreateMeteoriteLandingResp(mL models.MeteoriteLanding) createMeteoriteLandingResp {
	return createMeteoriteLandingResp{
		ID:       mL.ID,
		Name:     mL.Name,
		NameType: mL.NameType,
		Year:     response.Date(mL.Year),
		Mass:     mL.Mass,
		Recclass: mL.Recclass,
		Fall:     mL.Fall,
		Reclat:   mL.Reclat,
		Reclong:  mL.Reclong,
		GeoLocation: geoLocation{
			Latitude:      mL.GeoLocation.Latitude,
			Longitude:     mL.GeoLocation.Longitude,
			NeedsRecoding: mL.GeoLocation.NeedsRecoding,
		},
		CreatedAt: response.DateTime(mL.CreatedAt),
		UpdatedAt: response.DateTime(mL.UpdatedAt),
	}
}

type updateMeteoriteLandingReq struct {
	ID          uint        `uri:"id"`
	Name        string      `json:"name"`
	NameType    string      `json:"name_type"`
	Year        string      `json:"year"`
	Mass        float64     `json:"mass"`
	Recclass    string      `json:"recclass"`
	Fall        string      `json:"fall"`
	Reclat      float64     `json:"reclat"`
	Reclong     float64     `json:"reclong"`
	GeoLocation geoLocation `json:"geo_location"`
}

func (req updateMeteoriteLandingReq) validate() error {
	_, err := time.Parse(utils.StandardDateTime, req.Year)
	if err != nil {
		return errWrongBody
	}

	return nil
}

func (req updateMeteoriteLandingReq) toInput() usecase.UpdateMeteoriteLandingInput {
	input := usecase.UpdateMeteoriteLandingInput{
		ID:       req.ID,
		Name:     req.Name,
		NameType: req.NameType,
		Mass:     req.Mass,
		Recclass: req.Recclass,
		Fall:     req.Fall,
		Reclat:   req.Reclat,
		Reclong:  req.Reclong,
		GeoLocation: usecase.GeoLocation{
			Latitude:      req.GeoLocation.Latitude,
			Longitude:     req.GeoLocation.Longitude,
			NeedsRecoding: req.GeoLocation.NeedsRecoding,
		},
	}

	input.Year, _ = time.Parse(utils.StandardDateTime, req.Year)

	return input
}

type updateMeteoriteLandingResp struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	NameType    string            `json:"name_type"`
	Year        response.Date     `json:"year"`
	Mass        float64           `json:"mass"`
	Recclass    string            `json:"recclass"`
	Fall        string            `json:"fall"`
	Reclat      float64           `json:"reclat"`
	Reclong     float64           `json:"reclong"`
	GeoLocation geoLocation       `json:"geo_location"`
	CreatedAt   response.DateTime `json:"created_at"`
	UpdatedAt   response.DateTime `json:"updated_at"`
}

func (h handler) newUpdateMeteoriteLandingResp(mL models.MeteoriteLanding) updateMeteoriteLandingResp {
	return updateMeteoriteLandingResp{
		ID:       mL.ID,
		Name:     mL.Name,
		NameType: mL.NameType,
		Year:     response.Date(mL.Year),
		Mass:     mL.Mass,
		Recclass: mL.Recclass,
		Fall:     mL.Fall,
		Reclat:   mL.Reclat,
		Reclong:  mL.Reclong,
		GeoLocation: geoLocation{
			Latitude:      mL.GeoLocation.Latitude,
			Longitude:     mL.GeoLocation.Longitude,
			NeedsRecoding: mL.GeoLocation.NeedsRecoding,
		},
		CreatedAt: response.DateTime(mL.CreatedAt),
		UpdatedAt: response.DateTime(mL.UpdatedAt),
	}
}

type deleteMeteoriteLandingReq struct {
	ID uint `uri:"id"`
}
