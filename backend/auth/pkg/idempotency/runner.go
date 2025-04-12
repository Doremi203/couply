package idempotency

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/tx"
	"google.golang.org/grpc/codes"
	"log/slog"
)

func RunGRPCHandler[T any](
	ctx context.Context,
	log *slog.Logger,
	txProvider tx.Provider,
	repo Repo,
	useCase func(context.Context) (Response[T], error),
) (Response[T], error) {
	key, ok := FromGRPCCtx(ctx)
	if !ok {
		return Response[T]{
			Code:    codes.InvalidArgument,
			Message: "idempotency-key header not provided",
		}, nil
	}

	ctx, err := txProvider.ContextWithTx(ctx, tx.IsolationReadCommitted)
	if err != nil {
		return Response[T]{}, errors.WrapFail(err, "inject tx into context")
	}
	defer func() {
		if err != nil {
			rollbackErr := txProvider.RollbackTx(ctx)
			if rollbackErr != nil {
				log.Error("failed to rollback idempotency tx", "error", err)
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
		var resp Response[T]
		err = repo.GetData(ctx, key, &resp)
		if err != nil {
			return Response[T]{}, errors.WrapFail(err, "get existing idempotent request result")
		}

		return resp, nil
	case err != nil:
		return Response[T]{}, errors.WrapFail(err, "create idempotency request")
	}

	resp, err := useCase(ctx)
	if err != nil {
		return Response[T]{}, errors.WrapFail(err, "run idempotent useCase")
	}

	err = repo.UpdateData(ctx, key, resp)
	if err != nil {
		return Response[T]{}, errors.WrapFail(err, "save idempotent request result")
	}

	return resp, nil
}

type Response[T any] struct {
	Data    *T         `json:"data,omitempty"`
	Code    codes.Code `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
}
