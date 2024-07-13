package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h handler) GetMeteoriteLandings(c *gin.Context) {
	ctx := c.Request.Context()

	mLs, err := h.uc.GetMeteoriteLandings(ctx)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, mLs)
}
