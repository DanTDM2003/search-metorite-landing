package database

import (
	"log"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func migrateUsers(db *gorm.DB, table interface{}) error {
	err := db.AutoMigrate(&table)
	if err != nil {
		log.Fatalf("failed to auto migrate table: %v", err)
	}

	password := "D@nTDM22122003"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	admin := models.User{
		Username: "John Doe",
		Email:    "admin@gmail.com",
		Password: string(hashedPassword),
		Role:     "superadmin",
		Tag:      "2212",
	}

	err = db.Create(&admin).Error
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	return err
}
