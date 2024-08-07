package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"
	pkgError "github.com/DanTDM2003/search-api-docker-redis/pkg/errors"
)

var (
	errWrongQuery                     = pkgError.NewHTTPError(10000, "Wrong query")
	errWrongBody                      = pkgError.NewHTTPError(10001, "Wrong body")
	errMeteoriteLandingsNotFound      = pkgError.NewHTTPError(10002, "Meteorite landings not found")
	errMeteoriteLandingsAlreadyExists = pkgError.NewHTTPError(10003, "Meteorite landings already exists")
)

func (h handler) mapError(err error) error {
	switch err {
	case usecase.ErrMeteoriteLandingsNotFound:
		return errMeteoriteLandingsNotFound
	case usecase.ErrMeteoriteLandingAlreadyExists:
		return errMeteoriteLandingsAlreadyExists
	default:
		return err
	}
}
