package hash

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

//go:generate mockgen -source=hash.go -destination=../../mocks/hash/hash_mock.go -typed

type Provider interface {
	Hash(value string) ([]byte, error)

	// Verify Сравнивает строку с хешированной строкой и возвращает ErrNoMatch, если они не совпадают.
	Verify(value string, hash []byte) error
}

func NewDefaultProvider(
	saltProvider salt.Provider,
	argonProvider argon.Provider,
) defaultProvider {
	return defaultProvider{
		saltProvider:  saltProvider,
		argonProvider: argonProvider,
	}
}

type defaultProvider struct {
	saltProvider  salt.Provider
	argonProvider argon.Provider
}

func (h defaultProvider) Hash(value string) ([]byte, error) {
	s, err := h.saltProvider.Generate(saltLength)
	if err != nil {
		return nil, errors.WrapFail(err, "generate salt")
	}

	return []byte(
		fmt.Sprintf(
			"%s$%s",
			base64.RawStdEncoding.EncodeToString(h.argonProvider.Hash([]byte(value), s)),
			base64.RawStdEncoding.EncodeToString(s),
		),
	), nil
}

func (h defaultProvider) Verify(value string, saltedHash []byte) error {
	parts := strings.Split(string(saltedHash), "$")
	if len(parts) != 2 {
		return errors.Errorf("incorrect format for %v", errors.Token("salted_hash", saltedHash))
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return errors.WrapFail(err, "decode base64 hash")
	}

	s, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return errors.WrapFail(err, "decode base64 salt")
	}

	providedHash := h.argonProvider.Hash([]byte(value), s)
	if subtle.ConstantTimeCompare(providedHash, hash) != 1 {
		return ErrNoMatch
	}

	return nil
}
