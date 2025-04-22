package matching

import desc "github.com/Doremi203/couply/backend/matcher/gen/api/matching-service/v1"

type Match struct {
	MainUserID   int64
	ChosenUserID int64
	Approved     bool
}

func MatchToPB(match *Match) *desc.Match {
	return &desc.Match{
		MainUserId:   match.MainUserID,
		ChosenUserId: match.ChosenUserID,
		Approved:     match.Approved,
	}
}

func PBToMatch(match *desc.Match) *Match {
	return &Match{
		MainUserID:   match.MainUserId,
		ChosenUserID: match.ChosenUserId,
		Approved:     match.Approved,
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
