package user_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
)

type GetUserV1Request struct{}

type GetUserV1Response struct {
	User *user.User
}

func (x *GetUserV1Response) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func GetUserRequestToPB(_ *GetUserV1Request) *desc.GetUserV1Request {
	return &desc.GetUserV1Request{}
}

func PBToGetUserRequest(_ *desc.GetUserV1Request) *GetUserV1Request {
	return &GetUserV1Request{}
}

func GetUserResponseToPB(resp *GetUserV1Response) *desc.GetUserV1Response {
	return &desc.GetUserV1Response{
		User: user.UserToPB(resp.GetUser()),
	}
}

func PBToGetUserResponse(resp *desc.GetUserV1Response) *GetUserV1Response {
	return &GetUserV1Response{
		User: user.PBToUser(resp.GetUser()),
	}
}
