package usecase

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/session"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	pkgPassword "github.com/DanTDM2003/search-api-docker-redis/pkg/password"
)

type impleUsecase struct {
	l               pkgLog.Logger
	jwtManager      pkgJWT.Manager
	locator         *serviceLocator.ServiceLocator
	passwordManager pkgPassword.Password
}

func New(
	l pkgLog.Logger,
	jwtManager pkgJWT.Manager,
	passwordManager pkgPassword.Password,
) session.Usecase {
	return &impleUsecase{
		l:               l,
		jwtManager:      jwtManager,
		locator:         serviceLocator.GetServiceLocator(),
		passwordManager: passwordManager,
	}
}
