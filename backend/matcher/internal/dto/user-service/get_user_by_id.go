package user_service

import (
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
)

type GetUserByIDV1Request struct {
	UserID uuid.UUID
}

func (x *GetUserByIDV1Request) GetUserID() uuid.UUID {
	if x != nil {
		return x.UserID
	}
	return uuid.Nil
}

type GetUserByIDV1Response struct {
	User *user.User
}

func (x *GetUserByIDV1Response) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func PBToGetUserByIDRequest(req *desc.GetUserByIDV1Request) (*GetUserByIDV1Request, error) {
	userID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	return &GetUserByIDV1Request{
		UserID: userID,
	}, nil
}

func GetUserByIDResponseToPB(resp *GetUserByIDV1Response) *desc.GetUserByIDV1Response {
	return &desc.GetUserByIDV1Response{
		User: user.UserToPB(resp.GetUser()),
	}
}
