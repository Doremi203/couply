package user_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

type DeleteUserV1Request struct{}

type DeleteUserV1Response struct {
}

func PBToDeleteUserRequest(_ *desc.DeleteUserV1Request) *DeleteUserV1Request {
	return &DeleteUserV1Request{}
}
func DeleteUserResponseToPB(_ *DeleteUserV1Response) *desc.DeleteUserV1Response {
	return &desc.DeleteUserV1Response{}
}
