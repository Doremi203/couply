package search_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type SearchUsersV1Request struct {
	Offset uint64
	Limit  uint64
}

func (x *SearchUsersV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SearchUsersV1Request) GetLimit() uint64 {
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

func PBToSearchUsersRequest(req *desc.SearchUsersV1Request) *SearchUsersV1Request {
	return &SearchUsersV1Request{
		Offset: req.GetOffset(),
		Limit:  req.GetLimit(),
	}
}

func SearchUsersResponseToPB(resp *SearchUsersV1Response) *desc.SearchUsersV1Response {
	return &desc.SearchUsersV1Response{
		Users: user.UsersToPB(resp.GetUsers()),
	}
}
