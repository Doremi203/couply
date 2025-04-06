package user

import (
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
)

type Goal int

const (
	GoalUnspecified Goal = iota
	GoalDating
	GoalRelationship
	GoalFriendship
	GoalJustChatting
)

func PBToGoal(goal desc.Goal) Goal {
	switch goal {
	case desc.Goal_GOAL_UNSPECIFIED:
		return GoalUnspecified
	case desc.Goal_GOAL_DATING:
		return GoalDating
	case desc.Goal_GOAL_RELATIONSHIP:
		return GoalRelationship
	case desc.Goal_GOAL_FRIENDSHIP:
		return GoalFriendship
	case desc.Goal_GOAL_JUST_CHATTING:
		return GoalJustChatting
	default:
		return Goal(0)
	}
}

func GoalToPB(goal Goal) desc.Goal {
	switch goal {
	case GoalUnspecified:
		return desc.Goal_GOAL_UNSPECIFIED
	case GoalDating:
		return desc.Goal_GOAL_DATING
	case GoalRelationship:
		return desc.Goal_GOAL_RELATIONSHIP
	case GoalFriendship:
		return desc.Goal_GOAL_FRIENDSHIP
	case GoalJustChatting:
		return desc.Goal_GOAL_JUST_CHATTING
	default:
		return desc.Goal(0)
	}
}
