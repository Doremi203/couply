package telegram

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	UpsertTelegramData(
		ctx context.Context,
		data Data,
	) error
	GetTelegramData(
		ctx context.Context,
		userID uuid.UUID,
	) (Data, error)
}
