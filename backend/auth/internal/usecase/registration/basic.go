package registration

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/password"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/repo"
)

type Basic struct {
	UserRepository user.Repo
	Hasher         password.Hasher
	UIDGenerator   user.UIDGenerator
}

func (r Basic) Run(ctx context.Context, email user.Email, password user.Password) error {
	hash, err := r.Hasher.Hash(password)
	if err != nil {
		return errors.Wrapf(err, "failed to hash password")
	}

	uid, err := r.UIDGenerator.Generate()
	if err != nil {
		return errors.Wrap(err, "failed to generate uid")
	}

	err = r.UserRepository.Save(ctx, user.User{
		UID:      uid,
		Email:    email,
		Password: hash,
	})
	switch {
	case errors.Is(err, repo.ErrAlreadyExists):
		return ErrAlreadyRegistered
	case err != nil:
		return errors.Wrapf(err, "failed to save user")
	}

	return nil
}
