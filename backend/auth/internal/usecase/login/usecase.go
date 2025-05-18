package login

import (
	"github.com/Doremi203/couply/backend/auth/internal/domain/hash"
	"github.com/Doremi203/couply/backend/auth/internal/domain/oauth"
	"github.com/Doremi203/couply/backend/auth/internal/domain/token"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/auth/pkg/tx"
	"github.com/Doremi203/couply/backend/auth/pkg/uuid"
)

func NewUseCase(
	userRepo user.Repo,
	userOAuthAccountRepo user.OAuthAccountRepo,
	oauthInfoFetcherFactory oauth.InfoFetcherFactory,
	hashProvider hash.Provider,
	tokenIssuer token.Issuer,
	txProvider tx.Provider,
	logger log.Logger,
	uuidProvider uuid.Provider,
) UseCase {
	return UseCase{
		userRepo:                userRepo,
		userOAuthAccountRepo:    userOAuthAccountRepo,
		oauthInfoFetcherFactory: oauthInfoFetcherFactory,
		hashProvider:            hashProvider,
		tokenIssuer:             tokenIssuer,
		txProvider:              txProvider,
		logger:                  logger,
		uuidProvider:            uuidProvider,
	}
}

type UseCase struct {
	userRepo                user.Repo
	userOAuthAccountRepo    user.OAuthAccountRepo
	oauthInfoFetcherFactory oauth.InfoFetcherFactory
	hashProvider            hash.Provider
	tokenIssuer             token.Issuer
	txProvider              tx.Provider
	logger                  log.Logger
	uuidProvider            uuid.Provider
}
