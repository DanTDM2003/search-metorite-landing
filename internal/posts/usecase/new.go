package usecase

import (
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type UseCase interface {
}

type impleUsecase struct {
	l pkgLog.Logger
}

func New(l pkgLog.Logger) UseCase {
	return &impleUsecase{
		l: l,
	}
}
