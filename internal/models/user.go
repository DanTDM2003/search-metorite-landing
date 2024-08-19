package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const (
	UserSuperAdmin = "superadmin"
	UserRoleAdmin  = "admin"
	UserRoleUser   = "user"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"not null;uniqueIndex:idx_username_tag"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Tag       string         `json:"tag" gorm:"not null;uniqueIndex:idx_username_tag"`
	Password  string         `json:"password" gorm:"not null"`
	Role      string         `json:"role" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if len(u.Tag) != 4 {
		return errors.New("tag must be exactly 4 characters long")
	}
	return nil
}
