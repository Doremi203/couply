package token

import (
	"errors"
	"time"

	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
)

var ErrRefreshTokenNotFound = errors.New("refresh token not found")

type RefreshValue string

type Refresh struct {
	Token     RefreshValue
	UserID    user.ID
	ExpiresAt time.Time
	ExpiresIn time.Duration
}
