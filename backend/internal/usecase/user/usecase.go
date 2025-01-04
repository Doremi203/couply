package user

type userStorageFacade interface {
}

type UseCase struct {
	userStorageFacade userStorageFacade
}

func NewUseCase(userStorageFacade userStorageFacade) *UseCase {
	return &UseCase{userStorageFacade: userStorageFacade}
}
