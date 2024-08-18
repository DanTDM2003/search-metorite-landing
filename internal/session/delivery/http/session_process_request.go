package http

import "github.com/gin-gonic/gin"

func (h handler) processSignInReq(c *gin.Context) (signInReq, error) {
	ctx := c.Request.Context()

	var req signInReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processSignInReq.ShouldBindJSON: %v", err)
		return signInReq{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "users.http.processSignInReq.validate: %v", err)
		return signInReq{}, err
	}

	return req, nil
}

func (h handler) processSignUpReq(c *gin.Context) (signUpReq, error) {
	ctx := c.Request.Context()

	var req signUpReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processSignUpReq.ShouldBindJSON: %v", err)
		return signUpReq{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "users.http.processSignUpReq.validate: %v", err)
		return signUpReq{}, err
	}

	return req, nil
}

func (h handler) processRefreshReq(c *gin.Context) (refreshReq, error) {
	ctx := c.Request.Context()

	var req refreshReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "users.http.processRefreshReq.ShouldBindJSON: %v", err)
		return refreshReq{}, errWrongBody
	}

	return req, nil
}
