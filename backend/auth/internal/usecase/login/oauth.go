package login

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/tx"
)

type OAuthV1Response struct {
	TokenPair    token.Pair
	IsFirstLogin bool
}

func (u UseCase) OAuthV1(
	ctx context.Context,
	req oauth.Request,
) (OAuthV1Response, error) {
	oauthProvider, err := u.oauthProviderFactory.New(req.Provider)
	if err != nil {
		return OAuthV1Response{}, errors.WrapFailf(
			err,
			"create oauth info fetcher for %v",
			errors.Token("provider", req.Provider),
		)
	}

	accessToken, err := oauthProvider.ExchangeCodeForToken(
		ctx,
		req.Code,
		req.State,
		req.CodeVerifier,
		req.DeviceID,
	)
	if err != nil {
		return OAuthV1Response{}, errors.WrapFailf(
			err,
			"exchange code for token from %v",
			errors.Token("provider", req.Provider),
		)
	}

	oauthInfo, err := oauthProvider.FetchUserInfo(ctx, accessToken)
	if err != nil {
		return OAuthV1Response{}, errors.WrapFailf(
			err,
			"fetch oauth info for user from %v",
			errors.Token("provider", req.Provider),
		)
	}

	ret := OAuthV1Response{}

	ctx, err = u.txProvider.ContextWithTx(ctx, tx.IsolationReadCommitted)
	if err != nil {
		return OAuthV1Response{}, errors.WrapFail(err, "run create user transaction")
	}
	defer func() {
		if err != nil {
			if txErr := u.txProvider.RollbackTx(ctx); txErr != nil {
				u.logger.Error(errors.WrapFail(txErr, "rollback create user transaction"))
			}
		}
	}()

	usr, err := u.userRepo.GetByAny(
		ctx,
		user.GetByAnyParams{
			Email: user.Email(oauthInfo.Email),
			Phone: user.Phone(oauthInfo.Phone),
			OAuthUserAccount: &oauth.UserAccount{
				Provider:       req.Provider,
				ProviderUserID: oauthInfo.ProviderUserID,
			},
		})
	switch {
	case errors.Is(err, user.ErrNotFound):
		id, err := u.uuidProvider.GenerateV7()
		if err != nil {
			return OAuthV1Response{}, errors.WrapFail(err, "generate user id")
		}
		usr = user.User{
			ID:    user.ID(id),
			Email: user.Email(oauthInfo.Email),
		}
		ret.IsFirstLogin = true

	case err != nil:
		return OAuthV1Response{}, errors.WrapFailf(
			err,
			"get user by %v",
			errors.Token("provider_user_id", oauthInfo.ProviderUserID),
		)
	}
	if usr.Phone == "" {
		usr.Phone = user.Phone(oauthInfo.Phone)
	}

	err = u.userRepo.Upsert(ctx, usr)
	if err != nil {
		return OAuthV1Response{}, errors.WrapFailf(err, "upsert user")
	}

	err = u.userOAuthAccountRepo.Upsert(ctx, user.OAuthAccount{
		UserID:         usr.ID,
		Provider:       req.Provider,
		ProviderUserID: oauthInfo.ProviderUserID,
	})
	if err != nil {
		return OAuthV1Response{}, errors.WrapFailf(
			err,
			"create oauth account with %v and %v",
			errors.Token("provider", req.Provider),
			errors.Token("provider_user_id", oauthInfo.ProviderUserID),
		)
	}

	ret.TokenPair, err = u.tokenIssuer.IssuePair(ctx, usr.ID)
	if err != nil {
		return OAuthV1Response{}, errors.WrapFailf(err, "issue token")
	}

	err = u.txProvider.CommitTx(ctx)
	if err != nil {
		return OAuthV1Response{}, errors.WrapFail(err, "commit oauth transaction")
	}

	return ret, nil
}
