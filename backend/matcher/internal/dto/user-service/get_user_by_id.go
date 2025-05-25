package user_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
)

type GetUserByIDV1Request struct {
	UserID uuid.UUID
}

type GetUserByIDV1Response struct {
	User *user.User
}

func PBToGetUserByIDRequest(req *desc.GetUserByIDV1Request) (*GetUserByIDV1Request, error) {
	userID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "uuid.Parse")
	}

	return &GetUserByIDV1Request{
		UserID: userID,
	}, nil
}

func GetUserByIDResponseToPB(resp *GetUserByIDV1Response) *desc.GetUserByIDV1Response {
	return &desc.GetUserByIDV1Response{
		User: user.UserToPB(resp.User),
	}
}
