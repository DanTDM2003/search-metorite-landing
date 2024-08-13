package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/articles/usecase"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Handler interface {
}

type handler struct {
	l  pkgLog.Logger
	uc usecase.Usecase
}

func New(
	l pkgLog.Logger,
	uc usecase.Usecase,
) Handler {
	return &handler{
		l:  l,
		uc: uc,
	}
}
