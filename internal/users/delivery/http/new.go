package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/users"
	serviceLocator "github.com/DanTDM2003/search-api-docker-redis/pkg/locator"
	pkgLog "github.com/DanTDM2003/search-api-docker-redis/pkg/log"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetUsers(c *gin.Context)
	GetOneUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	PromoteToAdmin(c *gin.Context)
	DemoteToUser(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type handler struct {
	l       pkgLog.Logger
	uc      users.Usecase
	locator *serviceLocator.ServiceLocator
}

func New(
	l pkgLog.Logger,
	uc users.Usecase,
) Handler {
	return &handler{
		l:       l,
		uc:      uc,
		locator: serviceLocator.GetServiceLocator(),
	}
}
