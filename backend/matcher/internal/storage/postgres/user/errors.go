package user

import "github.com/Doremi203/couply/backend/auth/pkg/errors"

var (
	ErrUserNotFound      = errors.Error("user not found")
	ErrPhotoNotFound     = errors.Error("photo not found")
	ErrInterestNotFound  = errors.Error("interest not found")
	ErrDuplicateInterest = errors.Error("interest already exists for this user")
	ErrDuplicatePhoto    = errors.Error("photo with this order number already exists")
	ErrInvalidInterest   = errors.Error("invalid interest type")
)
