package user

import (
	"context"
)

//go:generate mockgen -source=repo.go -destination=../../mocks/user/repo_mock.go -typed

type Repo interface {
	Create(context.Context, User) error
	GetByEmail(context.Context, Email) (User, error)
}
