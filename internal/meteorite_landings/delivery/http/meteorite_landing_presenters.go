package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
)

type geoLocation struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	NeedsRecoding bool    `json:"needs_recoding"`
}

type GetMeteoriteLandingsRespItem struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Year        response.Date     `json:"year"`
	Mass        float64           `json:"mass"`
	Reclass     string            `json:"reclass"`
	Fall        string            `json:"fall"`
	Reclat      float64           `json:"reclat"`
	Reclong     float64           `json:"reclong"`
	GeoLocation geoLocation       `json:"geolocation"`
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
			ID:      item.ID,
			Name:    item.Name,
			Year:    response.Date(item.Year),
			Mass:    item.Mass,
			Reclass: item.Recclass,
			Fall:    item.Fall,
			Reclat:  item.Reclat,
			Reclong: item.Reclong,
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
