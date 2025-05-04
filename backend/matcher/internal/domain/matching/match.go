package matching

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"

type Match struct {
	MainUserID   int64
	ChosenUserID int64
	Approved     bool
}

func (x *Match) GetMainUserID() int64 {
	if x != nil {
		return x.MainUserID
	}
	return 0
}

func (x *Match) GetChosenUserID() int64 {
	if x != nil {
		return x.ChosenUserID
	}
	return 0
}

func (x *Match) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

func MatchToPB(match *Match) *desc.Match {
	return &desc.Match{
		MainUserId:   match.GetMainUserID(),
		ChosenUserId: match.GetChosenUserID(),
		Approved:     match.GetApproved(),
	}
}

func PBToMatch(match *desc.Match) *Match {
	return &Match{
		MainUserID:   match.GetMainUserId(),
		ChosenUserID: match.GetChosenUserId(),
		Approved:     match.GetApproved(),
	}
}

func MatchSliceToPB(matches []*Match) []*desc.Match {
	pbMatches := make([]*desc.Match, len(matches))

	for i, match := range matches {
		pbMatches[i] = MatchToPB(match)
	}

	return pbMatches
}

func PBToMatchSlice(pbMatches []*desc.Match) []*Match {
	matches := make([]*Match, len(pbMatches))

	for i, match := range pbMatches {
		matches[i] = PBToMatch(match)
	}

	return matches
}
