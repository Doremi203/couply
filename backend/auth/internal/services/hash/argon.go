package hash

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/Doremi203/couply/backend/auth/internal/domain/password"
	user2 "github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/services/hash/salt"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"golang.org/x/crypto/argon2"
	"log/slog"
	"strings"
)

const (
	timeCost   = 1
	memoryCost = 64 * 1024 // 64 MB
	threads    = 4
	keyLen     = 32 // 32 байта
)

func NewArgon(log *slog.Logger) password.Hasher {
	return &argon{log: log}
}

type argon struct {
	log *slog.Logger
}

func (a *argon) Hash(password user2.Password) (user2.HashedPassword, error) {
	s, err := salt.Generate(16)
	if err != nil {
		return user2.HashedPassword{}, errors.Wrap(err, "failed to generate salt")
	}

	h := argon2.IDKey([]byte(password), s, timeCost, memoryCost, threads, keyLen)

	return user2.HashedPassword(
		fmt.Sprintf(
			"%s$%s",
			base64.RawStdEncoding.EncodeToString(h),
			base64.RawStdEncoding.EncodeToString(s),
		),
	), nil
}

func (a *argon) Verify(password user2.Password, hashedPassword user2.HashedPassword) error {
	parts := strings.Split(string(hashedPassword), "$")
	if len(parts) != 2 {
		return fmt.Errorf("incorrect hashed password format for %v", hashedPassword)
	}

	h, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return errors.Wrap(err, "failed hash base64 decoding")
	}

	s, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return errors.Wrap(err, "failed salt base64 decoding")
	}

	providedHash := argon2.IDKey([]byte(password), s, timeCost, memoryCost, threads, keyLen)
	if subtle.ConstantTimeCompare(providedHash, h) != 1 {
		return user2.ErrInvalidPassword
	}

	return nil
}
