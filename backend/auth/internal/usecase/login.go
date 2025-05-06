package usecase

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

func NewLogin(
	userRepo user.Repo,
	hasher pswrd.Hasher,
	tokenIssuer token.Issuer,
) Login {
	return Login{
		userRepo:    userRepo,
		hasher:      hasher,
		tokenIssuer: tokenIssuer,
	}
}

type Login struct {
	userRepo    user.Repo
	hasher      pswrd.Hasher
	tokenIssuer token.Issuer
}

var ErrInvalidCredentials = errors.Error("invalid credentials")

func (u Login) BasicV1(
	ctx context.Context,
	email user.Email,
	password pswrd.Password,
) (token.Token, error) {
	usr, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return token.Token{}, errors.WrapFailf(err, "get user by %v", errors.Token("error", email))
	}

	err = u.hasher.Verify(password, usr.Password)
	switch {
	case errors.Is(err, pswrd.ErrInvalidPassword):
		return token.Token{}, ErrInvalidCredentials
	case err != nil:
		return token.Token{}, errors.WrapFailf(err, "hash password")
	}

	t, err := u.tokenIssuer.Issue(usr)
	if err != nil {
		return token.Token{}, errors.WrapFailf(err, "issue token")
	}

	return t, nil
}
