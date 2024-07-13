package postgres

import (
	"context"
	"encoding/json"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	_ "github.com/lib/pq"
)

const (
	meteoriteLandingsTable = "meteorite_landings"
)

func (repo impleRepository) GetMetoriteLandings(ctx context.Context) ([]models.MeteoriteLanding, error) {
	rows, err := repo.db.Query("SELECT * FROM meteorite_landings")
	if err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.postgres.GetMetoriteLandings.db.Query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var metoriteLandings []models.MeteoriteLanding
	for rows.Next() {
		mL := new(models.MeteoriteLanding)
		var geolocationJSON []byte
		err := rows.Scan(&mL.ID, &mL.Year, &mL.Name, &mL.Nametype, &mL.Recclass, &mL.Mass, &mL.Fall, &mL.Reclat, &mL.Reclong, &geolocationJSON)
		if err != nil {
			repo.l.Errorf(ctx, "meteorite_landings.repository.postgres.GetMetoriteLandings.rows.Scan: %v", err)
			return nil, err
		}

		if err := json.Unmarshal(geolocationJSON, &mL.GeoLocation); err != nil {
			repo.l.Errorf(ctx, "meteorite_landings.repository.postgres.GetMetoriteLandings.json.Unmarshal: %v", err)
			return nil, err
		}

		metoriteLandings = append(metoriteLandings, *mL)
	}

	if err := rows.Err(); err != nil {
		repo.l.Errorf(ctx, "meteorite_landings.repository.postgres.GetMetoriteLandings.rows.Err: %v", err)
		return nil, err
	}

	return metoriteLandings, nil
}
