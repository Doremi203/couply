package user

import (
	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
)

type OAuthAccount struct {
	UserID         ID
	Provider       oauth.ProviderType
	ProviderUserID oauth.ProviderUserID
}
