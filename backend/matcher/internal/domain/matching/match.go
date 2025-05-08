package matching

import (
	"fmt"
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

func (x *Match) GetFirstUserID() uuid.UUID {
	if x != nil {
		return x.FirstUserID
	}
	return uuid.Nil
}

func (x *Match) GetSecondUserID() uuid.UUID {
	if x != nil {
		return x.SecondUserID
	}
	return uuid.Nil
}

func (x *Match) GetCreatedAt() time.Time {
	if x != nil {
		return x.CreatedAt
	}
	return time.Time{}
}

func MatchToPB(match *Match) *desc.Match {
	return &desc.Match{
		FirstUserId:  match.GetFirstUserID().String(),
		SecondUserId: match.GetSecondUserID().String(),
		CreatedAt:    timestamppb.New(match.GetCreatedAt()),
	}
}

func PBToMatch(match *desc.Match) (*Match, error) {
	firstUserID, err := uuid.Parse(match.GetFirstUserId())
	if err != nil {
		return nil, fmt.Errorf("failed to parse FirstUserID from PB: %v", err)
	}

	secondUserID, err := uuid.Parse(match.GetSecondUserId())
	if err != nil {
		return nil, fmt.Errorf("failed to parse SecondUserID from PB: %v", err)
	}

	return &Match{
		FirstUserID:  firstUserID,
		SecondUserID: secondUserID,
		CreatedAt:    match.GetCreatedAt().AsTime(),
	}, nil
}

func MatchSliceToPB(matches []*Match) []*desc.Match {
	pbMatches := make([]*desc.Match, len(matches))

	for i, match := range matches {
		pbMatches[i] = MatchToPB(match)
	}

	return pbMatches
}

func PBToMatchSlice(pbMatches []*desc.Match) ([]*Match, error) {
	matches := make([]*Match, len(pbMatches))

	for i, match := range pbMatches {
		convertedMatch, err := PBToMatch(match)
		if err != nil {
			return nil, err
		}

		matches[i] = convertedMatch
	}

	return matches, nil
}
