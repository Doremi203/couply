package postgres

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type txKey struct{}

type txData struct {
	Tx      pgx.Tx
	Options pgx.TxOptions
}

func ContextWithTx(ctx context.Context, tx pgx.Tx, options pgx.TxOptions) context.Context {
	return context.WithValue(ctx, txKey{}, txData{Tx: tx, Options: options})
}
func TxFromContext(ctx context.Context) (txData, bool) {
	tx, ok := ctx.Value(txKey{}).(txData)
	return tx, ok
}

type Client interface {
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
}

func NewClient(ctx context.Context, cfg Config) (*client, error) {
	pgPool, err := pgxpool.New(ctx, cfg.ConnectionString())
	if err != nil {
		return nil, errors.WrapFail(err, "create postgres client")
	}

	return &client{
		Pool: pgPool,
	}, nil
}

type client struct {
	*pgxpool.Pool
}

func (c *client) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	tx, ok := TxFromContext(ctx)
	if ok {
		return tx.Tx.Exec(ctx, query, args...)
	}

	return c.Pool.Exec(ctx, query, args...)
}

func (c *client) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	tx, ok := TxFromContext(ctx)
	if ok {
		return tx.Tx.Query(ctx, query, args...)
	}

	return c.Pool.Query(ctx, query, args...)
}

func (c *client) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	tx, ok := TxFromContext(ctx)
	if ok {
		return tx.Tx.QueryRow(ctx, query, args...)
	}

	return c.Pool.QueryRow(ctx, query, args...)
}

func (c *client) Close() error {
	c.Pool.Close()
	return nil
}
