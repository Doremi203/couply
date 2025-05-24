package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	CodeUniqueViolation     = "23505"
	CodeForeignKeyViolation = "23503"
)

func IsUniqueViolationError(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == CodeUniqueViolation
}

func IsForeignKeyViolationError(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == CodeForeignKeyViolation
}
