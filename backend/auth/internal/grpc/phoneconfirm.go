package grpc

import (
	"context"

	phoneconfirm "github.com/Doremi203/couply/backend/auth/gen/api/phone-confirm"
	phoneconfirmdom "github.com/Doremi203/couply/backend/auth/internal/domain/phoneconfirm"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewPhoneConfirmationService(
	phoneConfirmationUseCase usecase.PhoneConfirmation,
	logger log.Logger,
) *phoneConfirmationService {
	return &phoneConfirmationService{
		phoneConfirmationUseCase: phoneConfirmationUseCase,
		logger:                   logger,
	}
}

type phoneConfirmationService struct {
	phoneConfirmationUseCase usecase.PhoneConfirmation

	logger log.Logger
	phoneconfirm.UnimplementedPhoneConfirmationServer
}

func (s *phoneConfirmationService) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return phoneconfirm.RegisterPhoneConfirmationHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *phoneConfirmationService) RegisterToServer(gRPC *grpc.Server) {
	phoneconfirm.RegisterPhoneConfirmationServer(gRPC, s)
}

func (s *phoneConfirmationService) SendCodeV1(
	ctx context.Context,
	req *phoneconfirm.SendCodeV1Request,
) (*phoneconfirm.SendCodeV1Response, error) {
	phone, err := user.NewPhone(req.GetPhone())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}

	confirmReq, err := s.phoneConfirmationUseCase.SendCodeV1(ctx, s.logger, user.ID(t.GetUserID()), phone)
	switch {
	case errors.Is(err, usecase.ErrUnsupportedPhone):
		return nil, status.Error(codes.FailedPrecondition, err.Error())

	case errors.Is(err, usecase.ErrPendingConfirmationRequestAlreadyExists):
		return nil, status.Error(codes.ResourceExhausted, err.Error())

	case errors.Is(err, usecase.ErrPhoneAlreadyConfirmed):
		return nil, status.Error(codes.AlreadyExists, err.Error())

	case err != nil:
		s.logger.Error(errors.Wrap(err, "send code v1 failed"))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &phoneconfirm.SendCodeV1Response{
		SendAgainIn: int32(confirmReq.Code.ExpiresIn.Seconds()),
	}, nil
}

func (s *phoneConfirmationService) ConfirmV1(
	ctx context.Context,
	req *phoneconfirm.ConfirmV1Request,
) (*phoneconfirm.ConfirmV1Response, error) {
	phone, err := user.NewPhone(req.GetPhone())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}
	code, err := phoneconfirmdom.NewCodeValue(req.GetCode())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.phoneConfirmationUseCase.ConfirmV1(ctx, s.logger, user.ID(t.GetUserID()), phone, code)
	switch {
	case errors.Is(err, usecase.ErrNoPendingConfirmRequestExists):
		return nil, status.Error(codes.NotFound, err.Error())

	case errors.Is(err, usecase.ErrConfirmCodeExpired),
		errors.Is(err, usecase.ErrIncorrectConfirmationCode):
		return nil, status.Error(codes.FailedPrecondition, err.Error())

	case errors.Is(err, usecase.ErrPhoneAlreadyConfirmed):
		return nil, status.Error(codes.AlreadyExists, err.Error())

	case err != nil:
		s.logger.Error(errors.Wrap(err, "confirm v1 failed"))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &phoneconfirm.ConfirmV1Response{}, nil
}
