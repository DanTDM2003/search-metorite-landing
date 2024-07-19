package utils

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/DanTDM2003/search-api-docker-redis/internal/appconfig/database"
	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
)

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS meteorite_landings (
		id SERIAL PRIMARY KEY,
		year DATE,
		name TEXT,
		name_type TEXT,
		recclass TEXT,
		mass FLOAT8,
		fall TEXT,
		reclat FLOAT8,
		reclong FLOAT8,
		geo_location JSONB,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP DEFAULT NULL
	)`
	_, err := db.Exec(query)
	return err
}

func insertMetoriteLandings(db *sql.DB, metoriteLandings []models.MeteoriteLanding) error {
	for _, mL := range metoriteLandings {
		geolocationJSON, err := json.Marshal(mL.GeoLocation)
		if err != nil {
			return err
		}

		query := `
		INSERT INTO meteorite_landings (year, name, name_type, recclass, mass, fall, reclat, reclong, geo_location)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
		_, err = db.Exec(query, mL.Year, mL.Name, mL.NameType, mL.Recclass, mL.Mass, mL.Fall, mL.Reclat, mL.Reclong, geolocationJSON)
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

func InitDatabase(p *database.PostgresConnection) {
	db, err := p.DB.DB()
	if err != nil {
		log.Fatalf("Could not get database connection: %v", err)
		panic(err)
	}

	metoriteLandings, err := loadMetoriteLandingsFromFile("data/meteorite-landings.json")
	if err != nil {
		log.Fatalf("Could not load meteorite landings from file: %v", err)
		panic(err)
	}

	err = createTable(db)
	if err != nil {
		log.Fatalf("Could not create table: %v", err)
		panic(err)
	}

	err = insertMetoriteLandings(db, metoriteLandings)
	if err != nil {
		log.Fatalf("Could not insert meteorite landings: %v", err)
		panic(err)
	}
}
