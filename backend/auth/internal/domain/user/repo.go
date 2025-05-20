package user

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
)

//go:generate mockgen -source=repo.go -destination=../../mocks/user/repo_mock.go -typed

type GetByAnyParams struct {
	ID               *ID
	Email            Email
	Phone            Phone
	OAuthUserAccount *oauth.UserAccount
}

func (p GetByAnyParams) AllEmpty() bool {
	return p.ID == nil && p.Email == "" && p.Phone == "" && p.OAuthUserAccount == nil
}

type Repo interface {
	Create(context.Context, User) error
	Upsert(context.Context, User) error
	UpdatePhone(context.Context, ID, Phone) error
	GetByAny(context.Context, GetByAnyParams) (User, error)
}

type OAuthAccountRepo interface {
	Create(context.Context, OAuthAccount) error
	Upsert(context.Context, OAuthAccount) error
	GetByProviderUserID(context.Context, oauth.ProviderType, oauth.ProviderUserID) (OAuthAccount, error)
	GetByUserID(context.Context, ID) (OAuthAccount, error)
}
