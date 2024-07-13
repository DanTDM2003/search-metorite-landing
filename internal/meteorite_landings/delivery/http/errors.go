package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"
	pkgError "github.com/DanTDM2003/search-api-docker-redis/pkg/errors"
)

var (
	errWrongQuery                = pkgError.NewHTTPError(10000, "Wrong query")
	errWrongBody                 = pkgError.NewHTTPError(10001, "Wrong body")
	errMeteoriteLandingsNotFound = pkgError.NewHTTPError(10002, "Meteorite landings not found")
)

func (h handler) mapError(err error) error {
	switch err {
	case errWrongQuery:
		return errWrongQuery
	case errWrongBody:
		return errWrongBody
	case usecase.ErrMeteoriteLandingsNotFound:
		return errMeteoriteLandingsNotFound
	default:
		return err
	}
}
