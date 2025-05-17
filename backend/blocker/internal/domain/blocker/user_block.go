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
	CreatedAt time.Time `db:"created_at"`
}

func (x *UserBlock) GetID() uuid.UUID {
	if x != nil {
		return x.ID
	}
	return uuid.Nil
}

func (x *UserBlock) GetBlockedID() uuid.UUID {
	if x != nil {
		return x.BlockedID
	}
	return uuid.Nil
}

func (x *UserBlock) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserBlock) GetReasons() []ReportReason {
	if x != nil {
		return x.Reasons
	}
	return nil
}

func (x *UserBlock) GetCreatedAt() time.Time {
	if x != nil {
		return x.CreatedAt
	}
	return time.Time{}
}
