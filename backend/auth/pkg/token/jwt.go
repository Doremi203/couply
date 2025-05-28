package token

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func NewJWTProvider(cfg Config) *jwtProvider {
	return &jwtProvider{
		cfg: cfg,
	}
}

type jwtProvider struct {
	cfg Config
}

type сustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func (p *jwtProvider) Parse(tokenString string) (Token, error) {
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Error("unexpected signing method")
		}
		return []byte(p.cfg.SecretKey), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(time.Second*5))
	if err != nil {
		return Token{}, errors.WrapFail(err, "parse token")
	}
	if !jwtToken.Valid {
		return Token{}, errors.Error("token is not valid")
	}

	claims, err := getJWTClaims(jwtToken)
	if err != nil {
		return Token{}, errors.WrapFail(err, "get claims")
	}

	return Token{
		UserID: claims.UserID,
	}, nil
}

func getJWTClaims(t *jwt.Token) (сustomClaims, error) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return сustomClaims{}, errors.Error("invalid claims format")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return сustomClaims{}, errors.Error("user_id claim not found or invalid format")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return сustomClaims{}, errors.Wrap(err, "user id is not uuid")
	}

	var registeredClaims jwt.RegisteredClaims

	expTime, err := claims.GetExpirationTime()
	if err != nil {
		return сustomClaims{}, errors.Error("invalid expiration time")
	}
	registeredClaims.ExpiresAt = expTime

	issuedAt, err := claims.GetIssuedAt()
	if err != nil {
		return сustomClaims{}, errors.Error("invalid issued at time")
	}
	registeredClaims.IssuedAt = issuedAt

	notBefore, err := claims.GetNotBefore()
	if err != nil {
		return сustomClaims{}, errors.Error("invalid not before time")
	}
	registeredClaims.NotBefore = notBefore
	if issuer, err := claims.GetIssuer(); err == nil {
		registeredClaims.Issuer = issuer
	}
	if subject, err := claims.GetSubject(); err == nil {
		registeredClaims.Subject = subject
	}

	return сustomClaims{
		UserID:           userID,
		RegisteredClaims: registeredClaims,
	}, nil
}
