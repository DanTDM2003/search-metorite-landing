package middleware

import (
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"

	userUC "github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
)

type Middleware struct {
	l          pkgLog.Logger
	jwtManager jwt.Manager
	userUC     userUC.Usecase
}

func New(
	l pkgLog.Logger,
	jwtManager jwt.Manager,
	userUC userUC.Usecase,
) Middleware {
	return Middleware{
		l:          l,
		jwtManager: jwtManager,
		userUC:     userUC,
	}
}
