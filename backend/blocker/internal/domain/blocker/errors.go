package blocker

import "github.com/Doremi203/couply/backend/auth/pkg/errors"

var (
	ErrDuplicateUserBlock       = errors.Error("duplicate user block")
	ErrUserBlockNotFound        = errors.Error("user block not found")
	ErrUserBlockReasonsNotFound = errors.Error("user block reasons not found")
)
