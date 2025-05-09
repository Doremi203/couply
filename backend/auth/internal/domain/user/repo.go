package user

import (
	"context"
)

//go:generate mockgen -source=repo.go -destination=../../mocks/user/repo_mock.go -typed

type Repo interface {
	Create(context.Context, User) error
	UpdatePhone(context.Context, ID, Phone) error
	GetByEmail(context.Context, Email) (User, error)
	GetByPhone(context.Context, Phone) (User, error)
}
