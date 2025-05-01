package grpc

import (
	"context"
	"fmt"

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
	resp, err := idempotency.RunGRPCHandler(
		ctx,
		s.log,
		s.txProvider,
		s.idempotencyRepo,
		func(ctx context.Context) (idempotency.Response[registration.BasicRegisterResponseV1], error) {
			err := s.registrationUseCase.BasicV1(ctx, user.Email(req.GetEmail()), pswrd.Password(req.GetPassword()))
			switch {
			case errors.Is(err, usecase.ErrAlreadyRegistered):
				s.log.Warn(errors.Wrap(err, "user already registered"))
				return idempotency.Response[registration.BasicRegisterResponseV1]{
					Data:    nil,
					Code:    codes.AlreadyExists,
					Message: fmt.Sprintf("user with email %v already registered", req.GetEmail()),
				}, nil
			case err != nil:
				return idempotency.Response[registration.BasicRegisterResponseV1]{}, err
			}

			return idempotency.Response[registration.BasicRegisterResponseV1]{
				Data: &registration.BasicRegisterResponseV1{},
			}, nil
		})
	if err != nil {
		s.log.Error(errors.Wrap(err, "basic register v1 failed"))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return resp.Data, status.New(resp.Code, resp.Message).Err()
}
