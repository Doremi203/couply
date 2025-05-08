package matching

import (
	"time"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
)

type Like struct {
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	Message    string
	Status     Status
	CreatedAt  time.Time
}

func NewLike(senderID uuid.UUID, receiverID uuid.UUID, message string, status Status) *Like {
	return &Like{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Message:    message,
		Status:     status,
		CreatedAt:  time.Now(),
	}
}

func (x *Like) GetSenderID() uuid.UUID {
	if x != nil {
		return x.SenderID
	}
	return uuid.Nil
}

func (x *Like) GetReceiverID() uuid.UUID {
	if x != nil {
		return x.ReceiverID
	}
	return uuid.Nil
}

func (x *Like) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Like) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status(0)
}

func (x *Like) GetCreatedAt() time.Time {
	if x != nil {
		return x.CreatedAt
	}
	return time.Time{}
}

func LikeToPB(like *Like) *desc.Like {
	return &desc.Like{
		SenderId:   like.SenderID.String(),
		ReceiverId: like.ReceiverID.String(),
		Message:    like.Message,
		Status:     StatusToPB(like.GetStatus()),
	}
}
