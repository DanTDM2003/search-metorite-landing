package middleware

import (
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"

	"github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
)

type Middleware struct {
	l          pkgLog.Logger
	jwtManager jwt.Manager
}

func New(
	l pkgLog.Logger,
	jwtManager jwt.Manager,
) Middleware {
	return Middleware{
		l:          l,
		jwtManager: jwtManager,
	}
}
