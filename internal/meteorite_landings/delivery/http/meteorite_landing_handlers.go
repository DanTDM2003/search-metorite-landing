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
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newGetMeteoriteLandingsResp(o))
}

func (h handler) processGetOneMeteoriteLandingsReq(c *gin.Context) (GetOneMeteoriteLandingReq, error) {
	ctx := c.Request.Context()

	var req GetOneMeteoriteLandingReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Errorf(ctx, "http.handler.GetOneMeteoriteLanding.ShouldBindUri: %v", errWrongQuery)
		return GetOneMeteoriteLandingReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) GetOneMeteoriteLanding(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processGetOneMeteoriteLandingsReq(c)
	if err != nil {
		h.l.Errorf(ctx, "http.handler.GetOneMeteoriteLanding.processGetOneMeteoriteLandingsReq: %v", err)
		response.Error(c, err)
		return
	}

	mL, err := h.uc.GetOneMeteoriteLanding(ctx, usecase.GetOneMeteoriteLandingInput{
		ID: req.ID,
	})
	if err != nil {
		h.l.Errorf(ctx, "http.handler.GetOneMeteoriteLanding.uc.GetOneMeteoriteLanding: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newGetOneMeteoriteLandingResp(mL))
}

func (h handler) processCreateMeteoriteLanding(c *gin.Context) (CreateMeteoriteLandingReq, error) {
	ctx := c.Request.Context()

	var req CreateMeteoriteLandingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "http.handler.CreateMeteoriteLanding.ShouldBindJSON: %v", err)
		return CreateMeteoriteLandingReq{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "http.handler.CreateMeteoriteLanding.validate: %v", err)
		return CreateMeteoriteLandingReq{}, err
	}

	return req, nil

}

func (h handler) CreateMeteoriteLanding(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processCreateMeteoriteLanding(c)
	if err != nil {
		h.l.Errorf(ctx, "http.handler.CreateMeteoriteLanding.processCreateMeteoriteLanding: %v", err)
		response.Error(c, err)
		return
	}

	mL, err := h.uc.CreateMeteoriteLanding(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.handler.CreateMeteoriteLanding.uc.CreateMeteoriteLanding: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newCreateMeteoriteLandingResp(mL))
}

func (h handler) processUpdateMeteoriteLanding(c *gin.Context) (UpdateMeteoriteLandingReq, error) {
	ctx := c.Request.Context()

	var req UpdateMeteoriteLandingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "http.handler.UpdateMeteoriteLanding.ShouldBindJSON: %v", err)
		return UpdateMeteoriteLandingReq{}, errWrongQuery
	}

	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Errorf(ctx, "http.handler.UpdateMeteoriteLanding.ShouldBindUri: %v", err)
		return UpdateMeteoriteLandingReq{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "http.handler.UpdateMeteoriteLanding.validate: %v", err)
		return UpdateMeteoriteLandingReq{}, err
	}

	return req, nil
}

func (h handler) UpdateMeteoriteLanding(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processUpdateMeteoriteLanding(c)
	if err != nil {
		h.l.Errorf(ctx, "http.handler.UpdateMeteoriteLanding.processUpdateMeteoriteLanding: %v", err)
		response.Error(c, err)
		return
	}

	mL, err := h.uc.UpdateMeteoriteLanding(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "http.handler.UpdateMeteoriteLanding.uc.UpdateMeteoriteLanding: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newUpdateMeteoriteLandingResp(mL))
}

func (h handler) processDeleteMeteoriteLanding(c *gin.Context) (DeleteMeteoriteLandingReq, error) {
	ctx := c.Request.Context()

	var req DeleteMeteoriteLandingReq
	if err := c.ShouldBindUri(&req); err != nil {
		h.l.Errorf(ctx, "http.handler.DeleteMeteoriteLanding.ShouldBindUri: %v", err)
		return DeleteMeteoriteLandingReq{}, errWrongQuery
	}

	return req, nil
}

func (h handler) DeleteMeteoriteLanding(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processDeleteMeteoriteLanding(c)
	if err != nil {
		h.l.Errorf(ctx, "http.handler.DeleteMeteoriteLanding.processDeleteMeteoriteLanding: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.DeleteMeteoriteLanding(ctx, req.ID)
	if err != nil {
		h.l.Errorf(ctx, "http.handler.DeleteMeteoriteLanding.uc.DeleteMeteoriteLanding: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, nil)
}
