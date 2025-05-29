package telegram

import (
	"context"

	telegram2 "github.com/Doremi203/couply/backend/auth/internal/repo/telegram"

	"github.com/Doremi203/couply/backend/auth/gen/api/telegram"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcService) GetTelegramDataV1(
	ctx context.Context,
	req *telegram.GetTelegramDataV1Request,
) (*telegram.GetTelegramDataV1Response, error) {
	_, ok := token.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}

	url, err := s.useCase.GetTelegramData(ctx, req.GetUserId())
	switch {
	case errors.Is(err, telegram2.ErrTelegramDataNotFound):
		return nil, status.Error(codes.NotFound, telegram2.ErrTelegramDataNotFound.Error())
	case err != nil:
		s.logger.Error(errors.Wrap(err, "get telegram data v1 failed"))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &telegram.GetTelegramDataV1Response{
		TelegramUrl: url,
	}, nil
}
