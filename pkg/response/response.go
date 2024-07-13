package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Resp{
		ErrorCode: 0,
		Message:   "Success",
		Data:      data,
	})
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Resp{
		ErrorCode: 1,
		Message:   err.Error(),
		Data:      nil,
	})
}
