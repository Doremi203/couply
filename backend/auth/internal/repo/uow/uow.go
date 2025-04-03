package uow

import (
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
)

type Provider interface {
	StartUnitOfWork() (UnitOfWork, error)
}

type UnitOfWork interface {
	UserRepo() user.Repo
}
