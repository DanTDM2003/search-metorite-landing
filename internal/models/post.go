package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	AuthorID  uint           `json:"author_id" gorm:"index"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	ViewCount uint           `json:"view_count"`
	Rating    float64        `json:"rating"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
