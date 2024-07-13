package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MeteoriteLanding struct {
	gorm.Model
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

type GeoLocation struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	NeedsRecoding bool    `json:"needs_recoding"`
}

func (g GeoLocation) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// Scan implements the Scanner interface for GeoLocation.
func (g *GeoLocation) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, g)
}
