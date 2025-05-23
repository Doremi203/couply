package usecase

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/internal/domain/hash"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/uuid"
)

func NewRegistration(
	userRepository user.Repo,
	hashProvider hash.Provider,
	uidGenerator uuid.Provider,
) Registration {
	return Registration{
		userRepository: userRepository,
		hashProvider:   hashProvider,
		uuidProvider:   uidGenerator,
	}
}

type Registration struct {
	userRepository user.Repo
	hashProvider   hash.Provider
	uuidProvider   uuid.Provider
}

var ErrAlreadyRegistered = errors.Error("user already registered")

// BasicV1 создает аккаунт пользователя с переданным user.Email и pswrd.Password.
//
// Если пользователь с таким user.Email уже существует, возвращает ошибку ErrAlreadyRegistered.
func (u Registration) BasicV1(
	ctx context.Context,
	email user.Email,
	password pswrd.Password,
) error {
	_, err := u.userRepository.GetByAny(ctx, user.GetByAnyParams{Email: email})
	switch {
	case err == nil:
		return errors.Wrapf(ErrAlreadyRegistered, "%v already used", errors.Token("email", email))
	case errors.Is(err, user.ErrNotFound):
	// continue
	default:
		return errors.WrapFailf(err, "get existing user with %v", errors.Token("email", email))
	}

	hash, err := u.hashProvider.Hash(string(password))
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
