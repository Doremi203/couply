package user_service

import (
	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
)

type GetUsersV1Request struct {
	UserIDs []uuid.UUID
}

type GetUsersV1Response struct {
	Users []*user.User
}

func PBToGetUsersRequest(req *desc.GetUsersV1Request) (*GetUsersV1Request, error) {
	userIDs := make([]uuid.UUID, 0, len(req.GetUserIds()))
	for _, id := range req.GetUserIds() {
		parsedUserID, err := uuid.Parse(id)
		if err != nil {
			return nil, errors.Wrap(err, "uuid.Parse")
		}
		userIDs = append(userIDs, parsedUserID)
	}

	return &GetUsersV1Request{
		UserIDs: userIDs,
	}, nil
}

func GetUsersResponseToPB(resp *GetUsersV1Response) *desc.GetUsersV1Response {
	pbUsers := make([]*desc.User, 0, len(resp.Users))
	for _, domainUser := range resp.Users {
		pbUsers = append(pbUsers, user.UserToPB(domainUser))
	}

	return &desc.GetUsersV1Response{
		Users: pbUsers,
	}
}
