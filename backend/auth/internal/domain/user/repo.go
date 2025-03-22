package user

import "context"

//go:generate mockgen -source=repo.go -destination=../../mocks/user/repo_mock.go -typed

type Repo interface {
	Save(ctx context.Context, user User) error
	GetByEmail(ctx context.Context, email Email) (User, error)
}
