package models

import "time"

type GeoLocation struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	NeedsRecoding bool    `json:"needs_recoding"`
}

type MeteoriteLanding struct {
	ID          int         `json:"id"`
	Year        time.Time   `json:"year"`
	Name        string      `json:"name"`
	Nametype    string      `json:"nametype"`
	Recclass    string      `json:"recclass"`
	Mass        float64     `json:"mass"`
	Fall        string      `json:"fall"`
	Reclat      float64     `json:"reclat"`
	Reclong     float64     `json:"reclong"`
	GeoLocation GeoLocation `json:"geolocation"`
}
