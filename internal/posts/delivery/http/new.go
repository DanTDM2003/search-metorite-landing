package http

import (
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
)

type Handler interface {
}

type handler struct {
	log pkgLog.Logger
}

func New(l pkgLog.Logger) Handler {
	return &handler{
		log: l,
	}
}
