package postgres

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/tx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewTxProvider(db *pgxpool.Pool) *txProvider {
	return &txProvider{
		db: db,
	}
}

type txProvider struct {
	db *pgxpool.Pool
}

func (p *txProvider) ContextWithTx(ctx context.Context, isolation tx.Isolation) (context.Context, error) {
	existingTx, ok := TxFromContext(ctx)
	if ok {
		if p.toGeneralIsolation(existingTx.Options.IsoLevel) < isolation {
			return nil, errors.Error("existing tx isolation level should be stricter or same as nested tx")
		}
		return ctx, nil
	}

	txOptions := pgx.TxOptions{
		IsoLevel:   p.mapIsolation(isolation),
		AccessMode: pgx.ReadWrite,
	}
	tx, err := p.db.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, errors.WrapFailf(
			err,
			"begin tx with %v",
			errors.Token("isolation", isolation),
		)
	}

	return ContextWithTx(ctx, tx, txOptions), nil
}

func (p *txProvider) CommitTx(ctx context.Context) error {
	tx, ok := TxFromContext(ctx)
	if !ok {
		return errors.Error("no tx to commit")
	}

	return tx.Tx.Commit(ctx)
}

func (p *txProvider) RollbackTx(ctx context.Context) error {
	tx, ok := TxFromContext(ctx)
	if !ok {
		return errors.Error("no tx to rollback")
	}

	return tx.Tx.Rollback(ctx)
}

func (p *txProvider) mapIsolation(isolation tx.Isolation) pgx.TxIsoLevel {
	switch isolation {
	case tx.IsolationSerializable:
		return pgx.Serializable
	case tx.IsolationRepeatableRead:
		return pgx.RepeatableRead
	case tx.IsolationReadCommitted:
		return pgx.ReadCommitted
	case tx.IsolationReadUncommitted:
		return pgx.ReadUncommitted
	default:
		return pgx.ReadCommitted
	}
}

func (p *txProvider) toGeneralIsolation(isolation pgx.TxIsoLevel) tx.Isolation {
	switch isolation {
	case pgx.ReadUncommitted:
		return tx.IsolationReadUncommitted
	case pgx.ReadCommitted:
		return tx.IsolationReadCommitted
	case pgx.RepeatableRead:
		return tx.IsolationRepeatableRead
	case pgx.Serializable:
		return tx.IsolationSerializable
	default:
		return tx.IsolationReadCommitted
	}
}
