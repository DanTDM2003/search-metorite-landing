package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h handler) processGetMeteoriteLandingsReq(c *gin.Context) (paginator.PaginatorQuery, error) {
	ctx := c.Request.Context()

	var pagQuery paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pagQuery); err != nil {
		h.l.Errorf(ctx, "http.handler.GetMeteoriteLandings.ShouldBindQuery: %v", errWrongQuery)
		return paginator.PaginatorQuery{}, errWrongQuery
	}

	pagQuery.Adjust()

	return pagQuery, nil
}

func (h handler) GetMeteoriteLandings(c *gin.Context) {
	ctx := c.Request.Context()

	pagQuery, err := h.processGetMeteoriteLandingsReq(c)
	if err != nil {
		h.l.Errorf(ctx, "http.handler.GetMeteoriteLandings.processGetMeteoriteLandingsReq: %v", err)
		response.Error(c, err)
		return
	}

	o, err := h.uc.GetMeteoriteLandings(ctx, usecase.GetMeteoriteLandingsInput{
		PaginatorQuery: pagQuery,
	})
	if err != nil {
		h.l.Errorf(ctx, "http.handler.GetMeteoriteLandings.uc.GetMeteoriteLandings: %v", err)
		response.Error(c, err)
		return
	}

	response.Success(c, h.newGetMeteoriteLandingsResp(o))
}
