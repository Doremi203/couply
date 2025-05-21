package token

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/timeprovider"
)

func NewUseCase(
	tokenRepo token.Repo,
	tokenIssuer token.Issuer,
	timeProvider timeprovider.Provider,
) UseCase {
	return UseCase{
		tokenRepo:    tokenRepo,
		tokenIssuer:  tokenIssuer,
		timeProvider: timeProvider,
	}
}

type UseCase struct {
	tokenRepo    token.Repo
	tokenIssuer  token.Issuer
	timeProvider timeprovider.Provider
}

var ErrInvalidRefreshToken = errors.Error("invalid refresh token")

// Refresh выдает обновленную пару токенов token.Pair, если token.RefreshValue является действительным refresh токеном.
//
// Если refresh токен недействителен, возвращает ошибку ErrInvalidRefreshToken.
func (u UseCase) Refresh(
	ctx context.Context,
	refreshTokenValue token.RefreshValue,
) (token.Pair, error) {
	refreshToken, err := u.tokenRepo.Get(ctx, refreshTokenValue)
	switch {
	case errors.Is(err, token.ErrRefreshTokenNotFound):
		return token.Pair{}, ErrInvalidRefreshToken
	case err != nil:
		return token.Pair{}, errors.WrapFail(err, "get refresh token")
	}

	if refreshToken.ExpiresAt.Before(u.timeProvider.Now()) {
		err = u.tokenRepo.Delete(ctx, refreshToken)
		if err != nil {
			return token.Pair{}, errors.WrapFail(err, "delete expired refresh token")
		}
		return token.Pair{}, ErrInvalidRefreshToken
	}
	err = u.tokenRepo.Delete(ctx, refreshToken)
	if err != nil {
		return token.Pair{}, errors.WrapFail(err, "delete old refresh token")
	}

	pair, err := u.tokenIssuer.IssuePair(ctx, refreshToken.UserID)
	if err != nil {
		return token.Pair{}, errors.WrapFail(err, "issue new refresh and access token pair")
	}

	return pair, nil
}
