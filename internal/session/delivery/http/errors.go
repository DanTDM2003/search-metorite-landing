package http

import (
	userUC "github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	pkgError "github.com/DanTDM2003/search-api-docker-redis/pkg/errors"
)

var (
	errWrongBody       = pkgError.NewHTTPError(50000, "wrong body")
	errInvalidPassword = pkgError.NewHTTPError(50001, "password must be at least 8 characters long with special characters")
	errInvalidEmail    = pkgError.NewHTTPError(50002, "invalid email")
	errWrongPassword   = pkgError.NewHTTPError(50003, "wrong password")
)

func (h handler) mapError(err error) error {
	switch err {
	case userUC.ErrWrongPassword:
		return errWrongPassword
	default:
		return err
	}
}
