package user_service

import (
	"github.com/Doremi203/Couply/backend/internal/domain/user"
	desc "github.com/Doremi203/Couply/backend/pkg/user-service/v1"
)

type GetUserV1Request struct {
	ID int64
}

type GetUserV1Response struct {
	User *user.User
}

func GetUserRequestToPB(req *GetUserV1Request) *desc.GetUserV1Request {
	return &desc.GetUserV1Request{
		Id: req.ID,
	}
}

func PBToGetUserRequest(req *desc.GetUserV1Request) *GetUserV1Request {
	return &GetUserV1Request{
		ID: req.Id,
	}
}

func GetUserResponseToPB(resp *GetUserV1Response) *desc.GetUserV1Response {
	return &desc.GetUserV1Response{
		User: user.UserToPB(resp.User),
	}
}

func PBToGetUserResponse(resp *desc.GetUserV1Response) *GetUserV1Response {
	return &GetUserV1Response{
		User: user.PBToUser(resp.User),
	}
}
