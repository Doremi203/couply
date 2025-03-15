package user

import "context"

type Repository interface {
	Save(ctx context.Context, user User) error
	GetByEmail(ctx context.Context, email Email) (User, error)
}
