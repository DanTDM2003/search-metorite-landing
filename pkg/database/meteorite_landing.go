package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"gorm.io/gorm"
)

func migrateMetoriteLandings(db *gorm.DB, table interface{}) error {
	err := db.AutoMigrate(&table)
	if err != nil {
		log.Fatalf("failed to auto migrate table: %v", err)
	}

	metoriteLandings, err := loadMetoriteLandingsFromFile("data/meteorite-landings.json")
	if err != nil {
		log.Fatalf("Could not load meteorite landings from file: %v", err)
	}

	err = insertMetoriteLandings(db, metoriteLandings)
	if err != nil {
		log.Fatalf("Could not insert meteorite landings: %v", err)
	}

	return err
}

func insertMetoriteLandings(db *gorm.DB, mLs []models.MeteoriteLanding) error {
	for i, mL := range mLs {
		mL.ID = uint(i + 1)
		err := db.Create(&mL).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func loadMetoriteLandingsFromFile(filePath string) ([]models.MeteoriteLanding, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var metoriteLandings []models.MeteoriteLanding
	err = json.Unmarshal(byteValue, &metoriteLandings)
	return metoriteLandings, err
}
