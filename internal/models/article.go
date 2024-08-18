package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	AuthorID  uint           `json:"author_id" gorm:"not null;index"`
	Title     string         `json:"title" gorm:"not null;unique"`
	Slug      string         `json:"slug" gorm:"not null;unique;index"`
	Content   string         `json:"content" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUopdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
