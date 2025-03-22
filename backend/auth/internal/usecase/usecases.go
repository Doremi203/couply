package usecase

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
)

type Registration interface {
	BasicRegister(ctx context.Context, email user.Email, password user.Password) error
}
