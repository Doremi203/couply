package tx

import (
	"context"
)

type Isolation int

const (
	IsolationReadUncommitted Isolation = iota
	IsolationReadCommitted
	IsolationRepeatableRead
	IsolationSerializable
)

type Provider interface {
	ContextWithTx(context.Context, Isolation) (context.Context, error)
	CommitTx(context.Context) error
	RollbackTx(context.Context) error
}
