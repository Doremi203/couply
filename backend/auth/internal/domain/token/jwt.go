package token

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
)

func NewJWTIssuer(cfg JWTConfig) (*jwtIssuer, error) {
	if len(cfg.SecretKey) == 0 {
		return nil, errors.Error("jwt secret key must not be empty")
	}

	return &jwtIssuer{
		cfg: cfg,
	}, nil
}

type jwtIssuer struct {
	cfg JWTConfig
}

type JWTConfig struct {
	SecretKey     string `env:"SECRET_KEY"`
	TokenLifetime time.Duration
}

type сustomClaims struct {
	UserID    string     `json:"user_id"`
	UserEmail user.Email `json:"user_email"`
	jwt.RegisteredClaims
}

func (i *jwtIssuer) Issue(usr user.User) (Token, error) {
	claims := сustomClaims{
		UserID:    usr.ID.String(),
		UserEmail: usr.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(i.cfg.TokenLifetime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
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
