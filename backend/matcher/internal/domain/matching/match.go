package matching

import (
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrMatchAlreadyExists = errors.Error("match already exists")
	ErrMatchesNotFound    = errors.Error("matches not found")
	ErrMatchNotFound      = errors.Error("match not found")
)

type Match struct {
	FirstUserID  uuid.UUID `db:"first_user_id"`
	SecondUserID uuid.UUID `db:"second_user_id"`
	CreatedAt    time.Time `db:"created_at"`
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
