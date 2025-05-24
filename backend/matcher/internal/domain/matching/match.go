package matching

import (
	"time"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Match struct {
	FirstUserID  uuid.UUID
	SecondUserID uuid.UUID
	CreatedAt    time.Time
}

func NewMatch(firstUserID, secondUserID uuid.UUID) *Match {
	return &Match{
		FirstUserID:  firstUserID,
		SecondUserID: secondUserID,
		CreatedAt:    time.Now(),
	}
}

func MatchToPB(match *Match) *desc.Match {
	return &desc.Match{
		FirstUserId:  match.FirstUserID.String(),
		SecondUserId: match.SecondUserID.String(),
		CreatedAt:    timestamppb.New(match.CreatedAt),
	}
}
