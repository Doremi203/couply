package salt

import (
	"crypto/rand"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

//go:generate mockgen -source=provider.go -destination=mock/provider_mock.go -typed

type Provider interface {
	Generate(size int) ([]byte, error)
}

type DefaultProvider struct{}

func (DefaultProvider) Generate(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate random bytes")
	}

	return salt, nil
}
