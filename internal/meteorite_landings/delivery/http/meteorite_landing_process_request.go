package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/gin-gonic/gin"
)

func (h handler) processGetMeteoriteLandingsReq(c *gin.Context) (paginator.PaginatorQuery, error) {
	ctx := c.Request.Context()

	var pagQuery paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pagQuery); err != nil {
		h.l.Warnf(ctx, "meteorite_landings.http.processGetMeteoriteLandingsReq.ShouldBindQuery: %v", errWrongQuery)
		return paginator.PaginatorQuery{}, errWrongQuery
	}

	pagQuery.Adjust()

	return pagQuery, nil
}

func (h handler) processGetOneMeteoriteLandingReq(c *gin.Context) (getOneMeteoriteLandingReq, error) {
	ctx := c.Request.Context()

	var req getOneMeteoriteLandingReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "meteorite_landings.http.processGetOneMeteoriteLandingReq.ShouldBindUri: %v", errWrongQuery)
		return getOneMeteoriteLandingReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) processCreateMeteoriteLandingReq(c *gin.Context) (createMeteoriteLandingReq, error) {
	ctx := c.Request.Context()

	var req createMeteoriteLandingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "meteorite_landings.http.processCreateMeteoriteLandingReq.ShouldBindJSON: %v", err)
		return createMeteoriteLandingReq{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "meteorite_landings.http.processCreateMeteoriteLandingReq.validate: %v", err)
		return createMeteoriteLandingReq{}, err
	}

	return req, nil
}

func (h handler) processUpdateMeteoriteLandingReq(c *gin.Context) (updateMeteoriteLandingReq, error) {
	ctx := c.Request.Context()

	var req updateMeteoriteLandingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "meteorite_landings.http.processUpdateMeteoriteLandingReq.ShouldBindJSON: %v", err)
		return updateMeteoriteLandingReq{}, errWrongQuery
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "meteorite_landings.http.processUpdateMeteoriteLandingReq.ShouldBindUri: %v", err)
		return updateMeteoriteLandingReq{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "meteorite_landings.http.processUpdateMeteoriteLandingReq.validate: %v", err)
		return updateMeteoriteLandingReq{}, err
	}

	return req, nil
}

func (h handler) processDeleteMeteoriteLandingReq(c *gin.Context) (deleteMeteoriteLandingReq, error) {
	ctx := c.Request.Context()

	var req deleteMeteoriteLandingReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Warnf(ctx, "meteorite_landings.http.processDeleteMeteoriteLandingReq.ShouldBindUri: %v", err)
		return deleteMeteoriteLandingReq{}, errWrongQuery
	}

	return req, nil
}
