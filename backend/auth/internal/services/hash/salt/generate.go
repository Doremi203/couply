package salt

import (
	"crypto/rand"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

func Generate(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate random bytes")
	}

	return salt, nil
}
