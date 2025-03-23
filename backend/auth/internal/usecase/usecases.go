package usecase

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
)

type Registration interface {
	// BasicRegister создает аккаунт пользователя с переданным user.Email и pswrd.Password.
	// Операция идемпотентна по user.Email.
	//
	// Если пользователь с таким user.Email уже существует, возвращает ошибку ErrAlreadyRegistered.
	BasicRegister(
		context.Context,
		user.Email,
		pswrd.Password,
	) error
}
