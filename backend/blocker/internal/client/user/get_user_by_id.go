package user

import (
	"context"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"

	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
)

func (c *Client) GetUserByIDV1(ctx context.Context, userID string) (*userservicegrpc.User, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	resp, err := c.client.GetUserByIDV1(timeoutCtx, &userservicegrpc.GetUserByIDV1Request{
		Id: userID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "client.GetUserByIDV1")
	}

	return resp.GetUser(), nil
}
