package http

import (
	pkgErr "github.com/DanTDM2003/search-api-docker-redis/pkg/errors"
)

var (
	errWrongQuery = pkgErr.NewHTTPError(10000, "wrong query")
)
