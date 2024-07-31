package response

import (
	"net/http"

	pkgErrors "github.com/DanTDM2003/search-api-docker-redis/pkg/errors"
	"github.com/gin-gonic/gin"
)

type Resp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Resp{
		ErrorCode: 0,
		Message:   "Success",
		Data:      data,
	})
}

func parseError(err error) (int, Resp) {
	switch parsedErr := err.(type) {
	case pkgErrors.HTTPError:
		statusCode := parsedErr.StatusCode

		if statusCode == 0 {
			statusCode = http.StatusBadRequest
		}

		return statusCode, Resp{
			ErrorCode: parsedErr.Code,
			Message:   parsedErr.Message,
		}
	default:
		return http.StatusInternalServerError, Resp{
			ErrorCode: 500,
			Message:   DefaultErrorMessage,
		}
	}
}

func Error(c *gin.Context, err error) {
	c.JSON(parseError(err))
}

func PanicError(c *gin.Context, err any) {
	if err == nil {
		c.JSON(parseError(nil))
	} else {
		c.JSON(parseError(err.(error)))
	}
}

func Unauthorized(c *gin.Context) {
	c.JSON(parseError(pkgErrors.NewUnauthorizedHTTPError()))
}
