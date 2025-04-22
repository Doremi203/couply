package pswrd

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Doremi203/couply/backend/auth/pkg/argon"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/salt"
)

const (
	saltLength = 16
)

//go:generate mockgen -source=hash.go -destination=../../mocks/password/hash_mock.go -typed

type Hasher interface {
	Hash(password Password) (HashedPassword, error)

	// Verify Сравнивает пароль с хешированным паролем и возвращает ErrInvalidPassword, если они не совпадают.
	Verify(password Password, hashedPassword HashedPassword) error
}

func NewDefaultHasher(
	saltProvider salt.Provider,
	argonProvider argon.Provider,
) DefaultHasher {
	return DefaultHasher{
		saltProvider:  saltProvider,
		argonProvider: argonProvider,
	}
}

type DefaultHasher struct {
	saltProvider  salt.Provider
	argonProvider argon.Provider
}

func (h DefaultHasher) Hash(password Password) (HashedPassword, error) {
	s, err := h.saltProvider.Generate(saltLength)
	if err != nil {
		return HashedPassword{}, errors.Wrap(err, "failed to generate salt")
	}

	return HashedPassword(
		fmt.Sprintf(
			"%s$%s",
			base64.RawStdEncoding.EncodeToString(h.argonProvider.Hash([]byte(password), s)),
			base64.RawStdEncoding.EncodeToString(s),
		),
	), nil
}

func (h DefaultHasher) Verify(password Password, hashedPassword HashedPassword) error {
	parts := strings.Split(string(hashedPassword), "$")
	if len(parts) != 2 {
		return fmt.Errorf("incorrect hashed password format for %v", hashedPassword)
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return errors.Wrap(err, "failed hash base64 decoding")
	}

	s, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return errors.Wrap(err, "failed salt base64 decoding")
	}

	providedHash := h.argonProvider.Hash([]byte(password), s)
	if subtle.ConstantTimeCompare(providedHash, hash) != 1 {
		return ErrInvalidPassword
	}

	return nil
}
