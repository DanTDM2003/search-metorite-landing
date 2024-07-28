package usecase

import "errors"

var (
	ErrPostNotFound   = errors.New("post not found")
	ErrAuthorNotFound = errors.New("author not found")
)
