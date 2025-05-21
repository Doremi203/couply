package token

import "context"

type Repo interface {
	Get(context.Context, RefreshValue) (Refresh, error)
	Create(context.Context, Refresh) error
	Delete(context.Context, Refresh) error
}
