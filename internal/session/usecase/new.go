package usecase

import (
	"context"

	userUC "github.com/DanTDM2003/search-api-docker-redis/internal/users/usecase"
	pkgJWT "github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Usecase interface {
	SignIn(ctx context.Context, input SignInInput) (SignInOutput, error)
	SignUp(ctx context.Context, input SignUpInput) (SignUpOutput, error)
	SignOut(ctx context.Context, userID uint) error
	Refresh(ctx context.Context, input RefreshInput) (RefreshOutput, error)
}

type impleUsecase struct {
	l          pkgLog.Logger
	jwtManager pkgJWT.Manager
	userUC     userUC.Usecase
}

func New(
	l pkgLog.Logger,
	userUC userUC.Usecase,
	jwtManager pkgJWT.Manager,
) Usecase {
	return &impleUsecase{
		l:          l,
		jwtManager: jwtManager,
		userUC:     userUC,
	}
}
