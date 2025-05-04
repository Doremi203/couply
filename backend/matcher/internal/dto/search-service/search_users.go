package search_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type SearchUsersV1Request struct {
	UserID int64
	Offset int32
	Limit  int32
}

func (x *SearchUsersV1Request) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *SearchUsersV1Request) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SearchUsersV1Request) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchUsersV1Response struct {
	Users []*user.User
}

func (x *SearchUsersV1Response) GetUsers() []*user.User {
	if x != nil {
		return x.Users
	}
	return nil
}

func SearchUsersRequestToPB(req *SearchUsersV1Request) *desc.SearchUsersV1Request {
	return &desc.SearchUsersV1Request{
		UserId: req.GetUserID(),
		Offset: req.GetOffset(),
		Limit:  req.GetLimit(),
	}
}

func PBToSearchUsersRequest(req *desc.SearchUsersV1Request) *SearchUsersV1Request {
	return &SearchUsersV1Request{
		UserID: req.GetUserId(),
		Offset: req.GetOffset(),
		Limit:  req.GetLimit(),
	}
}

func SearchUsersResponseToPB(resp *SearchUsersV1Response) *desc.SearchUsersV1Response {
	return &desc.SearchUsersV1Response{
		Users: user.UsersToPB(resp.GetUsers()),
	}
}

func PBToSearchUsersResponse(resp *desc.SearchUsersV1Response) *SearchUsersV1Response {
	return &SearchUsersV1Response{
		Users: user.PBToUsers(resp.GetUsers()),
	}
}
