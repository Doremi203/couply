package matching

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"time"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
)

var (
	ErrLikeNotFound  = errors.Error("like not found")
	ErrLikesNotFound = errors.Error("likes not found")
)

type Like struct {
	SenderID   uuid.UUID `db:"sender_id"`
	ReceiverID uuid.UUID `db:"receiver_id"`
	Message    string    `db:"message"`
	Status     Status    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
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

func LikeToPB(like *Like) *desc.Like {
	return &desc.Like{
		SenderId:   like.SenderID.String(),
		ReceiverId: like.ReceiverID.String(),
		Message:    like.Message,
		Status:     StatusToPB(like.Status),
	}
}
