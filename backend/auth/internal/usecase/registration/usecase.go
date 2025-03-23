package registration

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/uuid"
)

func NewUseCase(
	userRepository user.Repo,
	hasher pswrd.Hasher,
	uidGenerator uuid.Provider,
) UseCase {
	return UseCase{
		userRepository: userRepository,
		hasher:         hasher,
		uuidProvider:   uidGenerator,
	}
}

type UseCase struct {
	userRepository user.Repo
	hasher         pswrd.Hasher
	uuidProvider   uuid.Provider
}

func (u UseCase) BasicRegister(
	ctx context.Context,
	email user.Email,
	password pswrd.Password,
) error {
	_, err := u.userRepository.GetByEmail(ctx, email)
	switch {
	case err == nil:
		return errors.Wrapf(usecase.ErrAlreadyRegistered, "email already used: %s", email)
	case errors.Is(err, user.ErrNotFound):
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
	switch {
	case errors.Is(err, user.ErrAlreadyExists):
		return nil
	case err != nil:
		return errors.WrapFailf(err, "save user")
	}

	return nil
}
