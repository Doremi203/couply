package idempotency

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/auth/pkg/tx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RunGRPCHandler[T any](
	ctx context.Context,
	log log.Logger,
	txProvider tx.Provider,
	repo Repo,
	useCase func(context.Context) (T, error),
) (ret T, err error) {
	key, err := FromGRPCCtx(ctx)
	if err != nil {
		return ret, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx, err = txProvider.ContextWithTx(ctx, tx.IsolationReadCommitted)
	if err != nil {
		return ret, errors.WrapFail(err, "inject tx into context")
	}
	defer func() {
		if err != nil {
			rollbackErr := txProvider.RollbackTx(ctx)
			if rollbackErr != nil {
				log.Error(errors.WrapFail(err, "rollback idempotency tx"))
			}
		}
	}()
	defer func() {
		if err == nil {
			err = txProvider.CommitTx(ctx)
			if err != nil {
				err = errors.WrapFail(err, "commit idempotency tx")
			}
		}
	}()

	err = repo.Create(ctx, key)
	switch {
	case errors.Is(err, ErrAlreadyProcessed):
		err = repo.GetData(ctx, key, &ret)
		if err != nil {
			return ret, errors.WrapFail(err, "get existing idempotent request result")
		}
		return ret, nil

	case err != nil:
		return ret, errors.WrapFail(err, "create idempotency request")
	}

	ret, err = useCase(ctx)
	if err != nil {
		return ret, errors.WrapFail(err, "run idempotent useCase")
	}

	err = repo.UpdateData(ctx, key, ret)
	if err != nil {
		return ret, errors.WrapFail(err, "save idempotent request result")
	}

	return ret, nil
}
