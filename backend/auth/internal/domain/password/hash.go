package password

import (
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
)

type Hasher interface {
	Hash(password user.Password) (user.HashedPassword, error)

	// Verify Сравнивает пароль с хешированным паролем и возвращает ErrInvalidPassword, если они не совпадают.
	Verify(password user.Password, hashedPassword user.HashedPassword) error
}
