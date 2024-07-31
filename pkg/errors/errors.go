package errors

var (
	ErrUnauthorized = NewHTTPError(401, "Unauthorized")
)
