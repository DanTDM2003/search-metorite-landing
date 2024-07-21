package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	pkgError "github.com/DanTDM2003/search-api-docker-redis/pkg/errors"
)

var (
	errWrongQuery       = pkgError.NewHTTPError(20000, "Wrong query")
	errWrongBody        = pkgError.NewHTTPError(20001, "Wrong body")
	errInvalidPassword  = pkgError.NewHTTPError(20002, "Invalid password")
	errInvalidEmail     = pkgError.NewHTTPError(20003, "Invalid email")
	errUserNotFound     = pkgError.NewHTTPError(20004, "User not found")
	errWrongPassword    = pkgError.NewHTTPError(20005, "Wrong password")
	errUserAlreadyAdmin = pkgError.NewHTTPError(20006, "User already admin")
)

func (h handler) mapError(err error) error {
	switch err {
	case usecase.ErrUserNotFound:
		return errUserNotFound
	case usecase.ErrWrongPassword:
		return errWrongPassword
	case usecase.ErrUserAlreadyAdmin:
		return errUserAlreadyAdmin
	default:
		return err
	}
}
