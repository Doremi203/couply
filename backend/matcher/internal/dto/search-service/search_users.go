package search_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/search-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/search"
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
	UsersSearchInfo []*search.UserSearchInfo
}

func (x *SearchUsersV1Response) GetUsersSearchInfo() []*search.UserSearchInfo {
	if x != nil {
		return x.UsersSearchInfo
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
	pbUserSearchInfo := make([]*desc.UserSearchInfo, len(resp.GetUsersSearchInfo()))
	for i, info := range resp.GetUsersSearchInfo() {
		pbUserSearchInfo[i] = search.UserSearchInfoToPB(info)
	}

	return &desc.SearchUsersV1Response{
		UsersSearchInfo: pbUserSearchInfo,
	}
}
