package grpc

import (
	"context"

	"github.com/Doremi203/couply/backend/auth/gen/api/registration"
	"github.com/Doremi203/couply/backend/auth/internal/domain/pswrd"
	"github.com/Doremi203/couply/backend/auth/internal/domain/user"
	"github.com/Doremi203/couply/backend/auth/internal/usecase"
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/idempotency"
	"github.com/Doremi203/couply/backend/auth/pkg/log"
	"github.com/Doremi203/couply/backend/auth/pkg/tx"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRegistrationService(
	useCase usecase.Registration,
	log log.Logger,
	txProvider tx.Provider,
	idempotencyRepo idempotency.Repo,
) *registrationService {
	return &registrationService{
		registrationUseCase: useCase,
		log:                 log,
		txProvider:          txProvider,
		idempotencyRepo:     idempotencyRepo,
	}
}

type registrationService struct {
	registrationUseCase usecase.Registration
	txProvider          tx.Provider
	idempotencyRepo     idempotency.Repo

	log log.Logger
	registration.UnimplementedRegistrationServer
}

func (s *registrationService) RegisterToGateway(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return registration.RegisterRegistrationHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (s *registrationService) RegisterToServer(gRPC *grpc.Server) {
	registration.RegisterRegistrationServer(gRPC, s)
}

func (s *registrationService) BasicRegisterV1(
	ctx context.Context,
	req *registration.BasicRegisterRequestV1,
) (*registration.BasicRegisterResponseV1, error) {
	if err := req.Validate(); err != nil {
		s.log.Warn(errors.Wrap(err, "invalid request"))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	pass, err := pswrd.NewPassword(req.GetPassword())
	if err != nil {
		s.log.Warn(errors.Wrap(err, "invalid password"))
		return nil, status.Errorf(codes.InvalidArgument, "invalid password: %v", err.Error())
	}

	resp, err := idempotency.RunGRPCHandler(
		ctx,
		s.log,
		s.txProvider,
		s.idempotencyRepo,
		func(ctx context.Context) (*registration.BasicRegisterResponseV1, error) {
			err := s.registrationUseCase.BasicV1(ctx, user.Email(req.GetEmail()), pass)
			if err != nil {
				return nil, err
			}

			return &registration.BasicRegisterResponseV1{}, nil
		})
	switch {
	case errors.Is(err, usecase.ErrAlreadyRegistered):
		s.log.Warn(errors.Wrap(err, "user already registered"))
		return nil, status.Errorf(codes.AlreadyExists, "user with email %v already registered", req.GetEmail())
	case err != nil:
		return nil, errors.WrapFail(err, "register user with email idempotently")
	}

	return resp, nil
}
