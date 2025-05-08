package matching

import "github.com/Doremi203/couply/backend/auth/pkg/errors"

var (
	ErrDuplicateLike      = errors.Error("like already exists between these users")
	ErrMatchAlreadyExists = errors.Error("match already exists between these users")
	ErrUserNotFound       = errors.Error("one or both users not found")
	ErrLikeNotFound       = errors.Error("like not found")
)
