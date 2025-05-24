package blocker

import (
	"time"

	"github.com/google/uuid"
)

type UserBlock struct {
	ID        uuid.UUID `db:"id"`
	BlockedID uuid.UUID `db:"blocked_id"`
	Message   string    `db:"message"`
	Reasons   []ReportReason
	Status    BlockStatus
	CreatedAt time.Time `db:"created_at"`
}
