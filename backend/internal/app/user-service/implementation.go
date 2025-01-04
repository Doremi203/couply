package user_service

import desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"

type userServiceUseCase interface {
}

type Implementation struct {
	desc.UnimplementedUserServiceServer
	usecase userServiceUseCase
}

func NewImplementation(usecase userServiceUseCase) *Implementation {
	return &Implementation{usecase: usecase}
}
