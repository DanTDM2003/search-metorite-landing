package postgres

import (
	"github.com/lib/pq"
)

var (
	ErrDuplicatedKeyCode   = pq.ErrorCode("23505")
	ErrForeignKeyViolation = pq.ErrorCode("23503")
)
