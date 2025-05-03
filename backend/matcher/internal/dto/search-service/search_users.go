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

type SearchUsersV1Response struct {
	Users []*user.User
}

func SearchUsersRequestToPB(req *SearchUsersV1Request) *desc.SearchUsersV1Request {
	return &desc.SearchUsersV1Request{
		UserId: req.UserID,
		Offset: req.Offset,
		Limit:  req.Limit,
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
		Users: user.UsersToPB(resp.Users),
	}
}

func PBToSearchUsersResponse(resp *desc.SearchUsersV1Response) *SearchUsersV1Response {
	return &SearchUsersV1Response{
		Users: user.PBToUsers(resp.GetUsers()),
	}
}
