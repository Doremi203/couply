package user_service

import (
	"fmt"

	desc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/Doremi203/couply/backend/matcher/internal/domain/user"
	"github.com/google/uuid"
)

type GetUsersV1Request struct {
	UserIDs []uuid.UUID
}

func (x *GetUsersV1Request) GetUserIDs() []uuid.UUID {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

type GetUsersV1Response struct {
	Users []*user.User
}

func (x *GetUsersV1Response) GetUsers() []*user.User {
	if x != nil {
		return x.Users
	}
	return nil
}

func PBToGetUsersRequest(req *desc.GetUsersV1Request) (*GetUsersV1Request, error) {
	userIDs := make([]uuid.UUID, 0, len(req.GetUserIds()))
	for _, id := range req.GetUserIds() {
		parsedUserID, err := uuid.Parse(id)
		if err != nil {
			return nil, fmt.Errorf("parse user id %s error: %v", id, err)
		}
		userIDs = append(userIDs, parsedUserID)
	}

	return &GetUsersV1Request{
		UserIDs: userIDs,
	}, nil
}

func GetUsersResponseToPB(resp *GetUsersV1Response) *desc.GetUsersV1Response {
	pbUsers := make([]*desc.User, 0, len(resp.GetUsers()))
	for _, domainUser := range resp.GetUsers() {
		pbUsers = append(pbUsers, user.UserToPB(domainUser))
	}

	return &desc.GetUsersV1Response{
		Users: pbUsers,
	}
}
