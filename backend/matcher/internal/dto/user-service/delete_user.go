package user_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

type DeleteUserV1Request struct{}

type DeleteUserV1Response struct {
}

func DeleteUserRequestToPB(_ *DeleteUserV1Request) *desc.DeleteUserV1Request {
	return &desc.DeleteUserV1Request{}
}

func PBToDeleteUserRequest(_ *desc.DeleteUserV1Request) *DeleteUserV1Request {
	return &DeleteUserV1Request{}
}

func DeleteUserResponseToPB(_ *DeleteUserV1Response) *desc.DeleteUserV1Response {
	return &desc.DeleteUserV1Response{}
}

func PBToDeleteUserResponse(_ *desc.DeleteUserV1Response) *DeleteUserV1Response {
	return &DeleteUserV1Response{}
}
