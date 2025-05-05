package token

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/user"
	"github.com/golang-jwt/jwt/v5"
)

func NewJWTProvider(cfg Config) *jwtProvider {
	return &jwtProvider{
		cfg: cfg,
	}
}

type jwtProvider struct {
	cfg Config
}

type Config struct {
	SecretKey string
}

type сustomClaims struct {
	UserID    user.ID    `json:"user_id"`
	UserEmail user.Email `json:"useremail"`
	jwt.RegisteredClaims
}

// Parse parses a token string and returns the token object
func (p *jwtProvider) Parse(tokenString string) (Token, error) {
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Error("unexpected signing method")
		}
		return p.cfg.SecretKey, nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(time.Second*5))
	if err != nil {
		return Token{}, errors.WrapFail(err, "parse token")
	}
	if !jwtToken.Valid {
		return Token{}, ErrInvalidToken
	}

	claims, err := getJWTClaims(jwtToken)
	if err != nil {
		return Token{}, errors.WrapFailf(ErrInvalidToken, "%v", errors.Token("cause", err.Error()))
	}

	return Token{
		userID:    claims.UserID,
		userEmail: claims.UserEmail,
	}, nil
}

func getJWTClaims(t *jwt.Token) (сustomClaims, error) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return сustomClaims{}, errors.Error("invalid claims format")
	}

	userID, ok := claims["user_id"].(user.ID)
	if !ok {
		return сustomClaims{}, errors.Error("user_id claim not found or invalid format")
	}

	email, ok := claims["user_email"].(user.Email)
	if !ok {
		return сustomClaims{}, errors.Error("email claim not found or invalid format")
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
		UserEmail:        email,
		RegisteredClaims: registeredClaims,
	}, nil
}
