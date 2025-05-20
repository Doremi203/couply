package login

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/hash"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

var ErrInvalidCredentials = errors.Error("invalid credentials")
var ErrUserNotRegistered = errors.Error("user not registered")

// BasicV1 выдает токен авторизации для пользователя, если пользователь с таким user.Email существует и pswrd.Password совпадает с паролем в базе данных.
//
// Если пользователь не существует, возвращает ошибку ErrUserNotRegistered.
//
// Если pswrd.Password не совпадает с паролем в базе данных, возвращает ошибку ErrInvalidCredentials.
func (u UseCase) BasicV1(
	ctx context.Context,
	email user.Email,
	password pswrd.Password,
) (token.Token, error) {
	usr, err := u.userRepo.GetByAny(ctx, user.GetByAnyParams{Email: email})
	switch {
	case errors.Is(err, user.ErrNotFound):
		return token.Token{}, ErrUserNotRegistered
	case err != nil:
		return token.Token{}, errors.WrapFailf(err, "get user by %v", errors.Token("error", email))
	}

	err = u.hashProvider.Verify(string(password), usr.Password)
	switch {
	case errors.Is(err, hash.ErrNoMatch):
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
