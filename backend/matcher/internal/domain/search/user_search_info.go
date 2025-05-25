package search

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type UserSearchInfo struct {
	User           *user.User
	DistanceToUser int32
}

func UserSearchInfoToPB(info *UserSearchInfo) *desc.UserSearchInfo {
	return &desc.UserSearchInfo{
		User:           user.UserToPB(info.User),
		DistanceToUser: info.DistanceToUser,
	}
}
