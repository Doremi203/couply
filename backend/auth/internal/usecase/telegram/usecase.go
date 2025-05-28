package telegram

import (
	"github.com/Doremi203/couply/backend/auth/internal/domain/telegram"
	"github.com/Doremi203/couply/backend/auth/pkg/tx"
)

func NewUseCase(
	telegramRepo telegram.Repo,
	txProvider tx.Provider,
) UseCase {
	return UseCase{
		telegramRepo: telegramRepo,
		txProvider:   txProvider,
	}
}

type UseCase struct {
	telegramRepo telegram.Repo
	txProvider   tx.Provider
}
