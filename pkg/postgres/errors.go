package postgres

import (
	"github.com/lib/pq"
)

var (
	ErrDuplicatedKeyCode = pq.ErrorCode("23505")
)
