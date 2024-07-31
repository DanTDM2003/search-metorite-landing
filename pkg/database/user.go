package database

import (
	"log"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/utils"
	"gorm.io/gorm"
)

func migrateUsers(db *gorm.DB, table interface{}) error {
	err := db.AutoMigrate(&table)
	if err != nil {
		log.Fatalf("failed to auto migrate table: %v", err)
	}

	hashedPassword, err := utils.HashPassword("DanTDM2003")
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	admin := models.User{
		ID:       1,
		Username: "John Doe",
		Email:    "admin@gmail.com",
		Password: hashedPassword,
		Role:     "superadmin",
	}

	err = db.Create(&admin).Error
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	return err
}
