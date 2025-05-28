package telegram

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/gen/api/telegram"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcService) SetTelegramDataV1(
	ctx context.Context,
	req *telegram.SetTelegramDataV1Request,
) (*telegram.SetTelegramDataV1Response, error) {
	_, ok := token.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}

	code, err := s.useCase.SetTelegramData(ctx, req.GetAuthData())
	switch {
	case err != nil:
		s.logger.Error(errors.Wrap(err, "set telegram data v1 failed"))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &telegram.SetTelegramDataV1Response{
		Id:       code.TelegramID,
		Username: code.TelegramUsername,
	}, nil
}
