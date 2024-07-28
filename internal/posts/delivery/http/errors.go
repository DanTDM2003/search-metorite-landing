package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/posts/usecase"
	pkgError "github.com/DanTDM2003/search-api-docker-redis/pkg/errors"
)

var (
	errWrongQuery     = pkgError.NewHTTPError(30000, "Wrong query")
	errWrongBody      = pkgError.NewHTTPError(30001, "Wrong body")
	errPostNotFound   = pkgError.NewHTTPError(30002, "Post not found")
	errAuthorNotFound = pkgError.NewHTTPError(30003, "Author not found")
)

func (h handler) mapError(err error) error {
	switch err {
	case usecase.ErrPostNotFound:
		return errPostNotFound
	case usecase.ErrAuthorNotFound:
		return errAuthorNotFound
	default:
		return err
	}
}
