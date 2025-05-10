package search

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type UserSearchInfo struct {
	User           *user.User
	DistanceToUser int32
}

func (x *UserSearchInfo) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserSearchInfo) GetDistanceToUser() int32 {
	if x != nil {
		return x.DistanceToUser
	}
	return 0
}

func UserSearchInfoToPB(info *UserSearchInfo) *desc.UserSearchInfo {
	return &desc.UserSearchInfo{
		User:           user.UserToPB(info.GetUser()),
		DistanceToUser: info.GetDistanceToUser(),
	}
}
