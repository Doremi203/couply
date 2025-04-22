package usecase

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/uuid"
)

func NewRegistration(
	userRepository user.Repo,
	hasher pswrd.Hasher,
	uidGenerator uuid.Provider,
) Registration {
	return Registration{
		userRepository: userRepository,
		hasher:         hasher,
		uuidProvider:   uidGenerator,
	}
}

type Registration struct {
	userRepository user.Repo
	hasher         pswrd.Hasher
	uuidProvider   uuid.Provider
}

var ErrAlreadyRegistered = errors.New("user already registered")

// BasicV1 создает аккаунт пользователя с переданным user.Email и pswrd.Password.
//
// Если пользователь с таким user.Email уже существует, возвращает ошибку ErrAlreadyRegistered.
func (u Registration) BasicV1(
	ctx context.Context,
	email user.Email,
	password pswrd.Password,
) error {
	_, err := u.userRepository.GetByEmail(ctx, email)
	switch {
	case err == nil:
		return errors.Wrapf(ErrAlreadyRegistered, "email already used: %s", email)
	case errors.As(err, &user.NotFoundError{}):
	// continue
	default:
		return errors.WrapFailf(err, "get existing user with email: %s", email)
	}

	hash, err := u.hasher.Hash(password)
	if err != nil {
		return errors.WrapFailf(err, "hash password")
	}

	usr, err := user.New(u.uuidProvider, email, hash)
	if err != nil {
		return errors.WrapFailf(err, "create user")
	}

	err = u.userRepository.Create(ctx, usr)
	if err != nil {
		return errors.WrapFailf(err, "create %v", errors.Token("user", usr))
	}

	return nil
}
