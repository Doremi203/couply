package user

import (
	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
	"github.com/google/uuid"
)

type OAuthAccountID uuid.UUID

type OAuthAccount struct {
	ID             OAuthAccountID
	UserID         ID
	Provider       oauth.Provider
	ProviderUserID oauth.ProviderUserID
}
