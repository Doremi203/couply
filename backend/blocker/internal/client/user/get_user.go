package user

import (
	"context"

	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

func (c *Client) GetUserV1(ctx context.Context) (*userservicegrpc.User, error) {
	resp, err := c.client.GetUserV1(ctx, &userservicegrpc.GetUserV1Request{})
	if err != nil {
		return nil, err
	}
	return resp.GetUser(), nil
}
