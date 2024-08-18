package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	AuthorID  uint           `json:"author_id" gorm:"not null;index"`
	Title     string         `json:"title" gorm:"not null"`
	Content   string         `json:"content" gorm:"not null"`
	ViewCount uint           `json:"view_count" gorm:"not null"`
	Rating    float64        `json:"rating" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
