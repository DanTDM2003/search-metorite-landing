package errors

import "net/http"

type HTTPError struct {
	Code       int
	Message    string
	StatusCode int
}

func NewHTTPError(code int, message string) HTTPError {
	return HTTPError{
		Code:    code,
		Message: message,
	}
}

func NewUnauthorizedHTTPError() HTTPError {
	return HTTPError{
		Code:       401,
		Message:    "Unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
}

func (e HTTPError) Error() string {
	return e.Message
}
