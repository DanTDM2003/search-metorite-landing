package usecase

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrWrongPassword    = errors.New("wrong password")
	ErrUserAlreadyAdmin = errors.New("user already admin")
	ErrUserAlreadyUser  = errors.New("user already user")
)
