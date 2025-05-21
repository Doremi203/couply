package token

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/timeprovider"
	"github.com/golang-jwt/jwt/v5"
)

func NewJWTIssuer(
	cfg JWTConfig,
	repo Repo,
	timeProvider timeprovider.Provider,
) (*jwtIssuer, error) {
	if len(cfg.SecretKey) == 0 {
		return nil, errors.Error("jwt secret key must not be empty")
	}

	return &jwtIssuer{
		cfg:          cfg,
		repo:         repo,
		timeProvider: timeProvider,
	}, nil
}

type jwtIssuer struct {
	cfg          JWTConfig
	repo         Repo
	timeProvider timeprovider.Provider
}

type JWTConfig struct {
	SecretKey            string `env:"SECRET_KEY" secret:"secret-key"`
	TokenLifetime        time.Duration
	RefreshTokenLifetime time.Duration
}

type сustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (i *jwtIssuer) IssuePair(ctx context.Context, userID user.ID) (Pair, error) {
	token, err := i.Issue(userID)
	if err != nil {
		return Pair{}, errors.WrapFail(err, "issue token")
	}

	refresh, err := i.IssueRefresh(ctx, userID)
	if err != nil {
		return Pair{}, errors.WrapFail(err, "issue refresh token")
	}

	return Pair{
		AccessToken:  token,
		RefreshToken: refresh,
	}, nil
}

func (i *jwtIssuer) Issue(userID user.ID) (Token, error) {
	claims := сustomClaims{
		UserID: userID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(i.timeProvider.Now().Add(i.cfg.TokenLifetime)),
			IssuedAt:  jwt.NewNumericDate(i.timeProvider.Now()),
			NotBefore: jwt.NewNumericDate(i.timeProvider.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(i.cfg.SecretKey))
	if err != nil {
		return Token{}, err
	}

	return Token{
		signedString: tokenString,
		expiresIn:    i.cfg.TokenLifetime,
	}, nil
}

func (i *jwtIssuer) IssueRefresh(ctx context.Context, userID user.ID) (Refresh, error) {
	tokenValue, err := i.generateSecureToken(32)
	if err != nil {
		return Refresh{}, errors.WrapFail(err, "generate secure token")
	}

	refresh := Refresh{
		Token:     RefreshValue(tokenValue),
		UserID:    userID,
		ExpiresAt: i.timeProvider.Now().Add(i.cfg.RefreshTokenLifetime),
		ExpiresIn: i.cfg.RefreshTokenLifetime,
	}

	err = i.repo.Create(ctx, refresh)
	if err != nil {
		return Refresh{}, errors.WrapFail(err, "create refresh token")
	}

	return refresh, nil
}

func (i *jwtIssuer) generateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b), nil
}
