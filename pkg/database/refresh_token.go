package database

import (
	"log"

	"gorm.io/gorm"
)

func migrateRefreshTokens(db *gorm.DB, table interface{}) error {
	err := db.AutoMigrate(&table)
	if err != nil {
		log.Fatalf("failed to auto migrate table: %v", err)
	}

	return err
}
