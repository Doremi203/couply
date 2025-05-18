package user

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
)

//go:generate mockgen -source=repo.go -destination=../../mocks/user/repo_mock.go -typed

type Repo interface {
	Create(context.Context, User) error
	UpdatePhone(context.Context, ID, Phone) error
	GetByEmail(context.Context, Email) (User, error)
	GetByPhone(context.Context, Phone) (User, error)
	GetByOAuthProviderUserID(context.Context, oauth.ProviderType, oauth.ProviderUserID) (User, error)
}

type OAuthAccountRepo interface {
	Create(context.Context, OAuthAccount) error
	GetByProviderUserID(context.Context, oauth.ProviderType, oauth.ProviderUserID) (OAuthAccount, error)
	GetByUserID(context.Context, ID) (OAuthAccount, error)
}
