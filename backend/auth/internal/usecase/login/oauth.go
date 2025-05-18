package login

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/tx"
)

func (u UseCase) OAuthV1(
	ctx context.Context,
	req oauth.Request,
) (token.Token, error) {
	oauthInfoFetcher, err := u.oauthInfoFetcherFactory.New(req.Provider)
	if err != nil {
		return token.Token{}, errors.WrapFailf(
			err,
			"create oauth info fetcher for %v",
			errors.Token("provider", req.Provider),
		)
	}

	oauthInfo, err := oauthInfoFetcher.Fetch(ctx, req.AccessToken)
	if err != nil {
		return token.Token{}, errors.WrapFailf(
			err,
			"fetch oauth info for user from %v",
			errors.Token("provider", req.Provider),
		)
	}

	usr, err := u.userRepo.GetByOAuthProviderUserID(ctx, req.Provider, oauthInfo.ProviderUserID)
	switch {
	case errors.Is(err, user.ErrNotFound):
		usr, err = u.createUser(ctx, req.Provider, oauthInfo)
		if err != nil {
			return token.Token{}, errors.WrapFailf(
				err,
				"create user with %v",
				errors.Token("provider_user_id", oauthInfo.ProviderUserID),
			)
		}

	case err != nil:
		return token.Token{}, errors.WrapFailf(
			err,
			"get user by %v",
			errors.Token("provider_user_id", oauthInfo.ProviderUserID),
		)
	}

	t, err := u.tokenIssuer.Issue(usr)
	if err != nil {
		return token.Token{}, errors.WrapFailf(err, "issue token")
	}

	return t, nil
}

// Write method for creating a new user
func (u UseCase) createUser(
	ctx context.Context,
	provider oauth.Provider,
	oauthInfo oauth.UserInfo,
) (user.User, error) {
	ctx, err := u.txProvider.ContextWithTx(ctx, tx.IsolationReadCommitted)
	if err != nil {
		return user.User{}, errors.WrapFail(err, "run create user transaction")
	}
	defer func() {
		if err != nil {
			if txErr := u.txProvider.RollbackTx(ctx); txErr != nil {
				u.logger.Error(errors.WrapFail(txErr, "rollback create user transaction"))
			}
		}
	}()

	id, err := u.uuidProvider.GenerateV7()
	if err != nil {
		return user.User{}, errors.WrapFail(err, "generate user id")
	}

	usr := user.User{
		ID:    user.ID(id),
		Email: user.Email(oauthInfo.Email),
		Phone: user.Phone(oauthInfo.Phone),
	}

	err = u.userRepo.Create(ctx, usr)
	if err != nil {
		return user.User{}, errors.WrapFailf(err, "create user")
	}

	accountID, err := u.uuidProvider.GenerateV7()
	if err != nil {
		return user.User{}, errors.WrapFail(err, "generate oauth account id")
	}

	err = u.userOAuthAccountRepo.Create(ctx, user.OAuthAccount{
		ID:             user.OAuthAccountID(accountID),
		UserID:         usr.ID,
		Provider:       provider,
		ProviderUserID: oauthInfo.ProviderUserID,
	})
	if err != nil {
		return user.User{}, errors.WrapFailf(
			err,
			"create oauth account with %v and %v",
			errors.Token("provider", provider),
			errors.Token("provider_user_id", oauthInfo.ProviderUserID),
		)
	}

	err = u.txProvider.CommitTx(ctx)
	if err != nil {
		return user.User{}, errors.WrapFail(err, "commit create user transaction")
	}

	return usr, nil
}
