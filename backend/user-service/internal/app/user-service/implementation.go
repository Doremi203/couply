package user_service

import (
	"context"

	"github.com/Doremi203/Couply/backend/internal/dto"
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
)

type userServiceUseCase interface {
	CreateUser(ctx context.Context, in *dto.CreateUserV1Request) (*dto.CreateUserV1Response, error)
	UpdateUser(ctx context.Context, in *dto.UpdateUserV1Request) (*dto.UpdateUserV1Response, error)
	DeleteUser(ctx context.Context, in *dto.DeleteUserV1Request) (*dto.DeleteUserV1Response, error)
}

type Implementation struct {
	desc.UnimplementedUserServiceServer
	usecase userServiceUseCase
}

func NewImplementation(usecase userServiceUseCase) *Implementation {
	return &Implementation{usecase: usecase}
}
