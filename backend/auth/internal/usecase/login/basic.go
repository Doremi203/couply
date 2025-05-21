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

// BasicV1 выдает пару access и refresh токены token.Pair авторизации для пользователя, если пользователь с таким user.Email существует и pswrd.Password совпадает с паролем в базе данных.
//
// Если пользователь не существует, возвращает ошибку ErrUserNotRegistered.
//
// Если pswrd.Password не совпадает с паролем в базе данных, возвращает ошибку ErrInvalidCredentials.
func (u UseCase) BasicV1(
	ctx context.Context,
	email user.Email,
	password pswrd.Password,
) (token.Pair, error) {
	usr, err := u.userRepo.GetByAny(ctx, user.GetByAnyParams{Email: email})
	switch {
	case errors.Is(err, user.ErrNotFound):
		return token.Pair{}, ErrUserNotRegistered
	case err != nil:
		return token.Pair{}, errors.WrapFailf(err, "get user by %v", errors.Token("error", email))
	}

	err = u.hashProvider.Verify(string(password), usr.Password)
	switch {
	case errors.Is(err, hash.ErrNoMatch):
		return token.Pair{}, ErrInvalidCredentials
	case err != nil:
		return token.Pair{}, errors.WrapFailf(err, "hash password")
	}

	pair, err := u.tokenIssuer.IssuePair(ctx, usr.ID)
	if err != nil {
		return token.Pair{}, errors.WrapFail(err, "issue access and refresh token pair")
	}

	return pair, nil
}
