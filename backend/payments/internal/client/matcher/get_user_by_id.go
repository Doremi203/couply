package matcher

import (
	"context"
	"time"

	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

func (c *Client) GetUserByIDV1(ctx context.Context, userID string) (*userservicegrpc.User, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	resp, err := c.client.GetUserByIDV1(timeoutCtx, &userservicegrpc.GetUserByIDV1Request{
		Id: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetUser(), nil
}
