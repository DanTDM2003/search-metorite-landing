package middleware

import (
	"github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Middleware struct {
	l          pkgLog.Logger
	jwtManager jwt.Manager
	locator    *serviceLocator.ServiceLocator
}

func New(
	l pkgLog.Logger,
	jwtManager jwt.Manager,
) Middleware {
	return Middleware{
		l:          l,
		jwtManager: jwtManager,
		locator:    serviceLocator.GetServiceLocator(),
	}
}
