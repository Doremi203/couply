package password

import (
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
)

//go:generate mockgen -source=hash.go -destination=../../mocks/password/hash_mock.go -typed

type Hasher interface {
	Hash(password user.Password) (user.HashedPassword, error)

	// Verify Сравнивает пароль с хешированным паролем и возвращает ErrInvalidPassword, если они не совпадают.
	Verify(password user.Password, hashedPassword user.HashedPassword) error
}
