package usecase

import "errors"

var (
	ErrMeteoriteLandingsNotFound     = errors.New("meteorite landings not found")
	ErrMeteoriteLandingAlreadyExists = errors.New("meteorite landing already exists")
)
