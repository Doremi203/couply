package user

import (
	"context"
	userservicegrpc "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"google.golang.org/grpc/metadata"
)

func (c *Client) GetUserByIDV1(ctx context.Context, userID string, userToken string) (*userservicegrpc.User, error) {
	outgoingMD := metadata.New(map[string]string{
		"user-token": userToken,
	})
	ctx = metadata.NewOutgoingContext(ctx, outgoingMD)

	resp, err := c.client.GetUserByIDV1(ctx, &userservicegrpc.GetUserByIDV1Request{
		Id: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetUser(), nil
}
