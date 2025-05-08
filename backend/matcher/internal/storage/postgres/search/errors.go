package search

import "github.com/Doremi203/couply/backend/auth/pkg/errors"

var (
	ErrFilterNotFound      = errors.Error("filter not found")
	ErrDuplicateFilter     = errors.Error("filter already exists for this user")
	ErrInvalidInterestType = errors.Error("invalid interest type")
)
