package http

import (
	"github.com/DanTDM2003/search-api-docker-redis/internal/meteorite_landings/usecase"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

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
