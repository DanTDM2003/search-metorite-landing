package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"os"

// 	"github.com/DanTDM2003/search-api-docker-redis/config"
// 	"github.com/DanTDM2003/search-api-docker-redis/internal/appconfig/database"
// 	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
// )

// func createTable(p *database.PostgresConnection) error {
// 	query := `
// 	DROP TABLE IF EXISTS metorite_landings;
// 	CREATE TABLE IF NOT EXISTS meteorite_landings (
// 		id SERIAL PRIMARY KEY,
// 		year DATE,
// 		name TEXT,
// 		nametype TEXT,
// 		recclass TEXT,
// 		mass FLOAT8,
// 		fall TEXT,
// 		reclat FLOAT8,
// 		reclong FLOAT8,
// 		geolocation JSONB,
// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 		deleted_at TIMESTAMP DEFAULT NULL
// 	)`
// 	_, err := p.DB1.Exec(query)
// 	return err
// }

// func insertMetoriteLandings(p *database.PostgresConnection, metoriteLandings []models.MeteoriteLanding) error {
// 	for _, mL := range metoriteLandings {
// 		geolocationJSON, err := json.Marshal(mL.GeoLocation)
// 		if err != nil {
// 			return err
// 		}

// 		query := `
// 		INSERT INTO meteorite_landings (year, name, nametype, recclass, mass, fall, reclat, reclong, geolocation)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
// 		_, err = p.DB1.Exec(query, mL.Year, mL.Name, mL.Nametype, mL.Recclass, mL.Mass, mL.Fall, mL.Reclat, mL.Reclong, geolocationJSON)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func loadMetoriteLandingsFromFile(filePath string) ([]models.MeteoriteLanding, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	byteValue, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var metoriteLandings []models.MeteoriteLanding
// 	err = json.Unmarshal(byteValue, &metoriteLandings)
// 	return metoriteLandings, err
// }

// func main() {
// 	cfg, err := config.Load()
// 	if err != nil {
// 		log.Fatalf("Could not load the configuration: %v", err)
// 		panic(err)
// 	}

// 	conn, err := database.Connect(cfg.Postgres)
// 	if err != nil {
// 		log.Fatalf("Could not connect to the database: %v", err)
// 		panic(err)
// 	}
// 	defer database.Close(conn.DB1)

// 	metoriteLandings, err := loadMetoriteLandingsFromFile("data/meteorite-landings.json")
// 	if err != nil {
// 		log.Fatalf("Could not load meteorite landings from file: %v", err)
// 		panic(err)
// 	}

// 	err = createTable(conn)
// 	if err != nil {
// 		log.Fatalf("Could not create table: %v", err)
// 		panic(err)
// 	}

// 	err = insertMetoriteLandings(conn, metoriteLandings)
// 	if err != nil {
// 		log.Fatalf("Could not insert meteorite landings: %v", err)
// 		panic(err)
// 	}
// }
