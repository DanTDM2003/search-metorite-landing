package main

import (
	"encoding/json"

	_ "github.com/lib/pq"
)

type Repository interface {
	GetMetoriteLandings() ([]MetoriteLanding, error)
}

func (p *PostgresConnection) GetMetoriteLandings() ([]MetoriteLanding, error) {
	rows, err := p.db.Query("SELECT * FROM metorite_landings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metoriteLandings []MetoriteLanding
	for rows.Next() {
		mL := new(MetoriteLanding)
		var geolocationJSON []byte
		err := rows.Scan(&mL.ID, &mL.Year, &mL.Name, &mL.Nametype, &mL.Recclass, &mL.Mass, &mL.Fall, &mL.Reclat, &mL.Reclong, &geolocationJSON)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(geolocationJSON, &mL.GeoLocation); err != nil {
			return nil, err
		}

		metoriteLandings = append(metoriteLandings, *mL)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return metoriteLandings, nil
}
