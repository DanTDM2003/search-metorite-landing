package usecase

import (
	"context"

	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	pkgPassword "github.com/DanTDM2003/search-api-docker-redis/pkg/password"
)

type Usecase interface {
	SignIn(ctx context.Context, input SignInInput) (SignInOutput, error)
	SignUp(ctx context.Context, input SignUpInput) (SignUpOutput, error)
	SignOut(ctx context.Context, userID uint) error
	Refresh(ctx context.Context, input RefreshInput) (RefreshOutput, error)
}

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
) Usecase {
	return &impleUsecase{
		l:               l,
		jwtManager:      jwtManager,
		locator:         serviceLocator.GetServiceLocator(),
		passwordManager: passwordManager,
	}
}
