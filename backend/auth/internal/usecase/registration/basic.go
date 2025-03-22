package registration

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/password"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
)

func NewUseCase(
	userRepository user.Repo,
	hasher password.Hasher,
	uidGenerator user.UIDGenerator,
) UseCase {
	return UseCase{
		userRepository: userRepository,
		hasher:         hasher,
		uidGenerator:   uidGenerator,
	}
}

type UseCase struct {
	userRepository user.Repo
	hasher         password.Hasher
	uidGenerator   user.UIDGenerator
}

func (r UseCase) BasicRegister(ctx context.Context, email user.Email, password user.Password) error {
	hash, err := r.hasher.Hash(password)
	if err != nil {
		return errors.WrapFailf(err, "hash password")
	}

	uid, err := r.uidGenerator.Generate()
	if err != nil {
		return errors.WrapFail(err, "generate user uid")
	}

	err = r.userRepository.Save(ctx, user.User{
		UID:      uid,
		Email:    email,
		Password: hash,
	})
	switch {
	case errors.Is(err, user.ErrAlreadyExists):
		return usecase.ErrAlreadyRegistered
	case err != nil:
		return errors.WrapFailf(err, "save user")
	}

	return nil
}
