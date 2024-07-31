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
	ID          uint           `json:"id" gorm:"primaryKey"`
	Year        time.Time      `json:"year" gorm:"not null"`
	Name        string         `json:"name" gorm:"unique"`
	NameType    string         `json:"name_type" gorm:"not null"`
	Recclass    string         `json:"recclass" gorm:"not null"`
	Mass        float64        `json:"mass" gorm:"not null"`
	Fall        string         `json:"fall" gorm:"not null"`
	Reclat      float64        `json:"reclat" gorm:"not null"`
	Reclong     float64        `json:"reclong" gorm:"not null"`
	GeoLocation GeoLocation    `json:"geo_location" gorm:"type:jsonb"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
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
