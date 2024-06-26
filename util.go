package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func createTable(p *PostgresConnection) error {
	query := `
	DROP TABLE IF EXISTS metorite_landings;
	CREATE TABLE IF NOT EXISTS metorite_landings (
		id SERIAL PRIMARY KEY,
		year DATE,
		name TEXT,
		nametype TEXT,
		recclass TEXT,
		mass FLOAT8,
		fall TEXT,
		reclat FLOAT8,
		reclong FLOAT8,
		geolocation JSONB
	)`
	_, err := p.db.Exec(query)
	return err
}

func (p *PostgresConnection) insertMetoriteLandings(metoriteLandings []MetoriteLanding) error {
	for _, mL := range metoriteLandings {
		geolocationJSON, err := json.Marshal(mL.GeoLocation)
		if err != nil {
			return err
		}

		query := `
		INSERT INTO metorite_landings (year, name, nametype, recclass, mass, fall, reclat, reclong, geolocation)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
		_, err = p.db.Exec(query, mL.Year, mL.Name, mL.Nametype, mL.Recclass, mL.Mass, mL.Fall, mL.Reclat, mL.Reclong, geolocationJSON)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadMetoriteLandingsFromFile(filePath string) ([]MetoriteLanding, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var metoriteLandings []MetoriteLanding
	err = json.Unmarshal(byteValue, &metoriteLandings)
	return metoriteLandings, err
}
