package user

import (
	"context"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

func (c *Client) GetUserByIDV1(ctx context.Context, userID string) (*userservicegrpc.User, error) {
	resp, err := c.client.GetUserByIDV1(ctx, &userservicegrpc.GetUserByIDV1Request{
		Id: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetUser(), nil
}
