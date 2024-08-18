package usecase

import "errors"

var (
	ErrArticleTitleAlreadyUsed = errors.New("article' title is already used")
	ErrArticleNotFound         = errors.New("article not found")
)
