package http

import (
	"strconv"

	"github.com/DanTDM2003/search-api-docker-redis/internal/session"
	pkgError "github.com/DanTDM2003/search-api-docker-redis/pkg/errors"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/jwt"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/response"
	"github.com/gin-gonic/gin"
)

func (h handler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processSignInReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.SignIn.processSignInReq: %v", err)
		response.Error(c, err)
		return
	}

	o, err := h.uc.SignIn(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "users.http.SignIn.uc.SignIn: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newSignInResp(o))
}

func (h handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processSignUpReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.SignUp.processSignUpReq: %v", err)
		response.Error(c, err)
		return
	}

	u, err := h.uc.SignUp(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "users.http.SignUp.uc.SignUp: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newSignUpResp(u))
}

func (h handler) SignOut(c *gin.Context) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "users.http.SignOut.jwt.GetPayloadFromContext: %v", pkgError.NewUnauthorizedHTTPError())
		response.Error(c, pkgError.NewUnauthorizedHTTPError())
		return
	}

	userID, err := strconv.Atoi(payload.Subject)
	if err != nil {
		h.l.Warnf(ctx, "users.http.SignOut.strconv.Atoi: %v", err)
		response.Error(c, pkgError.NewUnauthorizedHTTPError())
		return
	}

	err = h.uc.SignOut(ctx, uint(userID))
	if err != nil {
		h.l.Errorf(ctx, "users.http.SignOut.uc.SignOut: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, nil)
}

func (h handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processRefreshReq(c)
	if err != nil {
		h.l.Warnf(ctx, "users.http.Refresh.processRefreshReq: %v", err)
		response.Error(c, err)
		return
	}

	o, err := h.uc.Refresh(ctx, session.RefreshInput{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		h.l.Errorf(ctx, "users.http.Refresh.uc.Refresh: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.Success(c, h.newRefreshResp(o))
}
