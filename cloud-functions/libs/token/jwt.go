package token

import (
	"fmt"
	"time"

	"errors"

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
			return nil, errors.New("unexpected signing method")
		}
		return []byte(p.cfg.SecretKey), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(time.Second*5))
	if err != nil {
		return Token{}, fmt.Errorf("failed to parse token %w", err)
	}
	if !jwtToken.Valid {
		return Token{}, errors.New("token is not valid")
	}

	claims, err := getJWTClaims(jwtToken)
	if err != nil {
		return Token{}, fmt.Errorf("failed to get claims %w", err)
	}

	return Token{
		userID: claims.UserID,
	}, nil
}

func getJWTClaims(t *jwt.Token) (сustomClaims, error) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return сustomClaims{}, errors.New("invalid claims format")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return сustomClaims{}, errors.New("user_id claim not found or invalid format")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return сustomClaims{}, fmt.Errorf("user id is not uuid %w", err)
	}

	var registeredClaims jwt.RegisteredClaims

	expTime, err := claims.GetExpirationTime()
	if err != nil {
		return сustomClaims{}, errors.New("invalid expiration time")
	}
	registeredClaims.ExpiresAt = expTime

	issuedAt, err := claims.GetIssuedAt()
	if err != nil {
		return сustomClaims{}, errors.New("invalid issued at time")
	}
	registeredClaims.IssuedAt = issuedAt

	notBefore, err := claims.GetNotBefore()
	if err != nil {
		return сustomClaims{}, errors.New("invalid not before time")
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
