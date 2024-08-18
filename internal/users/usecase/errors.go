package usecase

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrWrongPassword    = errors.New("wrong password")
	ErrUserAlreadyAdmin = errors.New("user's role is already admin")
	ErrUserAlreadyUser  = errors.New("user's role is already user")
	ErrUserEmailExists  = errors.New("user email already exists")
)
